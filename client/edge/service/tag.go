package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/snple/beacon/consts"
	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/edges"
)

func TagList(ctx context.Context, client edges.TagServiceClient) {
	page := pb.Page{
		Limit:   10,
		Offset:  0,
		OrderBy: "name",
		// Search:  "t",
	}

	request := &edges.TagListRequest{
		Page: &page,
		// NodeId: "017a053b3f7be81caa209b8e",
		// SourceId: "017a9b416ef270dbd799c1f5",
		// Tags: "aaa,bbb",
	}

	reply, err := client.List(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func TagView(ctx context.Context, client edges.TagServiceClient) {
	request := &pb.Id{Id: "017a9b416ef270dc2380ce54"}

	reply, err := client.View(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func TagName(ctx context.Context, client edges.TagServiceClient) {
	request := &pb.Name{Name: "TAG"}

	reply, err := client.Name(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func TagCreate(ctx context.Context, client edges.TagServiceClient) {
	request := &pb.Tag{
		SourceId: "0187712e361544594841b6fb",
		Name:     "tag1",
		Desc:     "",
		Address:  "test_address",
		DataType: "F32",
		Status:   consts.ON,
	}

	reply, err := client.Create(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func TagUpdate(ctx context.Context, client edges.TagServiceClient) {
	request := &pb.Tag{
		Id:     "017b213669a4984d98b160a6",
		Name:   "TAG",
		Desc:   "",
		Status: -1,
		Tags:   "aaa,bbb",
	}

	reply, err := client.Update(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func TagDelete(ctx context.Context, client edges.TagServiceClient) {
	request := &pb.Id{Id: "017a9b776bda0c0c7bcd3435"}

	reply, err := client.Delete(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func TagGetValue(ctx context.Context, client edges.TagServiceClient) {
	request := &pb.Id{Id: "01946a5aae65c0ceeaa257db"}

	reply, err := client.GetValue(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func TagSetValue(ctx context.Context, client edges.TagServiceClient) {
	request := &pb.TagValue{
		Id:    "01946a5aae65c0ceeaa257db",
		Value: fmt.Sprintf("%v", rand.Float64()*100),
		// Value: "1",
	}

	reply, err := client.SetValue(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func TagSetWrite(ctx context.Context, client edges.TagServiceClient) {
	request := &pb.TagValue{
		Id:    "01946a5aae65c0ceeaa257db",
		Value: fmt.Sprintf("%v", rand.Float64()*100),
	}

	reply, err := client.SetWrite(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func TagGetWrite(ctx context.Context, client edges.TagServiceClient) {
	request := &pb.Id{Id: "01946a5aae65c0ceeaa257db"}

	reply, err := client.GetWrite(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}
