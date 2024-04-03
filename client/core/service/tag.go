package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
)

func TagGetValue(ctx context.Context, client cores.TagServiceClient) {
	request := &pb.Id{Id: "01880166c70f451c041bb351"}

	reply, err := client.GetValue(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func TagSetValue(ctx context.Context, client cores.TagServiceClient) {
	request := &pb.TagValue{
		Id: "01880166c70f451c041bb351",
		// Value: fmt.Sprintf("%v", rand.Float64()*100),
		Value: fmt.Sprintf("%v", rand.Int31()),
		// Value: "1",
	}

	reply, err := client.SetValue(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}
