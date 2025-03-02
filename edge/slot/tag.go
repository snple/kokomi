package slot

import (
	"context"

	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/edges"
	"github.com/snple/beacon/pb/slots"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PinService struct {
	ss *SlotService

	slots.UnimplementedPinServiceServer
}

func newPinService(ss *SlotService) *PinService {
	return &PinService{
		ss: ss,
	}
}

func (s *PinService) Create(ctx context.Context, in *pb.Pin) (*pb.Pin, error) {
	var output pb.Pin
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

	return s.ss.Edge().GetPin().Create(ctx, in)
}

func (s *PinService) Update(ctx context.Context, in *pb.Pin) (*pb.Pin, error) {
	var output pb.Pin
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

	return s.ss.Edge().GetPin().Update(ctx, in)
}

func (s *PinService) View(ctx context.Context, in *pb.Id) (*pb.Pin, error) {
	var output pb.Pin
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

	return s.ss.Edge().GetPin().View(ctx, in)
}

func (s *PinService) Name(ctx context.Context, in *pb.Name) (*pb.Pin, error) {
	var output pb.Pin
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

	return s.ss.Edge().GetPin().Name(ctx, in)
}

func (s *PinService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetPin().Delete(ctx, in)
}

func (s *PinService) List(ctx context.Context, in *slots.PinListRequest) (*slots.PinListResponse, error) {
	var err error
	var output slots.PinListResponse

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

	request := &edges.PinListRequest{
		Page:     in.GetPage(),
		SourceId: in.GetSourceId(),
		Tags:     in.GetTags(),
	}

	reply, err := s.ss.Edge().GetPin().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Pin = reply.GetPin()

	return &output, nil
}

func (s *PinService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Pin, error) {
	var output pb.Pin
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

	reply, err := s.ss.Edge().GetPin().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	return reply, nil
}

func (s *PinService) Pull(ctx context.Context, in *slots.PinPullRequest) (*slots.PinPullResponse, error) {
	var err error
	var output slots.PinPullResponse

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

	request := &edges.PinPullRequest{
		After:    in.GetAfter(),
		Limit:    in.GetLimit(),
		SourceId: in.GetSourceId(),
	}

	reply, err := s.ss.Edge().GetPin().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Pin = reply.GetPin()

	return &output, nil
}

func (s *PinService) Sync(ctx context.Context, in *pb.Pin) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetPin().Sync(ctx, in)
}

// value

func (s *PinService) GetValue(ctx context.Context, in *pb.Id) (*pb.PinValue, error) {
	var err error
	var output pb.PinValue

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

	return s.ss.Edge().GetPin().GetValue(ctx, in)
}

func (s *PinService) SetValue(ctx context.Context, in *pb.PinValue) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetPin().SetValue(ctx, in)
}

func (s *PinService) GetValueByName(ctx context.Context, in *pb.Name) (*pb.PinNameValue, error) {
	var err error
	var output pb.PinNameValue

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

	return s.ss.Edge().GetPin().GetValueByName(ctx, in)
}

func (s *PinService) SetValueByName(ctx context.Context, in *pb.PinNameValue) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetPin().SetValueByName(ctx, in)
}

func (s *PinService) ViewValue(ctx context.Context, in *pb.Id) (*pb.PinValueUpdated, error) {
	var output pb.PinValueUpdated
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

	reply, err := s.ss.Edge().GetPin().ViewValue(ctx, in)
	if err != nil {
		return &output, err
	}

	return reply, nil
}

func (s *PinService) DeleteValue(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetPin().DeleteValue(ctx, in)
}

func (s *PinService) PullValue(ctx context.Context, in *slots.PinPullValueRequest) (*slots.PinPullValueResponse, error) {
	var err error
	var output slots.PinPullValueResponse

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

	request := &edges.PinPullValueRequest{
		After:    in.GetAfter(),
		Limit:    in.GetLimit(),
		SourceId: in.GetSourceId(),
	}

	reply, err := s.ss.Edge().GetPin().PullValue(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Pin = reply.GetPin()

	return &output, nil
}

func (s *PinService) SyncValue(ctx context.Context, in *pb.PinValue) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetPin().SyncValue(ctx, in)
}

// write

func (s *PinService) GetWrite(ctx context.Context, in *pb.Id) (*pb.PinValue, error) {
	var err error
	var output pb.PinValue

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

	return s.ss.Edge().GetPin().GetWrite(ctx, in)
}

func (s *PinService) SetWrite(ctx context.Context, in *pb.PinValue) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetPin().SetWrite(ctx, in)
}

func (s *PinService) GetWriteByName(ctx context.Context, in *pb.Name) (*pb.PinNameValue, error) {
	var err error
	var output pb.PinNameValue

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

	return s.ss.Edge().GetPin().GetWriteByName(ctx, in)
}

func (s *PinService) SetWriteByName(ctx context.Context, in *pb.PinNameValue) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetPin().SetWriteByName(ctx, in)
}

func (s *PinService) ViewWrite(ctx context.Context, in *pb.Id) (*pb.PinValueUpdated, error) {
	var output pb.PinValueUpdated
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

	reply, err := s.ss.Edge().GetPin().ViewWrite(ctx, in)
	if err != nil {
		return &output, err
	}

	return reply, nil
}

func (s *PinService) DeleteWrite(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetPin().DeleteWrite(ctx, in)
}

func (s *PinService) PullWrite(ctx context.Context, in *slots.PinPullValueRequest) (*slots.PinPullValueResponse, error) {
	var err error
	var output slots.PinPullValueResponse

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

	request := &edges.PinPullValueRequest{
		After:    in.GetAfter(),
		Limit:    in.GetLimit(),
		SourceId: in.GetSourceId(),
	}

	reply, err := s.ss.Edge().GetPin().PullWrite(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Pin = reply.GetPin()

	return &output, nil
}

func (s *PinService) SyncWrite(ctx context.Context, in *pb.PinValue) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetPin().SyncWrite(ctx, in)
}
