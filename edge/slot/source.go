package slot

import (
	"context"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/pb/slots"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SourceService struct {
	ss *SlotService

	slots.UnimplementedSourceServiceServer
}

func newSourceService(ss *SlotService) *SourceService {
	return &SourceService{
		ss: ss,
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetSource().Create(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetSource().Update(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetSource().View(ctx, in)
}

func (s *SourceService) Name(ctx context.Context, in *pb.Name) (*pb.Source, error) {
	var output pb.Source
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

	return s.ss.Edge().GetSource().Name(ctx, in)
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetSource().Delete(ctx, in)
}

func (s *SourceService) List(ctx context.Context, in *slots.SourceListRequest) (*slots.SourceListResponse, error) {
	var err error
	var output slots.SourceListResponse

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

	request := &edges.SourceListRequest{
		Page:   in.GetPage(),
		Tags:   in.GetTags(),
		Type:   in.GetType(),
		Source: in.GetSource(),
	}

	reply, err := s.ss.Edge().GetSource().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Source = reply.GetSource()

	return &output, nil
}

func (s *SourceService) Link(ctx context.Context, in *slots.SourceLinkRequest) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	slotID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request2 := &edges.SourceLinkRequest{Id: in.GetId(), Status: in.GetStatus()}

	reply, err := s.ss.Edge().GetSource().Link(ctx, request2)
	if err != nil {
		return &output, err
	}

	s.ss.Edge().GetControl().SourceLink(slotID, in.GetId(), in.GetStatus())

	return reply, nil
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ss.Edge().GetSource().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	return reply, nil
}

func (s *SourceService) Pull(ctx context.Context, in *slots.SourcePullRequest) (*slots.SourcePullResponse, error) {
	var err error
	var output slots.SourcePullResponse

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

	request := &edges.SourcePullRequest{
		After:  in.GetAfter(),
		Limit:  in.GetLimit(),
		Type:   in.GetType(),
		Source: in.GetSource(),
	}

	reply, err := s.ss.Edge().GetSource().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Source = reply.GetSource()

	return &output, nil
}

func (s *SourceService) Sync(ctx context.Context, in *pb.Source) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetSource().Sync(ctx, in)
}
