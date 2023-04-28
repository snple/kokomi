package slot

import (
	"context"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/pb/slots"
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

	return s.ss.Edge().GetSync().SetDeviceUpdated(ctx,
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

	reply, err := s.ss.Edge().GetSync().GetDeviceUpdated(ctx, in)
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

	return s.ss.Edge().GetSync().WaitDeviceUpdated(in, stream)
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

func (s *SyncService) SetOptionUpdated(ctx context.Context, in *slots.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid option updated")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetSync().SetOptionUpdated(ctx,
		&edges.SyncUpdated{Updated: in.GetUpdated()})
}

func (s *SyncService) GetOptionUpdated(ctx context.Context, in *pb.MyEmpty) (*slots.SyncUpdated, error) {
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

	reply, err := s.ss.Edge().GetSync().GetOptionUpdated(ctx, in)
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

	return s.ss.Edge().GetSync().SetSourceUpdated(ctx,
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

	reply, err := s.ss.Edge().GetSync().GetSourceUpdated(ctx, in)
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

	return s.ss.Edge().GetSync().SetTagUpdated(ctx,
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

	reply, err := s.ss.Edge().GetSync().GetTagUpdated(ctx, in)
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

	return s.ss.Edge().GetSync().SetVarUpdated(ctx,
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

	reply, err := s.ss.Edge().GetSync().GetVarUpdated(ctx, in)
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

	return s.ss.Edge().GetSync().SetCableUpdated(ctx,
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

	reply, err := s.ss.Edge().GetSync().GetCableUpdated(ctx, in)
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

func (s *SyncService) SetClassUpdated(ctx context.Context, in *slots.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid class updated")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetSync().SetClassUpdated(ctx,
		&edges.SyncUpdated{Updated: in.GetUpdated()})
}

func (s *SyncService) GetClassUpdated(ctx context.Context, in *pb.MyEmpty) (*slots.SyncUpdated, error) {
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

	reply, err := s.ss.Edge().GetSync().GetClassUpdated(ctx, in)
	if err != nil {
		return &output, err
	}

	output.Updated = reply.GetUpdated()

	return &output, nil
}

func (s *SyncService) SetAttrUpdated(ctx context.Context, in *slots.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid attr updated")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetSync().SetAttrUpdated(ctx,
		&edges.SyncUpdated{Updated: in.GetUpdated()})
}

func (s *SyncService) GetAttrUpdated(ctx context.Context, in *pb.MyEmpty) (*slots.SyncUpdated, error) {
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

	reply, err := s.ss.Edge().GetSync().GetAttrUpdated(ctx, in)
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

	return s.ss.Edge().GetSync().SetTagValueUpdated(ctx,
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

	reply, err := s.ss.Edge().GetSync().GetTagValueUpdated(ctx, in)
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

	return s.ss.Edge().GetSync().WaitTagValueUpdated(in, stream)
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

	return s.ss.Edge().GetSync().SetWireValueUpdated(ctx,
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

	reply, err := s.ss.Edge().GetSync().GetWireValueUpdated(ctx, in)
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

	return s.ss.Edge().GetSync().WaitWireValueUpdated(in, stream)
}
