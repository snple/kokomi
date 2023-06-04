package edge

import (
	"context"
	"sync"
	"time"

	"github.com/dgraph-io/badger/v4"
)

type BadgerService struct {
	es *EdgeService

	badger *badger.DB

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup
}

func newBadgerService(es *EdgeService) (*BadgerService, error) {
	badger, err := badger.OpenManaged(es.dopts.BadgerOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(es.Context())

	s := &BadgerService{
		es:     es,
		badger: badger,
		ctx:    ctx,
		cancel: cancel,
	}

	return s, nil
}

func (s *BadgerService) start() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	s.es.Logger().Sugar().Info("badger service started")

	ticker := time.NewTicker(s.es.dopts.BadgerGCOptions.GC)
	defer ticker.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			{
			again:
				err := s.badger.RunValueLogGC(s.es.dopts.BadgerGCOptions.GCDiscardRatio)
				if err == nil {
					goto again
				}
			}
		}
	}
}

func (s *BadgerService) stop() {
	s.cancel()
	s.closeWG.Wait()

	s.badger.Close()
}

func (s *BadgerService) GetDB() *badger.DB {
	return s.badger
}
