package edge

import (
	"context"
	"sync"

	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/pb/slots"
	"github.com/snple/kokomi/util/datatype"
	"github.com/snple/types"
	"github.com/snple/types/cache"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ControlService struct {
	es *EdgeService

	slots       map[string]*controlChannels
	sourceCache *cache.Cache[string]
	servers     map[string]*controlServers
	lock        sync.RWMutex

	edges.UnimplementedControlServiceServer
}

func newControlService(es *EdgeService) *ControlService {
	return &ControlService{
		es:          es,
		slots:       make(map[string]*controlChannels),
		sourceCache: cache.NewCache[string](nil),
		servers:     make(map[string]*controlServers),
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}
	}

	output.Id = in.GetId()

	// tag
	tag, err := s.es.GetTag().view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	if tag.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag.Status != ON")
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
				return &output, status.Errorf(codes.FailedPrecondition, "Source.Status != ON")
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Value")
		}
	}

	// tag
	tag, err := s.es.GetTag().view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	if tag.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag.Status != ON")
	}

	if tag.Access != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag.Access != ON")
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
				return &output, status.Errorf(codes.FailedPrecondition, "Source.Status != ON")
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

	if chans, ok := s.slots[slotID]; ok {
		chans.add(client)
	} else {
		chans := &controlChannels{}
		chans.add(client)
		s.slots[slotID] = chans
	}
}

func (s *ControlService) DeleteSlotClient(slotID string, client slots.ControlServiceClient) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if chans, ok := s.slots[slotID]; ok {
		chans.delete(client)
	}
}

func (s *ControlService) GetSlotClient(slotID string) types.Option[slots.ControlServiceClient] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if chans, ok := s.slots[slotID]; ok {
		return chans.pick()
	}

	return types.None[slots.ControlServiceClient]()
}

func (s *ControlService) LinkSource(slotID string, sourceID string, status int32) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if status == consts.ON {
		s.sourceCache.Set(sourceID, slotID, s.es.dopts.linkTTL)
	} else {
		s.sourceCache.Delete(sourceID)
	}
}

func (s *ControlService) AddSourceServer(sourceID string, server slots.ControlServiceServer) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if servers, ok := s.servers[sourceID]; ok {
		servers.add(server)
	} else {
		servers := &controlServers{}
		servers.add(server)
		s.servers[sourceID] = servers
	}
}

func (s *ControlService) RemoveSourceServer(sourceID string, server slots.ControlServiceServer) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if servers, ok := s.servers[sourceID]; ok {
		servers.delete(server)
	}
}

func (s *ControlService) GetSourceServer(sourceID string) types.Option[slots.ControlServiceServer] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if servers, ok := s.servers[sourceID]; ok {
		return servers.pick()
	}

	return types.None[slots.ControlServiceServer]()
}

type controlChannels struct {
	cs []slots.ControlServiceClient
}

func (c *controlChannels) add(client slots.ControlServiceClient) {
	c.cs = append(c.cs, client)
}

func (c *controlChannels) delete(client slots.ControlServiceClient) {
	for i := range c.cs {
		if c.cs[i] == client {
			c.cs = append(c.cs[:i], c.cs[i+1:]...)
			break
		}
	}
}

func (c *controlChannels) pick() types.Option[slots.ControlServiceClient] {
	if c == nil {
		return types.None[slots.ControlServiceClient]()
	}

	if len(c.cs) == 0 {
		return types.None[slots.ControlServiceClient]()
	}

	return types.Some(c.cs[len(c.cs)-1])
}

type controlServers struct {
	cs []slots.ControlServiceServer
}

func (c *controlServers) add(cc slots.ControlServiceServer) {
	c.cs = append(c.cs, cc)
}

func (c *controlServers) delete(cc slots.ControlServiceServer) {
	for i := range c.cs {
		if c.cs[i] == cc {
			c.cs = append(c.cs[:i], c.cs[i+1:]...)
			break
		}
	}
}

func (c *controlServers) pick() types.Option[slots.ControlServiceServer] {
	if c == nil {
		return types.None[slots.ControlServiceServer]()
	}

	if len(c.cs) == 0 {
		return types.None[slots.ControlServiceServer]()
	}

	return types.Some(c.cs[len(c.cs)-1])
}
