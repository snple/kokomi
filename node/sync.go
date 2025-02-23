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

func (s *SyncService) SetDeviceUpdated(ctx context.Context, in *nodes.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Device.Updated")
		}
	}

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ns.Core().GetSync().SetDeviceUpdated(ctx,
		&cores.SyncUpdated{Id: deviceID, Updated: in.GetUpdated()})
}

func (s *SyncService) GetDeviceUpdated(ctx context.Context, in *pb.MyEmpty) (*nodes.SyncUpdated, error) {
	var output nodes.SyncUpdated
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

	reply, err := s.ns.Core().GetSync().GetDeviceUpdated(ctx, &pb.Id{Id: deviceID})
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) WaitDeviceUpdated(in *pb.MyEmpty, stream nodes.SyncService_WaitDeviceUpdatedServer) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	deviceID, err := validateToken(stream.Context())
	if err != nil {
		return err
	}

	return s.ns.Core().GetSync().WaitDeviceUpdated(&pb.Id{Id: deviceID}, stream)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ns.Core().GetSync().SetTagValueUpdated(ctx,
		&cores.SyncUpdated{Id: deviceID, Updated: in.GetUpdated()})
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetSync().GetTagValueUpdated(ctx, &pb.Id{Id: deviceID})
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

	deviceID, err := validateToken(stream.Context())
	if err != nil {
		return err
	}

	return s.ns.Core().GetSync().WaitTagValueUpdated(&pb.Id{Id: deviceID}, stream)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ns.Core().GetSync().SetTagWriteUpdated(ctx,
		&cores.SyncUpdated{Id: deviceID, Updated: in.GetUpdated()})
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ns.Core().GetSync().GetTagWriteUpdated(ctx, &pb.Id{Id: deviceID})
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

	deviceID, err := validateToken(stream.Context())
	if err != nil {
		return err
	}

	return s.ns.Core().GetSync().WaitTagWriteUpdated(&pb.Id{Id: deviceID}, stream)
}
