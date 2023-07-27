package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/snple/kokomi/client/core/service"
	"github.com/snple/kokomi/pb/cores"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

const (
	address = "127.0.0.1:6006"
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

	device := cores.NewDeviceServiceClient(conn)
	service.DeviceList(ctx, device)
	// service.DeviceView(ctx, device)
	// service.DeviceViewByName(ctx, device)
	// service.DeviceCreate(ctx, device)
	// service.DeviceUpdate(ctx, device)
	// service.DeviceDelete(ctx, device)
	// service.DeviceDestory(ctx, device)
	// service.DeviceClone(ctx, device)

	// port := cores.NewPortServiceClient(conn)
	// service.PortList(ctx, port)
	// service.PortView(ctx, port)
	// service.PortCreate(ctx, port)
	// service.PortUpdate(ctx, port)
	// service.PortDelete(ctx, port)

	// proxy := cores.NewProxyServiceClient(conn)
	// service.ProxyList(ctx, proxy)
	// service.ProxyView(ctx, proxy)
	// service.ProxyCreate(ctx, proxy)
	// service.ProxyUpdate(ctx, proxy)
	// service.ProxyDelete(ctx, proxy)

	cable := cores.NewCableServiceClient(conn)
	service.CableList(ctx, cable)
	// service.CableView(ctx, cable)
	// service.CableViewByName(ctx, cable)
	// service.CableViewByNameFull(ctx, cable)
	// service.CableCreate(ctx, cable)
	// service.CableUpdate(ctx, cable)
	// service.CableDelete(ctx, cable)

	wire := cores.NewWireServiceClient(conn)
	service.WireList(ctx, wire)
	// service.WireView(ctx, wire)
	// service.WireViewByName(ctx, wire)
	// service.WireViewByNameFull(ctx, wire)
	// service.WireCreate(ctx, wire)
	// service.WireUpdate(ctx, wire)
	// service.WireDelete(ctx, wire)
	// service.WireGetValue(ctx, wire)
	// for {
	// 	service.WireSetValue(ctx, wire)
	// 	time.Sleep(time.Second * 5)
	// }
	// service.WireGetValueByName(ctx, wire)
	// service.WireSetValueByName(ctx, wire)

	route := cores.NewRouteServiceClient(conn)
	service.RouteList(ctx, route)
	// service.RouteView(ctx, route)
	// service.RouteViewByName(ctx, route)
	// service.RouteCreate(ctx, route)
	// service.RouteUpdate(ctx, route)
	// service.RouteDelete(ctx, route)
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
