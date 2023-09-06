package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/snple/kokomi/client/edge/service"
	"github.com/snple/kokomi/pb/edges"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

const (
	address = "127.0.0.1:6010"
)

func main() {
	rand.Seed(time.Now().Unix())

	logger, _ := zap.NewDevelopment()

	logger.Info("main : Started")
	defer logger.Info("main : Completed")

	cfg, err := loadCert()
	if err != nil {
		logger.Fatal(err.Error())
	}

	var kacp = keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(credentials.NewTLS(cfg)),
		// grpc.WithInsecure(),
		// grpc.WithBlock(),
		grpc.WithKeepaliveParams(kacp),
	)
	if err != nil {
		logger.Sugar().Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	_ = ctx

	device := edges.NewDeviceServiceClient(conn)
	service.DeviceView(ctx, device)
	// service.DeviceUpdate(ctx, device)

	// slot := edges.NewSlotServiceClient(conn)
	// service.SlotList(ctx, slot)
	// service.SlotView(ctx, slot)
	// service.SlotName(ctx, slot)
	// service.SlotCreate(ctx, slot)
	// service.SlotUpdate(ctx, slot)
	// service.SlotDelete(ctx, slot)

	source := edges.NewSourceServiceClient(conn)
	service.SourceList(ctx, source)
	// service.SourceView(ctx, source)
	// service.SourceName(ctx, source)
	// service.SourceCreate(ctx, source)
	// service.SourceUpdate(ctx, source)
	// service.SourceDelete(ctx, source)
	// service.SourceLink(ctx, source)

	tag := edges.NewTagServiceClient(conn)
	service.TagList(ctx, tag)
	// service.TagView(ctx, tag)
	// service.TagName(ctx, tag)
	// service.TagCreate(ctx, tag)
	// service.TagUpdate(ctx, tag)
	// service.TagDelete(ctx, tag)
	// t1 := time.Now()
	// for i := 0; i < 10000; i++ {
	service.TagSetValue(ctx, tag)
	// 	service.TagGetValue(ctx, tag)
	// }
	// t2 := time.Now()
	// fmt.Println("t2-t1", t2.Sub(t1))
}

func loadCert() (*tls.Config, error) {
	pool := x509.NewCertPool()

	ca, err := ioutil.ReadFile("certs/ca.crt")
	if err != nil {
		return nil, err
	}

	if ok := pool.AppendCertsFromPEM(ca); !ok {
		return nil, errors.New("pool.AppendCertsFromPEM err")
	}

	cert, err := tls.LoadX509KeyPair("certs/client.crt", "certs/client.key")
	if err != nil {
		return nil, err
	}

	cfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
		// ServerName:         "snple.com",
		RootCAs:            pool,
		InsecureSkipVerify: true,
	}

	return cfg, nil
}
