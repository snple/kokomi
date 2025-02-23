package edge

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/snple/beacon/consts"
	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/nodes"
	"github.com/snple/beacon/util/metadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NodeService struct {
	es *EdgeService

	NodeConn *grpc.ClientConn

	token string
	lock  sync.RWMutex

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup
}

func newNodeService(es *EdgeService) (*NodeService, error) {
	var err error

	es.Logger().Sugar().Infof("link node service: %v", es.dopts.NodeOptions.Addr)

	nodeConn, err := grpc.NewClient(es.dopts.NodeOptions.Addr, es.dopts.NodeOptions.GRPCOptions...)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(es.Context())

	s := &NodeService{
		es:       es,
		NodeConn: nodeConn,
		ctx:      ctx,
		cancel:   cancel,
	}

	return s, nil
}

func (s *NodeService) start() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	s.es.Logger().Sugar().Info("node service started")

	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			s.try()
		}
	}
}

func (s *NodeService) stop() {
	s.cancel()
	s.closeWG.Wait()
}

func (s *NodeService) push() error {
	// login
	operation := func() error {
		err := s.login(s.ctx)
		if err != nil {
			s.es.Logger().Sugar().Errorf("device login: %v", err)
		}

		return err
	}

	err := backoff.Retry(operation, backoff.WithContext(backoff.NewExponentialBackOff(), s.ctx))
	if err != nil {
		s.es.Logger().Sugar().Errorf("backoff.Retry: %v", err)
		return err
	}

	s.es.Logger().Sugar().Info("device login success")

	s.es.GetSync().setDeviceUpdatedLocalToRemote(s.ctx, time.Time{})

	if err := s.sync2(s.ctx); err != nil {
		return err
	}

	if err := s.syncTagValue2(s.ctx); err != nil {
		return err
	}

	s.es.Logger().Sugar().Info("push success")

	return nil
}

func (s *NodeService) pull() error {
	// login
	operation := func() error {
		err := s.login(s.ctx)
		if err != nil {
			s.es.Logger().Sugar().Errorf("device login: %v", err)
		}

		return err
	}

	err := backoff.Retry(operation, backoff.WithContext(backoff.NewExponentialBackOff(), s.ctx))
	if err != nil {
		s.es.Logger().Sugar().Errorf("backoff.Retry: %v", err)
		return err
	}

	s.es.Logger().Sugar().Info("device login success")

	s.es.GetSync().setDeviceUpdatedRemoteToLocal(s.ctx, time.Time{})

	if err := s.sync1(s.ctx); err != nil {
		return err
	}

	if err := s.syncTagValue1(s.ctx); err != nil {
		return err
	}

	s.es.Logger().Sugar().Info("pull success")

	return nil
}

func (s *NodeService) try() {
	// login
	operation := func() error {
		err := s.login(s.ctx)
		if err != nil {
			s.es.Logger().Sugar().Infof("device login: %v", err)
		}

		return err
	}

	err := backoff.Retry(operation, backoff.WithContext(backoff.NewExponentialBackOff(), s.ctx))
	if err != nil {
		s.es.Logger().Sugar().Errorf("backoff.Retry: %v", err)
		return
	}

	s.es.Logger().Sugar().Info("device login success")

	s.DeviceLink(s.ctx)
	s.link(true)
	defer s.link(false)

	if err := s.sync(s.ctx); err != nil {
		s.es.Logger().Sugar().Errorf("sync: %v", err)
		time.Sleep(time.Second * 3)
		return
	}

	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()

	// start ticker
	go s.ticker(ctx)

	// start realtime sync
	if s.es.dopts.SyncOptions.Realtime {
		go s.waitRemoteDeviceUpdated(ctx)
		go s.waitLocalDeviceUpdated(ctx)
		go s.waitRemoteTagValueUpdated(ctx)
		go s.waitLocalTagValueUpdated(ctx)
		go s.waitRemoteTagWriteUpdated(ctx)
		go s.waitLocalTagWriteUpdated(ctx)
	}

	// keep alive
	{
		ctx = metadata.SetToken(ctx, s.GetToken())

		stream, err := s.DeviceServiceClient().KeepAlive(ctx, &pb.MyEmpty{})
		if err != nil {
			s.es.Logger().Sugar().Errorf("KeepAlive: %v", err)
			return
		}

		for {
			reply, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}

				if code, ok := status.FromError(err); ok {
					if code.Code() == codes.Canceled {
						return
					}
				}

				s.es.Logger().Sugar().Errorf("KeepAlive.Recv(): %v", err)
				return
			}

			s.es.Logger().Sugar().Infof("keep alive reply: %+v", reply)
		}
	}
}

func (s *NodeService) IsLinked() bool {
	return s.es.GetStatus().GetDeviceLink() == consts.ON
}

func (s *NodeService) link(value bool) {
	if value {
		s.es.GetStatus().SetDeviceLink(consts.ON)
	} else {
		s.es.GetStatus().SetDeviceLink(consts.OFF)
	}
}

func (s *NodeService) ticker(ctx context.Context) {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	tokenRefreshTicker := time.NewTicker(s.es.dopts.SyncOptions.TokenRefresh)
	defer tokenRefreshTicker.Stop()

	linkStatusTicker := time.NewTicker(s.es.dopts.SyncOptions.Link)
	defer linkStatusTicker.Stop()

	syncTicker := time.NewTicker(s.es.dopts.SyncOptions.Interval)
	defer syncTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			s.es.Logger().Sugar().Info("ticker: ctx.Done()")
			return
		case <-tokenRefreshTicker.C:
			err := s.login(ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("device login: %v", err)
			}
		case <-linkStatusTicker.C:
			err := s.DeviceLink(ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("link device : %v", err)
			} else {
				s.es.Logger().Sugar().Info("link device ticker success")
			}

		case <-syncTicker.C:
			if err := s.sync(ctx); err != nil {
				s.es.Logger().Sugar().Errorf("sync: %v", err)
			}
		}
	}
}

func (s *NodeService) SyncServiceClient() nodes.SyncServiceClient {
	return nodes.NewSyncServiceClient(s.NodeConn)
}

func (s *NodeService) DeviceServiceClient() nodes.DeviceServiceClient {
	return nodes.NewDeviceServiceClient(s.NodeConn)
}

func (s *NodeService) SlotServiceClient() nodes.SlotServiceClient {
	return nodes.NewSlotServiceClient(s.NodeConn)
}

func (s *NodeService) SourceServiceClient() nodes.SourceServiceClient {
	return nodes.NewSourceServiceClient(s.NodeConn)
}

func (s *NodeService) TagServiceClient() nodes.TagServiceClient {
	return nodes.NewTagServiceClient(s.NodeConn)
}

func (s *NodeService) ConstServiceClient() nodes.ConstServiceClient {
	return nodes.NewConstServiceClient(s.NodeConn)
}

func (s *NodeService) GetToken() string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.token
}

func (s *NodeService) SetToken(ctx context.Context) context.Context {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return metadata.SetToken(ctx, s.token)
}

func (s *NodeService) DeviceLink(ctx context.Context) error {
	ctx = metadata.SetToken(ctx, s.GetToken())

	request := &nodes.DeviceLinkRequest{Status: consts.ON}
	_, err := s.DeviceServiceClient().Link(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func (s *NodeService) login(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var err error

	request := &nodes.DeviceLoginRequest{
		Id:     s.es.dopts.deviceID,
		Secret: s.es.dopts.secret,
	}

	// try login
	reply, err := s.DeviceServiceClient().Login(ctx, request)
	if err != nil {
		return err
	}

	if len(reply.GetToken()) == 0 {
		return errors.New("login: reply token is empty")
	}

	// set token
	s.lock.Lock()
	s.token = reply.GetToken()
	s.lock.Unlock()

	return nil
}

func (s *NodeService) waitRemoteDeviceUpdated(ctx context.Context) {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	for {
		ctx := metadata.SetToken(ctx, s.GetToken())

		stream, err := s.SyncServiceClient().WaitDeviceUpdated(ctx, &pb.MyEmpty{})
		if err != nil {
			if code, ok := status.FromError(err); ok {
				if code.Code() == codes.Canceled {
					return
				}
			}

			s.es.Logger().Sugar().Errorf("WaitDeviceUpdated: %v", err)

			// retry
			time.Sleep(s.es.dopts.SyncOptions.Retry)
			continue
		}

		for {
			_, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}

				if code, ok := status.FromError(err); ok {
					if code.Code() == codes.Canceled {
						return
					}
				}

				s.es.Logger().Sugar().Errorf("WaitDeviceUpdated.Recv(): %v", err)
				return
			}

			ctx = metadata.SetToken(ctx, s.GetToken())

			err = s.sync1(ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("sync1: %v", err)
			}
		}
	}
}

func (s *NodeService) waitLocalDeviceUpdated(ctx context.Context) {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	notify := s.es.GetSync().Notify(NOTIFY)
	defer notify.Close()

	for {
		select {
		case <-ctx.Done():
			return
		case <-notify.Wait():
			ctx = metadata.SetToken(ctx, s.GetToken())

			err := s.sync2(ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("sync2: %v", err)
			}
		}
	}
}

func (s *NodeService) waitRemoteTagValueUpdated(ctx context.Context) {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	for {
		ctx := metadata.SetToken(ctx, s.GetToken())

		stream, err := s.SyncServiceClient().WaitTagValueUpdated(ctx, &pb.MyEmpty{})
		if err != nil {
			if code, ok := status.FromError(err); ok {
				if code.Code() == codes.Canceled {
					return
				}
			}

			s.es.Logger().Sugar().Errorf("WaitTagValueUpdated: %v", err)

			// retry
			time.Sleep(s.es.dopts.SyncOptions.Retry)
			continue
		}

		for {
			_, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				if code, ok := status.FromError(err); ok {
					if code.Code() == codes.Canceled {
						return
					}
				}

				s.es.Logger().Sugar().Errorf("WaitTagValueUpdated.Recv(): %v", err)
				return
			}

			ctx = metadata.SetToken(ctx, s.GetToken())

			err = s.syncTagValue1(ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("syncTagValue1: %v", err)
			}
		}
	}
}

func (s *NodeService) waitLocalTagValueUpdated(ctx context.Context) {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	notify := s.es.GetSync().Notify(NOTIFY_TV)
	defer notify.Close()

	for {
		select {
		case <-ctx.Done():
			return
		case <-notify.Wait():
			ctx = metadata.SetToken(ctx, s.GetToken())

			err := s.syncTagValue2(ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("syncTagValue2: %v", err)
			}
		}
	}
}

func (s *NodeService) waitRemoteTagWriteUpdated(ctx context.Context) {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	for {
		ctx := metadata.SetToken(ctx, s.GetToken())

		stream, err := s.SyncServiceClient().WaitTagWriteUpdated(ctx, &pb.MyEmpty{})
		if err != nil {
			if code, ok := status.FromError(err); ok {
				if code.Code() == codes.Canceled {
					return
				}
			}

			s.es.Logger().Sugar().Errorf("WaitTagWriteUpdated: %v", err)

			// retry
			time.Sleep(s.es.dopts.SyncOptions.Retry)
			continue
		}

		for {
			_, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				if code, ok := status.FromError(err); ok {
					if code.Code() == codes.Canceled {
						return
					}
				}

				s.es.Logger().Sugar().Errorf("WaitTagWriteUpdated.Recv(): %v", err)
				return
			}

			ctx = metadata.SetToken(ctx, s.GetToken())

			err = s.syncTagWrite1(ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("syncTagWrite1: %v", err)
			}
		}
	}
}

func (s *NodeService) waitLocalTagWriteUpdated(ctx context.Context) {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	notify := s.es.GetSync().Notify(NOTIFY_TW)
	defer notify.Close()

	for {
		select {
		case <-ctx.Done():
			return
		case <-notify.Wait():
			ctx = metadata.SetToken(ctx, s.GetToken())

			err := s.syncTagWrite2(ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("syncTagWrite2: %v", err)
			}
		}
	}
}

func (s *NodeService) sync(ctx context.Context) error {
	ctx = metadata.SetToken(ctx, s.GetToken())

	if err := s.sync1(ctx); err != nil {
		return err
	}

	if err := s.sync2(ctx); err != nil {
		return err
	}

	if err := s.syncTagValue1(ctx); err != nil {
		return err
	}

	if err := s.syncTagValue2(ctx); err != nil {
		return err
	}

	if err := s.syncTagWrite1(ctx); err != nil {
		return err
	}

	return s.syncTagWrite2(ctx)
}

func (s *NodeService) sync1(ctx context.Context) error {
	return s.syncRemoteToLocal(ctx)
}

func (s *NodeService) sync2(ctx context.Context) error {
	return s.syncLocalToRemote(ctx)
}

func (s *NodeService) syncTagValue1(ctx context.Context) error {
	return s.syncTagValueRemoteToLocal(ctx)
}

func (s *NodeService) syncTagValue2(ctx context.Context) error {
	return s.syncTagValueLocalToRemote(ctx)
}

func (s *NodeService) syncTagWrite1(ctx context.Context) error {
	return s.syncTagWriteRemoteToLocal(ctx)
}

func (s *NodeService) syncTagWrite2(ctx context.Context) error {
	return s.syncTagWriteLocalToRemote(ctx)
}
