package node

import (
	"context"

	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/nodes"
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
