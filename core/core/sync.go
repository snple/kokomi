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

	lock      sync.RWMutex
	waits     map[string]map[chan struct{}]struct{}
	waitsTVal map[string]map[chan struct{}]struct{}
	waitsWVal map[string]map[chan struct{}]struct{}

	cores.UnimplementedSyncServiceServer
}

func newSyncService(cs *CoreService) *SyncService {
	return &SyncService{
		cs:        cs,
		waits:     make(map[string]map[chan struct{}]struct{}),
		waitsTVal: make(map[string]map[chan struct{}]struct{}),
		waitsWVal: make(map[string]map[chan struct{}]struct{}),
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

		if len(in.GetId()) == 0 {
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

		if len(in.GetId()) == 0 {
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

	return s.waitUpdated(in, stream, s.waits)
}

func (s *SyncService) SetTagValueUpdated(ctx context.Context, in *cores.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
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

		if len(in.GetId()) == 0 {
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

	return s.waitUpdated(in, stream, s.waitsTVal)
}

func (s *SyncService) SetWireValueUpdated(ctx context.Context, in *cores.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.ID")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.Value.Updated")
		}
	}

	err = s.setWireValueUpdated(ctx, s.cs.GetDB(), in.GetId(), time.UnixMicro(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncService) GetWireValueUpdated(ctx context.Context, in *pb.Id) (*cores.SyncUpdated, error) {
	var output cores.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.ID")
		}
	}

	output.Id = in.GetId()

	t, err := s.getWireValueUpdated(ctx, s.cs.GetDB(), in.GetId())
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMicro()

	return &output, nil
}

func (s *SyncService) WaitWireValueUpdated(in *pb.Id,
	stream cores.SyncService_WaitWireValueUpdatedServer) error {

	return s.waitUpdated(in, stream, s.waitsWVal)
}

func (s *SyncService) WaitDeviceUpdated2(ctx context.Context, id string) chan bool {
	return s.waitUpdated2(ctx, id, s.waits)
}

func (s *SyncService) WaitTagValueUpdated2(ctx context.Context, id string) chan bool {
	return s.waitUpdated2(ctx, id, s.waitsTVal)
}

func (s *SyncService) WaitWireValueUpdated2(ctx context.Context, id string) chan bool {
	return s.waitUpdated2(ctx, id, s.waitsWVal)
}

func (s *SyncService) getDeviceUpdated(ctx context.Context, db bun.IDB, id string) (time.Time, error) {
	return s.getUpdated(ctx, db, id+model.SYNC_DEVICE_SUFFIX)
}

func (s *SyncService) setDeviceUpdated(ctx context.Context, db bun.IDB, id string, updated time.Time) error {
	err := s.setUpdated(ctx, db, id+model.SYNC_DEVICE_SUFFIX, updated)
	if err != nil {
		return err
	}

	s.notifyUpdated(id, s.waits)

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

	s.notifyUpdated(id, s.waitsTVal)

	return nil
}

func (s *SyncService) getWireValueUpdated(ctx context.Context, db bun.IDB, id string) (time.Time, error) {
	return s.getUpdated(ctx, db, id+model.SYNC_WIRE_VALUE_SUFFIX)
}

func (s *SyncService) setWireValueUpdated(ctx context.Context, db bun.IDB, id string, updated time.Time) error {
	err := s.setUpdated(ctx, db, id+model.SYNC_WIRE_VALUE_SUFFIX, updated)
	if err != nil {
		return err
	}

	s.notifyUpdated(id, s.waitsWVal)

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

func (s *SyncService) notifyUpdated(id string, waits map[string]map[chan struct{}]struct{}) {
	s.lock.RLock()
	if chans, ok := waits[id]; ok {
		for wait := range chans {
			select {
			case wait <- struct{}{}:
			default:
			}
		}
	}
	s.lock.RUnlock()
}

type waitUpdatedStream interface {
	Send(*pb.MyBool) error
	grpc.ServerStream
}

func (s *SyncService) waitUpdated(in *pb.Id,
	stream waitUpdatedStream,
	waits map[string]map[chan struct{}]struct{}) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
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

func (s *SyncService) waitUpdated2(ctx context.Context,
	id string, waits map[string]map[chan struct{}]struct{}) chan bool {
	wait := make(chan struct{}, 1)
	output := make(chan bool, 2)

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

	go func() {
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

		select {
		case <-wait:
			output <- true
		case <-ctx.Done():
			output <- false
		}
	}()

	return output
}

func (s *SyncService) destory(ctx context.Context, db bun.IDB, id string) error {
	ids := []string{id, id + model.SYNC_TAG_VALUE_SUFFIX, id + model.SYNC_WIRE_VALUE_SUFFIX}

	for _, id := range ids {
		_, err := db.NewDelete().Model(&model.Sync{}).Where("id = ?", id).Exec(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Delete: %v", err)
		}
	}

	return nil
}
