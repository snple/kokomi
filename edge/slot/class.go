package slot

import (
	"context"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/pb/slots"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClassService struct {
	ss *SlotService

	slots.UnimplementedClassServiceServer
}

func newClassService(ss *SlotService) *ClassService {
	return &ClassService{
		ss: ss,
	}
}

func (s *ClassService) Create(ctx context.Context, in *pb.Class) (*pb.Class, error) {
	var output pb.Class
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

	return s.ss.Edge().GetClass().Create(ctx, in)
}

func (s *ClassService) Update(ctx context.Context, in *pb.Class) (*pb.Class, error) {
	var output pb.Class
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

	return s.ss.Edge().GetClass().Update(ctx, in)
}

func (s *ClassService) View(ctx context.Context, in *pb.Id) (*pb.Class, error) {
	var output pb.Class
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

	return s.ss.Edge().GetClass().View(ctx, in)
}

func (s *ClassService) ViewByName(ctx context.Context, in *pb.Name) (*pb.Class, error) {
	var output pb.Class
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

	return s.ss.Edge().GetClass().ViewByName(ctx, in)
}

func (s *ClassService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetClass().Delete(ctx, in)
}

func (s *ClassService) List(ctx context.Context, in *slots.ListClassRequest) (*slots.ListClassResponse, error) {
	var err error
	var output slots.ListClassResponse

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

	request := &edges.ListClassRequest{
		Page: in.GetPage(),
		Tags: in.GetTags(),
		Type: in.GetType(),
	}

	reply, err := s.ss.Edge().GetClass().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Class = reply.GetClass()

	return &output, nil
}

func (s *ClassService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Class, error) {
	var output pb.Class
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

	reply, err := s.ss.Edge().GetClass().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	return reply, nil
}

func (s *ClassService) Pull(ctx context.Context, in *slots.PullClassRequest) (*slots.PullClassResponse, error) {
	var err error
	var output slots.PullClassResponse

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

	request := &edges.PullClassRequest{
		After: in.GetAfter(),
		Limit: in.GetLimit(),
		Type:  in.GetType(),
	}

	reply, err := s.ss.Edge().GetClass().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Class = reply.GetClass()

	return &output, nil
}

func (s *ClassService) Sync(ctx context.Context, in *pb.Class) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetClass().Sync(ctx, in)
}
