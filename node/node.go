package node

import (
	"context"
	"sync"
	"time"

	"github.com/snple/beacon/core"
	"github.com/snple/beacon/pb/nodes"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type NodeService struct {
	cs *core.CoreService

	sync     *SyncService
	device   *DeviceService
	slot     *SlotService
	source   *SourceService
	tag      *TagService
	constant *ConstService

	auth *AuthService
	user *UserService

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup

	dopts nodeOptions
}

func Node(cs *core.CoreService, opts ...NodeOption) (*NodeService, error) {
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

	ns.sync = newSyncService(ns)
	ns.device = newDeviceService(ns)
	ns.slot = newSlotService(ns)
	ns.source = newSourceService(ns)
	ns.tag = newTagService(ns)
	ns.constant = newConstService(ns)

	ns.auth = newAuthService(ns)
	ns.user = newUserService(ns)

	return ns, nil
}

func (ns *NodeService) Start() {

}

func (ns *NodeService) Stop() {

	ns.cancel()
	ns.closeWG.Wait()
}

func (ns *NodeService) Core() *core.CoreService {
	return ns.cs
}

func (ns *NodeService) Context() context.Context {
	return ns.ctx
}

func (ns *NodeService) Logger() *zap.Logger {
	return ns.cs.Logger()
}

func (ns *NodeService) RegisterGrpc(server *grpc.Server) {
	nodes.RegisterSyncServiceServer(server, ns.sync)
	nodes.RegisterDeviceServiceServer(server, ns.device)
	nodes.RegisterSlotServiceServer(server, ns.slot)
	nodes.RegisterSourceServiceServer(server, ns.source)
	nodes.RegisterTagServiceServer(server, ns.tag)
	nodes.RegisterConstServiceServer(server, ns.constant)

	nodes.RegisterAuthServiceServer(server, ns.auth)
	nodes.RegisterUserServiceServer(server, ns.user)
}

type nodeOptions struct {
	Ping time.Duration
}

func defaultNodeOptions() nodeOptions {
	return nodeOptions{
		Ping: 60 * time.Second,
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

func WithPing(d time.Duration) NodeOption {
	return newFuncNodeOption(func(o *nodeOptions) {
		o.Ping = d
	})
}
