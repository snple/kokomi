package edge

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/danclive/nson-go"
	"github.com/quic-go/quic-go"
	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/util"
	"go.uber.org/zap"
)

type TunnelService struct {
	es *EdgeService

	listens map[string]*tunnelListener
	lock    sync.RWMutex

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup
}

func newTunnelService(es *EdgeService) *TunnelService {
	ctx, cancel := context.WithCancel(es.Context())

	return &TunnelService{
		listens: make(map[string]*tunnelListener),
		es:      es,
		ctx:     ctx,
		cancel:  cancel,
	}
}

func (s *TunnelService) Start() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	go s.waitDeviceUpdated()

	ticker := time.NewTicker(60 * 3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			err := s.ticker()
			if err != nil {
				s.es.Logger().Sugar().Errorf("TunnelService ticker: %v", err)
			}
		}
	}
}

func (s *TunnelService) Stop() {
	s.cancel()
	s.closeWG.Wait()
}

func (s *TunnelService) waitDeviceUpdated() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	updated := int64(0)

	for {
		output := s.es.GetSync().WaitDeviceUpdated2(s.ctx)

		<-output
		err := s.syncProxy(s.ctx, &updated)
		if err != nil {
			s.es.Logger().Sugar().Errorf("device sync: %v", err)
		}

		ok := <-output
		if ok {
			err := s.syncProxy(s.ctx, &updated)
			if err != nil {
				s.es.Logger().Sugar().Errorf("device sync: %v", err)
			}
		} else {
			return
		}

		time.Sleep(time.Second)
	}
}

func (s *TunnelService) syncProxy(ctx context.Context, updated *int64) error {
	proxyUpdated, err := s.es.GetSync().GetProxyUpdated(ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	if proxyUpdated.GetUpdated() <= *updated {
		return nil
	}

	{
		after := *updated
		limit := uint32(10)

		for {
			remotes, err := s.es.GetProxy().Pull(ctx, &edges.PullProxyRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetProxy() {
				after = remote.GetUpdated()

				if remote.GetDeleted() > 0 {
					s.lock.RLock()
					listen, ok := s.listens[remote.GetId()]
					s.lock.RUnlock()

					if ok {
						listen.stop()

						s.lock.Lock()
						delete(s.listens, remote.GetId())
						s.lock.Unlock()
					}
				} else {
					s.checkProxy(remote)
				}
			}

			if len(remotes.GetProxy()) < int(limit) {
				break
			}
		}
	}

	*updated = proxyUpdated.GetUpdated()

	return nil
}

func (s *TunnelService) ticker() error {
	request := edges.ListProxyRequest{
		Page: &pb.Page{
			Limit: 1000,
		},
	}

	reply, err := s.es.GetProxy().List(s.ctx, &request)
	if err != nil {
		return err
	}

	for _, proxy := range reply.GetProxy() {
		s.checkProxy(proxy)
	}

	return nil
}

func (s *TunnelService) openStreamSync() (quic.Stream, error) {
	if option := s.es.GetQuic(); option.IsSome() {
		return option.Unwrap().OpenStreamSync()
	}

	return nil, errors.New("quic service not enable")
}

func (s *TunnelService) checkProxy(proxy *pb.Proxy) {
	{
		s.lock.RLock()
		listen, ok := s.listens[proxy.GetId()]
		s.lock.RUnlock()

		if ok {
			index1 := genProxyIndex(proxy)
			index2 := genProxyIndex(listen.proxy)
			if index1 == index2 {
				return
			}

			listen.stop()

			s.lock.Lock()
			delete(s.listens, proxy.GetId())
			s.lock.Unlock()
		}

		if proxy.GetNetwork() == "" || proxy.GetTarget() == "" || proxy.GetStatus() != consts.ON {
			return
		}
	}

	listen, err := tunnelListen(s, proxy)
	if err != nil {
		s.es.Logger().Sugar().Errorf("tunnelListen error: %v", err)
		return
	}

	go func() {
		s.closeWG.Add(1)
		defer s.closeWG.Done()

		s.lock.Lock()
		s.listens[proxy.GetId()] = listen
		s.lock.Unlock()

		listen.accept()

		s.lock.Lock()
		delete(s.listens, proxy.GetId())
		s.lock.Unlock()

		listen.closeWG.Wait()
	}()
}

type tunnelListener struct {
	ts    *TunnelService
	proxy *pb.Proxy
	nl    net.Listener
	conns map[net.Conn]struct{}
	lock  sync.RWMutex

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup
}

func tunnelListen(ts *TunnelService, proxy *pb.Proxy) (*tunnelListener, error) {
	// The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket".
	nl, err := net.Listen(proxy.Network, proxy.Address)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(ts.ctx)

	tl := &tunnelListener{
		ts:     ts,
		proxy:  proxy,
		nl:     nl,
		conns:  make(map[net.Conn]struct{}),
		ctx:    ctx,
		cancel: cancel,
	}

	go func() {
		tl.closeWG.Add(1)
		defer tl.closeWG.Done()

		<-ctx.Done()
		tl.nl.Close()
	}()
	go tl.syncLinkStatus()

	return tl, nil
}

func (tl *tunnelListener) stop() {
	tl.cancel()
	tl.closeWG.Wait()
}

func (tl *tunnelListener) accept() {
	tl.closeWG.Add(1)
	defer tl.closeWG.Done()

	defer func() {
		tl.lock.Lock()
		for conn := range tl.conns {
			conn.Close()
		}
		tl.lock.Unlock()

		tl.cancel()
	}()

	for {
		conn, err := tl.nl.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return
			}

			tl.logger().Sugar().Errorf("net.Listener accept error: %v", err)
			return
		}

		go func() {
			err := tl.handleConn(conn)
			if err != nil {
				tl.logger().Sugar().Errorf("tunnelListener.handleConn(conn) error: %v", err)
			}
		}()
	}
}

func (tl *tunnelListener) handleConn(conn net.Conn) error {
	defer conn.Close()

	stream, err := tl.ts.openStreamSync()
	if err != nil {
		return err
	}
	defer stream.Close()

	err = tl.openProxy(stream)
	if err != nil {
		return err
	}

	tl.lock.Lock()
	tl.conns[conn] = struct{}{}
	tl.lock.Unlock()

	errChan := make(chan error)
	go quicStreamCopy1(stream, conn, errChan)
	go quicStreamCopy2(conn, stream, errChan)

	err = <-errChan
	if err != nil {
		tl.logger().Sugar().Errorf("stream.Copy error: %v", err)
	}

	<-errChan

	tl.lock.Lock()
	delete(tl.conns, conn)
	tl.lock.Unlock()

	return err
}

func (tl *tunnelListener) openProxy(stream quic.Stream) error {
	{
		wmessage := nson.Message{
			"method": nson.String("proxy"),
			"proxy":  nson.String(tl.proxy.GetId()),
		}

		err := util.WriteNsonMessage(stream, wmessage)
		if err != nil {
			return err
		}
	}

	err := stream.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))
	if err != nil {
		return err
	}

	{
		rmessage, err := util.ReadNsonMessage(stream)
		if err != nil {
			return err
		}

		code, err := rmessage.GetI32("code")
		if err != nil {
			return err
		}

		if code != 0 {
			return fmt.Errorf("rmessage: error code %v", code)
		}
	}

	err = stream.SetReadDeadline(time.Time{})
	if err != nil {
		return nil
	}

	return nil
}

func (tl *tunnelListener) syncLinkStatus() {
	ticker := time.NewTicker(tl.ts.es.dopts.syncLinkStatus)
	defer ticker.Stop()

	defer func() {
		request := edges.LinkProxyRequest{Id: tl.proxy.GetId(), Status: consts.OFF}

		ctx := context.Background()
		tl.ts.es.GetProxy().Link(ctx, &request)
	}()

	for {
		select {
		case <-tl.ctx.Done():
			return
		case <-ticker.C:
			{
				tl.lock.RLock()
				n := len(tl.conns)
				tl.lock.RUnlock()

				request := edges.LinkProxyRequest{Id: tl.proxy.GetId(), Status: int32(n)}

				tl.ts.es.GetProxy().Link(tl.ctx, &request)
			}
		}
	}
}

func (tl *tunnelListener) logger() *zap.Logger {
	return tl.ts.es.Logger()
}

func genProxyIndex(proxy *pb.Proxy) string {
	return fmt.Sprintf("%v-%v-%v-%v-%v",
		proxy.GetName(), proxy.GetNetwork(), proxy.GetAddress(), proxy.GetTarget(), proxy.GetStatus())
}
