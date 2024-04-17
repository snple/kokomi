package edge

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/danclive/nson-go"
	"github.com/quic-go/quic-go"
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
}

func newQuicService(es *EdgeService) (*QuicService, error) {
	ctx, cancel := context.WithCancel(es.Context())

	s := &QuicService{
		es:     es,
		ctx:    ctx,
		cancel: cancel,
	}

	return s, nil
}

func (s *QuicService) start() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	if option := s.es.GetNode(); option.IsNone() {
		return
	}

	s.es.Logger().Info("start quic service")

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
	option := s.es.GetNode()
	if option.IsNone() {
		panic("node not enable")
	}

	token := option.Unwrap().GetToken()
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
				s.es.Logger().Sugar().Infof("quic connect: %v", err)
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

	ctx := conn.Context()
	<-ctx.Done()
	// if !errors.Is(ctx.Err(), context.Canceled) {
	s.es.Logger().Sugar().Debugf("break conn error: %v", ctx.Err())
	// }

	s.lock.Lock()
	s.conn.Take()
	s.lock.Unlock()

	return nil
}

func (s *QuicService) connect(ctx context.Context) (quic.Connection, error) {
	conn, err := quic.DialAddr(ctx, s.es.dopts.QuicOptions.Addr,
		s.es.dopts.QuicOptions.TLSConfig, s.es.dopts.QuicOptions.QUICConfig)
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
	stream, err := conn.OpenStreamSync(s.ctx)
	if err != nil {
		return err
	}

	{
		option := s.es.GetNode()
		if option.IsNone() {
			panic("node not enable")
		}

		request := nson.Map{
			"method": nson.String("handshake"),
			"token":  nson.String(option.Unwrap().GetToken()),
		}

		err = util.WriteNsonMessage(stream, request)
		if err != nil {
			return err
		}
	}

	err = stream.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))
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

	request, err := util.ReadNsonMessage(stream)
	if err != nil {
		return err
	}

	err = stream.SetReadDeadline(time.Time{})
	if err != nil {
		return err
	}

	method, err := request.GetString("method")
	if err != nil {
		return fmt.Errorf("message format error: %v", err)
	}

	switch method {
	case "proxy":
		if option := s.es.GetQuicProxy(); option.IsSome() {
			return option.Unwrap().handleProxy(request, stream)
		}
	}

	response := nson.Map{
		"code": nson.I32(404),
	}
	util.WriteNsonMessage(stream, response)

	return nil
}

func (s *QuicService) OpenStreamSync() (quic.Stream, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if s.conn.IsSome() {
		return s.conn.Unwrap().OpenStreamSync(s.ctx)
	}

	return nil, errors.New("quic not connect")
}
