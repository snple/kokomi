package slot

import (
	"context"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/pb/slots"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TagService struct {
	ss *SlotService

	slots.UnimplementedTagServiceServer
}

func newTagService(ss *SlotService) *TagService {
	return &TagService{
		ss: ss,
	}
}

func (s *TagService) Create(ctx context.Context, in *pb.Tag) (*pb.Tag, error) {
	var output pb.Tag
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

	return s.ss.Edge().GetTag().Create(ctx, in)
}

func (s *TagService) Update(ctx context.Context, in *pb.Tag) (*pb.Tag, error) {
	var output pb.Tag
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

	return s.ss.Edge().GetTag().Update(ctx, in)
}

func (s *TagService) View(ctx context.Context, in *pb.Id) (*pb.Tag, error) {
	var output pb.Tag
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

	return s.ss.Edge().GetTag().View(ctx, in)
}

func (s *TagService) Name(ctx context.Context, in *pb.Name) (*pb.Tag, error) {
	var output pb.Tag
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

	return s.ss.Edge().GetTag().Name(ctx, in)
}

func (s *TagService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetTag().Delete(ctx, in)
}

func (s *TagService) List(ctx context.Context, in *slots.TagListRequest) (*slots.TagListResponse, error) {
	var err error
	var output slots.TagListResponse

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

	request := &edges.TagListRequest{
		Page:     in.GetPage(),
		SourceId: in.GetSourceId(),
		Tags:     in.GetTags(),
		Type:     in.GetType(),
	}

	reply, err := s.ss.Edge().GetTag().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Tag = reply.GetTag()

	return &output, nil
}

func (s *TagService) GetValue(ctx context.Context, in *pb.Id) (*pb.TagValue, error) {
	var err error
	var output pb.TagValue

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

	return s.ss.Edge().GetTag().GetValue(ctx, in)
}

func (s *TagService) SetValue(ctx context.Context, in *pb.TagValue) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetTag().SetValue(ctx, in)
}

func (s *TagService) SetValueForce(ctx context.Context, in *pb.TagValue) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetTag().SetValueForce(ctx, in)
}

func (s *TagService) GetValueByName(ctx context.Context, in *pb.Name) (*pb.TagNameValue, error) {
	var err error
	var output pb.TagNameValue

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

	return s.ss.Edge().GetTag().GetValueByName(ctx, in)
}

func (s *TagService) SetValueByName(ctx context.Context, in *pb.TagNameValue) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetTag().SetValueByName(ctx, in)
}

func (s *TagService) SetValueByNameForce(ctx context.Context, in *pb.TagNameValue) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetTag().SetValueByNameForce(ctx, in)
}

func (s *TagService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Tag, error) {
	var output pb.Tag
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

	reply, err := s.ss.Edge().GetTag().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	return reply, nil
}

func (s *TagService) Pull(ctx context.Context, in *slots.TagPullRequest) (*slots.TagPullResponse, error) {
	var err error
	var output slots.TagPullResponse

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

	request := &edges.TagPullRequest{
		After:    in.GetAfter(),
		Limit:    in.GetLimit(),
		SourceId: in.GetSourceId(),
		Type:     in.GetType(),
	}

	reply, err := s.ss.Edge().GetTag().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Tag = reply.GetTag()

	return &output, nil
}

func (s *TagService) Sync(ctx context.Context, in *pb.Tag) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetTag().Sync(ctx, in)
}

func (s *TagService) ViewValue(ctx context.Context, in *pb.Id) (*pb.TagValueUpdated, error) {
	var output pb.TagValueUpdated
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

	reply, err := s.ss.Edge().GetTag().ViewValue(ctx, in)
	if err != nil {
		return &output, err
	}

	return reply, nil
}

func (s *TagService) DeleteValue(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetTag().DeleteValue(ctx, in)
}

func (s *TagService) PullValue(ctx context.Context, in *slots.TagPullValueRequest) (*slots.TagPullValueResponse, error) {
	var err error
	var output slots.TagPullValueResponse

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

	request := &edges.TagPullValueRequest{
		After:    in.GetAfter(),
		Limit:    in.GetLimit(),
		SourceId: in.GetSourceId(),
	}

	reply, err := s.ss.Edge().GetTag().PullValue(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Tag = reply.GetTag()

	return &output, nil
}

func (s *TagService) SyncValue(ctx context.Context, in *pb.TagValue) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetTag().SyncValue(ctx, in)
}
