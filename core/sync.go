package core

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/snple/kokomi/core/model"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
	"github.com/uptrace/bun"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SyncService struct {
	cs *CoreService

	lock    sync.RWMutex
	waits   map[string]map[chan struct{}]struct{}
	waitsTV map[string]map[chan struct{}]struct{}
	waitsTW map[string]map[chan struct{}]struct{}

	cores.UnimplementedSyncServiceServer
}

func newSyncService(cs *CoreService) *SyncService {
	return &SyncService{
		cs:      cs,
		waits:   make(map[string]map[chan struct{}]struct{}),
		waitsTV: make(map[string]map[chan struct{}]struct{}),
		waitsTW: make(map[string]map[chan struct{}]struct{}),
	}
}

func (s *SyncService) SetDeviceUpdated(ctx context.Context, in *cores.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Device.ID")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Device.Updated")
		}
	}

	err = s.setDeviceUpdated(ctx, s.cs.GetDB(), in.GetId(), time.UnixMicro(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncService) GetDeviceUpdated(ctx context.Context, in *pb.Id) (*cores.SyncUpdated, error) {
	var output cores.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Device.ID")
		}
	}

	output.Id = in.GetId()

	t, err := s.getDeviceUpdated(ctx, s.cs.GetDB(), in.GetId())
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMicro()

	return &output, nil
}

func (s *SyncService) WaitDeviceUpdated(in *pb.Id,
	stream cores.SyncService_WaitDeviceUpdatedServer) error {

	return s.waitUpdated(in, stream, NOTIFY)
}

func (s *SyncService) SetTagValueUpdated(ctx context.Context, in *cores.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Value.Updated")
		}
	}

	err = s.setTagValueUpdated(ctx, s.cs.GetDB(), in.GetId(), time.UnixMicro(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncService) GetTagValueUpdated(ctx context.Context, in *pb.Id) (*cores.SyncUpdated, error) {
	var output cores.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}
	}

	output.Id = in.GetId()

	t, err := s.getTagValueUpdated(ctx, s.cs.GetDB(), in.GetId())
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMicro()

	return &output, nil
}

func (s *SyncService) WaitTagValueUpdated(in *pb.Id,
	stream cores.SyncService_WaitTagValueUpdatedServer) error {

	return s.waitUpdated(in, stream, NOTIFY_TV)
}

func (s *SyncService) SetTagWriteUpdated(ctx context.Context, in *cores.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Value.Updated")
		}
	}

	err = s.setTagValueUpdated(ctx, s.cs.GetDB(), in.GetId(), time.UnixMicro(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncService) GetTagWriteUpdated(ctx context.Context, in *pb.Id) (*cores.SyncUpdated, error) {
	var output cores.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}
	}

	output.Id = in.GetId()

	t, err := s.getTagValueUpdated(ctx, s.cs.GetDB(), in.GetId())
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMicro()

	return &output, nil
}

func (s *SyncService) WaitTagWriteUpdated(in *pb.Id,
	stream cores.SyncService_WaitTagWriteUpdatedServer) error {

	return s.waitUpdated(in, stream, NOTIFY_TV)
}

func (s *SyncService) getDeviceUpdated(ctx context.Context, db bun.IDB, id string) (time.Time, error) {
	return s.getUpdated(ctx, db, id+model.SYNC_DEVICE_SUFFIX)
}

func (s *SyncService) setDeviceUpdated(ctx context.Context, db bun.IDB, id string, updated time.Time) error {
	err := s.setUpdated(ctx, db, id+model.SYNC_DEVICE_SUFFIX, updated)
	if err != nil {
		return err
	}

	s.notifyUpdated(id, NOTIFY)

	return nil
}

func (s *SyncService) getTagValueUpdated(ctx context.Context, db bun.IDB, id string) (time.Time, error) {
	return s.getUpdated(ctx, db, id+model.SYNC_TAG_VALUE_SUFFIX)
}

func (s *SyncService) setTagValueUpdated(ctx context.Context, db bun.IDB, id string, updated time.Time) error {
	err := s.setUpdated(ctx, db, id+model.SYNC_TAG_VALUE_SUFFIX, updated)
	if err != nil {
		return err
	}

	s.notifyUpdated(id, NOTIFY_TV)

	return nil
}

func (s *SyncService) getTagWriteUpdated(ctx context.Context, db bun.IDB, id string) (time.Time, error) {
	return s.getUpdated(ctx, db, id+model.SYNC_TAG_WRITE_SUFFIX)
}

func (s *SyncService) setTagWriteUpdated(ctx context.Context, db bun.IDB, id string, updated time.Time) error {
	err := s.setUpdated(ctx, db, id+model.SYNC_TAG_WRITE_SUFFIX, updated)
	if err != nil {
		return err
	}

	s.notifyUpdated(id, NOTIFY_TW)

	return nil
}

func (s *SyncService) getUpdated(ctx context.Context, db bun.IDB, id string) (time.Time, error) {
	item := model.Sync{
		ID: id,
	}

	err := db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err = db.NewInsert().Model(&item).Exec(ctx)
			if err != nil {
				return time.Time{}, status.Errorf(codes.Internal, "Insert: %v", err)
			}

			return time.Time{}, nil
		}

		return time.Time{}, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item.Updated, nil
}

func (s *SyncService) setUpdated(ctx context.Context, db bun.IDB, id string, updated time.Time) error {
	item := model.Sync{
		ID:      id,
		Updated: updated,
	}

	ret, err := db.NewUpdate().Model(&item).WherePK().Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Update: %v", err)
	}

	n, err := ret.RowsAffected()
	if err != nil {
		return status.Errorf(codes.Internal, "RowsAffected: %v", err)
	}

	if n < 1 {
		_, err = db.NewInsert().Model(&item).Exec(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

type NotifyType int

const (
	NOTIFY    NotifyType = 0
	NOTIFY_TV NotifyType = 1
	NOTIFY_TW NotifyType = 2
)

func (s *SyncService) notifyUpdated(id string, nt NotifyType) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	switch nt {
	case NOTIFY:
		if waits, ok := s.waits[id]; ok {
			for wait := range waits {
				select {
				case wait <- struct{}{}:
				default:
				}
			}
		}
	case NOTIFY_TV:
		if waits, ok := s.waitsTV[id]; ok {
			for wait := range waits {
				select {
				case wait <- struct{}{}:
				default:
				}
			}
		}
	case NOTIFY_TW:
		if waits, ok := s.waitsTW[id]; ok {
			for wait := range waits {
				select {
				case wait <- struct{}{}:
				default:
				}
			}
		}
	}
}

func (s *SyncService) Notify(id string, nt NotifyType) *Notify {
	ch := make(chan struct{}, 1)

	s.lock.Lock()

	switch nt {
	case NOTIFY:
		if waits, ok := s.waits[id]; ok {
			waits[ch] = struct{}{}
		} else {
			waits := map[chan struct{}]struct{}{
				ch: {},
			}
			s.waits[id] = waits
		}
	case NOTIFY_TV:
		if waits, ok := s.waitsTV[id]; ok {
			waits[ch] = struct{}{}
		} else {
			waits := map[chan struct{}]struct{}{
				ch: {},
			}
			s.waitsTV[id] = waits
		}
	case NOTIFY_TW:
		if waits, ok := s.waitsTW[id]; ok {
			waits[ch] = struct{}{}
		} else {
			waits := map[chan struct{}]struct{}{
				ch: {},
			}
			s.waitsTW[id] = waits
		}
	}

	s.lock.Unlock()

	n := &Notify{
		id,
		s,
		ch,
		nt,
	}

	n.notify()

	return n
}

type Notify struct {
	id string
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
		if waits, ok := n.ss.waits[n.id]; ok {
			delete(waits, n.ch)

			if len(waits) == 0 {
				delete(n.ss.waits, n.id)
			}
		}
	case NOTIFY_TV:
		if waits, ok := n.ss.waitsTV[n.id]; ok {
			delete(waits, n.ch)

			if len(waits) == 0 {
				delete(n.ss.waitsTV, n.id)
			}
		}
	case NOTIFY_TW:
		if waits, ok := n.ss.waitsTW[n.id]; ok {
			delete(waits, n.ch)

			if len(waits) == 0 {
				delete(n.ss.waitsTW, n.id)
			}
		}
	}
}

func (n *Notify) Id() string {
	return n.id
}

type waitUpdatedStream interface {
	Send(*pb.MyBool) error
	grpc.ServerStream
}

func (s *SyncService) waitUpdated(in *pb.Id, stream waitUpdatedStream, nt NotifyType) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return status.Error(codes.InvalidArgument, "Please supply valid DeviceID")
		}
	}

	notify := s.Notify(in.GetId(), nt)
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

func (s *SyncService) waitUpdated2(in *pb.Id,
	stream waitUpdatedStream,
	waits map[string]map[chan struct{}]struct{}) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return status.Error(codes.InvalidArgument, "Please supply valid DeviceID")
		}
	}

	id := in.GetId()

	wait := make(chan struct{}, 1)

	s.lock.Lock()
	if chans, ok := waits[id]; ok {
		chans[wait] = struct{}{}
	} else {
		chans := map[chan struct{}]struct{}{
			wait: {},
		}
		waits[id] = chans
	}
	s.lock.Unlock()

	defer func() {
		s.lock.Lock()
		if chans, ok := waits[id]; ok {
			delete(chans, wait)

			if len(chans) == 0 {
				delete(waits, id)
			}
		}
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

func (s *SyncService) destory(ctx context.Context, db bun.IDB, id string) error {
	ids := []string{id, id + model.SYNC_TAG_VALUE_SUFFIX, id + model.SYNC_TAG_WRITE_SUFFIX}

	for _, id := range ids {
		_, err := db.NewDelete().Model(&model.Sync{}).Where("id = ?", id).Exec(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Delete: %v", err)
		}
	}

	return nil
}
