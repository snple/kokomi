package main

import (
	"flag"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/quic-go/quic-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"snple.com/kokomi/bin/core/config"
	"snple.com/kokomi/bin/core/log"
	"snple.com/kokomi/core/core"
	"snple.com/kokomi/core/node"
	"snple.com/kokomi/db"
	"snple.com/kokomi/util"
)

func main() {
	rand.Seed(time.Now().Unix())

	config.Parse()

	log.Init(config.Config.Debug)

	log.Logger.Info("main : Started")
	defer log.Logger.Info("main : Completed")

	command := flag.Arg(0)
	switch command {
	case "seed":
		if err := cli(command); err != nil {
			log.Logger.Sugar().Errorf("error: shutting down: %s", err)
		}

		return
	}

	sqlitedb, err := db.ConnectSqlite(config.Config.DB.File, config.Config.DB.Debug)
	if err != nil {
		log.Logger.Sugar().Fatalf("connecting to db: %v", err)
	}

	defer sqlitedb.Close()

	if err = core.CreateSchema(sqlitedb); err != nil {
		log.Logger.Sugar().Fatalf("create schema: %v", err)
	}

	if err = core.Seed(sqlitedb); err != nil {
		log.Logger.Sugar().Fatalf("seed: %v", err)
	}

	opts := make([]core.CoreOption, 0)

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

		opts = append(opts, core.WithInfluxDB(influxdb))
	}

	cs, err := core.Core(sqlitedb, opts...)
	if err != nil {
		log.Logger.Sugar().Fatalf("NewCoreService: %v", err)
	}

	cs.Start()
	defer cs.Stop()

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
			log.Logger.Sugar().Infof("core grpc start port: %v", config.Config.CoreService.Addr)
			if err := s.Serve(lis); err != nil {
				log.Logger.Sugar().Errorf("failed to serve: %v", err)
			}
		}()
	}

	if config.Config.NodeService.Enable {
		opts := make([]node.NodeOption, 0)
		{
			tlsConfig, err := util.LoadServerCert(config.Config.QuicService.CA, config.Config.QuicService.Cert, config.Config.QuicService.Key)
			if err != nil {
				log.Logger.Sugar().Fatal(err)
			}

			quicConfig := &quic.Config{
				EnableDatagrams: true,
				MaxIdleTimeout:  time.Minute * 3,
			}

			opts = append(opts, node.WithQuic(config.Config.QuicService.Addr, tlsConfig, quicConfig))
		}

		ns, err := node.Node(cs, opts...)
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
			log.Logger.Sugar().Infof("node grpc start port: %v", config.Config.NodeService.Addr)
			if err := s.Serve(lis); err != nil {
				log.Logger.Sugar().Fatalf("failed to serve: %v", err)
			}
		}()
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh
}

func cli(command string) error {
	log.Logger.Sugar().Infof("cli %v: Started", command)
	defer log.Logger.Sugar().Infof("cli %v : Completed", command)

	db, err := db.ConnectSqlite(config.Config.DB.File, config.Config.DB.Debug)
	if err != nil {
		log.Logger.Sugar().Fatalf("connecting to db: %v", err)
	}

	defer db.Close()

	switch command {
	case "seed":
		if err = core.CreateSchema(db); err != nil {
			return err
		}

		if err = core.Seed(db); err != nil {
			return err
		}
	}

	return nil
}
