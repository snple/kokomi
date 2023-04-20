package edge

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/danclive/nson-go"
	"github.com/quic-go/quic-go"
	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/util"
	"github.com/snple/types"
)

type QuicService struct {
	es *EdgeService

	conn types.Option[quic.Connection]
	lock sync.RWMutex

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup

	linkNum map[string]int
}

func newQuicService(es *EdgeService) (*QuicService, error) {
	ctx, cancel := context.WithCancel(context.Background())

	s := &QuicService{
		es:      es,
		ctx:     ctx,
		cancel:  cancel,
		linkNum: make(map[string]int),
	}

	return s, nil
}

func (s *QuicService) start() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	s.es.Logger().Info("start quic service")

	go s.syncLinkStatus()

	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			s.loop()
		}
	}
}

func (s *QuicService) stop() {
	s.cancel()
	s.closeWG.Wait()
}

func (s *QuicService) loop() error {
	token := s.es.GetNode().GetToken()

	if token == "" {
		time.Sleep(time.Second)
		return nil
	}

	var conn quic.Connection

	{
		operation := func() error {
			var err error
			conn, err = s.connect(s.ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("quic connect: %v", err)
			}

			return err
		}

		err := backoff.Retry(operation, backoff.WithContext(backoff.NewExponentialBackOff(), s.ctx))
		if err != nil {
			s.es.Logger().Sugar().Errorf("backoff.Retry: %v", err)
			return err
		}
	}

	s.es.Logger().Sugar().Info("quic connect success")

	s.lock.Lock()
	s.conn = types.Some(conn)
	s.lock.Unlock()

	go s.accept(conn)

	context := conn.Context()
	<-context.Done()
	s.es.Logger().Sugar().Errorf("break conn error: %v", context.Err())

	s.lock.Lock()
	s.conn.Take()
	s.lock.Unlock()

	return nil
}

func (s *QuicService) connect(ctx context.Context) (quic.Connection, error) {
	conn, err := quic.DialAddrContext(ctx, s.es.dopts.quicOptions.Addr,
		s.es.dopts.quicOptions.TLSConfig, s.es.dopts.quicOptions.QUICConfig)
	if err != nil {
		return nil, err
	}

	err = s.handshake(conn)
	if err != nil {
		conn.CloseWithError(1, "handshake error")
		return nil, err
	}

	return conn, nil
}

func (s *QuicService) handshake(conn quic.Connection) error {
	stream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		return err
	}

	{
		wmessage := nson.Message{
			"method": nson.String("handshake"),
			"token":  nson.String(s.es.GetNode().GetToken()),
		}

		err = util.WriteNsonMessage(stream, wmessage)
		if err != nil {
			return err
		}
	}

	err = stream.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))
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
			return fmt.Errorf("handshake: error code %v", code)
		}
	}

	err = stream.SetReadDeadline(time.Time{})
	if err != nil {
		return err
	}

	go func(stream quic.Stream) {
		defer stream.Close()

		_, err := io.Copy(stream, stream)
		if err != nil {
			s.es.Logger().Sugar().Errorf("stream(%v) error: %v", stream.StreamID(), err)
		}
	}(stream)

	return nil
}

func (s *QuicService) accept(conn quic.Connection) error {
	defer conn.CloseWithError(1, "exit")

	for {
		stream, err := conn.AcceptStream(s.ctx)
		if err != nil {
			return err
		}

		go func() {
			err := s.handleStream(stream)
			if err != nil {
				s.es.Logger().Sugar().Errorf("QuicService.handleStream(stream) error: %v", err)
			}
		}()
	}
}

func (s *QuicService) handleStream(stream quic.Stream) error {
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

	portId, conn, err := s.openPort(rmessage)
	if err != nil {
		wmessage := nson.Message{}
		wmessage.Insert("code", nson.I32(400))
		util.WriteNsonMessage(stream, wmessage)
		return err
	}

	defer conn.Close()

	wmessage := nson.Message{}
	wmessage.Insert("code", nson.I32(0))
	err = util.WriteNsonMessage(stream, wmessage)
	if err != nil {
		return err
	}

	s.incLinkNum(portId)

	errChan := make(chan error)
	go quicStreamCopy1(stream, conn, errChan)
	go quicStreamCopy2(conn, stream, errChan)

	err = <-errChan
	if err != nil {
		s.es.Logger().Sugar().Errorf("stream.Copy error: %v", err)
	}

	<-errChan

	s.decLinkNum(portId)

	return err
}

func (s *QuicService) openPort(rmessage nson.Message) (string, net.Conn, error) {
	method, err := rmessage.GetString("method")
	if err != nil {
		return "", nil, err
	}

	portId, err := rmessage.GetString("port")
	if err != nil {
		return "", nil, err
	}

	if method != "proxy" || portId == "" {
		return "", nil, errors.New(`method != "proxy" || portId == ""`)
	}

	port, err := s.es.GetPort().View(context.Background(), &pb.Id{Id: portId})
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

func (s *QuicService) OpenStreamSync() (quic.Stream, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if s.conn.IsSome() {
		return s.conn.Unwrap().OpenStreamSync(context.Background())
	}

	return nil, errors.New("quic not connect")
}

func (s *QuicService) incLinkNum(portId string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if num, ok := s.linkNum[portId]; ok {
		s.linkNum[portId] = num + 1
	} else {
		s.linkNum[portId] = 1
	}
}

func (s *QuicService) decLinkNum(portId string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if num, ok := s.linkNum[portId]; ok {
		if num-1 == 0 {
			delete(s.linkNum, portId)
		} else {
			s.linkNum[portId] = num - 1
		}
	}
}

func (s *QuicService) syncLinkStatus() {
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
				for portId, num := range s.linkNum {
					go func(portId string, num int) {
						request := edges.LinkPortRequest{Id: portId, Status: int32(num)}

						ctx := context.Background()
						s.es.GetPort().Link(ctx, &request)
					}(portId, num)
				}
				s.lock.RUnlock()
			}
		}
	}
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
