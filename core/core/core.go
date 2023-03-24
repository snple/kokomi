package core

import (
	"context"
	"sync"

	"github.com/snple/types"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"snple.com/kokomi/core/model"
	"snple.com/kokomi/db"
	"snple.com/kokomi/pb/cores"
)

type CoreService struct {
	db *bun.DB

	status      *StatusService
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
	control     *ControlService

	clone *cloneService

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup

	dopts coreOptions
}

func Core(db *bun.DB, opts ...CoreOption) (*CoreService, error) {
	return CoreContext(context.Background(), db)
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
	cs.option = newOptionService(cs)
	cs.port = newPortService(cs)
	cs.proxy = newProxyService(cs)
	cs.source = newSourceService(cs)
	cs.tag = newTagService(cs)
	cs.variable = newVarService(cs)
	cs.cable = newCableService(cs)
	cs.wire = newWireService(cs)
	cs.control = newControlService(cs)

	cs.clone = newCloneService(cs)

	return cs, nil
}

func (cs *CoreService) Start() {

}

func (cs *CoreService) Stop() {
	cs.cancel()
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

func (cs *CoreService) GetOption() *OptionService {
	return cs.option
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

func (cs *CoreService) GetVar() *VarService {
	return cs.variable
}

func (cs *CoreService) GetCable() *CableService {
	return cs.cable
}

func (cs *CoreService) GetWire() *WireService {
	return cs.wire
}

func (cs *CoreService) GetControl() *ControlService {
	return cs.control
}

func (cs *CoreService) getClone() *cloneService {
	return cs.clone
}

func (cs *CoreService) Context() context.Context {
	return cs.ctx
}

func (cs *CoreService) Logger() *zap.Logger {
	return cs.dopts.logger
}

func (cs *CoreService) Register(server *grpc.Server) {
	cores.RegisterSyncServiceServer(server, cs.sync)
	cores.RegisterSyncGlobalServiceServer(server, cs.sync_global)
	cores.RegisterDeviceServiceServer(server, cs.device)
	cores.RegisterSlotServiceServer(server, cs.slot)
	cores.RegisterOptionServiceServer(server, cs.option)
	cores.RegisterPortServiceServer(server, cs.port)
	cores.RegisterProxyServiceServer(server, cs.proxy)
	cores.RegisterSourceServiceServer(server, cs.source)
	cores.RegisterTagServiceServer(server, cs.tag)
	cores.RegisterVarServiceServer(server, cs.variable)
	cores.RegisterCableServiceServer(server, cs.cable)
	cores.RegisterWireServiceServer(server, cs.wire)
	cores.RegisterControlServiceServer(server, cs.control)
}

func CreateSchema(db bun.IDB) error {
	models := []interface{}{
		(*model.Sync)(nil),
		(*model.SyncGlobal)(nil),
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
		(*model.TagValue)(nil),
		(*model.WireValue)(nil),
	}

	for _, model := range models {
		_, err := db.NewCreateTable().Model(model).IfNotExists().Exec(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}
