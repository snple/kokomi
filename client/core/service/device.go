package service

import (
	"context"
	"log"

	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
)

func DeviceList(ctx context.Context, client cores.DeviceServiceClient) {
	page := pb.Page{
		Limit:   10,
		Offset:  0,
		OrderBy: "",
		Search:  "",
	}

	request := &cores.DeviceListRequest{
		Page: &page,
		Tags: "",
	}

	reply, err := client.List(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func DeviceView(ctx context.Context, client cores.DeviceServiceClient) {
	request := &pb.Id{Id: "017a053b3f7be81caa209b8e"}

	reply, err := client.View(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func DeviceName(ctx context.Context, client cores.DeviceServiceClient) {
	request := &pb.Name{Name: "device"}

	reply, err := client.Name(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func DeviceCreate(ctx context.Context, client cores.DeviceServiceClient) {
	request := &pb.Device{
		Name:   "test_device2",
		Desc:   "test",
		Secret: "123456",
		Status: consts.ON,
	}

	reply, err := client.Create(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func DeviceUpdate(ctx context.Context, client cores.DeviceServiceClient) {
	request := &pb.Device{
		Id:     "01946a0cabdabc925941e98a",
		Name:   "device",
		Desc:   "hahaha",
		Secret: "123456.",
		Status: consts.ON,
		Tags:   "tag1,tag2",
	}

	reply, err := client.Update(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func DeviceDelete(ctx context.Context, client cores.DeviceServiceClient) {
	request := &pb.Id{Id: "017a0980738df69c17287eab"}

	reply, err := client.Delete(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func DeviceDestory(ctx context.Context, client cores.DeviceServiceClient) {
	request := &pb.Id{Id: "0189f3d94f0d1579c4e2a817"}

	reply, err := client.Destory(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func DeviceClone(ctx context.Context, client cores.DeviceServiceClient) {
	request := &pb.Id{Id: "0187f0bb5e6cfdd553884496"}

	reply, err := client.Clone(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}
