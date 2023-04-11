package edge

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/snple/kokomi/edge/model"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SyncService struct {
	es *EdgeService

	lock      sync.RWMutex
	waits     map[chan struct{}]struct{}
	waitsTVal map[chan struct{}]struct{}
	waitsWVal map[chan struct{}]struct{}

	edges.UnimplementedSyncServiceServer
}

func newSyncService(es *EdgeService) *SyncService {
	return &SyncService{
		es:        es,
		waits:     make(map[chan struct{}]struct{}),
		waitsTVal: make(map[chan struct{}]struct{}),
		waitsWVal: make(map[chan struct{}]struct{}),
	}
}

func (s *SyncService) SetDeviceUpdated(ctx context.Context, in *edges.SyncUpdated) (*pb.MyBool, error) {
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

	err = s.setDeviceUpdated(ctx, time.UnixMilli(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncService) GetDeviceUpdated(ctx context.Context, in *pb.MyEmpty) (*edges.SyncUpdated, error) {
	var output edges.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	t, err := s.getDeviceUpdated(ctx)
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMilli()

	return &output, nil
}

func (s *SyncService) WaitDeviceUpdated(in *pb.MyEmpty,
	stream edges.SyncService_WaitDeviceUpdatedServer) error {

	return s.waitUpdated(in, stream, s.waits)
}

func (s *SyncService) SetSlotUpdated(ctx context.Context, in *edges.SyncUpdated) (*pb.MyBool, error) {
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

	err = s.setSlotUpdated(ctx, time.UnixMilli(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncService) GetSlotUpdated(ctx context.Context, in *pb.MyEmpty) (*edges.SyncUpdated, error) {
	var output edges.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	t, err := s.getSlotUpdated(ctx)
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMilli()

	return &output, nil
}

func (s *SyncService) SetOptionUpdated(ctx context.Context, in *edges.SyncUpdated) (*pb.MyBool, error) {
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

	err = s.setOptionUpdated(ctx, time.UnixMilli(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncService) GetOptionUpdated(ctx context.Context, in *pb.MyEmpty) (*edges.SyncUpdated, error) {
	var output edges.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	t, err := s.getOptionUpdated(ctx)
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMilli()

	return &output, nil
}

func (s *SyncService) SetPortUpdated(ctx context.Context, in *edges.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid port updated")
		}
	}

	err = s.setPortUpdated(ctx, time.UnixMilli(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncService) GetPortUpdated(ctx context.Context, in *pb.MyEmpty) (*edges.SyncUpdated, error) {
	var output edges.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	t, err := s.getPortUpdated(ctx)
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMilli()

	return &output, nil
}

func (s *SyncService) SetProxyUpdated(ctx context.Context, in *edges.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid proxy updated")
		}
	}

	err = s.setProxyUpdated(ctx, time.UnixMilli(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncService) GetProxyUpdated(ctx context.Context, in *pb.MyEmpty) (*edges.SyncUpdated, error) {
	var output edges.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	t, err := s.getProxyUpdated(ctx)
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMilli()

	return &output, nil
}

func (s *SyncService) SetSourceUpdated(ctx context.Context, in *edges.SyncUpdated) (*pb.MyBool, error) {
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

	err = s.setSourceUpdated(ctx, time.UnixMilli(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncService) GetSourceUpdated(ctx context.Context, in *pb.MyEmpty) (*edges.SyncUpdated, error) {
	var output edges.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	t, err := s.getSourceUpdated(ctx)
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMilli()

	return &output, nil
}

func (s *SyncService) SetTagUpdated(ctx context.Context, in *edges.SyncUpdated) (*pb.MyBool, error) {
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

	err = s.setTagUpdated(ctx, time.UnixMilli(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncService) GetTagUpdated(ctx context.Context, in *pb.MyEmpty) (*edges.SyncUpdated, error) {
	var output edges.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	t, err := s.getTagUpdated(ctx)
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMilli()

	return &output, nil
}

func (s *SyncService) SetVarUpdated(ctx context.Context, in *edges.SyncUpdated) (*pb.MyBool, error) {
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

	err = s.setVarUpdated(ctx, time.UnixMilli(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncService) GetVarUpdated(ctx context.Context, in *pb.MyEmpty) (*edges.SyncUpdated, error) {
	var output edges.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	t, err := s.getVarUpdated(ctx)
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMilli()

	return &output, nil
}

func (s *SyncService) SetCableUpdated(ctx context.Context, in *edges.SyncUpdated) (*pb.MyBool, error) {
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

	err = s.setCableUpdated(ctx, time.UnixMilli(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncService) GetCableUpdated(ctx context.Context, in *pb.MyEmpty) (*edges.SyncUpdated, error) {
	var output edges.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	t, err := s.getCableUpdated(ctx)
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMilli()

	return &output, nil
}

func (s *SyncService) SetWireUpdated(ctx context.Context, in *edges.SyncUpdated) (*pb.MyBool, error) {
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

	err = s.setWireUpdated(ctx, time.UnixMilli(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncService) GetWireUpdated(ctx context.Context, in *pb.MyEmpty) (*edges.SyncUpdated, error) {
	var output edges.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	t, err := s.getWireUpdated(ctx)
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMilli()

	return &output, nil
}

func (s *SyncService) SetClassUpdated(ctx context.Context, in *edges.SyncUpdated) (*pb.MyBool, error) {
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

	err = s.setClassUpdated(ctx, time.UnixMilli(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncService) GetClassUpdated(ctx context.Context, in *pb.MyEmpty) (*edges.SyncUpdated, error) {
	var output edges.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	t, err := s.getClassUpdated(ctx)
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMilli()

	return &output, nil
}

func (s *SyncService) SetAttrUpdated(ctx context.Context, in *edges.SyncUpdated) (*pb.MyBool, error) {
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

	err = s.setAttrUpdated(ctx, time.UnixMilli(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncService) GetAttrUpdated(ctx context.Context, in *pb.MyEmpty) (*edges.SyncUpdated, error) {
	var output edges.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	t, err := s.getAttrUpdated(ctx)
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMilli()

	return &output, nil
}

func (s *SyncService) SetTagValueUpdated(ctx context.Context, in *edges.SyncUpdated) (*pb.MyBool, error) {
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

	err = s.setTagValueUpdated(ctx, time.UnixMilli(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncService) GetTagValueUpdated(ctx context.Context, in *pb.MyEmpty) (*edges.SyncUpdated, error) {
	var output edges.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	t, err := s.getTagValueUpdated(ctx)
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMilli()

	return &output, nil
}

func (s *SyncService) WaitTagValueUpdated(in *pb.MyEmpty,
	stream edges.SyncService_WaitTagValueUpdatedServer) error {

	return s.waitUpdated(in, stream, s.waitsTVal)
}

func (s *SyncService) SetWireValueUpdated(ctx context.Context, in *edges.SyncUpdated) (*pb.MyBool, error) {
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

	err = s.setWireValueUpdated(ctx, time.UnixMilli(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncService) GetWireValueUpdated(ctx context.Context, in *pb.MyEmpty) (*edges.SyncUpdated, error) {
	var output edges.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	t, err := s.getWireValueUpdated(ctx)
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMilli()

	return &output, nil
}

func (s *SyncService) WaitWireValueUpdated(in *pb.MyEmpty,
	stream edges.SyncService_WaitWireValueUpdatedServer) error {

	return s.waitUpdated(in, stream, s.waitsWVal)
}

func (s *SyncService) WaitDeviceUpdated2(ctx context.Context) chan bool {
	return s.waitUpdated2(ctx, s.waits)
}

func (s *SyncService) WaitTagValueUpdated2(ctx context.Context) chan bool {
	return s.waitUpdated2(ctx, s.waitsTVal)
}

func (s *SyncService) WaitWireValueUpdated2(ctx context.Context) chan bool {
	return s.waitUpdated2(ctx, s.waitsWVal)
}

func (s *SyncService) getDeviceUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_DEVICE)
}

func (s *SyncService) setDeviceUpdated(ctx context.Context, updated time.Time) error {
	err := s.setUpdated(ctx, model.SYNC_DEVICE, updated)
	if err != nil {
		return err
	}

	s.notifyUpdated(s.waits)

	return nil
}

func (s *SyncService) getSlotUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_SLOT)
}

func (s *SyncService) setSlotUpdated(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_SLOT, updated)
}

func (s *SyncService) getOptionUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_OPTION)
}

func (s *SyncService) setOptionUpdated(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_OPTION, updated)
}

func (s *SyncService) getPortUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_PORT)
}

func (s *SyncService) setPortUpdated(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_PORT, updated)
}

func (s *SyncService) getProxyUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_PROXY)
}

func (s *SyncService) setProxyUpdated(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_PROXY, updated)
}

func (s *SyncService) getSourceUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_SOURCE)
}

func (s *SyncService) setSourceUpdated(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_SOURCE, updated)
}

func (s *SyncService) getTagUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_TAG)
}

func (s *SyncService) setTagUpdated(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_TAG, updated)
}

func (s *SyncService) getVarUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_VAR)
}

func (s *SyncService) setVarUpdated(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_VAR, updated)
}

func (s *SyncService) getCableUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_CABLE)
}

func (s *SyncService) setCableUpdated(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_CABLE, updated)
}

func (s *SyncService) getWireUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_WIRE)
}

func (s *SyncService) setWireUpdated(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_WIRE, updated)
}

func (s *SyncService) getClassUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_CLASS)
}

func (s *SyncService) setClassUpdated(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_CLASS, updated)
}

func (s *SyncService) getAttrUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_ATTR)
}

func (s *SyncService) setAttrUpdated(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_ATTR, updated)
}

func (s *SyncService) getTagValueUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_TAG_VALUE)
}

func (s *SyncService) setTagValueUpdated(ctx context.Context, updated time.Time) error {
	err := s.setUpdated(ctx, model.SYNC_TAG_VALUE, updated)
	if err != nil {
		return err
	}

	s.notifyUpdated(s.waitsTVal)

	return nil
}

func (s *SyncService) getWireValueUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_WIRE_VALUE)
}

func (s *SyncService) setWireValueUpdated(ctx context.Context, updated time.Time) error {
	err := s.setUpdated(ctx, model.SYNC_WIRE_VALUE, updated)
	if err != nil {
		return err
	}

	s.notifyUpdated(s.waitsWVal)

	return nil
}

func (s *SyncService) getLocalDeviceUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_LOCAL_DEVICE)
}

func (s *SyncService) setLocalDeviceUpdated(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_LOCAL_DEVICE, updated)
}

func (s *SyncService) getRemoteDeviceUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_REMOTE_DEVICE)
}

func (s *SyncService) setRemoteDeviceUpdated(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_REMOTE_DEVICE, updated)
}

func (s *SyncService) getLocalTagValueUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_LOCAL_TAG_VALUE)
}

func (s *SyncService) setLocalTagValueUpdated(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_LOCAL_TAG_VALUE, updated)
}

func (s *SyncService) getRemoteTagValueUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_REMOTE_TAG_VALUE)
}

func (s *SyncService) setRemoteTagValueUpdated(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_REMOTE_TAG_VALUE, updated)
}

func (s *SyncService) getLocalWireValueUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_LOCAL_WIRE_VALUE)
}

func (s *SyncService) setLocalWireValueUpdated(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_LOCAL_WIRE_VALUE, updated)
}

func (s *SyncService) getRemoteWireValueUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_REMOTE_WIRE_VALUE)
}

func (s *SyncService) setRemoteWireValueUpdated(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_REMOTE_WIRE_VALUE, updated)
}

func (s *SyncService) getUpdated(ctx context.Context, key string) (time.Time, error) {
	item := model.Sync{
		Key: key,
	}

	err := s.es.GetDB().NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err = s.es.GetDB().NewInsert().Model(&item).Exec(ctx)
			if err != nil {
				return time.Time{}, status.Errorf(codes.Internal, "Insert: %v", err)
			}

			return time.Time{}, nil
		}

		return time.Time{}, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item.Updated, nil
}

func (s *SyncService) setUpdated(ctx context.Context, key string, updated time.Time) error {
	item := model.Sync{
		Key:     key,
		Updated: updated,
	}

	ret, err := s.es.GetDB().NewUpdate().Model(&item).WherePK().Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Update: %v", err)
	}

	n, err := ret.RowsAffected()
	if err != nil {
		return status.Errorf(codes.Internal, "RowsAffected: %v", err)
	}

	if n < 1 {
		_, err = s.es.GetDB().NewInsert().Model(&item).Exec(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func (s *SyncService) notifyUpdated(chans map[chan struct{}]struct{}) {
	s.lock.RLock()
	for wait := range chans {
		select {
		case wait <- struct{}{}:
		default:
		}
	}
	s.lock.RUnlock()
}

type waitUpdatedStream interface {
	Send(*pb.MyBool) error
	grpc.ServerStream
}

func (s *SyncService) waitUpdated(in *pb.MyEmpty,
	stream waitUpdatedStream, chans map[chan struct{}]struct{}) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	wait := make(chan struct{}, 1)

	s.lock.Lock()
	chans[wait] = struct{}{}
	s.lock.Unlock()

	defer func() {
		s.lock.Lock()
		delete(chans, wait)
		s.lock.Unlock()
	}()

	err = stream.Send(&pb.MyBool{Bool: true})
	if err != nil {
		return err
	}

	select {
	case <-wait:
	case <-stream.Context().Done():
		return nil
	}

	return stream.Send(&pb.MyBool{Bool: true})
}

func (s *SyncService) waitUpdated2(ctx context.Context,
	chans map[chan struct{}]struct{}) chan bool {
	wait := make(chan struct{}, 1)
	output := make(chan bool, 2)

	s.lock.Lock()
	chans[wait] = struct{}{}
	s.lock.Unlock()

	defer func() {
		s.lock.Lock()
		delete(chans, wait)
		s.lock.Unlock()
	}()

	output <- true

	go func() {
		select {
		case <-wait:
			output <- true
		case <-ctx.Done():
			output <- false
		}
	}()

	return output
}
