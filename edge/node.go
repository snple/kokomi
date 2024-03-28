package edge

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/pb/nodes"
	"github.com/snple/kokomi/util/metadata"
	"github.com/snple/rgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NodeService struct {
	es *EdgeService

	NodeConn *grpc.ClientConn

	token string
	lock  sync.RWMutex

	handlerMap rgrpc.HandlerMap

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup
}

func newNodeService(es *EdgeService) (*NodeService, error) {
	var err error

	es.Logger().Sugar().Infof("link node service: %v", es.dopts.NodeOptions.Addr)

	nodeConn, err := grpc.Dial(es.dopts.NodeOptions.Addr, es.dopts.NodeOptions.GRPCOptions...)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(es.Context())

	s := &NodeService{
		es:         es,
		NodeConn:   nodeConn,
		handlerMap: rgrpc.HandlerMap{},
		ctx:        ctx,
		cancel:     cancel,
	}

	edges.RegisterControlServiceServer(s.handlerMap, s.es.GetControl())

	return s, nil
}

func (s *NodeService) start() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	s.es.Logger().Sugar().Info("node service started")

	go s.ticker()

	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			s.loop()
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

func (s *NodeService) loop() {
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
		return
	}

	s.es.Logger().Sugar().Info("device login success")

	s.DeviceLink(s.ctx)
	s.es.GetStatus().SetDeviceLink(consts.ON)
	defer s.es.GetStatus().SetDeviceLink(consts.OFF)

	if err := s.sync(s.ctx); err != nil {
		s.es.Logger().Sugar().Errorf("sync: %v", err)
		time.Sleep(time.Second * 3)
		return
	}

	if s.es.dopts.SyncOptions.Realtime {
		go s.waitRemoteDeviceUpdated()
		go s.waitLocalDeviceUpdated()
		go s.waitRemoteTagValueUpdated()
		go s.waitLocalTagValueUpdated()
	}

	err = s.linkRgrpc(s.ctx)
	if err != nil {
		if code, ok := status.FromError(err); ok {
			if code.Code() == codes.Canceled {
				return
			}
		}

		s.es.Logger().Sugar().Errorf("link rgrpc: %v", err)
	}
}

func (s *NodeService) ticker() {
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
		case <-s.ctx.Done():
			return
		case <-tokenRefreshTicker.C:
			if s.es.GetStatus().GetDeviceLink() == consts.ON {
				err := s.login(s.ctx)
				if err != nil {
					s.es.Logger().Sugar().Errorf("device login: %v", err)
				}
			}
		case <-linkStatusTicker.C:
			if option := s.es.GetQuic(); option.IsNone() {
				if s.es.GetStatus().GetDeviceLink() == consts.ON {
					err := s.DeviceLink(s.ctx)
					if err != nil {
						s.es.Logger().Sugar().Errorf("link device : %v", err)
					} else {
						s.es.Logger().Sugar().Info("link device ticker success")
					}
				}
			}
		case <-syncTicker.C:
			if s.es.GetStatus().GetDeviceLink() == consts.ON {
				if err := s.sync(s.ctx); err != nil {
					s.es.Logger().Sugar().Errorf("sync: %v", err)
				}
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

func (s *NodeService) PortServiceClient() nodes.PortServiceClient {
	return nodes.NewPortServiceClient(s.NodeConn)
}

func (s *NodeService) ProxyServiceClient() nodes.ProxyServiceClient {
	return nodes.NewProxyServiceClient(s.NodeConn)
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

func (s *NodeService) RrpcServiceClient() rgrpc.RgrpcServiceClient {
	return rgrpc.NewRgrpcServiceClient(s.NodeConn)
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

func (s *NodeService) linkRgrpc(ctx context.Context) error {
	ctx = metadata.SetToken(ctx, s.GetToken())

	stream, err := s.RrpcServiceClient().OpenRgrpc(ctx)
	if err != nil {
		return err
	}
	defer stream.CloseSend()

	s.es.Logger().Sugar().Info("link rgrpc success")

	return rgrpc.Serve(stream, s.handlerMap, func() bool { return false })
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

func (s *NodeService) waitRemoteDeviceUpdated() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	for {
		ctx := metadata.SetToken(s.ctx, s.GetToken())

		stream, err := s.SyncServiceClient().WaitDeviceUpdated(ctx, &pb.MyEmpty{})
		if err != nil {
			if code, ok := status.FromError(err); ok {
				if code.Code() == codes.Canceled {
					return
				}
			}

			s.es.Logger().Sugar().Errorf("WaitDeviceUpdated: %v", err)
			return
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

			err = s.sync1(s.ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("sync1: %v", err)
			}
		}
	}
}

func (s *NodeService) waitLocalDeviceUpdated() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	notify := s.es.GetSync().Notify(NOTIFY)
	defer notify.Close()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-notify.Wait():
			err := s.sync2(s.ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("sync2: %v", err)
			}
		}
	}
}

func (s *NodeService) waitRemoteTagValueUpdated() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	for {
		ctx := metadata.SetToken(s.ctx, s.GetToken())

		stream, err := s.SyncServiceClient().WaitTagValueUpdated(ctx, &pb.MyEmpty{})
		if err != nil {
			if code, ok := status.FromError(err); ok {
				if code.Code() == codes.Canceled {
					return
				}
			}

			s.es.Logger().Sugar().Errorf("WaitTagValueUpdated: %v", err)
			return
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

			err = s.syncTagValue1(s.ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("syncTagValue1: %v", err)
			}
		}
	}
}

func (s *NodeService) waitLocalTagValueUpdated() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	notify := s.es.GetSync().Notify(NOTIFY_TVAL)
	defer notify.Close()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-notify.Wait():
			err := s.syncTagValue2(s.ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("syncTagValue2: %v", err)
			}
		}
	}
}

func (s *NodeService) sync(ctx context.Context) error {
	if err := s.sync1(ctx); err != nil {
		return err
	}

	if err := s.sync2(ctx); err != nil {
		return err
	}

	if err := s.syncTagValue1(ctx); err != nil {
		return err
	}

	return s.syncTagValue2(ctx)
}

func (s *NodeService) sync1(ctx context.Context) error {
	ctx = metadata.SetToken(ctx, s.GetToken())

	return s.syncRemoteToLocal(ctx)
}

func (s *NodeService) sync2(ctx context.Context) error {
	ctx = metadata.SetToken(ctx, s.GetToken())

	return s.syncLocalToRemote(ctx)
}

func (s *NodeService) syncTagValue1(ctx context.Context) error {
	ctx = metadata.SetToken(ctx, s.GetToken())

	return s.syncTagValueRemoteToLocal(ctx)
}

func (s *NodeService) syncTagValue2(ctx context.Context) error {
	ctx = metadata.SetToken(ctx, s.GetToken())

	return s.syncTagValueLocalToRemote(ctx)
}
