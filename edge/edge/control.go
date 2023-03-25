package edge

import (
	"context"
	"sync"

	"github.com/snple/types"
	"github.com/snple/types/cache"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"snple.com/kokomi/consts"
	"snple.com/kokomi/pb"
	"snple.com/kokomi/pb/edges"
	"snple.com/kokomi/pb/slots"
	"snple.com/kokomi/util/datatype"
)

type ControlService struct {
	es *EdgeService

	lock        sync.RWMutex
	slots       map[string]*clients
	sourceCache *cache.Cache[string]
	sources     map[string]*servers

	edges.UnimplementedControlServiceServer
}

func newControlService(es *EdgeService) *ControlService {
	return &ControlService{
		es:          es,
		slots:       make(map[string]*clients),
		sourceCache: cache.NewCache[string](nil),
		sources:     make(map[string]*servers),
	}
}

func (s *ControlService) GetTagValue(ctx context.Context, in *pb.Id) (*pb.TagValue, error) {
	var err error
	var output pb.TagValue

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag id")
		}
	}

	output.Id = in.GetId()

	// tag
	tag, err := s.es.GetTag().view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	if tag.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag Status != ON")
	}

	// validation source
	{
		// source
		{
			source, err := s.es.GetSource().view(ctx, tag.SourceID)
			if err != nil {
				return &output, err
			}

			if source.Status != consts.ON {
				return &output, status.Errorf(codes.FailedPrecondition, "Source Status != ON")
			}
		}
	}

	if slot := s.sourceCache.Get(tag.SourceID); slot.IsSome() {
		if option := s.GetSlotClient(slot.Unwrap()); option.IsSome() {
			client := option.Unwrap()

			return client.GetTagValue(ctx, in)
		}
	}

	if option := s.GetSourceServer(tag.SourceID); option.IsSome() {
		server := option.Unwrap()

		return server.GetTagValue(ctx, in)
	}

	return &output, status.Errorf(codes.Unavailable, "Source not connect")
}

func (s *ControlService) SetTagValue(ctx context.Context, in *pb.TagValue) (*pb.TagValue, error) {
	var err error
	var output pb.TagValue

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag id")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid value")
		}
	}

	// tag
	tag, err := s.es.GetTag().view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	if tag.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag Status != ON")
	}

	if tag.Access != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag Access != ON")
	}

	_, err = datatype.DecodeNsonValue(in.GetValue(), tag.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	// validation source
	{
		// source
		{
			source, err := s.es.GetSource().view(ctx, tag.SourceID)
			if err != nil {
				return &output, err
			}

			if source.Status != consts.ON {
				return &output, status.Errorf(codes.FailedPrecondition, "Source Status != ON")
			}
		}
	}

	if slot := s.sourceCache.Get(tag.SourceID); slot.IsSome() {
		if option := s.GetSlotClient(slot.Unwrap()); option.IsSome() {
			client := option.Unwrap()

			return client.SetTagValue(ctx, in)
		}
	}

	if option := s.GetSourceServer(tag.SourceID); option.IsSome() {
		server := option.Unwrap()

		return server.SetTagValue(ctx, in)
	}

	return &output, status.Errorf(codes.Unavailable, "Source not connect")
}

func (s *ControlService) AddSlotClient(slotID string, client slots.ControlServiceClient) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if items, ok := s.slots[slotID]; ok {
		items.add(client)
	} else {
		items := &clients{}
		items.add(client)
		s.slots[slotID] = items
	}
}

func (s *ControlService) DeleteSlotClient(slotID string, client slots.ControlServiceClient) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if items, ok := s.slots[slotID]; ok {
		items.delete(client)
	}
}

func (s *ControlService) GetSlotClient(slotID string) types.Option[slots.ControlServiceClient] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if items, ok := s.slots[slotID]; ok {
		return items.pick()
	}

	return types.None[slots.ControlServiceClient]()
}

func (s *ControlService) LinkSource(slotID string, sourceID string, status int32) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if status == consts.ON {
		s.sourceCache.Set(sourceID, slotID, s.es.dopts.linkStatusTTL)
	} else {
		s.sourceCache.Delete(sourceID)
	}
}

func (s *ControlService) AddSourceServer(sourceID string, server slots.ControlServiceServer) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if items, ok := s.sources[sourceID]; ok {
		items.add(server)
	} else {
		item := &servers{}
		item.add(server)
		s.sources[sourceID] = item
	}
}

func (s *ControlService) RemoveSourceServer(sourceID string, server slots.ControlServiceServer) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if servers, ok := s.sources[sourceID]; ok {
		servers.delete(server)
	}
}

func (s *ControlService) GetSourceServer(sourceID string) types.Option[slots.ControlServiceServer] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if servers, ok := s.sources[sourceID]; ok {
		return servers.pick()
	}

	return types.None[slots.ControlServiceServer]()
}

type clients struct {
	inner []slots.ControlServiceClient
}

func (cs *clients) add(cc slots.ControlServiceClient) {
	cs.inner = append(cs.inner, cc)
}

func (cs *clients) delete(cc slots.ControlServiceClient) {
	for i := range cs.inner {
		if cs.inner[i] == cc {
			cs.inner = append(cs.inner[:i], cs.inner[i+1:]...)
			break
		}
	}
}

func (cs *clients) pick() types.Option[slots.ControlServiceClient] {
	if cs == nil {
		return types.None[slots.ControlServiceClient]()
	}

	if len(cs.inner) == 0 {
		return types.None[slots.ControlServiceClient]()
	}

	return types.Some(cs.inner[len(cs.inner)-1])
}

type servers struct {
	inner []slots.ControlServiceServer
}

func (cs *servers) add(cc slots.ControlServiceServer) {
	cs.inner = append(cs.inner, cc)
}

func (cs *servers) delete(cc slots.ControlServiceServer) {
	for i := range cs.inner {
		if cs.inner[i] == cc {
			cs.inner = append(cs.inner[:i], cs.inner[i+1:]...)
			break
		}
	}
}

func (cs *servers) pick() types.Option[slots.ControlServiceServer] {
	if cs == nil {
		return types.None[slots.ControlServiceServer]()
	}

	if len(cs.inner) == 0 {
		return types.None[slots.ControlServiceServer]()
	}

	return types.Some(cs.inner[len(cs.inner)-1])
}
