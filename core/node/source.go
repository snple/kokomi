package node

import (
	"context"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
	"github.com/snple/kokomi/pb/nodes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SourceService struct {
	ns *NodeService

	nodes.UnimplementedSourceServiceServer
}

func newSourceService(ns *NodeService) *SourceService {
	return &SourceService{
		ns: ns,
	}
}

func (s *SourceService) Create(ctx context.Context, in *pb.Source) (*pb.Source, error) {
	var output pb.Source
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

	return s.ns.cs.GetSource().Create(ctx, in)
}

func (s *SourceService) Update(ctx context.Context, in *pb.Source) (*pb.Source, error) {
	var output pb.Source
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

	reply, err := s.ns.cs.GetSource().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.cs.GetSource().Update(ctx, in)
}

func (s *SourceService) View(ctx context.Context, in *pb.Id) (*pb.Source, error) {
	var output pb.Source
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

	reply, err := s.ns.cs.GetSource().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *SourceService) ViewByName(ctx context.Context, in *pb.Name) (*pb.Source, error) {
	var output pb.Source
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

	request := &cores.ViewSourceByNameRequest{DeviceId: deviceID, Name: in.GetName()}

	reply, err := s.ns.cs.GetSource().ViewByName(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *SourceService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
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

	reply, err := s.ns.cs.GetSource().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.cs.GetSource().Delete(ctx, in)
}

func (s *SourceService) List(ctx context.Context, in *nodes.ListSourceRequest) (*nodes.ListSourceResponse, error) {
	var err error
	var output nodes.ListSourceResponse

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

	request := &cores.ListSourceRequest{
		Page:     in.GetPage(),
		DeviceId: deviceID,
		Tags:     in.GetTags(),
		Type:     in.GetType(),
		Source:   in.GetSource(),
	}

	reply, err := s.ns.cs.GetSource().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Source = reply.GetSource()

	return &output, nil
}

func (s *SourceService) Link(ctx context.Context, in *nodes.LinkSourceRequest) (*pb.MyBool, error) {
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

	reply, err := s.ns.cs.GetSource().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	request2 := &cores.LinkSourceRequest{Id: in.GetId(), Status: in.GetStatus()}

	return s.ns.cs.GetSource().Link(ctx, request2)
}

func (s *SourceService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Source, error) {
	var output pb.Source
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

	reply, err := s.ns.cs.GetSource().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *SourceService) Pull(ctx context.Context, in *nodes.PullSourceRequest) (*nodes.PullSourceResponse, error) {
	var err error
	var output nodes.PullSourceResponse

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

	request := &cores.PullSourceRequest{
		After:    in.GetAfter(),
		Limit:    in.GetLimit(),
		DeviceId: deviceID,
		Type:     in.GetType(),
		Source:   in.GetSource(),
	}

	reply, err := s.ns.cs.GetSource().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Source = reply.GetSource()

	return &output, nil
}
