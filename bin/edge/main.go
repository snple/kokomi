package main

import (
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
	"snple.com/kokomi/bin/edge/config"
	"snple.com/kokomi/bin/edge/log"
	"snple.com/kokomi/db"
	"snple.com/kokomi/edge/edge"
	"snple.com/kokomi/util"
)

func main() {
	rand.Seed(time.Now().Unix())

	config.Parse()

	log.Init(config.Config.Debug)

	log.Logger.Info("main : Started")
	defer log.Logger.Info("main : Completed")

	sqlitedb, err := db.ConnectSqlite(config.Config.DB.File, config.Config.DB.Debug)
	if err != nil {
		log.Logger.Sugar().Fatalf("connecting to db: %v", err)
	}

	defer sqlitedb.Close()

	if err = edge.CreateSchema(sqlitedb); err != nil {
		log.Logger.Sugar().Fatalf("create schema: %v", err)
	}

	opts := make([]edge.EdgeOption, 0)

	{
		opts = append(opts, edge.WithDeviceID(config.Config.DeviceID, config.Config.Secret))
	}

	{
		kacp := keepalive.ClientParameters{
			Time:                120 * time.Second, // send pings every 120 seconds if there is no activity
			Timeout:             10 * time.Second,  // wait 10 second for ping ack before considering the connection dead
			PermitWithoutStream: true,              // send pings even without active streams
		}

		grpcOpts := []grpc.DialOption{
			grpc.WithKeepaliveParams(kacp),
		}

		if config.Config.NodeClient.TLS {
			tlsConfig, err := util.LoadClientCert(
				config.Config.NodeClient.CA,
				config.Config.NodeClient.Cert,
				config.Config.NodeClient.Key,
				config.Config.NodeClient.ServerName,
				config.Config.NodeClient.InsecureSkipVerify,
			)
			if err != nil {
				log.Logger.Sugar().Fatalf("LoadClientCert: %v", err)
			}

			grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
		} else {
			grpcOpts = append(grpcOpts, grpc.WithInsecure())
		}

		opts = append(opts, edge.WithNode(config.Config.NodeClient.Addr, grpcOpts))
	}

	{
		tlsConfig, err := util.LoadClientCert(
			config.Config.QuicClient.CA,
			config.Config.QuicClient.Cert,
			config.Config.QuicClient.Key,
			config.Config.QuicClient.ServerName,
			config.Config.QuicClient.InsecureSkipVerify,
		)
		if err != nil {
			log.Logger.Sugar().Fatalf("LoadClientCert: %v", err)
		}

		quicConfig := &quic.Config{
			EnableDatagrams: true,
			MaxIdleTimeout:  time.Minute * 3,
		}

		opts = append(opts, edge.WithQuic(config.Config.QuicClient.Addr, tlsConfig, quicConfig))
	}

	es, err := edge.Edge(sqlitedb, opts...)
	if err != nil {
		log.Logger.Sugar().Fatalf("NewEdgeService: %v", err)
	}

	es.Start()
	defer es.Stop()

	if config.Config.EdgeService.Enable {
		opts := []grpc.ServerOption{
			grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{PermitWithoutStream: true}),
		}

		if config.Config.EdgeService.TLS {
			tlsConfig, err := util.LoadServerCert(config.Config.EdgeService.CA, config.Config.EdgeService.Cert, config.Config.EdgeService.Key)
			if err != nil {
				log.Logger.Sugar().Fatal(err)
			}

			opts = append(opts, grpc.Creds(credentials.NewTLS(tlsConfig)))
		}

		lis, err := net.Listen("tcp", config.Config.EdgeService.Addr)
		if err != nil {
			log.Logger.Sugar().Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer(opts...)

		es.Register(s)

		go func() {
			log.Logger.Sugar().Infof("edge grpc start: %v, tls: %v", config.Config.EdgeService.Addr, config.Config.EdgeService.TLS)
			if err := s.Serve(lis); err != nil {
				log.Logger.Sugar().Fatalf("failed to serve: %v", err)
			}
		}()
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh
}
