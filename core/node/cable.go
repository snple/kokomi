package node

import (
	"context"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
	"github.com/snple/kokomi/pb/nodes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CableService struct {
	ns *NodeService

	nodes.UnimplementedCableServiceServer
}

func newCableService(ns *NodeService) *CableService {
	return &CableService{
		ns: ns,
	}
}

func (s *CableService) Create(ctx context.Context, in *pb.Cable) (*pb.Cable, error) {
	var output pb.Cable
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

	in.DeviceId = deviceID

	return s.ns.Core().GetCable().Create(ctx, in)
}

func (s *CableService) Update(ctx context.Context, in *pb.Cable) (*pb.Cable, error) {
	var output pb.Cable
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

	reply, err := s.ns.Core().GetCable().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetCable().Update(ctx, in)
}

func (s *CableService) View(ctx context.Context, in *pb.Id) (*pb.Cable, error) {
	var output pb.Cable
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

	reply, err := s.ns.Core().GetCable().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *CableService) Name(ctx context.Context, in *pb.Name) (*pb.Cable, error) {
	var output pb.Cable
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

	request := &cores.CableNameRequest{DeviceId: deviceID, Name: in.GetName()}

	reply, err := s.ns.Core().GetCable().Name(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *CableService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

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

	reply, err := s.ns.Core().GetCable().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetCable().Delete(ctx, in)
}

func (s *CableService) List(ctx context.Context, in *nodes.CableListRequest) (*nodes.CableListResponse, error) {
	var err error
	var output nodes.CableListResponse

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

	request := &cores.CableListRequest{
		Page:     in.GetPage(),
		DeviceId: deviceID,
		Tags:     in.GetTags(),
		Type:     in.GetType(),
	}

	reply, err := s.ns.Core().GetCable().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Cable = reply.GetCable()

	return &output, nil
}

func (s *CableService) Link(ctx context.Context, in *nodes.CableLinkRequest) (*pb.MyBool, error) {
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

	reply, err := s.ns.Core().GetCable().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	request2 := &cores.CableLinkRequest{Id: in.GetId(), Status: in.GetStatus()}

	return s.ns.Core().GetCable().Link(ctx, request2)
}

func (s *CableService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Cable, error) {
	var output pb.Cable
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

	reply, err := s.ns.Core().GetCable().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *CableService) Pull(ctx context.Context, in *nodes.CablePullRequest) (*nodes.CablePullResponse, error) {
	var err error
	var output nodes.CablePullResponse

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

	request := &cores.CablePullRequest{
		After:    in.GetAfter(),
		Limit:    in.GetLimit(),
		DeviceId: deviceID,
		Type:     in.GetType(),
	}

	reply, err := s.ns.Core().GetCable().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Cable = reply.GetCable()

	return &output, nil
}

func (s *CableService) Sync(ctx context.Context, in *pb.Cable) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

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

	in.DeviceId = deviceID

	return s.ns.Core().GetCable().Sync(ctx, in)
}
