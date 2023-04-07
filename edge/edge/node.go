package edge

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/edge/model"
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

	token   string
	optLock sync.RWMutex

	handlerMap rgrpc.HandlerMap

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup
}

func newNodeService(es *EdgeService) (*NodeService, error) {
	var err error

	es.Logger().Sugar().Infof("link node service: %v", es.dopts.nodeOptions.Addr)

	nodeConn, err := grpc.Dial(es.dopts.nodeOptions.Addr, es.dopts.nodeOptions.Options...)
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

func (s *NodeService) Start() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	s.es.Logger().Sugar().Info("start node service")

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

func (s *NodeService) Stop() {
	s.cancel()
	s.closeWG.Wait()
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

	s.linkDevice()
	s.es.GetStatus().SetDeviceLink(consts.ON)
	defer s.es.GetStatus().SetDeviceLink(consts.OFF)

	go s.waitDeviceUpdated()
	go s.waitDeviceUpdatedLocal()
	go s.waitTagValueUpdated()
	go s.waitWireValueUpdated()
	go s.waitWireValueUpdatedLocal()

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

	tokenRefreshTicker := time.NewTicker(s.es.dopts.tokenRefresh)
	defer tokenRefreshTicker.Stop()

	syncLinkStatusTicker := time.NewTicker(s.es.dopts.syncLinkStatus)
	defer syncLinkStatusTicker.Stop()

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
		case <-syncLinkStatusTicker.C:
			if option := s.es.GetQuic(); option.IsNone() {
				if s.es.GetStatus().GetDeviceLink() == consts.ON {
					err := s.linkDevice()
					if err != nil {
						s.es.Logger().Sugar().Errorf("link device : %v", err)
					} else {
						s.es.Logger().Sugar().Info("link device ticker success")
					}
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

func (s *NodeService) OptionServiceClient() nodes.OptionServiceClient {
	return nodes.NewOptionServiceClient(s.NodeConn)
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

func (s *NodeService) VarServiceClient() nodes.VarServiceClient {
	return nodes.NewVarServiceClient(s.NodeConn)
}

func (s *NodeService) CableServiceClient() nodes.CableServiceClient {
	return nodes.NewCableServiceClient(s.NodeConn)
}

func (s *NodeService) WireServiceClient() nodes.WireServiceClient {
	return nodes.NewWireServiceClient(s.NodeConn)
}

func (s *NodeService) DataServiceClient() nodes.DataServiceClient {
	return nodes.NewDataServiceClient(s.NodeConn)
}

func (s *NodeService) RrpcServiceClient() rgrpc.RgrpcServiceClient {
	return rgrpc.NewRgrpcServiceClient(s.NodeConn)
}

func (s *NodeService) GetToken() string {
	s.optLock.RLock()
	defer s.optLock.RUnlock()
	return s.token
}

func (s *NodeService) setToken(token string) {
	s.optLock.RLock()
	defer s.optLock.RUnlock()
	s.token = token
}

func (s *NodeService) linkDevice() error {
	ctx := metadata.SetToken(context.Background(), s.GetToken())

	request := &nodes.LinkDeviceRequest{Status: consts.ON}
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

	option_device_id := s.es.dopts.deviceID

	{
		has, err := s.es.GetOption().Has(ctx, model.OPTION_DEVICE_ID)
		if err != nil {
			return err
		}

		if has {
			option_device_id, err = s.es.GetOption().Get(ctx, model.OPTION_DEVICE_ID)
			if err != nil {
				return err
			}
		}
	}

	option_secret := s.es.dopts.secret

	{
		has, err := s.es.GetOption().Has(ctx, model.OPTION_SECRET)
		if err != nil {
			return err
		}

		if has {
			option_secret, err = s.es.GetOption().Get(ctx, model.OPTION_SECRET)
			if err != nil {
				return err
			}
		}
	}

	request := &nodes.LoginDeviceRequest{
		Id:     option_device_id,
		Secret: option_secret,
	}

	// try login
	reply, err := s.DeviceServiceClient().Login(ctx, request)
	if err != nil {
		return err
	}

	if len(reply.GetToken()) == 0 {
		return errors.New("login: reply token is empty")
	}

	ctx = metadata.SetToken(ctx, reply.GetToken())

	// query remote database
	remoteDevice, err := s.DeviceServiceClient().View(ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	// query local database and insert
	_, err = s.es.GetDevice().View(ctx, &pb.MyEmpty{})
	if err != nil {
		if code, ok := status.FromError(err); ok {
			if code.Code() == codes.NotFound {
				ctx = metadata.SetSync(ctx)
				_, err = s.es.GetDevice().Create(ctx, remoteDevice)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			return err
		}
	}

	// set token
	s.setToken(reply.GetToken())

	return nil
}

func (s *NodeService) waitDeviceUpdated() {
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

			err = s.sync(s.ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("device sync: %v", err)
			}
		}

		time.Sleep(time.Second)
	}
}

func (s *NodeService) waitDeviceUpdatedLocal() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	for {
		output := s.es.GetSync().WaitDeviceUpdated2(s.ctx)

		<-output
		err := s.syncLocal(s.ctx)
		if err != nil {
			s.es.Logger().Sugar().Errorf("device sync: %v", err)
		}

		ok := <-output
		if ok {
			err := s.syncLocal(s.ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("device sync: %v", err)
			}
		} else {
			return
		}

		time.Sleep(time.Second)
	}
}

func (s *NodeService) waitTagValueUpdated() {
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

			err = s.syncTagValue()
			if err != nil {
				s.es.Logger().Sugar().Errorf("device sync tag value: %v", err)
			}
		}

		time.Sleep(time.Second)
	}
}

func (s *NodeService) waitWireValueUpdated() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	for {
		ctx := metadata.SetToken(s.ctx, s.GetToken())

		stream, err := s.SyncServiceClient().WaitWireValueUpdated(ctx, &pb.MyEmpty{})
		if err != nil {
			if code, ok := status.FromError(err); ok {
				if code.Code() == codes.Canceled {
					return
				}
			}

			s.es.Logger().Sugar().Errorf("WaitWireValueUpdated: %v", err)
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

				s.es.Logger().Sugar().Errorf("WaitWireValueUpdated.Recv(): %v", err)
				return
			}

			err = s.syncWireValue(s.ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("device sync wire value: %v", err)
			}
		}

		time.Sleep(time.Second)
	}
}

func (s *NodeService) waitWireValueUpdatedLocal() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	for {
		output := s.es.GetSync().WaitWireValueUpdated2(s.ctx)

		<-output
		err := s.syncWireValueLocal(s.ctx)
		if err != nil {
			s.es.Logger().Sugar().Errorf("device sync: %v", err)
		}

		ok := <-output
		if ok {
			err := s.syncWireValueLocal(s.ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("device sync: %v", err)
			}
		} else {
			return
		}

		time.Sleep(time.Second)
	}
}

func (s *NodeService) sync(ctx context.Context) error {
	ctx = metadata.SetToken(ctx, s.GetToken())
	ctx = metadata.SetSync(ctx)

	return s.syncRemoteToLocal(ctx)
}

func (s *NodeService) syncLocal(ctx context.Context) error {
	ctx = metadata.SetToken(ctx, s.GetToken())
	ctx = metadata.SetSync(ctx)

	return s.syncLocalToRemote(ctx)
}

func (s *NodeService) syncTagValue() error {
	var err error

	ctx := context.Background()
	ctx = metadata.SetToken(ctx, s.GetToken())

	err = s.syncRemoteToLocalTagValue(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *NodeService) syncWireValue(ctx context.Context) error {
	ctx = metadata.SetToken(ctx, s.GetToken())
	ctx = metadata.SetSync(ctx)

	return s.syncRemoteToLocalWireValue(ctx)
}

func (s *NodeService) syncWireValueLocal(ctx context.Context) error {
	ctx = metadata.SetToken(ctx, s.GetToken())
	ctx = metadata.SetSync(ctx)

	return s.syncLocalToRemoteWireValue(ctx)
}
