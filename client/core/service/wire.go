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

func WireList(ctx context.Context, client cores.WireServiceClient) {
	page := pb.Page{
		Limit:   10,
		Offset:  0,
		OrderBy: "name",
	}

	request := &cores.WireListRequest{
		Page: &page,
		// DeviceId: "018383b2f2782ae48c0dabc0",
		// CableId: "01876f39d884856f39ecc199",
		// Tags:     "aaa,bbb",
	}

	reply, err := client.List(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func WireView(ctx context.Context, client cores.WireServiceClient) {
	request := &pb.Id{Id: "017da377b1584a53ea734403"}

	reply, err := client.View(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func WireViewByName(ctx context.Context, client cores.WireServiceClient) {
	request := &cores.WireNameRequest{Name: "Wire"}

	reply, err := client.Name(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func WireViewByNameFull(ctx context.Context, client cores.WireServiceClient) {
	request := &pb.Name{Name: "device1.source1.TX"}

	reply, err := client.NameFull(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func WireCreate(ctx context.Context, client cores.WireServiceClient) {
	request := &pb.Wire{
		CableId:  "01876f3a40078570c40886d1",
		Name:     "wire1",
		Desc:     "desc",
		DataType: "F32",
		Status:   consts.ON,
		Access:   1,
	}

	reply, err := client.Create(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func WireUpdate(ctx context.Context, client cores.WireServiceClient) {
	request := &pb.Wire{
		Id:       "0182b9a0e5b9bde2756f03cb",
		Name:     "Wire",
		Desc:     "",
		DataType: "F32",
		Status:   1,
		Access:   1,
		Save:     1,
		Tags:     "aaa,bbb",
	}

	reply, err := client.Update(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func WireDelete(ctx context.Context, client cores.WireServiceClient) {
	request := &pb.Id{Id: "017a43e9074c21c4ea52e6e1"}

	reply, err := client.Delete(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func WireGetValue(ctx context.Context, client cores.WireServiceClient) {
	request := &pb.Id{Id: "0186d04a6a4f736eb5ac138c"}

	reply, err := client.GetValue(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func WireSetValue(ctx context.Context, client cores.WireServiceClient) {
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

func WireGetValueByName(ctx context.Context, client cores.WireServiceClient) {
	request := &cores.WireGetValueByNameRequest{
		DeviceId: "018383b2f2782ae48c0dabc0",
		Name:     "source1.P5_ME17_VAL",
	}

	reply, err := client.GetValueByName(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func WireSetValueByName(ctx context.Context, client cores.WireServiceClient) {
	request := &cores.WireNameValue{
		DeviceId: "018383b2f2782ae48c0dabc0",
		Name:     "source1.P5_ME17_VAL",
		Value:    fmt.Sprintf("%v", rand.Float64()*100),
		// Value: "1",
	}

	reply, err := client.SetValueByName(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}
