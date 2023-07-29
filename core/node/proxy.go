package node

import (
	"context"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
	"github.com/snple/kokomi/pb/nodes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProxyService struct {
	ns *NodeService

	nodes.UnimplementedProxyServiceServer
}

func newProxyService(ns *NodeService) *ProxyService {
	return &ProxyService{
		ns: ns,
	}
}

func (s *ProxyService) View(ctx context.Context, in *pb.Id) (*pb.Proxy, error) {
	var output pb.Proxy
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetProxy().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *ProxyService) Name(ctx context.Context, in *pb.Name) (*pb.Proxy, error) {
	var output pb.Proxy
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &cores.ProxyNameRequest{DeviceId: deviceID, Name: in.GetName()}

	reply, err := s.ns.Core().GetProxy().Name(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *ProxyService) List(ctx context.Context, in *nodes.ProxyListRequest) (*nodes.ProxyListResponse, error) {
	var err error
	var output nodes.ProxyListResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &cores.ProxyListRequest{
		Page:     in.GetPage(),
		DeviceId: deviceID,
		Tags:     in.GetTags(),
		Type:     in.GetType(),
	}

	reply, err := s.ns.Core().GetProxy().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Proxy = reply.GetProxy()

	return &output, nil
}

func (s *ProxyService) Link(ctx context.Context, in *nodes.ProxyLinkRequest) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &pb.Id{Id: in.GetId()}

	reply, err := s.ns.Core().GetProxy().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	request2 := &cores.ProxyLinkRequest{Id: in.GetId(), Status: in.GetStatus()}

	reply2, err := s.ns.Core().GetProxy().Link(ctx, request2)
	if err != nil {
		return &output, err
	}

	return reply2, nil
}

func (s *ProxyService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Proxy, error) {
	var output pb.Proxy
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetProxy().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *ProxyService) Pull(ctx context.Context, in *nodes.ProxyPullRequest) (*nodes.ProxyPullResponse, error) {
	var err error
	var output nodes.ProxyPullResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &cores.ProxyPullRequest{
		After:    in.GetAfter(),
		Limit:    in.GetLimit(),
		DeviceId: deviceID,
		Type:     in.GetType(),
	}

	reply, err := s.ns.Core().GetProxy().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Proxy = reply.GetProxy()

	return &output, nil
}
