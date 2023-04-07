package slot

import (
	"context"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/pb/slots"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WireService struct {
	ss *SlotService

	slots.UnimplementedWireServiceServer
}

func newWireService(ss *SlotService) *WireService {
	return &WireService{
		ss: ss,
	}
}

func (s *WireService) Create(ctx context.Context, in *pb.Wire) (*pb.Wire, error) {
	var output pb.Wire
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetWire().Create(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetWire().Update(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetWire().View(ctx, in)
}

func (s *WireService) ViewByName(ctx context.Context, in *pb.Name) (*pb.Wire, error) {
	var output pb.Wire
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetWire().ViewByName(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetWire().Delete(ctx, in)
}

func (s *WireService) List(ctx context.Context, in *slots.ListWireRequest) (*slots.ListWireResponse, error) {
	var err error
	var output slots.ListWireResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &edges.ListWireRequest{
		Page:    in.GetPage(),
		CableId: in.GetCableId(),
		Tags:    in.GetTags(),
		Type:    in.GetType(),
	}

	reply, err := s.ss.es.GetWire().List(ctx, request)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetWire().GetValue(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetWire().SetValue(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetWire().SetValueUnchecked(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetWire().GetValueByName(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetWire().SetValueByName(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetWire().SetValueByNameUnchecked(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ss.es.GetWire().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	return reply, nil
}

func (s *WireService) Pull(ctx context.Context, in *slots.PullWireRequest) (*slots.PullWireResponse, error) {
	var err error
	var output slots.PullWireResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &edges.PullWireRequest{
		After:   in.GetAfter(),
		Limit:   in.GetLimit(),
		CableId: in.GetCableId(),
	}

	reply, err := s.ss.es.GetWire().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Wire = reply.GetWire()

	return &output, nil
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ss.es.GetWire().ViewValue(ctx, in)
	if err != nil {
		return &output, err
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetWire().DeleteValue(ctx, in)
}

func (s *WireService) PullValue(ctx context.Context, in *slots.PullWireValueRequest) (*slots.PullWireValueResponse, error) {
	var err error
	var output slots.PullWireValueResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &edges.PullWireValueRequest{
		After:   in.GetAfter(),
		Limit:   in.GetLimit(),
		CableId: in.GetCableId(),
	}

	reply, err := s.ss.es.GetWire().PullValue(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Wire = reply.GetWire()

	return &output, nil
}
