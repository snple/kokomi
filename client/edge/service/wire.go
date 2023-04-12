package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
)

func WireGetValue(ctx context.Context, client edges.WireServiceClient) {
	request := &pb.Id{Id: "01876f8c37ac8571f434be3a"}

	reply, err := client.GetValue(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func WireSetValue(ctx context.Context, client edges.WireServiceClient) {
	request := &pb.WireValue{
		Id:    "01876f8c37ac8571f434be3a",
		Value: fmt.Sprintf("%v", rand.Float64()*100),
		// Value: "1",
	}

	reply, err := client.SetValue(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}
