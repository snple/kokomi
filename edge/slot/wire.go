package slot

import (
	"context"

	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/edges"
	"github.com/snple/beacon/pb/slots"
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

	return s.ss.Edge().GetWire().Create(ctx, in)
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

	return s.ss.Edge().GetWire().Update(ctx, in)
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

	return s.ss.Edge().GetWire().View(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetWire().Name(ctx, in)
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

	return s.ss.Edge().GetWire().Delete(ctx, in)
}

func (s *WireService) List(ctx context.Context, in *slots.WireListRequest) (*slots.WireListResponse, error) {
	var err error
	var output slots.WireListResponse

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

	request := &edges.WireListRequest{
		Page:   in.GetPage(),
		Tags:   in.GetTags(),
		Source: in.GetSource(),
	}

	reply, err := s.ss.Edge().GetWire().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Wire = reply.GetWire()

	return &output, nil
}

func (s *WireService) Link(ctx context.Context, in *slots.WireLinkRequest) (*pb.MyBool, error) {
	var output pb.MyBool
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

	request2 := &edges.WireLinkRequest{Id: in.GetId(), Status: in.GetStatus()}

	reply, err := s.ss.Edge().GetWire().Link(ctx, request2)
	if err != nil {
		return &output, err
	}

	return reply, nil
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

	reply, err := s.ss.Edge().GetWire().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	return reply, nil
}

func (s *WireService) Pull(ctx context.Context, in *slots.WirePullRequest) (*slots.WirePullResponse, error) {
	var err error
	var output slots.WirePullResponse

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

	request := &edges.WirePullRequest{
		After:  in.GetAfter(),
		Limit:  in.GetLimit(),
		Source: in.GetSource(),
	}

	reply, err := s.ss.Edge().GetWire().Pull(ctx, request)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetWire().Sync(ctx, in)
}
