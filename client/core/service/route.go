package service

import (
	"context"
	"log"

	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
)

func RouteList(ctx context.Context, client cores.RouteServiceClient) {
	page := pb.Page{
		Limit:   10,
		Offset:  0,
		OrderBy: "name",
	}

	request := &cores.ListRouteRequest{
		Page: &page,
	}

	reply, err := client.List(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func RouteView(ctx context.Context, client cores.RouteServiceClient) {
	request := &pb.Id{Id: "017a053b3f7be81d700e0976"}

	reply, err := client.View(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func RouteViewByName(ctx context.Context, client cores.RouteServiceClient) {
	request := &pb.Name{Name: "Route"}

	reply, err := client.ViewByName(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func RouteCreate(ctx context.Context, client cores.RouteServiceClient) {
	request := &cores.Route{
		Name:   "Route2",
		Desc:   "Route2",
		Src:    "01876f39d884856f39ecc199",
		Dst:    "01876f3a40078570c40886d1",
		Status: consts.ON,
	}

	reply, err := client.Create(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}

func RouteUpdate(ctx context.Context, client cores.RouteServiceClient) {
	request := &cores.Route{
		Id:     "017bb54ddc3ab7314de068d8",
		Name:   "Route",
		Desc:   "Route",
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

func RouteDelete(ctx context.Context, client cores.RouteServiceClient) {
	request := &pb.Id{Id: "017a09b3af1f47a21d6b84c0"}

	reply, err := client.Delete(ctx, request)

	if err != nil {
		log.Fatalf("Error when calling grpc service: %s", err)
	}
	log.Printf("Resp received: %v", reply)
}
