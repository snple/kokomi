package service

import (
	"context"
	"log"

	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
)

func PortList(ctx context.Context, client cores.PortServiceClient) {
	page := pb.Page{
		Limit:   10,
		Offset:  0,
		OrderBy: "",
		Search:  "",
	}

	request := &cores.PortListRequest{
		Page: &page,
		Tags: "",
	}

	reply, err := client.List(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func PortView(ctx context.Context, client cores.PortServiceClient) {
	request := &pb.Id{Id: "017a053b3f7be81caa209b8e"}

	reply, err := client.View(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func PortCreate(ctx context.Context, client cores.PortServiceClient) {
	request := &pb.Port{
		DeviceId: "01867d48ef0e780faad4058e",
		Name:     "test_port",
		Desc:     "test",
		Status:   consts.ON,
		Network:  "tcp",
		Address:  "127.0.0.1:22",
	}

	reply, err := client.Create(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func PortUpdate(ctx context.Context, client cores.PortServiceClient) {
	request := &pb.Port{
		Id:     "017b8d58b001f3aad847d47e",
		Name:   "Port1",
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

func PortDelete(ctx context.Context, client cores.PortServiceClient) {
	request := &pb.Id{Id: "017a0980738df69c17287eab"}

	reply, err := client.Delete(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}
