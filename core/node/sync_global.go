package node

import (
	"github.com/snple/kokomi/pb/nodes"
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

// func (s *SyncGlobalService) SetUserUpdated(ctx context.Context, in *nodes.SyncUpdated) (*pb.MyBool, error) {
// 	var output pb.MyBool
// 	var err error

// 	// basic validation
// 	{
// 		if in == nil {
// 			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
// 		}

// 		if in.GetUpdated() == 0 {
// 			return &output, status.Error(codes.InvalidArgument, "Please supply valid user updated")
// 		}
// 	}

// 	_, err = validateToken(ctx)
// 	if err != nil {
// 		return &output, err
// 	}

// 	return s.ns.cs.GetSyncGlobal().SetUserUpdated(ctx,
// 		&clouds.SyncUpdated{Updated: in.GetUpdated()})
// }

// func (s *SyncGlobalService) GetUserUpdated(ctx context.Context, in *pb.MyEmpty) (*nodes.SyncUpdated, error) {
// 	var output nodes.SyncUpdated
// 	var err error

// 	// basic validation
// 	{
// 		if in == nil {
// 			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
// 		}
// 	}

// 	_, err = validateToken(ctx)
// 	if err != nil {
// 		return &output, err
// 	}

// 	reply, err := s.ns.cs.GetSyncGlobal().GetUserUpdated(ctx, in)
// 	if err != nil {
// 		return &output, err
// 	}

// 	output.Updated = reply.GetUpdated()

// 	return &output, nil
// }

// func (s *SyncGlobalService) WaitUserUpdated(in *pb.MyEmpty, stream nodes.SyncService_WaitDeviceUpdatedServer) error {
// 	var err error

// 	// basic validation
// 	{
// 		if in == nil {
// 			return status.Error(codes.InvalidArgument, "Please supply valid argument")
// 		}
// 	}

// 	_, err = validateToken(stream.Context())
// 	if err != nil {
// 		return err
// 	}

// 	return s.ns.cs.GetSyncGlobal().WaitUserUpdated(&pb.MyEmpty{}, stream)
// }
