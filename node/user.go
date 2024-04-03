package node

import (
	"context"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
	"github.com/snple/kokomi/pb/nodes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	ns *NodeService

	nodes.UnimplementedUserServiceServer
}

func newUserService(ns *NodeService) *UserService {
	return &UserService{
		ns: ns,
	}
}

func (s *UserService) View(ctx context.Context, in *pb.Id) (*pb.User, error) {
	var err error
	var output pb.User

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

	return s.ns.Core().GetUser().View(ctx, in)
}

func (s *UserService) Name(ctx context.Context, in *pb.Name) (*pb.User, error) {
	var err error
	var output pb.User

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

	return s.ns.Core().GetUser().Name(ctx, in)
}

func (s *UserService) List(ctx context.Context, in *nodes.UserListRequest) (*nodes.UserListResponse, error) {
	var err error
	var output nodes.UserListResponse

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

	request := &cores.UserListRequest{
		Page: in.GetPage(),
		Tags: in.GetTags(),
		Type: in.GetType(),
	}

	reply, err := s.ns.Core().GetUser().List(ctx, request)
	if err != nil {
		return &output, err
	}

	output.Count = reply.Count
	output.Page = reply.GetPage()
	output.User = reply.GetUser()

	return &output, nil
}

func (s *UserService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.User, error) {
	var output pb.User
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

	reply, err := s.ns.Core().GetUser().ViewWithDeleted(ctx, in)
	if err != nil {
		return &output, err
	}

	return reply, nil
}

func (s *UserService) Pull(ctx context.Context, in *nodes.UserPullRequest) (*nodes.UserPullResponse, error) {
	var err error
	var output nodes.UserPullResponse

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

	request := &cores.UserPullRequest{
		After: in.GetAfter(),
		Limit: in.GetLimit(),
		Type:  in.GetType(),
	}

	reply, err := s.ns.Core().GetUser().Pull(ctx, request)
	if err != nil {
		return &output, err
	}

	output.User = reply.GetUser()

	return &output, nil
}
