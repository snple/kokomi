package node

import (
	"context"

	"github.com/snple/beacon/pb/cores"
	"github.com/snple/beacon/pb/nodes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	ns *NodeService

	nodes.UnimplementedAuthServiceServer
}

func newAuthService(ns *NodeService) *AuthService {
	return &AuthService{
		ns: ns,
	}
}

func (s *AuthService) Login(ctx context.Context, in *nodes.LoginRequest) (*nodes.LoginResponse, error) {
	var err error
	var output nodes.LoginResponse

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

	request := &cores.LoginRequest{
		Name: in.GetName(),
		Pass: in.GetPass(),
	}

	reply, err := s.ns.Core().GetAuth().Login(ctx, request)

	output.User = reply.GetUser()

	return &output, nil
}
