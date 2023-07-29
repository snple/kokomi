package slot

import (
	"context"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/pb/slots"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CableService struct {
	ss *SlotService

	slots.UnimplementedCableServiceServer
}

func newCableService(ss *SlotService) *CableService {
	return &CableService{
		ss: ss,
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetCable().Create(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetCable().Update(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetCable().View(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetCable().Name(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetCable().Delete(ctx, in)
}

func (s *CableService) List(ctx context.Context, in *slots.CableListRequest) (*slots.CableListResponse, error) {
	var err error
	var output slots.CableListResponse

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

	request := &edges.CableListRequest{
		Page: in.GetPage(),
		Tags: in.GetTags(),
		Type: in.GetType(),
	}

	reply, err := s.ss.Edge().GetCable().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Cable = reply.GetCable()

	return &output, nil
}

func (s *CableService) Link(ctx context.Context, in *slots.CableLinkRequest) (*pb.MyBool, error) {
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

	request2 := &edges.CableLinkRequest{Id: in.GetId(), Status: in.GetStatus()}

	reply, err := s.ss.Edge().GetCable().Link(ctx, request2)
	if err != nil {
		return &output, err
	}

	return reply, nil
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ss.Edge().GetCable().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	return reply, nil
}

func (s *CableService) Pull(ctx context.Context, in *slots.CablePullRequest) (*slots.CablePullResponse, error) {
	var err error
	var output slots.CablePullResponse

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

	request := &edges.CablePullRequest{
		After: in.GetAfter(),
		Limit: in.GetLimit(),
		Type:  in.GetType(),
	}

	reply, err := s.ss.Edge().GetCable().Pull(ctx, request)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetCable().Sync(ctx, in)
}
