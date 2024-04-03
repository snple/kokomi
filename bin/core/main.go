package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quic-go/quic-go"
	"github.com/snple/kokomi"
	"github.com/snple/kokomi/bin/core/config"
	"github.com/snple/kokomi/bin/core/log"
	"github.com/snple/kokomi/core"
	"github.com/snple/kokomi/core/plugins/mqtt"
	"github.com/snple/kokomi/db"
	"github.com/snple/kokomi/http"
	"github.com/snple/kokomi/http/core/api"
	"github.com/snple/kokomi/http/core/web"
	"github.com/snple/kokomi/node"
	"github.com/snple/kokomi/util"
	"github.com/snple/kokomi/util/compress/zstd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/keepalive"
)

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "version", "-V":
			fmt.Printf("kokomi core version: %v\n", kokomi.Version)
			return
		}
	}

	rand.Seed(time.Now().Unix())

	config.Parse()

	log.Init(config.Config.Debug)

	log.Logger.Info("main: Started")
	defer log.Logger.Info("main: Completed")

	bundb, err := db.ConnectSqlite(config.Config.DB.File, config.Config.DB.Debug)
	if err != nil {
		log.Logger.Sugar().Fatalf("connecting to db: %v", err)
	}

	defer bundb.Close()

	if err = core.CreateSchema(bundb); err != nil {
		log.Logger.Sugar().Fatalf("create schema: %v", err)
	}

	if err = core.Seed(bundb); err != nil {
		log.Logger.Sugar().Fatalf("seed: %v", err)
	}

	command := flag.Arg(0)
	switch command {
	case "seed":
		log.Logger.Sugar().Infof("seed: Completed")
		return
	}

	coreOpts := make([]core.CoreOption, 0)

	if config.Config.InfluxDB.Enable {
		influxdb, err := db.ConnectInfluxDB(
			config.Config.InfluxDB.Url,
			config.Config.InfluxDB.Org,
			config.Config.InfluxDB.Bucket,
			config.Config.InfluxDB.Token,
		)
		if err != nil {
			log.Logger.Sugar().Fatalf("connecting to db: %v", err)
		}
		defer influxdb.Close()

		coreOpts = append(coreOpts, core.WithInfluxDB(influxdb))
	}

	cs, err := core.Core(bundb, coreOpts...)
	if err != nil {
		log.Logger.Sugar().Fatalf("NewCoreService: %v", err)
	}

	cs.Start()
	defer cs.Stop()

	zstd.Register()

	if config.Config.CoreService.Enable {
		tlsConfig, err := util.LoadServerCert(config.Config.CoreService.CA, config.Config.CoreService.Cert, config.Config.CoreService.Key)
		if err != nil {
			log.Logger.Sugar().Fatal(err)
		}

		lis, err := net.Listen("tcp", config.Config.CoreService.Addr)
		if err != nil {
			log.Logger.Sugar().Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer(
			grpc.Creds(credentials.NewTLS(tlsConfig)),
			grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{PermitWithoutStream: true}),
		)

		cs.Register(s)

		go func() {
			log.Logger.Sugar().Infof("core grpc start: %v", config.Config.CoreService.Addr)
			if err := s.Serve(lis); err != nil {
				log.Logger.Sugar().Errorf("failed to serve: %v", err)
			}
		}()
	}

	if config.Config.NodeService.Enable {
		nodeOpts := make([]node.NodeOption, 0)

		if config.Config.QuicService.Enable {
			tlsConfig, err := util.LoadServerCert(config.Config.QuicService.CA, config.Config.QuicService.Cert, config.Config.QuicService.Key)
			if err != nil {
				log.Logger.Sugar().Fatal(err)
			}

			quicConfig := &quic.Config{
				EnableDatagrams: true,
				MaxIdleTimeout:  time.Minute * 3,
			}

			nodeOpts = append(nodeOpts, node.WithQuic(node.QuicOptions{
				Addr:       config.Config.QuicService.Addr,
				TLSConfig:  tlsConfig,
				QUICConfig: quicConfig,
			}))
		}

		ns, err := node.Node(cs, nodeOpts...)
		if err != nil {
			log.Logger.Sugar().Fatalf("NewNodeService: %v", err)
		}

		ns.Start()
		defer ns.Stop()

		tlsConfig, err := util.LoadServerCert(config.Config.NodeService.CA, config.Config.NodeService.Cert, config.Config.NodeService.Key)
		if err != nil {
			log.Logger.Sugar().Fatal(err)
		}

		lis, err := net.Listen("tcp", config.Config.NodeService.Addr)
		if err != nil {
			log.Logger.Sugar().Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer(
			grpc.Creds(credentials.NewTLS(tlsConfig)),
			grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{PermitWithoutStream: true}),
		)

		ns.RegisterGrpc(s)

		go func() {
			log.Logger.Sugar().Infof("node grpc start: %v", config.Config.NodeService.Addr)
			if err := s.Serve(lis); err != nil {
				log.Logger.Sugar().Fatalf("failed to serve: %v", err)
			}
		}()
	}

	if !config.Config.Gin.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	if config.Config.WebService.Enable {
		opts := make([]http.HttpServerOption, 0)

		opts = append(opts, http.WithAppName("web"))
		opts = append(opts, http.WithAddr(config.Config.WebService.Addr))
		opts = append(opts, http.WithDebug(config.Config.WebService.Debug))

		if config.Config.WebService.TLS {
			if config.Config.WebService.CA != "" {
				pool := x509.NewCertPool()

				ca, err := os.ReadFile(config.Config.WebService.CA)
				if err != nil {
					log.Logger.Sugar().Fatal(err)
				}

				if ok := pool.AppendCertsFromPEM(ca); !ok {
					log.Logger.Sugar().Fatal(err)
				}

				tlsConfig := &tls.Config{
					ClientAuth: tls.RequireAndVerifyClientCert,
					ClientCAs:  pool,
				}

				opts = append(opts, http.WithTLSConfig(tlsConfig))
			}

			opts = append(opts, http.WithTLS(config.Config.WebService.Cert, config.Config.WebService.Key))
		}

		hs, err := http.NewHttpServer(cs.Context(), opts...)
		if err != nil {
			log.Logger.Sugar().Fatalf("NewHttpServer: %v", err)
		}

		{
			ws, err := web.NewWebService(cs)
			if err != nil {
				log.Logger.Sugar().Fatalf("NewWebService: %v", err)
			}

			ws.Register(hs.Engine())

			go ws.Start()
			defer ws.Stop()

			as, err := api.NewApiService(cs)
			if err != nil {
				log.Logger.Sugar().Fatalf("NewApiService: %v", err)
			}

			apiGroup := hs.Engine().Group("/api", ws.GetAuth().MiddlewareFunc())
			as.Register(apiGroup)

			go as.Start()
			defer as.Stop()
		}

		go hs.Start()
		defer hs.Stop()
	}

	if config.Config.ApiService.Enable {
		opts := make([]http.HttpServerOption, 0)

		opts = append(opts, http.WithAppName("api"))
		opts = append(opts, http.WithAddr(config.Config.ApiService.Addr))
		opts = append(opts, http.WithDebug(config.Config.ApiService.Debug))

		if config.Config.ApiService.TLS {
			if config.Config.ApiService.CA != "" {
				pool := x509.NewCertPool()

				ca, err := ioutil.ReadFile(config.Config.ApiService.CA)
				if err != nil {
					log.Logger.Sugar().Fatal(err)
				}

				if ok := pool.AppendCertsFromPEM(ca); !ok {
					log.Logger.Sugar().Fatal(err)
				}

				tlsConfig := &tls.Config{
					ClientAuth: tls.RequireAndVerifyClientCert,
					ClientCAs:  pool,
				}

				opts = append(opts, http.WithTLSConfig(tlsConfig))
			}

			opts = append(opts, http.WithTLS(config.Config.ApiService.Cert, config.Config.ApiService.Key))
		}

		hs, err := http.NewHttpServer(cs.Context(), opts...)
		if err != nil {
			log.Logger.Sugar().Fatalf("NewHttpServer: %v", err)
		}

		{
			as, err := api.NewApiService(cs)
			if err != nil {
				log.Logger.Sugar().Fatalf("NewApiService: %v", err)
			}

			as.Register(hs.Engine())

			go as.Start()
			defer as.Stop()
		}

		go hs.Start()
		defer hs.Stop()
	}

	for _, static := range config.Config.Statics {
		if !static.Enable {
			continue
		}

		engine := gin.New()
		engine.Use(gin.Recovery())

		if config.Config.Gin.Debug {
			engine.Use(gin.LoggerWithWriter(os.Stdout))
		}

		engine.Static("/", static.Path)
		engine.NoRoute(func(ctx *gin.Context) {
			ctx.File(static.Path + "/index.html")
		})

		log.Logger.Sugar().Infof("static server start: %v", static.Addr)

		if static.TLS {
			go engine.RunTLS(static.Addr, static.Cert, static.Key)
		} else {
			go engine.Run(static.Addr)
		}
	}

	if config.Config.MqttService.Enable {
		opts := make([]mqtt.MqttOption, 0)

		opts = append(opts, mqtt.WithCache(config.Config.MqttService.Cache))
		opts = append(opts, mqtt.WithSave(config.Config.MqttService.Save))

		for _, option := range config.Config.MqttListen {
			if option.Enable {
				opts = append(opts, mqtt.WithListenOption(mqtt.ListenOption{
					Addr: option.Addr,
					TLS:  option.TLS,
					CA:   option.CA,
					Cert: option.Cert,
					Key:  option.Key,
					Ws:   option.Ws,
				}))
			}
		}

		ms, err := mqtt.NewMqttService(cs, opts...)
		if err != nil {
			log.Logger.Sugar().Fatalf("NewMqttService: %v", err)
		}

		go ms.Start()
		defer ms.Stop()
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh
}
