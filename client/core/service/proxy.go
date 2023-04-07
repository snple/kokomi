package service

import (
	"context"
	"log"

	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
)

func ProxyList(ctx context.Context, client cores.ProxyServiceClient) {
	page := pb.Page{
		Limit:   10,
		Offset:  0,
		OrderBy: "",
		Search:  "",
	}

	request := &cores.ListProxyRequest{
		Page: &page,
		Tags: "",
	}

	reply, err := client.List(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func ProxyView(ctx context.Context, client cores.ProxyServiceClient) {
	request := &pb.Id{Id: "017a053b3f7be81caa209b8e"}

	reply, err := client.View(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func ProxyCreate(ctx context.Context, client cores.ProxyServiceClient) {
	request := &pb.Proxy{
		DeviceId: "01867d48be6b780ed1fb9afe",
		Name:     "test_proxy",
		Desc:     "test",
		Status:   consts.ON,
		Network:  "tcp",
		Address:  "127.0.0.1:2222",
		Target:   "01867d590e1a78108f887fc3",
	}

	reply, err := client.Create(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func ProxyUpdate(ctx context.Context, client cores.ProxyServiceClient) {
	request := &pb.Proxy{
		Id:     "017b8d58b001f3aad847d47e",
		Name:   "Proxy1",
		Desc:   "hahaha",
		Status: consts.ON,
		Tags:   "tag1,tag2",
	}

	reply, err := client.Update(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func ProxyDelete(ctx context.Context, client cores.ProxyServiceClient) {
	request := &pb.Id{Id: "017a0980738df69c17287eab"}

	reply, err := client.Delete(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}
