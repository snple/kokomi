package node

import (
	"context"

	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/cores"
	"github.com/snple/beacon/pb/nodes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ConstService struct {
	ns *NodeService

	nodes.UnimplementedConstServiceServer
}

func newConstService(ns *NodeService) *ConstService {
	return &ConstService{
		ns: ns,
	}
}

func (s *ConstService) Create(ctx context.Context, in *pb.Const) (*pb.Const, error) {
	var output pb.Const
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	nodeID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	in.NodeId = nodeID

	return s.ns.Core().GetConst().Create(ctx, in)
}

func (s *ConstService) Update(ctx context.Context, in *pb.Const) (*pb.Const, error) {
	var output pb.Const
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	nodeID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &pb.Id{Id: in.GetId()}

	reply, err := s.ns.Core().GetConst().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetNodeId() != nodeID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetNodeId() != nodeID")
	}

	return s.ns.Core().GetConst().Update(ctx, in)
}

func (s *ConstService) View(ctx context.Context, in *pb.Id) (*pb.Const, error) {
	var output pb.Const
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	nodeID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetConst().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetNodeId() != nodeID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetNodeId() != nodeID")
	}

	return reply, nil
}

func (s *ConstService) Name(ctx context.Context, in *pb.Name) (*pb.Const, error) {
	var output pb.Const
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	nodeID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &cores.ConstNameRequest{NodeId: nodeID, Name: in.GetName()}

	reply, err := s.ns.Core().GetConst().Name(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetNodeId() != nodeID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetNodeId() != nodeID")
	}

	return reply, nil
}

func (s *ConstService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	nodeID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetConst().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetNodeId() != nodeID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetNodeId() != nodeID")
	}

	return s.ns.Core().GetConst().Delete(ctx, in)
}

func (s *ConstService) List(ctx context.Context, in *nodes.ConstListRequest) (*nodes.ConstListResponse, error) {
	var err error
	var output nodes.ConstListResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	nodeID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &cores.ConstListRequest{
		Page:   in.GetPage(),
		NodeId: nodeID,
		Tags:   in.GetTags(),
	}

	reply, err := s.ns.Core().GetConst().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Const = reply.GetConst()

	return &output, nil
}

func (s *ConstService) GetValue(ctx context.Context, in *pb.Id) (*pb.ConstValue, error) {
	var err error
	var output pb.ConstValue

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	nodeID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &pb.Id{Id: in.GetId()}

	reply, err := s.ns.Core().GetConst().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetNodeId() != nodeID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetNodeId() != nodeID")
	}

	return s.ns.Core().GetConst().GetValue(ctx, in)
}

func (s *ConstService) SetValue(ctx context.Context, in *pb.ConstValue) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	nodeID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &pb.Id{Id: in.GetId()}

	reply, err := s.ns.Core().GetConst().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetNodeId() != nodeID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetNodeId() != nodeID")
	}

	return s.ns.Core().GetConst().SetValue(ctx, in)
}

func (s *ConstService) GetValueByName(ctx context.Context, in *pb.Name) (*pb.ConstNameValue, error) {
	var err error
	var output pb.ConstNameValue

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	nodeID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetConst().GetValueByName(ctx,
		&cores.ConstGetValueByNameRequest{NodeId: nodeID, Name: in.GetName()})
	if err != nil {
		return &output, err
	}

	output.Id = reply.GetId()
	output.Name = reply.GetName()
	output.Value = reply.GetValue()
	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *ConstService) SetValueByName(ctx context.Context, in *pb.ConstNameValue) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	nodeID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ns.Core().GetConst().SetValueByName(ctx,
		&cores.ConstNameValue{NodeId: nodeID, Name: in.GetName(), Value: in.GetValue()})
}

func (s *ConstService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Const, error) {
	var output pb.Const
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	nodeID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetConst().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetNodeId() != nodeID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetNodeId() != nodeID")
	}

	return reply, nil
}

func (s *ConstService) Pull(ctx context.Context, in *nodes.ConstPullRequest) (*nodes.ConstPullResponse, error) {
	var err error
	var output nodes.ConstPullResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	nodeID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &cores.ConstPullRequest{
		After:  in.GetAfter(),
		Limit:  in.GetLimit(),
		NodeId: nodeID,
	}

	reply, err := s.ns.Core().GetConst().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Const = reply.GetConst()

	return &output, nil
}

func (s *ConstService) Sync(ctx context.Context, in *pb.Const) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	nodeID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	in.NodeId = nodeID

	return s.ns.Core().GetConst().Sync(ctx, in)
}
