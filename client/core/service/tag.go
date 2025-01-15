package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
)

func TagList(ctx context.Context, client cores.TagServiceClient) {
	page := pb.Page{
		Limit:   10,
		Offset:  0,
		OrderBy: "name",
		// Search:  "t",
	}

	request := &cores.TagListRequest{
		DeviceId: "01946a0cabdabc925941e98a",
		Page:     &page,
	}

	reply, err := client.List(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func TagView(ctx context.Context, client cores.TagServiceClient) {
	request := &pb.Id{Id: "01880166c70f451c041bb351"}

	reply, err := client.View(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func TagName(ctx context.Context, client cores.TagServiceClient) {
	request := &cores.TagNameRequest{
		DeviceId: "01946a0cabdabc925941e98a",
		Name:     "tag",
	}

	reply, err := client.Name(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func TagCreate(ctx context.Context, client cores.TagServiceClient) {
	request := &pb.Tag{
		SourceId: "01946a51cd5bc0cd7a776f35",
		Name:     "tag1",
		Desc:     "",
		Address:  "",
		DataType: "F32",
		Access:   consts.WRITE,
		Status:   consts.ON,
	}

	reply, err := client.Create(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func TagUpdate(ctx context.Context, client cores.TagServiceClient) {
	request := &pb.Tag{
		Id:     "01880166c70f451c041bb351",
		Name:   "TAG",
		Desc:   "",
		Status: consts.ON,
		Tags:   "aaa,bbb",
	}

	reply, err := client.Update(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func TagDelete(ctx context.Context, client cores.TagServiceClient) {
	request := &pb.Id{Id: "01880166c70f451c041bb351"}

	reply, err := client.Delete(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

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
