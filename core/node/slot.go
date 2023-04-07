package node

import (
	"context"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
	"github.com/snple/kokomi/pb/nodes"
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	in.DeviceId = deviceID

	return s.ns.cs.GetSlot().Create(ctx, in)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &pb.Id{Id: in.GetId()}

	reply, err := s.ns.cs.GetSlot().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.cs.GetSlot().Update(ctx, in)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.cs.GetSlot().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *SlotService) ViewByName(ctx context.Context, in *pb.Name) (*pb.Slot, error) {
	var output pb.Slot
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

	request := &cores.ViewSlotByNameRequest{DeviceId: deviceID, Name: in.GetName()}

	reply, err := s.ns.cs.GetSlot().ViewByName(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.cs.GetSlot().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.cs.GetSlot().Delete(ctx, in)
}

func (s *SlotService) List(ctx context.Context, in *nodes.ListSlotRequest) (*nodes.ListSlotResponse, error) {
	var err error
	var output nodes.ListSlotResponse

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

	request := &cores.ListSlotRequest{
		Page:     in.GetPage(),
		DeviceId: deviceID,
		Tags:     in.GetTags(),
		Type:     in.GetType(),
	}

	reply, err := s.ns.cs.GetSlot().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Slot = reply.GetSlot()

	return &output, nil
}

func (s *SlotService) Link(ctx context.Context, in *nodes.LinkSlotRequest) (*pb.MyBool, error) {
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

	reply, err := s.ns.cs.GetSlot().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	request2 := &cores.LinkSlotRequest{Id: in.GetId(), Status: in.GetStatus()}

	reply2, err := s.ns.cs.GetSlot().Link(ctx, request2)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.cs.GetSlot().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *SlotService) Pull(ctx context.Context, in *nodes.PullSlotRequest) (*nodes.PullSlotResponse, error) {
	var err error
	var output nodes.PullSlotResponse

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

	request := &cores.PullSlotRequest{
		After:    in.GetAfter(),
		Limit:    in.GetLimit(),
		DeviceId: deviceID,
	}

	reply, err := s.ns.cs.GetSlot().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Slot = reply.GetSlot()

	return &output, nil
}
