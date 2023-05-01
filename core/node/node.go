package node

import (
	"context"
	"crypto/tls"
	"sync"
	"time"

	"github.com/quic-go/quic-go"
	"github.com/snple/kokomi/core/core"
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
	option      *OptionService
	port        *PortService
	proxy       *ProxyService
	source      *SourceService
	tag         *TagService
	variable    *VarService
	cable       *CableService
	wire        *WireService
	class       *ClassService
	attr        *AttrService
	rgrpc       *RgrpcService
	quic        types.Option[*QuicService]

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
	ns.option = newOptionService(ns)
	ns.port = newPortService(ns)
	ns.proxy = newProxyService(ns)
	ns.source = newSourceService(ns)
	ns.tag = newTagService(ns)
	ns.variable = newVarService(ns)
	ns.cable = newCableService(ns)
	ns.wire = newWireService(ns)
	ns.class = newClassService(ns)
	ns.attr = newAttrService(ns)
	ns.rgrpc = newRgrpcService(ns)

	if ns.dopts.quicOptions != nil {
		quic, err := newQuicService(ns)
		if err != nil {
			return ns, err
		}

		ns.quic = types.Some(quic)
	}

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
	nodes.RegisterOptionServiceServer(server, ns.option)
	nodes.RegisterPortServiceServer(server, ns.port)
	nodes.RegisterProxyServiceServer(server, ns.proxy)
	nodes.RegisterSourceServiceServer(server, ns.source)
	nodes.RegisterTagServiceServer(server, ns.tag)
	nodes.RegisterVarServiceServer(server, ns.variable)
	nodes.RegisterCableServiceServer(server, ns.cable)
	nodes.RegisterWireServiceServer(server, ns.wire)
	nodes.RegisterClassServiceServer(server, ns.class)
	nodes.RegisterAttrServiceServer(server, ns.attr)
	rgrpc.RegisterRgrpcServiceServer(server, ns.rgrpc)
}

type nodeOptions struct {
	quicOptions      *quicOptions
	quicPingInterval time.Duration
}

type quicOptions struct {
	addr       string
	tlsConfig  *tls.Config
	quicConfig *quic.Config
}

func defaultNodeOptions() nodeOptions {
	return nodeOptions{
		quicPingInterval: 60 * time.Second,
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

func WithQuic(addr string, tlsConfig *tls.Config, quicConfig *quic.Config) NodeOption {
	return newFuncNodeOption(func(o *nodeOptions) {
		tlsConfig2 := tlsConfig.Clone()
		if len(tlsConfig2.NextProtos) == 0 {
			tlsConfig2.NextProtos = []string{"kokomi"}
		}

		quicConfig2 := quicConfig.Clone()
		quicConfig2.EnableDatagrams = true

		o.quicOptions = &quicOptions{addr, tlsConfig2, quicConfig2}
	})
}

func WithQuicPingInterval(d time.Duration) NodeOption {
	return newFuncNodeOption(func(o *nodeOptions) {
		o.quicPingInterval = d
	})
}
