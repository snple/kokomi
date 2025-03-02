package node

import (
	"context"

	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/cores"
	"github.com/snple/beacon/pb/nodes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SyncService struct {
	ns *NodeService

	nodes.UnimplementedSyncServiceServer
}

func newSyncService(ns *NodeService) *SyncService {
	return &SyncService{
		ns: ns,
	}
}

func (s *SyncService) SetNodeUpdated(ctx context.Context, in *nodes.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Node.Updated")
		}
	}

	nodeID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ns.Core().GetSync().SetNodeUpdated(ctx,
		&cores.SyncUpdated{Id: nodeID, Updated: in.GetUpdated()})
}

func (s *SyncService) GetNodeUpdated(ctx context.Context, in *pb.MyEmpty) (*nodes.SyncUpdated, error) {
	var output nodes.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	nodeID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetSync().GetNodeUpdated(ctx, &pb.Id{Id: nodeID})
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) WaitNodeUpdated(in *pb.MyEmpty, stream nodes.SyncService_WaitNodeUpdatedServer) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	nodeID, err := validateToken(stream.Context())
	if err != nil {
		return err
	}

	return s.ns.Core().GetSync().WaitNodeUpdated(&pb.Id{Id: nodeID}, stream)
}

func (s *SyncService) SetTagValueUpdated(ctx context.Context, in *nodes.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Value.Updated")
		}
	}

	nodeID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ns.Core().GetSync().SetTagValueUpdated(ctx,
		&cores.SyncUpdated{Id: nodeID, Updated: in.GetUpdated()})
}

func (s *SyncService) GetTagValueUpdated(ctx context.Context, in *pb.MyEmpty) (*nodes.SyncUpdated, error) {
	var output nodes.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	nodeID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetSync().GetTagValueUpdated(ctx, &pb.Id{Id: nodeID})
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) WaitTagValueUpdated(in *pb.MyEmpty, stream nodes.SyncService_WaitTagValueUpdatedServer) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	nodeID, err := validateToken(stream.Context())
	if err != nil {
		return err
	}

	return s.ns.Core().GetSync().WaitTagValueUpdated(&pb.Id{Id: nodeID}, stream)
}

func (s *SyncService) SetTagWriteUpdated(ctx context.Context, in *nodes.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Value.Updated")
		}
	}

	nodeID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ns.Core().GetSync().SetTagWriteUpdated(ctx,
		&cores.SyncUpdated{Id: nodeID, Updated: in.GetUpdated()})
}

func (s *SyncService) GetTagWriteUpdated(ctx context.Context, in *pb.MyEmpty) (*nodes.SyncUpdated, error) {
	var output nodes.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	nodeID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetSync().GetTagWriteUpdated(ctx, &pb.Id{Id: nodeID})
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) WaitTagWriteUpdated(in *pb.MyEmpty, stream nodes.SyncService_WaitTagWriteUpdatedServer) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	nodeID, err := validateToken(stream.Context())
	if err != nil {
		return err
	}

	return s.ns.Core().GetSync().WaitTagWriteUpdated(&pb.Id{Id: nodeID}, stream)
}
