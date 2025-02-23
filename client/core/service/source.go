package service

import (
	"context"
	"log"

	"github.com/snple/beacon/consts"
	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/cores"
)

func SourceList(ctx context.Context, client cores.SourceServiceClient) {
	page := pb.Page{
		Limit:   10,
		Offset:  0,
		OrderBy: "name",
	}

	request := &cores.SourceListRequest{
		Page: &page,
	}

	reply, err := client.List(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func SourceView(ctx context.Context, client cores.SourceServiceClient) {
	request := &pb.Id{Id: "01946a0cabdabc925941e98a"}

	reply, err := client.View(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func SourceName(ctx context.Context, client cores.SourceServiceClient) {
	request := &cores.SourceNameRequest{
		DeviceId: "01946a0cabdabc925941e98a",
		Name:     "source",
	}

	reply, err := client.Name(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func SourceCreate(ctx context.Context, client cores.SourceServiceClient) {
	request := &pb.Source{
		DeviceId: "01946a0cabdabc925941e98a",
		Name:     "source",
		Desc:     "source",
		Type:     "type",
		Source:   "source",
		Params:   "params",
		Config:   "config",
		Status:   consts.ON,
		Tags:     "tag1,tag2",
	}

	reply, err := client.Create(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func SourceUpdate(ctx context.Context, client cores.SourceServiceClient) {
	request := &pb.Source{
		Id:     "01946a51cd5bc0cd7a776f35",
		Name:   "source",
		Desc:   "source",
		Type:   "type",
		Source: "source",
		Params: "params",
		Config: "config",
		Status: consts.ON,
		Tags:   "tag1,tag2",
	}

	reply, err := client.Update(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func SourceDelete(ctx context.Context, client cores.SourceServiceClient) {
	request := &pb.Id{Id: "01946a51cd5bc0cd7a776f35"}

	reply, err := client.Delete(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}
