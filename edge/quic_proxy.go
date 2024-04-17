package edge

import (
	"context"
	"errors"
	"fmt"
	"io"
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

type QuicProxyService struct {
	es *EdgeService

	listens map[string]*proxyListener
	updated int64

	portLinkCount map[string]int

	lock sync.RWMutex

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup
}

func newQuicProxyService(es *EdgeService) *QuicProxyService {
	ctx, cancel := context.WithCancel(es.Context())

	return &QuicProxyService{
		listens:       make(map[string]*proxyListener),
		es:            es,
		ctx:           ctx,
		cancel:        cancel,
		portLinkCount: make(map[string]int),
	}
}

func (s *QuicProxyService) start() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	if option := s.es.GetQuic(); option.IsNone() {
		return
	}

	s.es.Logger().Info("start quic proxy service")

	go s.waitDeviceUpdated()
	go s.syncLinkStatus()

	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	err := s.ticker()
	if err != nil {
		s.es.Logger().Sugar().Errorf("QuicProxyService ticker: %v", err)
	}

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			err := s.ticker()
			if err != nil {
				s.es.Logger().Sugar().Errorf("QuicProxyService ticker: %v", err)
			}
		}
	}
}

func (s *QuicProxyService) stop() {
	s.cancel()
	s.closeWG.Wait()
}

func (s *QuicProxyService) ticker() error {
	request := edges.ProxyListRequest{
		Page: &pb.Page{
			Limit: 1000,
		},
	}

	reply, err := s.es.GetProxy().List(s.ctx, &request)
	if err != nil {
		return err
	}

	for _, proxy := range reply.GetProxy() {
		s.startProxy(proxy)
	}

	return nil
}

func (s *QuicProxyService) waitDeviceUpdated() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	notify := s.es.GetSync().Notify(NOTIFY)
	defer notify.Close()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-notify.Wait():
			err := s.checkProxyUpdated()
			if err != nil {
				s.es.Logger().Sugar().Errorf("quic proxy checkProxyUpdated: %v", err)
			}
		}
	}
}

func (s *QuicProxyService) checkProxyUpdated() error {
	proxyUpdated, err := s.es.GetSync().GetProxyUpdated(s.ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	updated := s.getUpdated()

	if proxyUpdated.GetUpdated() <= updated {
		return nil
	}

	{
		after := updated
		limit := uint32(10)

		for {
			remotes, err := s.es.GetProxy().Pull(s.ctx, &edges.ProxyPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetProxy() {
				after = remote.GetUpdated()

				if remote.GetDeleted() > 0 {
					s.lock.Lock()
					if listen, ok := s.listens[remote.GetId()]; ok {
						listen.stop()
					}
					s.lock.Unlock()
				} else {
					s.checkProxy(remote)
				}
			}

			if len(remotes.GetProxy()) < int(limit) {
				break
			}
		}
	}

	s.setUpdated(proxyUpdated.GetUpdated())

	return nil
}

func (s *QuicProxyService) getUpdated() int64 {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.updated
}

func (s *QuicProxyService) setUpdated(updated int64) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.updated = updated
}

func (s *QuicProxyService) openStreamSync() (quic.Stream, error) {
	if option := s.es.GetQuic(); option.IsSome() {
		return option.Unwrap().OpenStreamSync()
	}

	return nil, errors.New("quic service not enable")
}

func (s *QuicProxyService) startProxy(proxy *pb.Proxy) {
	if proxy.GetNetwork() == "" || proxy.GetAddress() == "" ||
		proxy.GetTarget() == "" || proxy.GetStatus() != consts.ON {
		return
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.listens[proxy.GetId()]; ok {
		return
	}

	listen, err := newListen(s, proxy)
	if err != nil {
		s.es.Logger().Sugar().Errorf("newListen error: %v", err)
		return
	}

	s.listens[proxy.GetId()] = listen

	go func() {
		s.closeWG.Add(1)
		defer s.closeWG.Done()

		listen.run()

		s.lock.Lock()
		delete(s.listens, proxy.GetId())
		s.lock.Unlock()
	}()
}

func (s *QuicProxyService) checkProxy(proxy *pb.Proxy) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if listen, ok := s.listens[proxy.GetId()]; ok {
		if proxy.GetStatus() == consts.ON &&
			proxy.GetName() == listen.proxy.GetName() &&
			proxy.GetNetwork() == listen.proxy.GetNetwork() &&
			proxy.GetAddress() == listen.proxy.GetAddress() &&
			proxy.GetTarget() == listen.proxy.GetTarget() {
			return
		}

		listen.stop()
	}
}

func (s *QuicProxyService) handleProxy(request nson.Map, stream quic.Stream) error {
	portId, conn, err := s.openPort(request)
	if err != nil {
		response := nson.Map{
			"code": nson.I32(400),
		}
		util.WriteNsonMessage(stream, response)
		return err
	}

	defer conn.Close()

	response := nson.Map{
		"code": nson.I32(0),
	}
	err = util.WriteNsonMessage(stream, response)
	if err != nil {
		return err
	}

	s.incPortLinkCount(portId)

	errChan := make(chan error)
	go quicStreamCopy1(stream, conn, errChan)
	go quicStreamCopy2(conn, stream, errChan)

	err = <-errChan
	if err != nil {
		s.es.Logger().Sugar().Errorf("stream.Copy error: %v", err)
	}

	<-errChan

	s.decPortLinkCount(portId)

	return err
}

func (s *QuicProxyService) openPort(request nson.Map) (string, net.Conn, error) {
	method, err := request.GetString("method")
	if err != nil {
		return "", nil, err
	}

	portId, err := request.GetString("port")
	if err != nil {
		return "", nil, err
	}

	if method != "proxy" || portId == "" {
		return "", nil, errors.New(`method != "proxy" || portId == ""`)
	}

	port, err := s.es.GetPort().View(s.ctx, &pb.Id{Id: portId})
	if err != nil {
		return "", nil, fmt.Errorf("view port %v, err: %v", portId, err)
	}

	if port.GetStatus() != consts.ON {
		return port.GetId(), nil, errors.New("port.status != consts.ON")
	}

	if port.GetNetwork() == "" || port.GetAddress() == "" {
		return port.GetId(), nil, errors.New("port.network == '' || port.address == ''")
	}

	conn, err := net.DialTimeout(port.GetNetwork(), port.GetAddress(), time.Second*5)
	if err != nil {
		return port.GetId(), nil, err
	}

	return port.GetId(), conn, nil
}

func (s *QuicProxyService) syncLinkStatus() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	ticker := time.NewTicker(time.Duration(60) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			{
				s.lock.RLock()
				for portId, num := range s.portLinkCount {
					go func(portId string, num int) {
						request := edges.PortLinkRequest{Id: portId, Status: int32(num)}

						s.es.GetPort().Link(s.ctx, &request)
					}(portId, num)
				}
				s.lock.RUnlock()
			}
		}
	}
}

func (s *QuicProxyService) incPortLinkCount(portId string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if num, ok := s.portLinkCount[portId]; ok {
		s.portLinkCount[portId] = num + 1
	} else {
		s.portLinkCount[portId] = 1
	}
}

func (s *QuicProxyService) decPortLinkCount(portId string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if num, ok := s.portLinkCount[portId]; ok {
		if num-1 == 0 {
			delete(s.portLinkCount, portId)
		} else {
			s.portLinkCount[portId] = num - 1
		}
	}
}

type proxyListener struct {
	ts    *QuicProxyService
	proxy *pb.Proxy
	nl    net.Listener
	conns map[net.Conn]struct{}
	lock  sync.RWMutex

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup
}

func newListen(ts *QuicProxyService, proxy *pb.Proxy) (*proxyListener, error) {
	// The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket".
	nl, err := net.Listen(proxy.Network, proxy.Address)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(ts.ctx)

	pl := &proxyListener{
		ts:     ts,
		proxy:  proxy,
		nl:     nl,
		conns:  make(map[net.Conn]struct{}),
		ctx:    ctx,
		cancel: cancel,
	}

	go func() {
		pl.closeWG.Add(1)
		defer pl.closeWG.Done()

		<-ctx.Done()
		pl.nl.Close()
	}()
	go pl.syncLinkStatus()

	return pl, nil
}

func (pl *proxyListener) stop() {
	pl.cancel()
	pl.closeWG.Wait()
}

func (pl *proxyListener) run() {
	defer pl.closeWG.Wait()

	defer func() {
		pl.lock.Lock()
		for conn := range pl.conns {
			conn.Close()
		}
		pl.lock.Unlock()

		pl.cancel()
	}()

	for {
		conn, err := pl.nl.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return
			}

			pl.logger().Sugar().Errorf("net.Listener accept error: %v", err)
			return
		}

		go func() {
			err := pl.handleConn(conn)
			if err != nil {
				pl.logger().Sugar().Errorf("proxyListener.handleConn(conn) error: %v", err)
			}
		}()
	}
}

func (pl *proxyListener) handleConn(conn net.Conn) error {
	defer conn.Close()

	stream, err := pl.ts.openStreamSync()
	if err != nil {
		return err
	}
	defer stream.Close()

	err = pl.openProxy(stream)
	if err != nil {
		return err
	}

	pl.lock.Lock()
	pl.conns[conn] = struct{}{}
	pl.lock.Unlock()

	errChan := make(chan error)
	go quicStreamCopy1(stream, conn, errChan)
	go quicStreamCopy2(conn, stream, errChan)

	err = <-errChan
	if err != nil {
		pl.logger().Sugar().Errorf("stream.Copy error: %v", err)
	}

	<-errChan

	pl.lock.Lock()
	delete(pl.conns, conn)
	pl.lock.Unlock()

	return err
}

func (pl *proxyListener) openProxy(stream quic.Stream) error {
	{
		request := nson.Map{
			"method": nson.String("proxy"),
			"proxy":  nson.String(pl.proxy.GetId()),
		}

		err := util.WriteNsonMessage(stream, request)
		if err != nil {
			return err
		}
	}

	err := stream.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))
	if err != nil {
		return err
	}

	{
		response, err := util.ReadNsonMessage(stream)
		if err != nil {
			return err
		}

		code, err := response.GetI32("code")
		if err != nil {
			return err
		}

		if code != 0 {
			return fmt.Errorf("response: error code %v", code)
		}
	}

	err = stream.SetReadDeadline(time.Time{})
	if err != nil {
		return nil
	}

	return nil
}

func (pl *proxyListener) syncLinkStatus() {
	pl.closeWG.Add(1)
	defer pl.closeWG.Done()

	ticker := time.NewTicker(pl.ts.es.dopts.SyncOptions.Link)
	defer ticker.Stop()

	defer func() {
		request := edges.ProxyLinkRequest{Id: pl.proxy.GetId(), Status: consts.OFF}

		ctx := context.Background()
		pl.ts.es.GetProxy().Link(ctx, &request)
	}()

	for {
		select {
		case <-pl.ctx.Done():
			return
		case <-ticker.C:
			{
				pl.lock.RLock()
				n := len(pl.conns)
				pl.lock.RUnlock()

				if pl.ts.es.GetStatus().GetDeviceLink() == consts.ON {
					request := edges.ProxyLinkRequest{Id: pl.proxy.GetId(), Status: int32(n)}

					_, err := pl.ts.es.GetProxy().Link(pl.ctx, &request)
					if err != nil {
						pl.logger().Sugar().Errorf("proxy link error: %v", err)
					}
				}
			}
		}
	}
}

func (pl *proxyListener) logger() *zap.Logger {
	return pl.ts.es.Logger()
}

func quicStreamCopy1(dst quic.Stream, src net.Conn, errCh chan error) {
	_, err := io.Copy(dst, src)
	errCh <- err

	dst.CancelWrite(1)
}

func quicStreamCopy2(dst net.Conn, src quic.Stream, errCh chan error) {
	_, err := io.Copy(dst, src)
	errCh <- err

	if tc, ok := dst.(*net.TCPConn); ok {
		tc.CloseWrite()
	} else if tc, ok := dst.(*net.UnixConn); ok {
		tc.CloseWrite()
	}
}
