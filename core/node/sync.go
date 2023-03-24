package node

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"snple.com/kokomi/pb"
	"snple.com/kokomi/pb/cores"
	"snple.com/kokomi/pb/nodes"
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid device updated")
		}
	}

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ns.cs.GetSync().SetDeviceUpdated(ctx,
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

	reply, err := s.ns.cs.GetSync().GetDeviceUpdated(ctx, &pb.Id{Id: deviceID})
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

	return s.ns.cs.GetSync().WaitDeviceUpdated(&pb.Id{Id: deviceID}, stream)
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag value updated")
		}
	}

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ns.cs.GetSync().SetTagValueUpdated(ctx,
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

	reply, err := s.ns.cs.GetSync().GetTagValueUpdated(ctx, &pb.Id{Id: deviceID})
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

	return s.ns.cs.GetSync().WaitTagValueUpdated(&pb.Id{Id: deviceID}, stream)
}

func (s *SyncService) SetWireValueUpdated(ctx context.Context, in *nodes.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid wire value updated")
		}
	}

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ns.cs.GetSync().SetWireValueUpdated(ctx,
		&cores.SyncUpdated{Id: deviceID, Updated: in.GetUpdated()})
}

func (s *SyncService) GetWireValueUpdated(ctx context.Context, in *pb.MyEmpty) (*nodes.SyncUpdated, error) {
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

	reply, err := s.ns.cs.GetSync().GetWireValueUpdated(ctx, &pb.Id{Id: deviceID})
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) WaitWireValueUpdated(in *pb.MyEmpty, stream nodes.SyncService_WaitWireValueUpdatedServer) error {
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

	return s.ns.cs.GetSync().WaitWireValueUpdated(&pb.Id{Id: deviceID}, stream)
}
