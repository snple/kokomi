package core

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/snple/kokomi/core/model"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SyncGlobalService struct {
	cs *CoreService

	lock sync.RWMutex
	// waitsUser map[chan struct{}]struct{}

	cores.UnimplementedSyncGlobalServiceServer
}

func newSyncGlobalService(cs *CoreService) *SyncGlobalService {
	return &SyncGlobalService{
		cs: cs,
	}
}

// func (s *SyncGlobalService) SetUserUpdated(ctx context.Context, in *cores.SyncUpdated) (*pb.MyBool, error) {
// 	var output pb.MyBool
// 	var err error

// 	// basic validation
// 	{
// 		if in == nil {
// 			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
// 		}

// 		if in.GetUpdated() == 0 {
// 			return &output, status.Error(codes.InvalidArgument, "Please supply valid user updated")
// 		}
// 	}

// 	err = s.setUserUpdated(ctx, time.UnixMilli(in.GetUpdated()))
// 	if err != nil {
// 		return &output, err
// 	}

// 	output.Bool = true

// 	return &output, nil
// }

// func (s *SyncGlobalService) GetUserUpdated(ctx context.Context, in *pb.MyEmpty) (*cores.SyncUpdated, error) {
// 	var output cores.SyncUpdated
// 	var err error

// 	// basic validation
// 	{
// 		if in == nil {
// 			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
// 		}
// 	}

// 	t, err := s.getUserUpdated(ctx)
// 	if err != nil {
// 		return &output, err
// 	}

// 	output.Updated = t.UnixMilli()

// 	return &output, nil
// }

// func (s *SyncGlobalService) WaitUserUpdated(in *pb.MyEmpty,
// 	stream cores.SyncGlobalService_WaitUserUpdatedServer) error {

// 	return s.waitUpdated(in, stream, s.waitsUser)
// }

// func (s *SyncGlobalService) WaitUserUpdated2(ctx context.Context) chan bool {
// 	return s.waitUpdated2(ctx, s.waitsUser)
// }

// func (s *SyncGlobalService) getUserUpdated(ctx context.Context) (time.Time, error) {
// 	return s.getUpdated(ctx, model.SYNC_USER)
// }

// func (s *SyncGlobalService) setUserUpdated(ctx context.Context, updated time.Time) error {
// 	err := s.setUpdated(ctx, model.SYNC_USER, updated)
// 	if err != nil {
// 		return err
// 	}

// 	s.notifyUpdated(s.waitsUser)

// 	return nil
// }

func (s *SyncGlobalService) getUpdated(ctx context.Context, key string) (time.Time, error) {
	item := model.SyncGlobal{
		Key: key,
	}

	err := s.cs.GetDB().NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err = s.cs.GetDB().NewInsert().Model(&item).Exec(ctx)
			if err != nil {
				return time.Time{}, status.Errorf(codes.Internal, "Insert: %v", err)
			}

			return time.Time{}, nil
		}

		return time.Time{}, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item.Updated, nil
}

func (s *SyncGlobalService) setUpdated(ctx context.Context, key string, updated time.Time) error {
	item := model.SyncGlobal{
		Key:     key,
		Updated: updated,
	}

	ret, err := s.cs.GetDB().NewUpdate().Model(&item).WherePK().Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Update: %v", err)
	}

	n, err := ret.RowsAffected()
	if err != nil {
		return status.Errorf(codes.Internal, "RowsAffected: %v", err)
	}

	if n < 1 {
		_, err = s.cs.GetDB().NewInsert().Model(&item).Exec(ctx)
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
