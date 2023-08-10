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

	"github.com/quic-go/quic-go"
	"github.com/snple/kokomi"
	"github.com/snple/kokomi/bin/core/config"
	"github.com/snple/kokomi/bin/core/log"
	"github.com/snple/kokomi/core/core"
	"github.com/snple/kokomi/core/node"
	"github.com/snple/kokomi/db"
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

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh
}
