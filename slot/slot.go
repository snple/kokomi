package slot

import (
	"context"
	"sync"
	"time"

	"github.com/snple/beacon/consts"
	"github.com/snple/beacon/edge"
	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/edges"
	"github.com/snple/beacon/pb/slots"
	"github.com/snple/beacon/util/metadata"
	"github.com/snple/beacon/util/token"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SlotService struct {
	es *edge.EdgeService

	sync     *SyncService
	device   *DeviceService
	source   *SourceService
	tag      *TagService
	constant *ConstService

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
	ss.source = newSourceService(ss)
	ss.tag = newTagService(ss)
	ss.constant = newConstService(ss)

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
	slots.RegisterSourceServiceServer(server, ss.source)
	slots.RegisterTagServiceServer(server, ss.tag)
	slots.RegisterConstServiceServer(server, ss.constant)
}

type slotOptions struct {
	keepAlive time.Duration
}

func defaultSlotOptions() slotOptions {
	return slotOptions{
		keepAlive: 10 * time.Second,
	}
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

func WithKeepAlive(d time.Duration) SlotOption {
	return newFuncSlotOption(func(o *slotOptions) {
		o.keepAlive = d
	})
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

func (s *SlotService) Link(ctx context.Context, in *slots.SlotLinkRequest) (*pb.MyBool, error) {
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

	request2 := &edges.SlotLinkRequest{Id: slotID, Status: in.GetStatus()}

	return s.es.GetSlot().Link(ctx, request2)
}

func (s *SlotService) KeepAlive(in *pb.MyEmpty, stream slots.SlotService_KeepAliveServer) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(stream.Context())
	if err != nil {
		return err
	}

	for {
		err := stream.Send(&slots.SlotKeepAliveReply{Time: int32(time.Now().Unix())})
		if err != nil {
			return err
		}

		time.Sleep(s.dopts.keepAlive)
	}
}
