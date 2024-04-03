package core

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/snple/kokomi/core/model"
	"github.com/snple/kokomi/db"
	"github.com/snple/kokomi/pb/cores"
	"github.com/snple/types"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type CoreService struct {
	db *bun.DB

	status      *StatusService
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
	save        types.Option[*SaveService]
	control     *ControlService

	clone *cloneService

	auth *AuthService
	user *UserService

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup

	dopts coreOptions
}

func Core(db *bun.DB, opts ...CoreOption) (*CoreService, error) {
	return CoreContext(context.Background(), db, opts...)
}

func CoreContext(ctx context.Context, db *bun.DB, opts ...CoreOption) (*CoreService, error) {
	ctx, cancel := context.WithCancel(ctx)

	if db == nil {
		panic("db == nil")
	}

	cs := &CoreService{
		db:     db,
		ctx:    ctx,
		cancel: cancel,
		dopts:  defaultCoreOptions(),
	}

	for _, opt := range extraCoreOptions {
		opt.apply(&cs.dopts)
	}

	for _, opt := range opts {
		opt.apply(&cs.dopts)
	}

	cs.status = newStatusService(cs)
	cs.sync = newSyncService(cs)
	cs.sync_global = newSyncGlobalService(cs)
	cs.device = newDeviceService(cs)
	cs.slot = newSlotService(cs)
	cs.port = newPortService(cs)
	cs.proxy = newProxyService(cs)
	cs.source = newSourceService(cs)
	cs.tag = newTagService(cs)
	cs.constant = newConstService(cs)
	cs.data = newDateService(cs)

	if cs.dopts.save {
		cs.save = types.Some(newSaveService(cs))
	}

	cs.control = newControlService(cs)

	cs.clone = newCloneService(cs)

	cs.auth = newAuthService(cs)
	cs.user = newUserService(cs)

	return cs, nil
}

func (cs *CoreService) Start() {
	if cs.save.IsSome() {
		go func() {
			cs.closeWG.Add(1)
			defer cs.closeWG.Done()

			cs.save.Unwrap().start()
		}()
	}

	if cs.dopts.cache {
		go cs.cacheGC()
	}
}

func (cs *CoreService) Stop() {
	cs.cancel()

	if cs.save.IsSome() {
		cs.save.Unwrap().stop()
	}

	cs.closeWG.Wait()
	cs.dopts.logger.Sync()
}

func (cs *CoreService) GetDB() *bun.DB {
	return cs.db
}

func (cs *CoreService) GetInfluxDB() types.Option[*db.InfluxDB] {
	if cs.dopts.influxdb != nil {
		return types.Some(cs.dopts.influxdb)
	}

	return types.None[*db.InfluxDB]()
}

func (cs *CoreService) GetStatus() *StatusService {
	return cs.status
}

func (cs *CoreService) GetSync() *SyncService {
	return cs.sync
}

func (cs *CoreService) GetSyncGlobal() *SyncGlobalService {
	return cs.sync_global
}

func (cs *CoreService) GetDevice() *DeviceService {
	return cs.device
}

func (cs *CoreService) GetSlot() *SlotService {
	return cs.slot
}

func (cs *CoreService) GetPort() *PortService {
	return cs.port
}

func (cs *CoreService) GetProxy() *ProxyService {
	return cs.proxy
}

func (cs *CoreService) GetSource() *SourceService {
	return cs.source
}

func (cs *CoreService) GetTag() *TagService {
	return cs.tag
}

func (cs *CoreService) GetConst() *ConstService {
	return cs.constant
}

func (cs *CoreService) GetData() *DataService {
	return cs.data
}

func (cs *CoreService) GetSave() types.Option[*SaveService] {
	return cs.save
}

func (cs *CoreService) GetControl() *ControlService {
	return cs.control
}

func (cs *CoreService) getClone() *cloneService {
	return cs.clone
}

func (cs *CoreService) GetAuth() *AuthService {
	return cs.auth
}

func (cs *CoreService) GetUser() *UserService {
	return cs.user
}

func (cs *CoreService) Context() context.Context {
	return cs.ctx
}

func (cs *CoreService) Logger() *zap.Logger {
	return cs.dopts.logger
}

func (cs *CoreService) cacheGC() {
	cs.closeWG.Add(1)
	defer cs.closeWG.Done()

	cs.Logger().Sugar().Info("cache gc started")

	ticker := time.NewTicker(cs.dopts.cacheGCTTL)
	defer ticker.Stop()

	for {
		select {
		case <-cs.ctx.Done():
			return
		case <-ticker.C:
			{
				cs.GetDevice().GC()
				cs.GetSource().GC()
				cs.GetTag().GC()
				cs.GetConst().GC()
				cs.GetUser().GC()
			}
		}
	}
}

func (cs *CoreService) Register(server *grpc.Server) {
	cores.RegisterSyncServiceServer(server, cs.sync)
	cores.RegisterSyncGlobalServiceServer(server, cs.sync_global)
	cores.RegisterDeviceServiceServer(server, cs.device)
	cores.RegisterSlotServiceServer(server, cs.slot)
	cores.RegisterPortServiceServer(server, cs.port)
	cores.RegisterProxyServiceServer(server, cs.proxy)
	cores.RegisterSourceServiceServer(server, cs.source)
	cores.RegisterTagServiceServer(server, cs.tag)
	cores.RegisterConstServiceServer(server, cs.constant)
	cores.RegisterDataServiceServer(server, cs.data)
	cores.RegisterControlServiceServer(server, cs.control)

	cores.RegisterAuthServiceServer(server, cs.auth)
	cores.RegisterUserServiceServer(server, cs.user)
}

func CreateSchema(db bun.IDB) error {
	models := []interface{}{
		(*model.Sync)(nil),
		(*model.SyncGlobal)(nil),
		(*model.Device)(nil),
		(*model.Slot)(nil),
		(*model.Port)(nil),
		(*model.Proxy)(nil),
		(*model.Source)(nil),
		(*model.Tag)(nil),
		(*model.Const)(nil),
		(*model.TagValue)(nil),
		(*model.User)(nil),
	}

	for _, model := range models {
		_, err := db.NewCreateTable().Model(model).IfNotExists().Exec(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

type coreOptions struct {
	logger       *zap.Logger
	linkTTL      time.Duration
	cache        bool
	cacheTTL     time.Duration
	cacheGCTTL   time.Duration
	influxdb     *db.InfluxDB
	save         bool
	saveInterval time.Duration
}

func defaultCoreOptions() coreOptions {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("zap.NewDevelopment(): %v", err)
	}

	return coreOptions{
		logger:       logger,
		linkTTL:      3 * time.Minute,
		cache:        true,
		cacheTTL:     3 * time.Second,
		cacheGCTTL:   3 * time.Hour,
		save:         true,
		saveInterval: time.Minute,
	}
}

type CoreOption interface {
	apply(*coreOptions)
}

var extraCoreOptions []CoreOption

type funcCoreOption struct {
	f func(*coreOptions)
}

func (fdo *funcCoreOption) apply(do *coreOptions) {
	fdo.f(do)
}

func newFuncCoreOption(f func(*coreOptions)) *funcCoreOption {
	return &funcCoreOption{
		f: f,
	}
}

func WithLogger(logger *zap.Logger) CoreOption {
	return newFuncCoreOption(func(o *coreOptions) {
		o.logger = logger
	})
}

func WithLinkTTL(d time.Duration) CoreOption {
	return newFuncCoreOption(func(o *coreOptions) {
		o.linkTTL = d
	})
}

func WithCache(enable bool) CoreOption {
	return newFuncCoreOption(func(o *coreOptions) {
		o.cache = enable
	})
}

func WithCacheTTL(d time.Duration) CoreOption {
	return newFuncCoreOption(func(o *coreOptions) {
		o.cacheTTL = d
	})
}

func WithCacheGCTTL(d time.Duration) CoreOption {
	return newFuncCoreOption(func(o *coreOptions) {
		o.cacheGCTTL = d
	})
}

func WithInfluxDB(influxdb *db.InfluxDB) CoreOption {
	return newFuncCoreOption(func(o *coreOptions) {
		o.influxdb = influxdb
	})
}

func WithSave(enable bool) CoreOption {
	return newFuncCoreOption(func(o *coreOptions) {
		o.save = enable
	})
}

func WithSaveInterval(d time.Duration) CoreOption {
	return newFuncCoreOption(func(o *coreOptions) {
		o.saveInterval = d
	})
}
