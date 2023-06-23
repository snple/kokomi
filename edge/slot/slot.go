package slot

import (
	"context"
	"sync"

	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/edge/edge"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/pb/slots"
	"github.com/snple/kokomi/util/metadata"
	"github.com/snple/kokomi/util/token"
	"github.com/snple/rgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SlotService struct {
	es *edge.EdgeService

	sync     *SyncService
	device   *DeviceService
	option   *OptionService
	source   *SourceService
	tag      *TagService
	constant *ConstService
	cable    *CableService
	wire     *WireService
	rgrpc    *RgrpcService

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup

	dopts slotOptions

	slots.UnimplementedSlotServiceServer
}

func Slot(es *edge.EdgeService, opts ...SlotOption) (*SlotService, error) {
	ctx, cancel := context.WithCancel(es.Context())

	ss := &SlotService{
		es:     es,
		ctx:    ctx,
		cancel: cancel,
		dopts:  defaultSlotOptions(),
	}

	for _, opt := range extraSlotOptions {
		opt.apply(&ss.dopts)
	}

	for _, opt := range opts {
		opt.apply(&ss.dopts)
	}

	ss.sync = newSyncService(ss)
	ss.device = newDeviceService(ss)
	ss.option = newOptionService(ss)
	ss.source = newSourceService(ss)
	ss.tag = newTagService(ss)
	ss.constant = newConstService(ss)
	ss.cable = newCableService(ss)
	ss.wire = newWireService(ss)
	ss.rgrpc = newRgrpcService(ss)

	return ss, nil
}

func (ss *SlotService) Start() {

}

func (ss *SlotService) Stop() {
	ss.cancel()
	ss.closeWG.Wait()
}

func (ss *SlotService) Edge() *edge.EdgeService {
	return ss.es
}

func (ss *SlotService) Context() context.Context {
	return ss.ctx
}

func (ss *SlotService) Logger() *zap.Logger {
	return ss.es.Logger()
}

func (ss *SlotService) RegisterGrpc(server *grpc.Server) {
	slots.RegisterSyncServiceServer(server, ss.sync)
	slots.RegisterDeviceServiceServer(server, ss.device)
	slots.RegisterSlotServiceServer(server, ss)
	slots.RegisterOptionServiceServer(server, ss.option)
	slots.RegisterSourceServiceServer(server, ss.source)
	slots.RegisterTagServiceServer(server, ss.tag)
	slots.RegisterConstServiceServer(server, ss.constant)
	slots.RegisterCableServiceServer(server, ss.cable)
	slots.RegisterWireServiceServer(server, ss.wire)
	rgrpc.RegisterRgrpcServiceServer(server, ss.rgrpc)
}

type slotOptions struct {
}

func defaultSlotOptions() slotOptions {
	return slotOptions{}
}

type SlotOption interface {
	apply(*slotOptions)
}

var extraSlotOptions []SlotOption

type funcSlotOption struct {
	f func(*slotOptions)
}

func (fdo *funcSlotOption) apply(do *slotOptions) {
	fdo.f(do)
}

func newFuncSlotOption(f func(*slotOptions)) *funcSlotOption {
	return &funcSlotOption{
		f: f,
	}
}

func (s *SlotService) Login(ctx context.Context, in *slots.LoginSlotRequest) (*slots.LoginSlotReply, error) {
	var output slots.LoginSlotReply

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	s.es.Logger().Sugar().Infof("slot connect start, id: %v, ip: %v", in.GetId(), metadata.GetPeerAddr(ctx))

	request := &pb.Id{Id: in.GetId()}

	reply, err := s.es.GetSlot().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetStatus() != consts.ON {
		s.es.Logger().Sugar().Errorf("slot connect error: slot is not enable, id: %v, ip: %v",
			in.GetId(), metadata.GetPeerAddr(ctx))
		return &output, status.Error(codes.FailedPrecondition, "The slot is not enable")
	}

	if reply.GetSecret() != string(in.GetSecret()) {
		s.es.Logger().Sugar().Errorf("slot connect error: slot secret is not valid, id: %v, ip: %v",
			in.GetId(), metadata.GetPeerAddr(ctx))
		return &output, status.Error(codes.Unauthenticated, "Please supply valid secret")
	}

	tokens, err := token.ClaimSlotToken(reply.Id)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "claim token: %v", err)
	}

	s.es.Logger().Sugar().Infof("slot connect success, id: %v, ip: %v", in.GetId(), metadata.GetPeerAddr(ctx))

	output.Token = tokens

	return &output, nil
}

func validateToken(ctx context.Context) (slotID string, err error) {
	tokens, err := metadata.GetToken(ctx)
	if err != nil {
		return "", err
	}

	ok, slotID := token.ValidateSlotToken(tokens)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "Token validation failed")
	}

	return slotID, nil
}

func (s *SlotService) Update(ctx context.Context, in *pb.Slot) (*pb.Slot, error) {
	var output pb.Slot
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	slotID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &pb.Id{Id: slotID}

	reply, err := s.es.GetSlot().View(ctx, request)
	if err != nil {
		return &output, err
	}

	reply.Name = in.GetName()
	reply.Desc = in.GetDesc()
	// reply.Secret = in.GetSecret()
	reply.Type = in.GetType()
	reply.Link = in.GetLink()
	reply.Config = in.GetConfig()
	// reply.Status = in.GetStatus()
	reply.Tags = in.GetTags()

	return s.es.GetSlot().Update(ctx, in)
}

func (s *SlotService) View(ctx context.Context, in *pb.MyEmpty) (*pb.Slot, error) {
	var output pb.Slot
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	slotID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &pb.Id{Id: slotID}

	return s.es.GetSlot().View(ctx, request)
}

func (s *SlotService) Link(ctx context.Context, in *slots.LinkSlotRequest) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	slotID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request2 := &edges.LinkSlotRequest{Id: slotID, Status: in.GetStatus()}

	return s.es.GetSlot().Link(ctx, request2)
}
