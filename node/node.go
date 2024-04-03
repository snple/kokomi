package node

import (
	"context"
	"crypto/tls"
	"sync"
	"time"

	"github.com/quic-go/quic-go"
	"github.com/snple/kokomi/core"
	"github.com/snple/kokomi/pb/nodes"
	"github.com/snple/rgrpc"
	"github.com/snple/types"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type NodeService struct {
	cs *core.CoreService

	sync        *SyncService
	sync_global *SyncGlobalService
	device      *DeviceService
	slot        *SlotService
	port        *PortService
	proxy       *ProxyService
	source      *SourceService
	tag         *TagService
	constant    *ConstService
	data        *DataService
	rgrpc       *RgrpcService
	quic        types.Option[*QuicService]

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
	ns.sync_global = newSyncGlobalService(ns)
	ns.device = newDeviceService(ns)
	ns.slot = newSlotService(ns)
	ns.port = newPortService(ns)
	ns.proxy = newProxyService(ns)
	ns.source = newSourceService(ns)
	ns.tag = newTagService(ns)
	ns.constant = newConstService(ns)
	ns.data = newDataService(ns)
	ns.rgrpc = newRgrpcService(ns)

	if ns.dopts.QuicOptions.enable {
		quic, err := newQuicService(ns)
		if err != nil {
			return ns, err
		}

		ns.quic = types.Some(quic)
	}

	ns.auth = newAuthService(ns)
	ns.user = newUserService(ns)

	return ns, nil
}

func (ns *NodeService) Start() {
	if quic := ns.quic; quic.IsSome() {
		go func() {
			ns.closeWG.Add(1)
			defer ns.closeWG.Done()

			quic.Unwrap().Start()
		}()
	}
}

func (ns *NodeService) Stop() {
	if quic := ns.quic; quic.IsSome() {
		quic.Unwrap().Stop()
	}

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
	nodes.RegisterSyncGlobalServiceServer(server, ns.sync_global)
	nodes.RegisterDeviceServiceServer(server, ns.device)
	nodes.RegisterSlotServiceServer(server, ns.slot)
	nodes.RegisterPortServiceServer(server, ns.port)
	nodes.RegisterProxyServiceServer(server, ns.proxy)
	nodes.RegisterSourceServiceServer(server, ns.source)
	nodes.RegisterTagServiceServer(server, ns.tag)
	nodes.RegisterConstServiceServer(server, ns.constant)
	nodes.RegisterDataServiceServer(server, ns.data)
	rgrpc.RegisterRgrpcServiceServer(server, ns.rgrpc)

	nodes.RegisterAuthServiceServer(server, ns.auth)
	nodes.RegisterUserServiceServer(server, ns.user)
}

type nodeOptions struct {
	QuicOptions QuicOptions
	Ping        time.Duration
}

type QuicOptions struct {
	enable     bool
	Addr       string
	TLSConfig  *tls.Config
	QUICConfig *quic.Config
}

func defaultNodeOptions() nodeOptions {
	return nodeOptions{
		QuicOptions: QuicOptions{},
		Ping:        60 * time.Second,
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

func WithQuic(options QuicOptions) NodeOption {
	return newFuncNodeOption(func(o *nodeOptions) {
		if len(options.TLSConfig.NextProtos) == 0 {
			options.TLSConfig.NextProtos = []string{"kokomi"}
		}

		options.QUICConfig.EnableDatagrams = true

		o.QuicOptions = options
		o.QuicOptions.enable = true
	})
}

func WithPing(d time.Duration) NodeOption {
	return newFuncNodeOption(func(o *nodeOptions) {
		o.Ping = d
	})
}
