package service

import (
	"context"
	"log"

	"github.com/snple/beacon/pb/cores"
)

func SetNodeUpdated(ctx context.Context, client cores.SyncServiceClient) {
	request := &cores.SyncUpdated{Id: "0189f3d94f0d1579c4e2a817", Updated: 1}

	reply, err := client.SetNodeUpdated(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}
