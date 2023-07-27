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

	lock sync.RWMutex

	cores.UnimplementedSyncGlobalServiceServer
}

func newSyncGlobalService(cs *CoreService) *SyncGlobalService {
	return &SyncGlobalService{
		cs: cs,
	}
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

	return nil
}

func (s *SyncGlobalService) notifyUpdated(chans map[chan struct{}]struct{}) {
	s.lock.RLock()
	for wait := range chans {
		select {
		case wait <- struct{}{}:
		default:
		}
	}
	s.lock.RUnlock()
}

type globalWaitUpdatedStream interface {
	Send(*pb.MyBool) error
	grpc.ServerStream
}

func (s *SyncGlobalService) waitUpdated(in *pb.MyEmpty,
	stream globalWaitUpdatedStream, chans map[chan struct{}]struct{}) error {
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

func (s *SyncGlobalService) waitUpdated2(ctx context.Context,
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
