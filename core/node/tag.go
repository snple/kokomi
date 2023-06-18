package node

import (
	"context"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
	"github.com/snple/kokomi/pb/nodes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TagService struct {
	ns *NodeService

	nodes.UnimplementedTagServiceServer
}

func newTagService(ns *NodeService) *TagService {
	return &TagService{
		ns: ns,
	}
}

func (s *TagService) Create(ctx context.Context, in *pb.Tag) (*pb.Tag, error) {
	var output pb.Tag

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

		request := &pb.Id{Id: in.GetSourceId()}

		reply, err := s.ns.Core().GetSource().View(ctx, request)
		if err != nil {
			return &output, err
		}

		if reply.GetDeviceId() != deviceID {
			return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
		}
	}

	return s.ns.Core().GetTag().Create(ctx, in)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &pb.Id{Id: in.GetId()}

	reply, err := s.ns.Core().GetTag().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetTag().Update(ctx, in)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetTag().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *TagService) ViewByName(ctx context.Context, in *pb.Name) (*pb.Tag, error) {
	var output pb.Tag
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

	request := &cores.ViewTagByNameRequest{DeviceId: deviceID, Name: in.GetName()}

	reply, err := s.ns.Core().GetTag().ViewByName(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetTag().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetTag().Delete(ctx, in)
}

func (s *TagService) List(ctx context.Context, in *nodes.ListTagRequest) (*nodes.ListTagResponse, error) {
	var err error
	var output nodes.ListTagResponse

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

	request := &cores.ListTagRequest{
		Page:     in.GetPage(),
		DeviceId: deviceID,
		SourceId: in.GetSourceId(),
		Tags:     in.GetTags(),
		Type:     in.GetType(),
	}

	reply, err := s.ns.Core().GetTag().List(ctx, request)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetTag().View(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetTag().GetValue(ctx, in)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &pb.Id{Id: in.GetId()}

	reply, err := s.ns.Core().GetTag().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetTag().SetValue(ctx, in)
}

func (s *TagService) SetValueUnchecked(ctx context.Context, in *pb.TagValue) (*pb.MyBool, error) {
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

	reply, err := s.ns.Core().GetTag().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetTag().SetValueUnchecked(ctx, in)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetTag().GetValueByName(ctx,
		&cores.GetTagValueByNameRequest{DeviceId: deviceID, Name: in.GetName()})
	if err != nil {
		return &output, err
	}

	output.Id = reply.GetId()
	output.Name = reply.GetName()
	output.Value = reply.GetValue()
	output.Updated = reply.GetUpdated()

	return &output, nil
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ns.Core().GetTag().SetValueByName(ctx,
		&cores.TagNameValue{DeviceId: deviceID, Name: in.GetName(), Value: in.GetValue()})
}

func (s *TagService) SetValueByNameUnchecked(ctx context.Context, in *pb.TagNameValue) (*pb.MyBool, error) {
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

	return s.ns.Core().GetTag().SetValueByNameUnchecked(ctx,
		&cores.TagNameValue{DeviceId: deviceID, Name: in.GetName(), Value: in.GetValue()})
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetTag().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return reply, nil
}

func (s *TagService) Pull(ctx context.Context, in *nodes.PullTagRequest) (*nodes.PullTagResponse, error) {
	var err error
	var output nodes.PullTagResponse

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

	request := &cores.PullTagRequest{
		After:    in.GetAfter(),
		Limit:    in.GetLimit(),
		DeviceId: deviceID,
		SourceId: in.GetSourceId(),
		Type:     in.GetType(),
	}

	reply, err := s.ns.Core().GetTag().Pull(ctx, request)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	in.DeviceId = deviceID

	return s.ns.Core().GetTag().Sync(ctx, in)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetTag().ViewValue(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetTag().ViewValue(ctx, in)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetTag().DeleteValue(ctx, in)
}

func (s *TagService) PullValue(ctx context.Context, in *nodes.PullTagValueRequest) (*nodes.PullTagValueResponse, error) {
	var err error
	var output nodes.PullTagValueResponse

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

	request := &cores.PullTagValueRequest{
		After:    in.GetAfter(),
		Limit:    in.GetLimit(),
		DeviceId: deviceID,
		SourceId: in.GetSourceId(),
	}

	reply, err := s.ns.Core().GetTag().PullValue(ctx, request)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &pb.Id{Id: in.GetId()}

	reply, err := s.ns.Core().GetTag().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetDeviceId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: reply.GetDeviceId() != deviceID")
	}

	return s.ns.Core().GetTag().SyncValue(ctx, in)
}
