package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/quic-go/quic-go"
	"github.com/snple/kokomi"
	"github.com/snple/kokomi/bin/edge/config"
	"github.com/snple/kokomi/bin/edge/log"
	"github.com/snple/kokomi/db"
	"github.com/snple/kokomi/edge"
	"github.com/snple/kokomi/slot"
	"github.com/snple/kokomi/util"
	"github.com/snple/kokomi/util/compress/zstd"
	"github.com/uptrace/bun"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "version", "-V":
			fmt.Printf("kokomi edge version: %v\n", kokomi.Version)
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

	if err = edge.CreateSchema(bundb); err != nil {
		log.Logger.Sugar().Fatalf("create schema: %v", err)
	}

	command := flag.Arg(0)
	switch command {
	case "seed":
		log.Logger.Sugar().Infof("seed: Completed")
		return
	case "pull", "push":
		if err := cli(command, bundb); err != nil {
			log.Logger.Sugar().Errorf("error: shutting down: %s", err)
		}

		return
	}

	edgeOpts := make([]edge.EdgeOption, 0)

	{
		edgeOpts = append(edgeOpts, edge.WithDeviceID(config.Config.DeviceID, config.Config.Secret))
		edgeOpts = append(edgeOpts, edge.WithLinkTTL(time.Second*time.Duration(config.Config.Status.LinkTTL)))

		edgeOpts = append(edgeOpts, edge.WithSync(edge.SyncOptions{
			TokenRefresh: time.Second * time.Duration(config.Config.Sync.TokenRefresh),
			Link:         time.Second * time.Duration(config.Config.Sync.Link),
			Interval:     time.Second * time.Duration(config.Config.Sync.Interval),
			Realtime:     config.Config.Sync.Realtime,
		}))

		edgeOpts = append(edgeOpts, edge.WithBadger(badger.DefaultOptions("badger")))
	}

	if config.Config.NodeClient.Enable {
		kacp := keepalive.ClientParameters{
			Time:                120 * time.Second, // send pings every 120 seconds if there is no activity
			Timeout:             10 * time.Second,  // wait 10 second for ping ack before considering the connection dead
			PermitWithoutStream: true,              // send pings even without active streams
		}

		grpcOpts := []grpc.DialOption{
			grpc.WithKeepaliveParams(kacp),
			grpc.WithDefaultCallOptions(grpc.UseCompressor(zstd.Name)),
		}

		zstd.Register()

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
			grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
		}

		edgeOpts = append(edgeOpts, edge.WithNode(edge.NodeOptions{
			Enable:      true,
			Addr:        config.Config.NodeClient.Addr,
			GRPCOptions: grpcOpts,
		}))
	}

	if config.Config.QuicClient.Enable {
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

		edgeOpts = append(edgeOpts, edge.WithQuic(edge.QuicOptions{
			Enable:     true,
			Addr:       config.Config.QuicClient.Addr,
			TLSConfig:  tlsConfig,
			QUICConfig: quicConfig,
		}))
	}

	es, err := edge.Edge(bundb, edgeOpts...)
	if err != nil {
		log.Logger.Sugar().Fatalf("NewEdgeService: %v", err)
	}

	es.Start()
	defer es.Stop()

	if config.Config.EdgeService.Enable {
		edgeGrpcOpts := []grpc.ServerOption{
			grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{PermitWithoutStream: true}),
		}

		if config.Config.EdgeService.TLS {
			tlsConfig, err := util.LoadServerCert(config.Config.EdgeService.CA, config.Config.EdgeService.Cert, config.Config.EdgeService.Key)
			if err != nil {
				log.Logger.Sugar().Fatal(err)
			}

			edgeGrpcOpts = append(edgeGrpcOpts, grpc.Creds(credentials.NewTLS(tlsConfig)))
		}

		lis, err := net.Listen("tcp", config.Config.EdgeService.Addr)
		if err != nil {
			log.Logger.Sugar().Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer(edgeGrpcOpts...)

		es.Register(s)

		go func() {
			log.Logger.Sugar().Infof("edge grpc start: %v, tls: %v", config.Config.EdgeService.Addr, config.Config.EdgeService.TLS)
			if err := s.Serve(lis); err != nil {
				log.Logger.Sugar().Fatalf("failed to serve: %v", err)
			}
		}()
	}

	if config.Config.SlotService.Enable {
		slotOpts := make([]slot.SlotOption, 0)

		ns, err := slot.Slot(es, slotOpts...)
		if err != nil {
			log.Logger.Sugar().Fatalf("NewSlotService: %v", err)
		}

		ns.Start()
		defer ns.Stop()

		slotGrpcOpts := []grpc.ServerOption{
			grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{PermitWithoutStream: true}),
		}

		if config.Config.SlotService.TLS {
			tlsConfig, err := util.LoadServerCert(config.Config.SlotService.CA, config.Config.SlotService.Cert, config.Config.SlotService.Key)
			if err != nil {
				log.Logger.Sugar().Fatal(err)
			}

			slotGrpcOpts = append(slotGrpcOpts, grpc.Creds(credentials.NewTLS(tlsConfig)))
		}

		lis, err := net.Listen("tcp", config.Config.SlotService.Addr)
		if err != nil {
			log.Logger.Sugar().Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer(slotGrpcOpts...)

		ns.RegisterGrpc(s)

		go func() {
			log.Logger.Sugar().Infof("slot grpc start: %v", config.Config.SlotService.Addr)
			if err := s.Serve(lis); err != nil {
				log.Logger.Sugar().Fatalf("failed to serve: %v", err)
			}
		}()
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh
}

func cli(command string, bundb *bun.DB) error {
	log.Logger.Sugar().Infof("cli %v: Started", command)
	defer log.Logger.Sugar().Infof("cli %v : Completed", command)

	edgeOpts := make([]edge.EdgeOption, 0)
	edgeOpts = append(edgeOpts, edge.WithDeviceID(config.Config.DeviceID, config.Config.Secret))

	{
		kacp := keepalive.ClientParameters{
			Time:                120 * time.Second, // send pings every 120 seconds if there is no activity
			Timeout:             10 * time.Second,  // wait 10 second for ping ack before considering the connection dead
			PermitWithoutStream: true,              // send pings even without active streams
		}

		grpcOpts := []grpc.DialOption{
			grpc.WithKeepaliveParams(kacp),
			grpc.WithDefaultCallOptions(grpc.UseCompressor(zstd.Name)),
		}

		zstd.Register()

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

		edgeOpts = append(edgeOpts, edge.WithNode(edge.NodeOptions{
			Addr:        config.Config.NodeClient.Addr,
			GRPCOptions: grpcOpts,
		}))
	}

	es, err := edge.Edge(bundb, edgeOpts...)
	if err != nil {
		log.Logger.Sugar().Fatalf("NewEdgeService: %v", err)
	}

	switch command {
	case "push":
		return es.Push()
	case "pull":
		return es.Pull()
	}

	return nil
}
