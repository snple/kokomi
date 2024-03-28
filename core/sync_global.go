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

type SyncGlobalService struct {
	cs *CoreService

	lock  sync.RWMutex
	waits map[string]map[chan struct{}]struct{}

	cores.UnimplementedSyncGlobalServiceServer
}

func newSyncGlobalService(cs *CoreService) *SyncGlobalService {
	return &SyncGlobalService{
		cs:    cs,
		waits: make(map[string]map[chan struct{}]struct{}),
	}
}

func (s *SyncGlobalService) SetUpdated(ctx context.Context, in *cores.SyncUpdated) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid ID")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Updated")
		}
	}

	err = s.setUpdated(ctx, s.cs.GetDB(), in.GetId(), time.UnixMicro(in.GetUpdated()))
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SyncGlobalService) GetUpdated(ctx context.Context, in *pb.Id) (*cores.SyncUpdated, error) {
	var output cores.SyncUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid ID")
		}
	}

	output.Id = in.GetId()

	t, err := s.getUpdated(ctx, s.cs.GetDB(), in.GetId())
	if err != nil {
		return &output, err
	}

	output.Updated = t.UnixMicro()

	return &output, nil
}

func (s *SyncGlobalService) WaitDeviceUpdated(in *pb.Id,
	stream cores.SyncGlobalService_WaitUpdatedServer) error {

	return s.waitUpdated(in, stream)
}

func (s *SyncGlobalService) getUpdated(ctx context.Context, db bun.IDB, key string) (time.Time, error) {
	item := model.SyncGlobal{
		Key: key,
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

func (s *SyncGlobalService) setUpdated(ctx context.Context, db bun.IDB, key string, updated time.Time) error {
	item := model.SyncGlobal{
		Key:     key,
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

	s.notifyUpdated(key)

	return nil
}

func (s *SyncGlobalService) notifyUpdated(id string) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if waits, ok := s.waits[id]; ok {
		for wait := range waits {
			select {
			case wait <- struct{}{}:
			default:
			}
		}
	}
}

func (s *SyncGlobalService) Notify(id string) *GlobalNotify {
	ch := make(chan struct{}, 1)

	s.lock.Lock()
	if waits, ok := s.waits[id]; ok {
		waits[ch] = struct{}{}
	} else {
		waits := map[chan struct{}]struct{}{
			ch: {},
		}
		s.waits[id] = waits
	}
	s.lock.Unlock()

	n := &GlobalNotify{
		id,
		s,
		ch,
	}

	n.notify()

	return n
}

type GlobalNotify struct {
	id string
	ss *SyncGlobalService
	ch chan struct{}
}

func (n *GlobalNotify) notify() {
	select {
	case n.ch <- struct{}{}:
	default:
	}
}

func (w *GlobalNotify) Wait() <-chan struct{} {
	return w.ch
}

func (n *GlobalNotify) Close() {
	n.ss.lock.Lock()
	defer n.ss.lock.Unlock()

	if waits, ok := n.ss.waits[n.id]; ok {
		delete(waits, n.ch)

		if len(waits) == 0 {
			delete(n.ss.waits, n.id)
		}
	}
}

func (n *GlobalNotify) Id() string {
	return n.id
}

type globalWaitUpdatedStream interface {
	Send(*pb.MyBool) error
	grpc.ServerStream
}

func (s *SyncGlobalService) waitUpdated(in *pb.Id, stream globalWaitUpdatedStream) error {
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

	notify := s.Notify(in.GetId())
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
