package service

import (
	"context"
	"log"

	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
)

func DeviceView(ctx context.Context, client edges.DeviceServiceClient) {
	request := &pb.MyEmpty{}

	reply, err := client.View(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func DeviceUpdate(ctx context.Context, client edges.DeviceServiceClient) {
	request := &pb.Device{
		Id:     "0187f0bb5e6cfdd553884496",
		Name:   "device1",
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
