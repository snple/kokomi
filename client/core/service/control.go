package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
)

func ControlGetTagValue(ctx context.Context, client cores.ControlServiceClient) {
	request := &pb.Id{Id: "01880166c70f451c041bb351"}

	reply, err := client.GetTagValue(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func ControlSetTagValue(ctx context.Context, client cores.ControlServiceClient) {
	request := &pb.TagValue{
		Id: "01880166c70f451c041bb351",
		// Value: fmt.Sprintf("%v", rand.Float64()*100),
		Value: fmt.Sprintf("%v", rand.Int31()),
	}

	reply, err := client.SetTagValue(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}
