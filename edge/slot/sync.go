package slot

import (
	"context"

	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/edges"
	"github.com/snple/beacon/pb/slots"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SyncService struct {
	ss *SlotService

	slots.UnimplementedSyncServiceServer
}

func newSyncService(ss *SlotService) *SyncService {
	return &SyncService{
		ss: ss,
	}
}

func (s *SyncService) SetNodeUpdated(ctx context.Context, in *slots.SyncUpdated) (*pb.MyBool, error) {
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetSync().SetNodeUpdated(ctx,
		&edges.SyncUpdated{Updated: in.GetUpdated()})
}

func (s *SyncService) GetNodeUpdated(ctx context.Context, in *pb.MyEmpty) (*slots.SyncUpdated, error) {
	var output slots.SyncUpdated
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

	reply, err := s.ss.Edge().GetSync().GetNodeUpdated(ctx, in)
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) WaitNodeUpdated(in *pb.MyEmpty, stream slots.SyncService_WaitNodeUpdatedServer) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(stream.Context())
	if err != nil {
		return err
	}

	return s.ss.Edge().GetSync().WaitNodeUpdated(in, stream)
}

func (s *SyncService) SetSlotUpdated(ctx context.Context, in *slots.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Slot.Updated")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetSync().SetSlotUpdated(ctx,
		&edges.SyncUpdated{Updated: in.GetUpdated()})
}

func (s *SyncService) GetSlotUpdated(ctx context.Context, in *pb.MyEmpty) (*slots.SyncUpdated, error) {
	var output slots.SyncUpdated
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

	reply, err := s.ss.Edge().GetSync().GetSlotUpdated(ctx, in)
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) SetWireUpdated(ctx context.Context, in *slots.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Source.Updated")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetSync().SetWireUpdated(ctx,
		&edges.SyncUpdated{Updated: in.GetUpdated()})
}

func (s *SyncService) GetWireUpdated(ctx context.Context, in *pb.MyEmpty) (*slots.SyncUpdated, error) {
	var output slots.SyncUpdated
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

	reply, err := s.ss.Edge().GetSync().GetWireUpdated(ctx, in)
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) SetPinUpdated(ctx context.Context, in *slots.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Pin.Updated")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetSync().SetPinUpdated(ctx,
		&edges.SyncUpdated{Updated: in.GetUpdated()})
}

func (s *SyncService) GetPinUpdated(ctx context.Context, in *pb.MyEmpty) (*slots.SyncUpdated, error) {
	var output slots.SyncUpdated
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

	reply, err := s.ss.Edge().GetSync().GetPinUpdated(ctx, in)
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) SetConstUpdated(ctx context.Context, in *slots.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Const.Updated")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetSync().SetConstUpdated(ctx,
		&edges.SyncUpdated{Updated: in.GetUpdated()})
}

func (s *SyncService) GetConstUpdated(ctx context.Context, in *pb.MyEmpty) (*slots.SyncUpdated, error) {
	var output slots.SyncUpdated
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

	reply, err := s.ss.Edge().GetSync().GetConstUpdated(ctx, in)
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) SetPinValueUpdated(ctx context.Context, in *slots.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Pin.Value.Updated")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetSync().SetPinValueUpdated(ctx,
		&edges.SyncUpdated{Updated: in.GetUpdated()})
}

func (s *SyncService) GetPinValueUpdated(ctx context.Context, in *pb.MyEmpty) (*slots.SyncUpdated, error) {
	var output slots.SyncUpdated
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

	reply, err := s.ss.Edge().GetSync().GetPinValueUpdated(ctx, in)
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) WaitPinValueUpdated(in *pb.MyEmpty, stream slots.SyncService_WaitPinValueUpdatedServer) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(stream.Context())
	if err != nil {
		return err
	}

	return s.ss.Edge().GetSync().WaitPinValueUpdated(in, stream)
}

func (s *SyncService) SetPinWriteUpdated(ctx context.Context, in *slots.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Pin.Write.Updated")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetSync().SetPinWriteUpdated(ctx,
		&edges.SyncUpdated{Updated: in.GetUpdated()})
}

func (s *SyncService) GetPinWriteUpdated(ctx context.Context, in *pb.MyEmpty) (*slots.SyncUpdated, error) {
	var output slots.SyncUpdated
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

	reply, err := s.ss.Edge().GetSync().GetPinWriteUpdated(ctx, in)
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) WaitPinWriteUpdated(in *pb.MyEmpty, stream slots.SyncService_WaitPinWriteUpdatedServer) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(stream.Context())
	if err != nil {
		return err
	}

	return s.ss.Edge().GetSync().WaitPinWriteUpdated(in, stream)
}
