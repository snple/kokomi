package node

import (
	"context"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
	"github.com/snple/kokomi/pb/nodes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SyncGlobalService struct {
	ns *NodeService

	nodes.UnimplementedSyncGlobalServiceServer
}

func newSyncGlobalService(ns *NodeService) *SyncGlobalService {
	return &SyncGlobalService{
		ns: ns,
	}
}

func (s *SyncGlobalService) SetUpdated(ctx context.Context, in *nodes.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid ID")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Updated")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ns.Core().GetSyncGlobal().SetUpdated(ctx,
		&cores.SyncUpdated{Id: in.GetId(), Updated: in.GetUpdated()})
}

func (s *SyncGlobalService) GetUpdated(ctx context.Context, in *pb.Id) (*nodes.SyncUpdated, error) {
	var output nodes.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid ID")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetSyncGlobal().GetUpdated(ctx, in)
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncGlobalService) WaitUpdated(in *pb.Id, stream nodes.SyncGlobalService_WaitUpdatedServer) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return status.Error(codes.InvalidArgument, "Please supply valid ID")
		}
	}

	_, err = validateToken(stream.Context())
	if err != nil {
		return err
	}

	return s.ns.Core().GetSyncGlobal().WaitUpdated(in, stream)
}
