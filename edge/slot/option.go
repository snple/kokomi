package slot

import (
	"context"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/pb/slots"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OptionService struct {
	ss *SlotService

	slots.UnimplementedOptionServiceServer
}

func newOptionService(ss *SlotService) *OptionService {
	return &OptionService{
		ss: ss,
	}
}

func (s *OptionService) Create(ctx context.Context, in *pb.Option) (*pb.Option, error) {
	var output pb.Option
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

	return s.ss.Edge().GetOption().Create(ctx, in)
}

func (s *OptionService) Update(ctx context.Context, in *pb.Option) (*pb.Option, error) {
	var output pb.Option
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

	return s.ss.Edge().GetOption().Update(ctx, in)
}

func (s *OptionService) View(ctx context.Context, in *pb.Id) (*pb.Option, error) {
	var output pb.Option
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

	return s.ss.Edge().GetOption().View(ctx, in)
}

func (s *OptionService) ViewByName(ctx context.Context, in *pb.Name) (*pb.Option, error) {
	var output pb.Option
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

	return s.ss.Edge().GetOption().ViewByName(ctx, in)
}

func (s *OptionService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetOption().Delete(ctx, in)
}

func (s *OptionService) List(ctx context.Context, in *slots.ListOptionRequest) (*slots.ListOptionResponse, error) {
	var err error
	var output slots.ListOptionResponse

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

	request := &edges.ListOptionRequest{
		Page: in.GetPage(),
		Tags: in.GetTags(),
		Type: in.GetType(),
	}

	reply, err := s.ss.Edge().GetOption().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Option = reply.GetOption()

	return &output, nil
}

func (s *OptionService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Option, error) {
	var output pb.Option
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

	reply, err := s.ss.Edge().GetOption().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	return reply, nil
}

func (s *OptionService) Pull(ctx context.Context, in *slots.PullOptionRequest) (*slots.PullOptionResponse, error) {
	var err error
	var output slots.PullOptionResponse

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

	request := &edges.PullOptionRequest{
		After: in.GetAfter(),
		Limit: in.GetLimit(),
		Type:  in.GetType(),
	}

	reply, err := s.ss.Edge().GetOption().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Option = reply.GetOption()

	return &output, nil
}

func (s *OptionService) Sync(ctx context.Context, in *pb.Option) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetOption().Sync(ctx, in)
}
