package core

import (
	"context"
	"sync"

	"github.com/snple/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"snple.com/kokomi/consts"
	"snple.com/kokomi/pb"
	"snple.com/kokomi/pb/cores"
	"snple.com/kokomi/pb/edges"
	"snple.com/kokomi/util/datatype"
)

type ControlService struct {
	cs *CoreService

	lock    sync.RWMutex
	clients map[string]*channels

	cores.UnimplementedControlServiceServer
}

func newControlService(cs *CoreService) *ControlService {
	return &ControlService{
		cs:      cs,
		clients: make(map[string]*channels),
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
	item, err := s.cs.GetTag().view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	if item.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag Status != ON")
	}

	// validation device and source
	{
		// device
		{
			device, err := s.cs.GetDevice().view(ctx, item.DeviceID)
			if err != nil {
				return &output, err
			}

			if device.Status != consts.ON {
				return &output, status.Errorf(codes.FailedPrecondition, "Device Status != ON")
			}
		}

		// source
		{
			source, err := s.cs.GetSource().view(ctx, item.SourceID)
			if err != nil {
				return &output, err
			}

			if source.Status != consts.ON {
				return &output, status.Errorf(codes.FailedPrecondition, "Source Status != ON")
			}
		}
	}

	if option := s.GetClient(item.DeviceID); option.IsSome() {
		client := option.Unwrap()

		return client.GetTagValue(ctx, in)
	}

	return &output, status.Errorf(codes.Unavailable, "Device not connect")
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
	item, err := s.cs.GetTag().view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	if item.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag Status != ON")
	}

	if item.Access != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag Access != ON")
	}

	_, err = datatype.DecodeNsonValue(in.GetValue(), item.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	// validation device and source
	{
		// device
		{
			device, err := s.cs.GetDevice().view(ctx, item.DeviceID)
			if err != nil {
				return &output, err
			}

			if device.Status != consts.ON {
				return &output, status.Errorf(codes.FailedPrecondition, "Device Status != ON")
			}
		}

		// source
		{
			source, err := s.cs.GetSource().view(ctx, item.SourceID)
			if err != nil {
				return &output, err
			}

			if source.Status != consts.ON {
				return &output, status.Errorf(codes.FailedPrecondition, "Source Status != ON")
			}
		}
	}

	if option := s.GetClient(item.DeviceID); option.IsSome() {
		client := option.Unwrap()

		return client.SetTagValue(ctx, in)
	}

	return &output, status.Errorf(codes.Unavailable, "Device not connect")
}

func (s *ControlService) AddClient(deviceID string, client edges.ControlServiceClient) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if chans, ok := s.clients[deviceID]; ok {
		chans.add(client)
	} else {
		chans := &channels{}
		chans.add(client)
		s.clients[deviceID] = chans
	}
}

func (s *ControlService) RemoveClient(deviceID string, client edges.ControlServiceClient) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if chans, ok := s.clients[deviceID]; ok {
		chans.remove(client)
	}
}

func (s *ControlService) GetClient(deviceID string) types.Option[edges.ControlServiceClient] {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if chans, ok := s.clients[deviceID]; ok {
		return chans.pick()
	}

	return types.None[edges.ControlServiceClient]()
}

type channels struct {
	chans []edges.ControlServiceClient
}

func (c *channels) add(ch edges.ControlServiceClient) {
	c.chans = append(c.chans, ch)
}

func (c *channels) remove(ch edges.ControlServiceClient) {
	for i := range c.chans {
		if c.chans[i] == ch {
			c.chans = append(c.chans[:i], c.chans[i+1:]...)
			break
		}
	}
}

func (c *channels) pick() types.Option[edges.ControlServiceClient] {
	if c == nil {
		return types.None[edges.ControlServiceClient]()
	}

	if len(c.chans) == 0 {
		return types.None[edges.ControlServiceClient]()
	}

	return types.Some(c.chans[len(c.chans)-1])
}
