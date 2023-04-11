package service

import (
	"context"
	"log"

	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
)

func CableList(ctx context.Context, client cores.CableServiceClient) {
	page := pb.Page{
		Limit:   10,
		Offset:  0,
		OrderBy: "name",
	}

	request := &cores.ListCableRequest{
		Page:     &page,
		DeviceId: "01867d48be6b780ed1fb9afe",
	}

	reply, err := client.List(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func CableView(ctx context.Context, client cores.CableServiceClient) {
	request := &pb.Id{Id: "017a053b3f7be81d700e0976"}

	reply, err := client.View(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func CableViewByName(ctx context.Context, client cores.CableServiceClient) {
	request := &cores.ViewCableByNameRequest{Name: "Cable"}

	reply, err := client.ViewByName(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func CableViewByNameFull(ctx context.Context, client cores.CableServiceClient) {
	request := &pb.Name{Name: "device1.Cable1"}

	reply, err := client.ViewByNameFull(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func CableCreate(ctx context.Context, client cores.CableServiceClient) {
	request := &pb.Cable{
		DeviceId: "01867d48ef0e780faad4058e",
		Name:     "Cable1",
		Desc:     "Cable1",
		Status:   consts.ON,
	}

	reply, err := client.Create(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func CableUpdate(ctx context.Context, client cores.CableServiceClient) {
	request := &pb.Cable{
		Id:     "017bb54ddc3ab7314de068d8",
		Name:   "Cable",
		Desc:   "Cable",
		Type:   "type",
		Config: "config",
		Tags:   "tag1,tag2",
	}

	reply, err := client.Update(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func CableDelete(ctx context.Context, client cores.CableServiceClient) {
	request := &pb.Id{Id: "017a09b3af1f47a21d6b84c0"}

	reply, err := client.Delete(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}
