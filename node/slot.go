package node

import (
	"context"

	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/cores"
	"github.com/snple/beacon/pb/nodes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SlotService struct {
	ns *NodeService

	nodes.UnimplementedSlotServiceServer
}

func newSlotService(ns *NodeService) *SlotService {
	return &SlotService{
		ns: ns,
	}
}

func (s *SlotService) Create(ctx context.Context, in *pb.Slot) (*pb.Slot, error) {
	var output pb.Slot
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

	return s.ns.Core().GetSlot().Create(ctx, in)
}

func (s *SlotService) Update(ctx context.Context, in *pb.Slot) (*pb.Slot, error) {
	var output pb.Slot
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

	reply, err := s.ns.Core().GetSlot().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetNodeId() != nodeID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetNodeId() != nodeID")
	}

	return s.ns.Core().GetSlot().Update(ctx, in)
}

func (s *SlotService) View(ctx context.Context, in *pb.Id) (*pb.Slot, error) {
	var output pb.Slot
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

	reply, err := s.ns.Core().GetSlot().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetNodeId() != nodeID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetNodeId() != nodeID")
	}

	return reply, nil
}

func (s *SlotService) Name(ctx context.Context, in *pb.Name) (*pb.Slot, error) {
	var output pb.Slot
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

	request := &cores.SlotNameRequest{NodeId: nodeID, Name: in.GetName()}

	reply, err := s.ns.Core().GetSlot().Name(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetNodeId() != nodeID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetNodeId() != nodeID")
	}

	return reply, nil
}

func (s *SlotService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
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

	reply, err := s.ns.Core().GetSlot().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetNodeId() != nodeID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetNodeId() != nodeID")
	}

	return s.ns.Core().GetSlot().Delete(ctx, in)
}

func (s *SlotService) List(ctx context.Context, in *nodes.SlotListRequest) (*nodes.SlotListResponse, error) {
	var err error
	var output nodes.SlotListResponse

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

	request := &cores.SlotListRequest{
		Page:   in.GetPage(),
		NodeId: nodeID,
		Tags:   in.GetTags(),
	}

	reply, err := s.ns.Core().GetSlot().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Slot = reply.GetSlot()

	return &output, nil
}

func (s *SlotService) Link(ctx context.Context, in *nodes.SlotLinkRequest) (*pb.MyBool, error) {
	var output pb.MyBool
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

	reply, err := s.ns.Core().GetSlot().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetNodeId() != nodeID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetNodeId() != nodeID")
	}

	request2 := &cores.SlotLinkRequest{Id: in.GetId(), Status: in.GetStatus()}

	reply2, err := s.ns.Core().GetSlot().Link(ctx, request2)
	if err != nil {
		return &output, err
	}

	return reply2, nil
}

func (s *SlotService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Slot, error) {
	var output pb.Slot
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

	reply, err := s.ns.Core().GetSlot().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetNodeId() != nodeID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetNodeId() != nodeID")
	}

	return reply, nil
}

func (s *SlotService) Pull(ctx context.Context, in *nodes.SlotPullRequest) (*nodes.SlotPullResponse, error) {
	var err error
	var output nodes.SlotPullResponse

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

	request := &cores.SlotPullRequest{
		After:  in.GetAfter(),
		Limit:  in.GetLimit(),
		NodeId: nodeID,
	}

	reply, err := s.ns.Core().GetSlot().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Slot = reply.GetSlot()

	return &output, nil
}

func (s *SlotService) Sync(ctx context.Context, in *pb.Slot) (*pb.MyBool, error) {
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

	return s.ns.Core().GetSlot().Sync(ctx, in)
}
