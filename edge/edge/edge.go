package edge

import (
	"context"
	"crypto/tls"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/quic-go/quic-go"
	"github.com/snple/kokomi/edge/model"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/types"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type EdgeService struct {
	db       *bun.DB
	badger   *BadgerService
	status   *StatusService
	node     *NodeService
	sync     *SyncService
	device   *DeviceService
	slot     *SlotService
	option   *OptionService
	port     *PortService
	proxy    *ProxyService
	source   *SourceService
	tag      *TagService
	variable *VarService
	cable    *CableService
	wire     *WireService
	class    *ClassService
	attr     *AttrService
	control  *ControlService
	quic     types.Option[*QuicService]
	tunnel   types.Option[*TunnelService]

	clone *cloneService

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup

	dopts edgeOptions
}

func Edge(db *bun.DB, badger *badger.DB, opts ...EdgeOption) (*EdgeService, error) {
	ctx := context.Background()
	return EdgeContext(ctx, db, badger, opts...)
}

func EdgeContext(ctx context.Context, db *bun.DB, badger *badger.DB, opts ...EdgeOption) (*EdgeService, error) {
	ctx, cancel := context.WithCancel(ctx)

	if db == nil {
		panic("db == nil")
	}

	if badger == nil {
		panic("badger == nil")
	}

	es := &EdgeService{
		db:     db,
		ctx:    ctx,
		cancel: cancel,
		dopts:  defaultEdgeOptions(),
	}

	for _, opt := range extraEdgeOptions {
		opt.apply(&es.dopts)
	}

	for _, opt := range opts {
		opt.apply(&es.dopts)
	}

	if err := es.dopts.check(); err != nil {
		return nil, err
	}

	es.badger = newBadgerService(es, badger)
	es.status = newStatusService(es)

	node, err := newNodeService(es)
	if err != nil {
		return nil, err
	}
	es.node = node

	es.sync = newSyncService(es)
	es.device = newDeviceService(es)
	es.slot = newSlotService(es)
	es.option = newOptionService(es)
	es.port = newPortService(es)
	es.proxy = newProxyService(es)
	es.source = newSourceService(es)
	es.tag = newTagService(es)
	es.variable = newVarService(es)
	es.cable = newCableService(es)
	es.wire = newWireService(es)
	es.class = newClassService(es)
	es.attr = newAttrService(es)
	es.control = newControlService(es)

	if es.dopts.quicOptions != nil {
		quic, err := newQuicService(es)
		if err != nil {
			return es, err
		}

		es.quic = types.Some(quic)

		es.tunnel = types.Some(newTunnelService(es))
	}

	es.clone = newCloneService(es)

	return es, nil
}

func (es *EdgeService) Start() {
	go func() {
		es.closeWG.Add(1)
		defer es.closeWG.Done()

		es.GetNode().start()
	}()

	if es.quic.IsSome() {
		go func() {
			es.closeWG.Add(1)
			defer es.closeWG.Done()

			es.quic.Unwrap().start()
		}()
	}

	if es.tunnel.IsSome() {
		go func() {
			es.closeWG.Add(1)
			defer es.closeWG.Done()

			es.tunnel.Unwrap().start()
		}()
	}
}

func (es *EdgeService) Stop() {
	es.cancel()

	if es.tunnel.IsSome() {
		es.tunnel.Unwrap().stop()
	}
	if es.quic.IsSome() {
		es.quic.Unwrap().stop()
	}
	es.GetNode().stop()

	es.dopts.logger.Sync()
	es.closeWG.Wait()
}

func (es *EdgeService) GetDB() *bun.DB {
	return es.db
}

func (es *EdgeService) GetBadgerDB() *badger.DB {
	return es.badger.GetDB()
}

func (es *EdgeService) GetStatus() *StatusService {
	return es.status
}

func (es *EdgeService) GetNode() *NodeService {
	return es.node
}

func (es *EdgeService) GetSync() *SyncService {
	return es.sync
}

func (es *EdgeService) GetDevice() *DeviceService {
	return es.device
}

func (es *EdgeService) GetSlot() *SlotService {
	return es.slot
}

func (es *EdgeService) GetOption() *OptionService {
	return es.option
}

func (es *EdgeService) GetPort() *PortService {
	return es.port
}

func (es *EdgeService) GetProxy() *ProxyService {
	return es.proxy
}

func (es *EdgeService) GetSource() *SourceService {
	return es.source
}

func (es *EdgeService) GetTag() *TagService {
	return es.tag
}

func (es *EdgeService) GetVar() *VarService {
	return es.variable
}

func (es *EdgeService) GetCable() *CableService {
	return es.cable
}

func (es *EdgeService) GetWire() *WireService {
	return es.wire
}

func (es *EdgeService) GetClass() *ClassService {
	return es.class
}

func (es *EdgeService) GetAttr() *AttrService {
	return es.attr
}

func (es *EdgeService) GetControl() *ControlService {
	return es.control
}

func (es *EdgeService) GetQuic() types.Option[*QuicService] {
	return es.quic
}

func (es *EdgeService) getClone() *cloneService {
	return es.clone
}

func (es *EdgeService) Context() context.Context {
	return es.ctx
}

func (es *EdgeService) Logger() *zap.Logger {
	return es.dopts.logger
}

func (es *EdgeService) Register(server *grpc.Server) {
	edges.RegisterSyncServiceServer(server, es.sync)
	edges.RegisterDeviceServiceServer(server, es.device)
	edges.RegisterSlotServiceServer(server, es.slot)
	edges.RegisterOptionServiceServer(server, es.option)
	edges.RegisterPortServiceServer(server, es.port)
	edges.RegisterProxyServiceServer(server, es.proxy)
	edges.RegisterSourceServiceServer(server, es.source)
	edges.RegisterTagServiceServer(server, es.tag)
	edges.RegisterVarServiceServer(server, es.variable)
	edges.RegisterCableServiceServer(server, es.cable)
	edges.RegisterWireServiceServer(server, es.wire)
	edges.RegisterClassServiceServer(server, es.class)
	edges.RegisterAttrServiceServer(server, es.attr)
	edges.RegisterControlServiceServer(server, es.control)
}

func CreateSchema(db bun.IDB) error {
	models := []interface{}{
		(*model.Device)(nil),
		(*model.Slot)(nil),
		(*model.Option)(nil),
		(*model.Port)(nil),
		(*model.Proxy)(nil),
		(*model.Source)(nil),
		(*model.Tag)(nil),
		(*model.Var)(nil),
		(*model.Cable)(nil),
		(*model.Wire)(nil),
		(*model.Class)(nil),
		(*model.Attr)(nil),
	}

	for _, model := range models {
		_, err := db.NewCreateTable().Model(model).IfNotExists().Exec(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

type edgeOptions struct {
	deviceID string
	secret   string

	nodeOptions *nodeOptions
	quicOptions *quicOptions
	logger      *zap.Logger

	linkStatusTTL  time.Duration
	tokenRefresh   time.Duration
	syncLinkStatus time.Duration
	syncInterval   time.Duration
	syncRealtime   bool

	badgerGCInterval     time.Duration
	badgerGCDiscardRatio float64
}

type nodeOptions struct {
	Addr    string
	Options []grpc.DialOption
}

type quicOptions struct {
	Addr       string
	TLSConfig  *tls.Config
	QUICConfig *quic.Config
}

func defaultEdgeOptions() edgeOptions {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("zap.NewDevelopment(): %v", err)
	}

	return edgeOptions{
		logger:               logger,
		linkStatusTTL:        3 * time.Minute,
		tokenRefresh:         30 * time.Minute,
		syncLinkStatus:       time.Minute,
		syncInterval:         time.Minute,
		syncRealtime:         false,
		badgerGCInterval:     time.Hour,
		badgerGCDiscardRatio: 0.7,
	}
}

func (o *edgeOptions) check() error {
	if o.nodeOptions == nil {
		return errors.New("Please supply valid node options")
	}

	return nil
}

type EdgeOption interface {
	apply(*edgeOptions)
}

var extraEdgeOptions []EdgeOption

type funcEdgeOption struct {
	f func(*edgeOptions)
}

func (fdo *funcEdgeOption) apply(do *edgeOptions) {
	fdo.f(do)
}

func newFuncEdgeOption(f func(*edgeOptions)) *funcEdgeOption {
	return &funcEdgeOption{
		f: f,
	}
}

func WithDeviceID(id, secret string) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.deviceID = id
		o.secret = secret
	})
}

func WithNode(addr string, options []grpc.DialOption) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.nodeOptions = &nodeOptions{addr, options}
	})
}

func WithQuic(addr string, tlsConfig *tls.Config, quicConfig *quic.Config) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		tlsConfig2 := tlsConfig.Clone()
		if len(tlsConfig2.NextProtos) == 0 {
			tlsConfig2.NextProtos = []string{"kokomi"}
		}

		quicConfig2 := quicConfig.Clone()
		quicConfig2.EnableDatagrams = true

		o.quicOptions = &quicOptions{addr, tlsConfig2, quicConfig2}
	})
}

func WithLogger(logger *zap.Logger) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.logger = logger
	})
}

func WithLinkStatusTTL(d time.Duration) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.linkStatusTTL = d
	})
}

func WithTokenRefresh(d time.Duration) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.tokenRefresh = d
	})
}

func WithSyncLinkStatus(d time.Duration) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.syncLinkStatus = d
	})
}

func WithSyncInterval(d time.Duration) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.syncInterval = d
	})
}

func WithSyncRealtime(realtime bool) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.syncRealtime = realtime
	})
}

func WithBadgerGCInterval(d time.Duration) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.badgerGCInterval = d
	})
}

func WithBadgerGCDiscardRatio(discardRatio float64) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.badgerGCDiscardRatio = discardRatio
	})
}
