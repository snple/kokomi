package node

import (
	"context"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
	"github.com/snple/kokomi/pb/nodes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WireService struct {
	ns *NodeService

	nodes.UnimplementedWireServiceServer
}

func newWireService(ns *NodeService) *WireService {
	return &WireService{
		ns: ns,
	}
}

func (s *WireService) Create(ctx context.Context, in *pb.Wire) (*pb.Wire, error) {
	var output pb.Wire

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	// source validation
	{
		deviceID, err := validateToken(ctx)
		if err != nil {
			return &output, err
		}

		request := &pb.Id{Id: in.GetCableId()}

		reply, err := s.ns.Core().GetCable().View(ctx, request)
		if err != nil {
			return &output, err
		}

		if reply.GetDeviceId() != deviceID {
			return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
		}
	}

	return s.ns.Core().GetWire().Create(ctx, in)
}

func (s *WireService) Update(ctx context.Context, in *pb.Wire) (*pb.Wire, error) {
	var output pb.Wire
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

	reply, err := s.ns.Core().GetWire().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetWire().Update(ctx, in)
}

func (s *WireService) View(ctx context.Context, in *pb.Id) (*pb.Wire, error) {
	var output pb.Wire
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

	reply, err := s.ns.Core().GetWire().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *WireService) Name(ctx context.Context, in *pb.Name) (*pb.Wire, error) {
	var output pb.Wire
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

	request := &cores.WireNameRequest{DeviceId: deviceID, Name: in.GetName()}

	reply, err := s.ns.Core().GetWire().Name(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *WireService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
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

	reply, err := s.ns.Core().GetWire().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetWire().Delete(ctx, in)
}

func (s *WireService) List(ctx context.Context, in *nodes.WireListRequest) (*nodes.WireListResponse, error) {
	var err error
	var output nodes.WireListResponse

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

	request := &cores.WireListRequest{
		Page:     in.GetPage(),
		DeviceId: deviceID,
		CableId:  in.GetCableId(),
		Tags:     in.GetTags(),
		Type:     in.GetType(),
	}

	reply, err := s.ns.Core().GetWire().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Wire = reply.GetWire()

	return &output, nil
}

func (s *WireService) GetValue(ctx context.Context, in *pb.Id) (*pb.WireValue, error) {
	var err error
	var output pb.WireValue

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

	reply, err := s.ns.Core().GetWire().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetWire().GetValue(ctx, in)
}

func (s *WireService) SetValue(ctx context.Context, in *pb.WireValue) (*pb.MyBool, error) {
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

	request := &pb.Id{Id: in.GetId()}

	reply, err := s.ns.Core().GetWire().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetWire().SetValue(ctx, in)
}

func (s *WireService) SetValueUnchecked(ctx context.Context, in *pb.WireValue) (*pb.MyBool, error) {
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

	request := &pb.Id{Id: in.GetId()}

	reply, err := s.ns.Core().GetWire().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetWire().SetValueUnchecked(ctx, in)
}

func (s *WireService) GetValueByName(ctx context.Context, in *pb.Name) (*pb.WireNameValue, error) {
	var err error
	var output pb.WireNameValue

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

	reply, err := s.ns.Core().GetWire().GetValueByName(ctx,
		&cores.WireGetValueByNameRequest{DeviceId: deviceID, Name: in.GetName()})
	if err != nil {
		return &output, err
	}

	output.Id = reply.GetId()
	output.Name = reply.GetName()
	output.Value = reply.GetValue()
	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *WireService) SetValueByName(ctx context.Context, in *pb.WireNameValue) (*pb.MyBool, error) {
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

	return s.ns.Core().GetWire().SetValueByName(ctx,
		&cores.WireNameValue{DeviceId: deviceID, Name: in.GetName(), Value: in.GetValue()})
}

func (s *WireService) SetValueByNameUnchecked(ctx context.Context, in *pb.WireNameValue) (*pb.MyBool, error) {
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

	return s.ns.Core().GetWire().SetValueByNameUnchecked(ctx,
		&cores.WireNameValue{DeviceId: deviceID, Name: in.GetName(), Value: in.GetValue()})
}

func (s *WireService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Wire, error) {
	var output pb.Wire
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

	reply, err := s.ns.Core().GetWire().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *WireService) Pull(ctx context.Context, in *nodes.WirePullRequest) (*nodes.WirePullResponse, error) {
	var err error
	var output nodes.WirePullResponse

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

	request := &cores.WirePullRequest{
		After:    in.GetAfter(),
		Limit:    in.GetLimit(),
		DeviceId: deviceID,
		CableId:  in.GetCableId(),
		Type:     in.GetType(),
	}

	reply, err := s.ns.Core().GetWire().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Wire = reply.GetWire()

	return &output, nil
}

func (s *WireService) Sync(ctx context.Context, in *pb.Wire) (*pb.MyBool, error) {
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

	return s.ns.Core().GetWire().Sync(ctx, in)
}

func (s *WireService) ViewValue(ctx context.Context, in *pb.Id) (*pb.WireValueUpdated, error) {
	var output pb.WireValueUpdated
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

	reply, err := s.ns.Core().GetWire().ViewValue(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *WireService) DeleteValue(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
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

	reply, err := s.ns.Core().GetWire().ViewValue(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetWire().DeleteValue(ctx, in)
}

func (s *WireService) PullValue(ctx context.Context, in *nodes.WirePullValueRequest) (*nodes.WirePullValueResponse, error) {
	var err error
	var output nodes.WirePullValueResponse

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

	request := &cores.WirePullValueRequest{
		After:    in.GetAfter(),
		Limit:    in.GetLimit(),
		DeviceId: deviceID,
		CableId:  in.GetCableId(),
	}

	reply, err := s.ns.Core().GetWire().PullValue(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Wire = reply.GetWire()

	return &output, nil
}

func (s *WireService) SyncValue(ctx context.Context, in *pb.WireValue) (*pb.MyBool, error) {
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

	request := &pb.Id{Id: in.GetId()}

	reply, err := s.ns.Core().GetWire().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetWire().SyncValue(ctx, in)
}
