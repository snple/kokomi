package slot

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"snple.com/kokomi/pb"
	"snple.com/kokomi/pb/edges"
	"snple.com/kokomi/pb/slots"
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

func (s *SyncService) SetDeviceUpdated(ctx context.Context, in *slots.SyncUpdated) (*pb.MyBool, error) {
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetSync().SetDeviceUpdated(ctx,
		&edges.SyncUpdated{Updated: in.GetUpdated()})
}

func (s *SyncService) GetDeviceUpdated(ctx context.Context, in *pb.MyEmpty) (*slots.SyncUpdated, error) {
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

	reply, err := s.ss.es.GetSync().GetDeviceUpdated(ctx, in)
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) WaitDeviceUpdated(in *pb.MyEmpty, stream slots.SyncService_WaitDeviceUpdatedServer) error {
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

	return s.ss.es.GetSync().WaitDeviceUpdated(in, stream)
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid slot updated")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetSync().SetSlotUpdated(ctx,
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

	reply, err := s.ss.es.GetSync().GetSlotUpdated(ctx, in)
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) SetSourceUpdated(ctx context.Context, in *slots.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid source updated")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetSync().SetSourceUpdated(ctx,
		&edges.SyncUpdated{Updated: in.GetUpdated()})
}

func (s *SyncService) GetSourceUpdated(ctx context.Context, in *pb.MyEmpty) (*slots.SyncUpdated, error) {
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

	reply, err := s.ss.es.GetSync().GetSourceUpdated(ctx, in)
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) SetTagUpdated(ctx context.Context, in *slots.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag updated")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetSync().SetTagUpdated(ctx,
		&edges.SyncUpdated{Updated: in.GetUpdated()})
}

func (s *SyncService) GetTagUpdated(ctx context.Context, in *pb.MyEmpty) (*slots.SyncUpdated, error) {
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

	reply, err := s.ss.es.GetSync().GetTagUpdated(ctx, in)
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) SetVarUpdated(ctx context.Context, in *slots.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid var updated")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetSync().SetVarUpdated(ctx,
		&edges.SyncUpdated{Updated: in.GetUpdated()})
}

func (s *SyncService) GetVarUpdated(ctx context.Context, in *pb.MyEmpty) (*slots.SyncUpdated, error) {
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

	reply, err := s.ss.es.GetSync().GetVarUpdated(ctx, in)
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) SetCableUpdated(ctx context.Context, in *slots.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid cable updated")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetSync().SetCableUpdated(ctx,
		&edges.SyncUpdated{Updated: in.GetUpdated()})
}

func (s *SyncService) GetCableUpdated(ctx context.Context, in *pb.MyEmpty) (*slots.SyncUpdated, error) {
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

	reply, err := s.ss.es.GetSync().GetCableUpdated(ctx, in)
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid wire updated")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetSync().SetWireUpdated(ctx,
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

	reply, err := s.ss.es.GetSync().GetWireUpdated(ctx, in)
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) SetTagValueUpdated(ctx context.Context, in *slots.SyncUpdated) (*pb.MyBool, error) {
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetSync().SetTagValueUpdated(ctx,
		&edges.SyncUpdated{Updated: in.GetUpdated()})
}

func (s *SyncService) GetTagValueUpdated(ctx context.Context, in *pb.MyEmpty) (*slots.SyncUpdated, error) {
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

	reply, err := s.ss.es.GetSync().GetTagValueUpdated(ctx, in)
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) WaitTagValueUpdated(in *pb.MyEmpty, stream slots.SyncService_WaitTagValueUpdatedServer) error {
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

	return s.ss.es.GetSync().WaitTagValueUpdated(in, stream)
}

func (s *SyncService) SetWireValueUpdated(ctx context.Context, in *slots.SyncUpdated) (*pb.MyBool, error) {
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

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.es.GetSync().SetWireValueUpdated(ctx,
		&edges.SyncUpdated{Updated: in.GetUpdated()})
}

func (s *SyncService) GetWireValueUpdated(ctx context.Context, in *pb.MyEmpty) (*slots.SyncUpdated, error) {
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

	reply, err := s.ss.es.GetSync().GetWireValueUpdated(ctx, in)
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) WaitWireValueUpdated(in *pb.MyEmpty, stream slots.SyncService_WaitWireValueUpdatedServer) error {
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

	return s.ss.es.GetSync().WaitWireValueUpdated(in, stream)
}
