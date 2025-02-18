package node

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/snple/kokomi/core"
	"go.uber.org/zap"
)

type NodeService struct {
	cs *core.CoreService

	listener *net.TCPListener

	connMap   map[string]*Conn
	connMapMu sync.RWMutex

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup

	dopts nodeOptions
}

func NewNodeService(cs *core.CoreService, opts ...NodeOption) (*NodeService, error) {
	ctx, cancel := context.WithCancel(cs.Context())

	ns := &NodeService{
		cs:     cs,
		ctx:    ctx,
		cancel: cancel,
		dopts:  defaultNodeOptions(),
	}

	for _, opt := range extraNodeOptions {
		opt.apply(&ns.dopts)
	}

	for _, opt := range opts {
		opt.apply(&ns.dopts)
	}

	lisConfig := &net.ListenConfig{
		KeepAliveConfig: net.KeepAliveConfig{
			Enable:   true,
			Idle:     15 * time.Second,
			Interval: 15 * time.Second,
			Count:    9,
		},
	}

	lis, err := lisConfig.Listen(ctx, "tcp", ns.dopts.addr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}

	ns.listener = lis.(*net.TCPListener)

	return ns, nil
}

func (ns *NodeService) Start() {
	go ns.Listen()
}

func (ns *NodeService) Stop() {

	ns.cancel()
	ns.closeWG.Wait()
}

func (ns *NodeService) Core() *core.CoreService {
	return ns.cs
}

func (ns *NodeService) Logger() *zap.Logger {
	return ns.cs.Logger()
}

func (ns *NodeService) Context() context.Context {
	return ns.ctx
}

func (ns *NodeService) Listen() {
	for {
		conn, err := ns.listener.Accept()
		if err != nil {
			ns.Logger().Error("failed to accept connection", zap.Error(err))
			continue
		}

		if ns.dopts.tlsConfig != nil {
			conn = tls.Server(conn, ns.dopts.tlsConfig)
		}

		go ns.handleConn(conn)
	}
}

func (ns *NodeService) handleConn(conn net.Conn) {
	ns.Logger().Info("new connection", zap.String("remote_addr", conn.RemoteAddr().String()))

	nc, err := NewConn(ns, conn)
	if err != nil {
		ns.Logger().Error("failed to create conn", zap.Error(err))
		return
	}

	ns.connMapMu.Lock()
	ns.connMap[nc.ID()] = nc
	ns.connMapMu.Unlock()

	nc.handle()

	ns.connMapMu.Lock()
	delete(ns.connMap, nc.ID())
	ns.connMapMu.Unlock()

	ns.Logger().Info("connection closed", zap.String("remote_addr", conn.RemoteAddr().String()))
	nc.Close()
}

type nodeOptions struct {
	addr      string
	tlsConfig *tls.Config
}

func defaultNodeOptions() nodeOptions {
	return nodeOptions{
		addr: ":8008",
	}
}

type NodeOption interface {
	apply(*nodeOptions)
}

var extraNodeOptions []NodeOption

type funcNodeOption struct {
	f func(*nodeOptions)
}

func (fdo *funcNodeOption) apply(do *nodeOptions) {
	fdo.f(do)
}

func newFuncNodeOption(f func(*nodeOptions)) *funcNodeOption {
	return &funcNodeOption{
		f: f,
	}
}

func WithAddr(addr string) NodeOption {
	return newFuncNodeOption(func(o *nodeOptions) {
		o.addr = addr
	})
}

func WithTLSConfig(tlsConfig *tls.Config) NodeOption {
	return newFuncNodeOption(func(o *nodeOptions) {
		o.tlsConfig = tlsConfig
	})
}
