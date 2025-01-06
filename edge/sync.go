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

	edges.UnimplementedSyncServiceServer
}

func newSyncService(es *EdgeService) *SyncService {
	return &SyncService{
		es:        es,
		waits:     make(map[chan struct{}]struct{}),
		waitsTVal: make(map[chan struct{}]struct{}),
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

	return s.waitUpdated(in, stream, NOTIFY)
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

	return s.waitUpdated(in, stream, NOTIFY_TVAL)
}

func (s *SyncService) getDeviceUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_DEVICE)
}

func (s *SyncService) setDeviceUpdated(ctx context.Context, updated time.Time) error {
	err := s.setUpdated(ctx, model.SYNC_DEVICE, updated)
	if err != nil {
		return err
	}

	s.notifyUpdated(NOTIFY)

	return nil
}

func (s *SyncService) getSlotUpdated(ctx context.Context) (time.Time, error) {
	return s.getUpdated(ctx, model.SYNC_SLOT)
}

func (s *SyncService) setSlotUpdated(ctx context.Context, updated time.Time) error {
	return s.setUpdated(ctx, model.SYNC_SLOT, updated)
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

	s.notifyUpdated(NOTIFY_TVAL)

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

func (s *SyncService) getUpdated(_ context.Context, key string) (time.Time, error) {
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

func (s *SyncService) setUpdated(_ context.Context, key string, updated time.Time) error {
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

type NotifyType int

const (
	NOTIFY      NotifyType = 0
	NOTIFY_TVAL NotifyType = 1
)

func (s *SyncService) notifyUpdated(nt NotifyType) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	switch nt {
	case NOTIFY:
		for wait := range s.waits {
			select {
			case wait <- struct{}{}:
			default:
			}
		}
	case NOTIFY_TVAL:
		for wait := range s.waitsTVal {
			select {
			case wait <- struct{}{}:
			default:
			}
		}
	}
}

func (s *SyncService) Notify(nt NotifyType) *Notify {
	ch := make(chan struct{}, 1)

	s.lock.Lock()
	switch nt {
	case NOTIFY:
		s.waits[ch] = struct{}{}
	case NOTIFY_TVAL:
		s.waitsTVal[ch] = struct{}{}
	}
	s.lock.Unlock()

	n := &Notify{
		s,
		ch,
		nt,
	}

	n.notify()

	return n
}

type Notify struct {
	ss *SyncService
	ch chan struct{}
	nt NotifyType
}

func (n *Notify) notify() {
	select {
	case n.ch <- struct{}{}:
	default:
	}
}

func (w *Notify) Wait() <-chan struct{} {
	return w.ch
}

func (n *Notify) Close() {
	n.ss.lock.Lock()
	defer n.ss.lock.Unlock()

	switch n.nt {
	case NOTIFY:
		delete(n.ss.waits, n.ch)
	case NOTIFY_TVAL:
		delete(n.ss.waitsTVal, n.ch)
	}
}

type waitUpdatedStream interface {
	Send(*pb.MyBool) error
	grpc.ServerStream
}

func (s *SyncService) waitUpdated(in *pb.MyEmpty, stream waitUpdatedStream, nt NotifyType) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	notify := s.Notify(nt)
	defer notify.Close()

	for {
		select {
		case <-notify.Wait():
			err = stream.Send(&pb.MyBool{Bool: true})
			if err != nil {
				return err
			}
		case <-stream.Context().Done():
			return nil
		}
	}
}
