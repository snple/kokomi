package edge

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/snple/beacon/edge/model"
	"github.com/snple/beacon/pb/edges"
	"github.com/snple/types"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type EdgeService struct {
	db       *bun.DB
	badger   *BadgerService
	status   *StatusService
	sync     *SyncService
	device   *DeviceService
	slot     *SlotService
	source   *SourceService
	tag      *TagService
	constant *ConstService

	node types.Option[*NodeService]

	clone *cloneService

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup

	dopts edgeOptions
}

func Edge(db *bun.DB, opts ...EdgeOption) (*EdgeService, error) {
	ctx := context.Background()
	return EdgeContext(ctx, db, opts...)
}

func EdgeContext(ctx context.Context, db *bun.DB, opts ...EdgeOption) (*EdgeService, error) {
	ctx, cancel := context.WithCancel(ctx)

	if db == nil {
		panic("db == nil")
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

	badger, err := newBadgerService(es)
	if err != nil {
		return nil, err
	}
	es.badger = badger

	es.status = newStatusService(es)
	es.sync = newSyncService(es)
	es.device = newDeviceService(es)
	es.slot = newSlotService(es)
	es.source = newSourceService(es)
	es.tag = newTagService(es)
	es.constant = newConstService(es)

	if es.dopts.NodeOptions.Enable {
		node, err := newNodeService(es)
		if err != nil {
			return nil, err
		}
		es.node = types.Some(node)
	}

	es.clone = newCloneService(es)

	return es, nil
}

func (es *EdgeService) Start() {
	es.closeWG.Add(1)
	go func() {
		defer es.closeWG.Done()

		es.badger.start()
	}()

	if es.node.IsSome() {
		es.closeWG.Add(1)
		go func() {
			defer es.closeWG.Done()

			es.node.Unwrap().start()
		}()
	}

	if es.dopts.cache {
		es.closeWG.Add(1)
		go func() {
			defer es.closeWG.Done()

			es.cacheGC()
		}()
	}
}

func (es *EdgeService) Stop() {
	if es.node.IsSome() {
		es.node.Unwrap().stop()
	}

	es.badger.stop()

	es.cancel()
	es.closeWG.Wait()
	es.dopts.logger.Sync()
}

func (es *EdgeService) Push() error {
	if es.node.IsSome() {
		return es.node.Unwrap().push()
	}

	return errors.New("node not enable")
}

func (es *EdgeService) Pull() error {
	if es.node.IsSome() {
		return es.node.Unwrap().pull()
	}

	return errors.New("node not enable")
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

func (es *EdgeService) GetSync() *SyncService {
	return es.sync
}

func (es *EdgeService) GetDevice() *DeviceService {
	return es.device
}

func (es *EdgeService) GetSlot() *SlotService {
	return es.slot
}

func (es *EdgeService) GetSource() *SourceService {
	return es.source
}

func (es *EdgeService) GetTag() *TagService {
	return es.tag
}

func (es *EdgeService) GetConst() *ConstService {
	return es.constant
}

func (es *EdgeService) GetNode() types.Option[*NodeService] {
	return es.node
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

func (es *EdgeService) cacheGC() {
	es.Logger().Sugar().Info("cache gc started")

	ticker := time.NewTicker(es.dopts.cacheGCTTL)
	defer ticker.Stop()

	for {
		select {
		case <-es.ctx.Done():
			return
		case <-ticker.C:
			{
				es.GetSource().GC()
				es.GetTag().GC()
				es.GetConst().GC()
			}
		}
	}
}

func (es *EdgeService) Register(server *grpc.Server) {
	edges.RegisterSyncServiceServer(server, es.sync)
	edges.RegisterDeviceServiceServer(server, es.device)
	edges.RegisterSlotServiceServer(server, es.slot)
	edges.RegisterSourceServiceServer(server, es.source)
	edges.RegisterTagServiceServer(server, es.tag)
	edges.RegisterConstServiceServer(server, es.constant)
}

func CreateSchema(db bun.IDB) error {
	models := []any{
		(*model.Device)(nil),
		(*model.Slot)(nil),
		(*model.Source)(nil),
		(*model.Tag)(nil),
		(*model.Const)(nil),
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
	logger   *zap.Logger
	deviceID string
	secret   string

	NodeOptions     NodeOptions
	SyncOptions     SyncOptions
	BadgerOptions   badger.Options
	BadgerGCOptions BadgerGCOptions

	linkTTL    time.Duration
	cache      bool
	cacheTTL   time.Duration
	cacheGCTTL time.Duration
}

type NodeOptions struct {
	Enable      bool
	Addr        string
	GRPCOptions []grpc.DialOption
}

type SyncOptions struct {
	TokenRefresh time.Duration
	Link         time.Duration
	Interval     time.Duration
	Realtime     bool
	Retry        time.Duration
}

type BadgerGCOptions struct {
	GC             time.Duration
	GCDiscardRatio float64
}

func defaultEdgeOptions() edgeOptions {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("zap.NewDevelopment(): %v", err)
	}

	return edgeOptions{
		logger:      logger,
		NodeOptions: NodeOptions{},
		SyncOptions: SyncOptions{
			TokenRefresh: 3 * time.Minute,
			Link:         time.Minute,
			Interval:     time.Minute,
			Realtime:     false,
			Retry:        time.Second * 3,
		},
		BadgerOptions: badger.DefaultOptions("").WithInMemory(true),
		BadgerGCOptions: BadgerGCOptions{
			GC:             time.Hour,
			GCDiscardRatio: 0.7,
		},
		linkTTL:    3 * time.Minute,
		cache:      true,
		cacheTTL:   3 * time.Second,
		cacheGCTTL: 3 * time.Hour,
	}
}

func (o *edgeOptions) check() error {
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

func WithLogger(logger *zap.Logger) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.logger = logger
	})
}

func WithDeviceID(id, secret string) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.deviceID = id
		o.secret = secret
	})
}

func WithNode(options NodeOptions) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.NodeOptions = options
	})
}

func WithSync(options SyncOptions) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		if options.TokenRefresh > 0 {
			o.SyncOptions.TokenRefresh = options.TokenRefresh
		}

		if options.Link > 0 {
			o.SyncOptions.Link = options.Link
		}

		if options.Interval > 0 {
			o.SyncOptions.Interval = options.Interval
		}

		o.SyncOptions.Realtime = options.Realtime
	})
}

func WithBadger(options badger.Options) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.BadgerOptions = options
	})
}

func WithBadgerGC(options *BadgerGCOptions) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		if options.GC > 0 {
			o.BadgerGCOptions.GC = options.GC
		}

		if options.GCDiscardRatio > 0 {
			o.BadgerGCOptions.GCDiscardRatio = options.GCDiscardRatio
		}
	})
}

func WithLinkTTL(d time.Duration) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.linkTTL = d
	})
}

func WithCache(enable bool) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.cache = enable
	})
}

func WithCacheTTL(d time.Duration) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.cacheTTL = d
	})
}

func WithCacheGCTTL(d time.Duration) EdgeOption {
	return newFuncEdgeOption(func(o *edgeOptions) {
		o.cacheGCTTL = d
	})
}
