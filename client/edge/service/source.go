package service

import (
	"context"
	"log"

	"github.com/snple/beacon/consts"
	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/edges"
)

func SourceList(ctx context.Context, client edges.SourceServiceClient) {
	page := pb.Page{
		Limit:   10,
		Offset:  0,
		OrderBy: "name",
	}

	request := &edges.SourceListRequest{
		Page: &page,
	}

	reply, err := client.List(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func SourceView(ctx context.Context, client edges.SourceServiceClient) {
	request := &pb.Id{Id: "017a053b3f7be81d700e0976"}

	reply, err := client.View(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func SourceName(ctx context.Context, client edges.SourceServiceClient) {
	request := &pb.Name{Name: "Source"}

	reply, err := client.Name(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func SourceCreate(ctx context.Context, client edges.SourceServiceClient) {
	request := &pb.Source{
		Name:   "source1",
		Desc:   "source1",
		Source: "source",
		Params: "params",
		Config: "config",
		Status: consts.ON,
		Tags:   "tag1,tag2",
	}

	reply, err := client.Create(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func SourceUpdate(ctx context.Context, client edges.SourceServiceClient) {
	request := &pb.Id{Id: "018849d40fdd9a6bfb3e5d54"}

	reply, err := client.View(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)

	reply.Status = consts.ON

	reply2, err := client.Update(ctx, reply)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply2)
}

func SourceDelete(ctx context.Context, client edges.SourceServiceClient) {
	request := &pb.Id{Id: "017a09b3af1f47a21d6b84c0"}

	reply, err := client.Delete(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func SourceLink(ctx context.Context, client edges.SourceServiceClient) {
	request := &edges.SourceLinkRequest{Id: "01800806d7aa2d9ea7486720", Status: consts.ON}

	reply, err := client.Link(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}
