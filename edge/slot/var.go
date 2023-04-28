package slot

import (
	"context"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/pb/slots"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VarService struct {
	ss *SlotService

	slots.UnimplementedVarServiceServer
}

func newVarService(ss *SlotService) *VarService {
	return &VarService{
		ss: ss,
	}
}

func (s *VarService) Create(ctx context.Context, in *pb.Var) (*pb.Var, error) {
	var output pb.Var
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

	return s.ss.Edge().GetVar().Create(ctx, in)
}

func (s *VarService) Update(ctx context.Context, in *pb.Var) (*pb.Var, error) {
	var output pb.Var
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

	return s.ss.Edge().GetVar().Update(ctx, in)
}

func (s *VarService) View(ctx context.Context, in *pb.Id) (*pb.Var, error) {
	var output pb.Var
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

	return s.ss.Edge().GetVar().View(ctx, in)
}

func (s *VarService) ViewByName(ctx context.Context, in *pb.Name) (*pb.Var, error) {
	var output pb.Var
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

	return s.ss.Edge().GetVar().ViewByName(ctx, in)
}

func (s *VarService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetVar().Delete(ctx, in)
}

func (s *VarService) List(ctx context.Context, in *slots.ListVarRequest) (*slots.ListVarResponse, error) {
	var err error
	var output slots.ListVarResponse

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

	request := &edges.ListVarRequest{
		Page: in.GetPage(),
		Tags: in.GetTags(),
		Type: in.GetType(),
	}

	reply, err := s.ss.Edge().GetVar().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.Var = reply.GetVar()

	return &output, nil
}

func (s *VarService) GetValue(ctx context.Context, in *pb.Id) (*pb.VarValue, error) {
	var err error
	var output pb.VarValue

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

	return s.ss.Edge().GetVar().GetValue(ctx, in)
}

func (s *VarService) SetValue(ctx context.Context, in *pb.VarValue) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetVar().SetValue(ctx, in)
}

func (s *VarService) SetValueUnchecked(ctx context.Context, in *pb.VarValue) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetVar().SetValueUnchecked(ctx, in)
}

func (s *VarService) GetValueByName(ctx context.Context, in *pb.Name) (*pb.VarNameValue, error) {
	var err error
	var output pb.VarNameValue

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

	return s.ss.Edge().GetVar().GetValueByName(ctx, in)
}

func (s *VarService) SetValueByName(ctx context.Context, in *pb.VarNameValue) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetVar().SetValueByName(ctx, in)
}

func (s *VarService) SetValueByNameUnchecked(ctx context.Context, in *pb.VarNameValue) (*pb.MyBool, error) {
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

	return s.ss.Edge().GetVar().SetValueByNameUnchecked(ctx, in)
}

func (s *VarService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Var, error) {
	var output pb.Var
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

	reply, err := s.ss.Edge().GetVar().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	return reply, nil
}

func (s *VarService) Pull(ctx context.Context, in *slots.PullVarRequest) (*slots.PullVarResponse, error) {
	var err error
	var output slots.PullVarResponse

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

	request := &edges.PullVarRequest{
		After: in.GetAfter(),
		Limit: in.GetLimit(),
		Type:  in.GetType(),
	}

	reply, err := s.ss.Edge().GetVar().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Var = reply.GetVar()

	return &output, nil
}
