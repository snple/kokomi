package node

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/danclive/nson-go"
	"github.com/quic-go/quic-go"
	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/util"
	"github.com/snple/kokomi/util/token"
	"github.com/snple/types"
)

type QuicService struct {
	ns *NodeService

	listener *quic.Listener

	conns map[string]*quicChannels
	lock  sync.RWMutex

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup
}

func newQuicService(ns *NodeService) (*QuicService, error) {
	listener, err := quic.ListenAddr(ns.dopts.QuicOptions.Addr,
		ns.dopts.QuicOptions.TLSConfig, ns.dopts.QuicOptions.QUICConfig)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(ns.Context())

	s := &QuicService{
		ns:       ns,
		listener: listener,
		conns:    make(map[string]*quicChannels),
		ctx:      ctx,
		cancel:   cancel,
	}

	return s, nil
}

func (s *QuicService) Start() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	s.ns.Logger().Sugar().Infof("node quic start: %v", s.ns.dopts.QuicOptions.Addr)

	err := s.acceptConn()
	if err != nil {
		if !errors.Is(err, quic.ErrServerClosed) {
			s.ns.Logger().Sugar().Errorf("QuicService.acceptConn error: %v", err)
		}
	}
}

func (s *QuicService) Stop() {
	s.listener.Close()
	s.cancel()
	s.closeWG.Wait()
}

func (s *QuicService) addConn(deviceID string, conn quic.Connection) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if chans, ok := s.conns[deviceID]; ok {
		chans.add(conn)
	} else {
		chans := &quicChannels{}
		chans.add(conn)
		s.conns[deviceID] = chans
	}
}

func (s *QuicService) removeConn(deviceID string, conn quic.Connection) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if chans, ok := s.conns[deviceID]; ok {
		chans.remove(conn)
	}
}

func (s *QuicService) GetConn(deviceID string) types.Option[quic.Connection] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if chans, ok := s.conns[deviceID]; ok {
		return chans.pick()
	}

	return types.None[quic.Connection]()
}

func (s *QuicService) acceptConn() error {
	for {
		conn, err := s.listener.Accept(s.ctx)
		if err != nil {
			return err
		}

		go func() {
			err := s.handleConn(conn)
			if err != nil {
				if !errors.Is(err, context.Canceled) {
					s.ns.Logger().Sugar().Errorf("QuicService.handleConn(conn) error: %v", err)
				}
			}
		}()
	}
}

func (s *QuicService) handleConn(conn quic.Connection) error {
	defer conn.CloseWithError(1, "break")

	deviceID, err := s.handshake(conn)
	if err != nil {
		return err
	}

	s.addConn(deviceID, conn)

	go s.accept(conn, deviceID)

	ctx := conn.Context()
	<-ctx.Done()
	err = ctx.Err()

	s.removeConn(deviceID, conn)

	return err
}

func (s *QuicService) handshake(conn quic.Connection) (string, error) {
	stream, err := conn.AcceptStream(s.ctx)
	if err != nil {
		return "", err
	}

	wmessage := nson.Message{}

	deviceId, err := s.validate(stream)
	if err != nil {
		wmessage.Insert("code", nson.I32(400))
		util.WriteNsonMessage(stream, wmessage)
		return "", err
	}

	wmessage.Insert("code", nson.I32(0))

	err = util.WriteNsonMessage(stream, wmessage)
	if err != nil {
		return "", err
	}

	go s.ping(conn, stream, deviceId)

	return deviceId, nil
}

func (s *QuicService) validate(stream quic.Stream) (string, error) {
	err := stream.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))
	if err != nil {
		return "", err
	}

	rmessage, err := util.ReadNsonMessage(stream)
	if err != nil {
		return "", err
	}

	err = stream.SetReadDeadline(time.Time{})
	if err != nil {
		return "", err
	}

	method, err := rmessage.GetString("method")
	if err != nil {
		return "", fmt.Errorf("rmessage.GetString('method') error: %v", err)
	}

	if method != "handshake" {
		return "", fmt.Errorf("method != 'handshake'")
	}

	tks, err := rmessage.GetString("token")
	if err != nil {
		return "", fmt.Errorf("rmessage.GetString('token') error: %v", err)
	}

	ok, deviceId := token.ValidateDeviceToken(tks)
	if !ok {
		return "", errors.New("token validation failed")
	}

	device, err := s.ns.Core().GetDevice().View(s.ctx, &pb.Id{Id: deviceId})
	if err != nil {
		return "", fmt.Errorf("GetDevice().View() error: %v", err)
	}

	if device.GetStatus() != consts.ON {
		return "", fmt.Errorf("device.status != consts.ON")
	}

	return deviceId, nil
}

func (s *QuicService) ping(conn quic.Connection, stream quic.Stream, deviceId string) {
	defer conn.CloseWithError(1, "ping error")
	defer stream.Close()

	handle := func() error {
		wmessage := nson.Message{
			"method": nson.String("ping"),
			"t":      nson.I64(time.Now().UnixMilli()),
		}

		err := util.WriteNsonMessage(stream, wmessage)
		if err != nil {
			return err
		}

		err = stream.SetReadDeadline(time.Now().Add(time.Duration(20) * time.Second))
		if err != nil {
			return err
		}

		rmessage, err := util.ReadNsonMessage(stream)
		if err != nil {
			return err
		}

		err = stream.SetReadDeadline(time.Time{})
		if err != nil {
			return err
		}

		method, err := rmessage.GetString("method")
		if err != nil {
			return err
		}

		if method != "ping" {
			return errors.New("rtt method != 'ping'")
		}

		t, err := rmessage.GetI64("t")
		if err != nil {
			return err
		}

		rtt := time.Now().UnixMilli() - t
		if rtt == 0 {
			rtt = 1
		}
		s.ns.Core().GetStatus().SetLink(deviceId, int32(rtt))

		s.ns.Logger().Sugar().Debugf("quic deviceId: %v, rtt %v", deviceId, rtt)

		return nil
	}

	for {
		err := handle()
		if err != nil {
			s.ns.Logger().Sugar().Errorf("ping error: %v", err)
			return
		}

		time.Sleep(s.ns.dopts.Ping)
	}
}

func (s *QuicService) accept(conn quic.Connection, deviceId string) error {
	for {
		stream, err := conn.AcceptStream(s.ctx)
		if err != nil {
			return err
		}

		go func() {
			err := s.handleStream(stream, deviceId)
			if err != nil {
				s.ns.Logger().Sugar().Errorf("QuicService.handleStream(stream) error: %v", err)
			}
		}()
	}
}

func (s *QuicService) handleStream(stream quic.Stream, deviceId string) error {
	defer stream.Close()

	err := stream.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))
	if err != nil {
		return err
	}

	rmessage, err := util.ReadNsonMessage(stream)
	if err != nil {
		return err
	}

	err = stream.SetReadDeadline(time.Time{})
	if err != nil {
		return err
	}

	stream2, err := s.openStream(rmessage, deviceId)
	if err != nil {
		wmessage := nson.Message{}
		wmessage.Insert("code", nson.I32(400))
		util.WriteNsonMessage(stream, wmessage)
		return err
	}

	defer stream2.Close()

	errChan := make(chan error)
	go s.streamCopy(stream, stream2, errChan)
	go s.streamCopy(stream2, stream, errChan)

	err = <-errChan
	if err != nil {
		s.ns.Logger().Sugar().Errorf("stream.Copy error: %v", err)
	}

	<-errChan

	return err
}

func (s *QuicService) openStream(rmessage nson.Message, deviceId string) (quic.Stream, error) {
	method, err := rmessage.GetString("method")
	if err != nil {
		return nil, err
	}

	proxyId, err := rmessage.GetString("proxy")
	if err != nil {
		return nil, err
	}

	if method != "proxy" || proxyId == "" {
		return nil, errors.New(`method != "proxy" || proxyId == ""`)
	}

	proxy, err := s.ns.Core().GetProxy().View(s.ctx, &pb.Id{Id: proxyId})
	if err != nil {
		return nil, err
	}

	if proxy.GetDeviceId() != deviceId {
		return nil, errors.New("proxy.deviceId != deviceId'")
	}

	if proxy.GetStatus() != consts.ON {
		return nil, errors.New("proxy.status != consts.ON'")
	}

	if proxy.GetTarget() == "" {
		return nil, errors.New("proxy.target == ''")
	}

	port, err := s.ns.Core().GetPort().View(s.ctx, &pb.Id{Id: proxy.Target})
	if err != nil {
		return nil, err
	}

	if port.GetStatus() != consts.ON {
		return nil, errors.New("port.status != consts.ON")
	}

	if port.GetNetwork() == "" || port.GetAddress() == "" {
		return nil, errors.New("port.network == '' || port.address == ''")
	}

	if option := s.GetConn(port.GetDeviceId()); option.IsSome() {
		stream2, err := option.Unwrap().OpenStreamSync(s.ctx)
		if err != nil {
			return nil, err
		}

		wmessage := nson.Message{
			"method": nson.String("proxy"),
			"proxy":  nson.String(proxy.GetId()),
			"port":   nson.String(port.GetId()),
		}

		err = util.WriteNsonMessage(stream2, wmessage)
		if err != nil {
			stream2.Close()
			return nil, err
		}

		return stream2, nil
	}

	return nil, errors.New("target not link")
}

func (s *QuicService) streamCopy(dst, src quic.Stream, errCh chan error) {
	_, err := io.Copy(dst, src)
	errCh <- err

	dst.CancelWrite(1)
}

type quicChannels struct {
	cs []quic.Connection
}

func (c *quicChannels) add(conn quic.Connection) {
	c.cs = append(c.cs, conn)
}

func (c *quicChannels) remove(conn quic.Connection) {
	for i := range c.cs {
		if c.cs[i] == conn {
			c.cs = append(c.cs[:i], c.cs[i+1:]...)
			break
		}
	}
}

func (c *quicChannels) pick() types.Option[quic.Connection] {
	if c == nil {
		return types.None[quic.Connection]()
	}

	if len(c.cs) == 0 {
		return types.None[quic.Connection]()
	}

	return types.Some(c.cs[len(c.cs)-1])
}
