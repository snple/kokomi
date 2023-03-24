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
	"snple.com/kokomi/consts"
	"snple.com/kokomi/pb"
	"snple.com/kokomi/util"
	"snple.com/kokomi/util/token"
)

type QuicService struct {
	ns *NodeService

	listener quic.Listener

	conns   map[string]quic.Connection
	optLock sync.RWMutex
}

func newQuicService(ns *NodeService) (*QuicService, error) {
	listener, err := quic.ListenAddr(ns.dopts.quicOptions.Addr,
		ns.dopts.quicOptions.TLSConfig, ns.dopts.quicOptions.QUICConfig)
	if err != nil {
		return nil, err
	}

	s := &QuicService{
		ns:       ns,
		listener: listener,
		conns:    make(map[string]quic.Connection),
	}

	return s, nil
}

func (s *QuicService) Start() {
	s.ns.Logger().Info("start quic server")

	err := s.acceptConn()
	if err != nil {
		if !errors.Is(err, quic.ErrServerClosed) {
			s.ns.Logger().Sugar().Errorf("QuicService.acceptConn error: %v", err)
		}
	}
}

func (s *QuicService) Stop() {
	s.listener.Close()
}

func (s *QuicService) GetConn(deviceId string) (quic.Connection, bool) {
	s.optLock.RLock()
	defer s.optLock.RUnlock()

	conn, has := s.conns[deviceId]

	return conn, has
}

func (s *QuicService) acceptConn() error {
	for {
		conn, err := s.listener.Accept(s.ns.ctx)
		if err != nil {
			return err
		}

		go func() {
			err := s.handleConn(conn)
			if err != nil {
				s.ns.Logger().Sugar().Errorf("QuicService.handleConn(conn) error: %v", err)
			}
		}()
	}
}

func (s *QuicService) handleConn(conn quic.Connection) error {
	defer conn.CloseWithError(1, "break")

	deviceId, err := s.handshake(conn)
	if err != nil {
		return err
	}

	s.optLock.Lock()
	s.conns[deviceId] = conn
	s.optLock.Unlock()

	go s.accept(conn, deviceId)

	context := conn.Context()
	<-context.Done()
	err = context.Err()

	s.optLock.Lock()
	delete(s.conns, deviceId)
	s.optLock.Unlock()

	return err
}

func (s *QuicService) handshake(conn quic.Connection) (string, error) {
	stream, err := conn.AcceptStream(context.Background())
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

	device, err := s.ns.cs.GetDevice().View(context.Background(), &pb.Id{Id: deviceId})
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

	rttKey := deviceId + "_rtt"

	handle := func() error {
		wmessage := nson.Message{
			"method": nson.String("ping"),
			"t":      nson.I64(time.Now().UnixMilli()),
		}

		err := util.WriteNsonMessage(stream, wmessage)
		if err != nil {
			return err
		}

		err = stream.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))
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
		s.ns.cs.GetStatus().SetLink(rttKey, int32(rtt))

		s.ns.Logger().Sugar().Debugf("quic deviceId: %v, rtt %v", deviceId, rtt)

		return nil
	}

	for {
		err := handle()
		if err != nil {
			s.ns.Logger().Sugar().Errorf("ping error: %v", err)
			return
		}

		time.Sleep(s.ns.dopts.quicPingInterval)
	}
}

func (s *QuicService) accept(conn quic.Connection, deviceId string) error {
	for {
		stream, err := conn.AcceptStream(context.Background())
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

	proxy, err := s.ns.cs.GetProxy().View(context.Background(), &pb.Id{Id: proxyId})
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

	port, err := s.ns.cs.GetPort().View(context.Background(), &pb.Id{Id: proxy.Target})
	if err != nil {
		return nil, err
	}

	if port.GetStatus() != consts.ON {
		return nil, errors.New("port.status != consts.ON")
	}

	if port.GetNetwork() == "" || port.GetAddress() == "" {
		return nil, errors.New("port.network == '' || port.address == ''")
	}

	conn2, ok := s.GetConn(port.GetDeviceId())
	if ok == false {
		return nil, errors.New("target not link")
	}

	stream2, err := conn2.OpenStreamSync(context.Background())
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

func (s *QuicService) streamCopy(dst, src quic.Stream, errCh chan error) {
	_, err := io.Copy(dst, src)
	errCh <- err

	dst.CancelWrite(1)
}
