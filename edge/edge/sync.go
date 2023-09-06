package edge

import (
	"context"
	"sync"
	"time"

	"github.com/dgraph-io/badger/v4"
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Device.Updated")
		}
	}

	err = s.setDeviceUpdated(ctx, time.UnixMicro(in.GetUpdated()))
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

	output.Updated = t.UnixMicro()

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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Slot.Updated")
		}
	}

	err = s.setSlotUpdated(ctx, time.UnixMicro(in.GetUpdated()))
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

	output.Updated = t.UnixMicro()

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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Option.Updated")
		}
	}

	err = s.setOptionUpdated(ctx, time.UnixMicro(in.GetUpdated()))
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

	output.Updated = t.UnixMicro()

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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Port.Updated")
		}
	}

	err = s.setPortUpdated(ctx, time.UnixMicro(in.GetUpdated()))
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

	output.Updated = t.UnixMicro()

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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Proxy.Updated")
		}
	}

	err = s.setProxyUpdated(ctx, time.UnixMicro(in.GetUpdated()))
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

	output.Updated = t.UnixMicro()

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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Source.Updated")
		}
	}

	err = s.setSourceUpdated(ctx, time.UnixMicro(in.GetUpdated()))
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

	output.Updated = t.UnixMicro()

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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Updated")
		}
	}

	err = s.setTagUpdated(ctx, time.UnixMicro(in.GetUpdated()))
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

	output.Updated = t.UnixMicro()

	return &output, nil
}

func (s *SyncService) SetConstUpdated(ctx context.Context, in *edges.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Updated")
		}
	}

	err = s.setConstUpdated(ctx, time.UnixMicro(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncService) GetConstUpdated(ctx context.Context, in *pb.MyEmpty) (*edges.SyncUpdated, error) {
	var output edges.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	t, err := s.getConstUpdated(ctx)
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMicro()

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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Value.Updated")
		}
	}

	err = s.setTagValueUpdated(ctx, time.UnixMicro(in.GetUpdated()))
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

	output.Updated = t.UnixMicro()

	return &output, nil
}

func (s *SyncService) WaitTagValueUpdated(in *pb.MyEmpty,
	stream edges.SyncService_WaitTagValueUpdatedServer) error {

	return s.waitUpdated(in, stream, s.waitsTVal)
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

func (s *SyncService) getConstUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_CONST)
}

func (s *SyncService) setConstUpdated(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_CONST, updated)
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

func (s *SyncService) getDeviceUpdatedRemoteToLocal(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_DEVICE_REMOTE_TO_LOCAL)
}

func (s *SyncService) setDeviceUpdatedRemoteToLocal(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_DEVICE_REMOTE_TO_LOCAL, updated)
}

func (s *SyncService) getDeviceUpdatedLocalToRemote(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_DEVICE_LOCAL_TO_REMOTE)
}

func (s *SyncService) setDeviceUpdatedLocalToRemote(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_DEVICE_LOCAL_TO_REMOTE, updated)
}

func (s *SyncService) getTagValueUpdatedRemoteToLocal(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_TAG_VALUE_REMOTE_TO_LOCAL)
}

func (s *SyncService) setTagValueUpdatedRemoteToLocal(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_TAG_VALUE_REMOTE_TO_LOCAL, updated)
}

func (s *SyncService) getTagValueUpdatedLocalToRemote(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_TAG_VALUE_LOCAL_TO_REMOTE)
}

func (s *SyncService) setTagValueUpdatedLocalToRemote(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_TAG_VALUE_LOCAL_TO_REMOTE, updated)
}

func (s *SyncService) getUpdated(ctx context.Context, key string) (time.Time, error) {
	txn := s.es.GetBadgerDB().NewTransactionAt(uint64(time.Now().UnixMicro()), false)
	defer txn.Discard()

	dbitem, err := txn.Get([]byte(key))
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return time.Time{}, nil
		}
		return time.Time{}, status.Errorf(codes.Internal, "BadgerDB Get: %v", err)
	}

	return time.UnixMicro(int64(dbitem.Version())), nil
}

func (s *SyncService) setUpdated(ctx context.Context, key string, updated time.Time) error {
	ts := uint64(updated.UnixMicro())

	txn := s.es.GetBadgerDB().NewTransactionAt(ts, true)
	defer txn.Discard()

	err := txn.Set([]byte(key), []byte{})
	if err != nil {
		return status.Errorf(codes.Internal, "BadgerDB Set: %v", err)
	}

	err = txn.CommitAt(ts, nil)
	if err != nil {
		return status.Errorf(codes.Internal, "BadgerDB CommitAt: %v", err)
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

	output <- true

	go func() {
		defer func() {
			s.lock.Lock()
			delete(chans, wait)
			s.lock.Unlock()
		}()

		select {
		case <-wait:
			output <- true
		case <-ctx.Done():
			output <- false
		}
	}()

	return output
}
