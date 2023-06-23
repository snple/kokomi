package edge

import (
	"context"
	"strings"
	"time"

	"github.com/snple/kokomi/edge/model"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/pb/nodes"
)

func (s *NodeService) syncRemoteToLocal(ctx context.Context) error {
	deviceUpdated, err := s.SyncServiceClient().GetDeviceUpdated(ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	deviceUpdated2, err := s.es.GetSync().getDeviceUpdatedRemoteToLocal(ctx)
	if err != nil {
		return err
	}

	if deviceUpdated.GetUpdated() <= deviceUpdated2.UnixMicro() {
		return nil
	}

	// device
	{
		remote, err := s.DeviceServiceClient().ViewWithDeleted(ctx, &pb.MyEmpty{})
		if err != nil {
			return err
		}

		_, err = s.es.GetDevice().Sync(ctx, remote)
		if err != nil {
			return err
		}
	}

	// slot
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.SlotServiceClient().Pull(ctx, &nodes.PullSlotRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetSlot() {
				_, err := s.es.GetSlot().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetSlot()) < int(limit) {
				break
			}
		}
	}

	// option
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.OptionServiceClient().Pull(ctx, &nodes.PullOptionRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetOption() {
				if !strings.HasPrefix(remote.GetName(), model.OPTION_PRIVATE_PREFIX) {
					_, err := s.es.GetOption().Sync(ctx, remote)
					if err != nil {
						return err
					}
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetOption()) < int(limit) {
				break
			}
		}
	}

	// port
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.PortServiceClient().Pull(ctx, &nodes.PullPortRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetPort() {
				_, err := s.es.GetPort().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetPort()) < int(limit) {
				break
			}
		}
	}

	// proxy
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.ProxyServiceClient().Pull(ctx, &nodes.PullProxyRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetProxy() {
				_, err = s.es.GetProxy().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetProxy()) < int(limit) {
				break
			}
		}
	}

	// source
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.SourceServiceClient().Pull(ctx, &nodes.PullSourceRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetSource() {
				_, err = s.es.GetSource().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetSource()) < int(limit) {
				break
			}
		}
	}

	// tag
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.TagServiceClient().Pull(ctx, &nodes.PullTagRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetTag() {
				_, err := s.es.GetTag().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetTag()) < int(limit) {
				break
			}
		}
	}

	// const
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.ConstServiceClient().Pull(ctx, &nodes.PullConstRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetConst() {
				_, err := s.es.GetConst().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetConst()) < int(limit) {
				break
			}
		}
	}

	// cable
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.CableServiceClient().Pull(ctx, &nodes.PullCableRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetCable() {
				_, err := s.es.GetCable().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetCable()) < int(limit) {
				break
			}
		}
	}

	// wire
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.WireServiceClient().Pull(ctx, &nodes.PullWireRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetWire() {
				_, err := s.es.GetWire().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetWire()) < int(limit) {
				break
			}
		}
	}

	return s.es.GetSync().setDeviceUpdatedRemoteToLocal(ctx, time.UnixMicro(deviceUpdated.GetUpdated()))
}

func (s *NodeService) syncLocalToRemote(ctx context.Context) error {
	deviceUpdated, err := s.es.GetSync().getDeviceUpdated(ctx)
	if err != nil {
		return err
	}

	deviceUpdated2, err := s.es.GetSync().getDeviceUpdatedLocalToRemote(ctx)
	if err != nil {
		return err
	}

	if deviceUpdated.UnixMicro() <= deviceUpdated2.UnixMicro() {
		return nil
	}

	// device
	{
		local, err := s.es.GetDevice().ViewWithDeleted(ctx, &pb.MyEmpty{})
		if err != nil {
			return err
		}

		_, err = s.DeviceServiceClient().Sync(ctx, local)
		if err != nil {
			return err
		}
	}

	// slot
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetSlot().Pull(ctx, &edges.PullSlotRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetSlot() {
				_, err = s.SlotServiceClient().Sync(ctx, local)
				if err != nil {
					return err
				}

				after = local.GetUpdated()
			}

			if len(locals.GetSlot()) < int(limit) {
				break
			}
		}
	}

	// option
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetOption().Pull(ctx, &edges.PullOptionRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetOption() {
				if !strings.HasPrefix(local.GetName(), model.OPTION_PRIVATE_PREFIX) {
					_, err = s.OptionServiceClient().Sync(ctx, local)
					if err != nil {
						return err
					}
				}

				after = local.GetUpdated()
			}

			if len(locals.GetOption()) < int(limit) {
				break
			}
		}
	}

	// port
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetPort().Pull(ctx, &edges.PullPortRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetPort() {
				_, err = s.PortServiceClient().Sync(ctx, local)
				if err != nil {
					return err
				}

				after = local.GetUpdated()
			}

			if len(locals.GetPort()) < int(limit) {
				break
			}
		}
	}

	// source
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetSource().Pull(ctx, &edges.PullSourceRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetSource() {
				_, err = s.SourceServiceClient().Sync(ctx, local)
				if err != nil {
					return err
				}

				after = local.GetUpdated()
			}

			if len(locals.GetSource()) < int(limit) {
				break
			}
		}
	}

	// tag
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetTag().Pull(ctx, &edges.PullTagRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetTag() {
				_, err = s.TagServiceClient().Sync(ctx, local)
				if err != nil {
					return err
				}

				after = local.GetUpdated()
			}

			if len(locals.GetTag()) < int(limit) {
				break
			}
		}
	}

	// const
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetConst().Pull(ctx, &edges.PullConstRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetConst() {
				_, err = s.ConstServiceClient().Sync(ctx, local)
				if err != nil {
					return err
				}

				after = local.GetUpdated()
			}

			if len(locals.GetConst()) < int(limit) {
				break
			}
		}
	}

	// cable
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetCable().Pull(ctx, &edges.PullCableRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetCable() {
				_, err = s.CableServiceClient().Sync(ctx, local)
				if err != nil {
					return err
				}

				after = local.GetUpdated()
			}

			if len(locals.GetCable()) < int(limit) {
				break
			}
		}
	}

	// wire
	{
		after := deviceUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetWire().Pull(ctx, &edges.PullWireRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetWire() {
				_, err = s.WireServiceClient().Sync(ctx, local)
				if err != nil {
					return err
				}

				after = local.GetUpdated()
			}

			if len(locals.GetWire()) < int(limit) {
				break
			}
		}
	}

	return s.es.GetSync().setDeviceUpdatedLocalToRemote(ctx, deviceUpdated)
}

func (s *NodeService) syncTagValueRemoteToLocal(ctx context.Context) error {
	tagValueUpdated, err := s.SyncServiceClient().GetTagValueUpdated(ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	tagValueUpdated2, err := s.es.GetSync().getTagValueUpdatedRemoteToLocal(ctx)
	if err != nil {
		return err
	}

	if tagValueUpdated.GetUpdated() <= tagValueUpdated2.UnixMicro() {
		return nil
	}

	after := tagValueUpdated2.UnixMicro()
	limit := uint32(100)

PULL:
	for {
		remotes, err := s.TagServiceClient().PullValue(ctx, &nodes.PullTagValueRequest{After: after, Limit: limit})
		if err != nil {
			return err
		}

		for _, remote := range remotes.GetTag() {
			if remote.GetUpdated() > tagValueUpdated.GetUpdated() {
				break PULL
			}

			_, err = s.es.GetTag().SyncValue(ctx,
				&pb.TagValue{Id: remote.GetId(), Value: remote.GetValue(), Updated: remote.GetUpdated()})
			if err != nil {
				s.es.Logger().Sugar().Errorf("SyncValue: %v", err)
				return err
			}

			after = remote.GetUpdated()
		}

		if len(remotes.GetTag()) < int(limit) {
			break
		}
	}

	return s.es.GetSync().setTagValueUpdatedRemoteToLocal(ctx, time.UnixMicro(tagValueUpdated.GetUpdated()))
}

func (s *NodeService) syncTagValueLocalToRemote(ctx context.Context) error {
	tagValueUpdated, err := s.es.GetSync().getTagValueUpdated(ctx)
	if err != nil {
		return err
	}

	tagValueUpdated2, err := s.es.GetSync().getTagValueUpdatedLocalToRemote(ctx)
	if err != nil {
		return err
	}

	if tagValueUpdated.UnixMicro() <= tagValueUpdated2.UnixMicro() {
		return nil
	}

	after := tagValueUpdated2.UnixMicro()
	limit := uint32(100)

PULL:
	for {
		locals, err := s.es.GetTag().PullValue(ctx, &edges.PullTagValueRequest{After: after, Limit: limit})
		if err != nil {
			return err
		}

		for _, local := range locals.GetTag() {
			if local.GetUpdated() > tagValueUpdated.UnixMicro() {
				break PULL
			}

			_, err = s.TagServiceClient().SyncValue(ctx,
				&pb.TagValue{Id: local.GetId(), Value: local.GetValue(), Updated: local.GetUpdated()})
			if err != nil {
				s.es.Logger().Sugar().Errorf("SyncValue: %v", err)
				return err
			}

			after = local.GetUpdated()
		}

		if len(locals.GetTag()) < int(limit) {
			break
		}
	}

	return s.es.GetSync().setTagValueUpdatedLocalToRemote(ctx, tagValueUpdated)
}

func (s *NodeService) syncWireValueRemoteToLocal(ctx context.Context) error {
	wireValueUpdated, err := s.SyncServiceClient().GetWireValueUpdated(ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	wireValueUpdated2, err := s.es.GetSync().getWireValueUpdatedRemoteToLocal(ctx)
	if err != nil {
		return err
	}

	if wireValueUpdated.GetUpdated() <= wireValueUpdated2.UnixMicro() {
		return nil
	}

	after := wireValueUpdated2.UnixMicro()
	limit := uint32(100)

PULL:
	for {
		remotes, err := s.WireServiceClient().PullValue(ctx, &nodes.PullWireValueRequest{After: after, Limit: limit})
		if err != nil {
			return err
		}

		for _, remote := range remotes.GetWire() {
			if remote.GetUpdated() > wireValueUpdated.GetUpdated() {
				break PULL
			}

			_, err = s.es.GetWire().SyncValue(ctx,
				&pb.WireValue{Id: remote.GetId(), Value: remote.GetValue(), Updated: remote.GetUpdated()})
			if err != nil {
				s.es.Logger().Sugar().Errorf("SyncValue: %v", err)
				return err
			}

			after = remote.GetUpdated()
		}

		if len(remotes.GetWire()) < int(limit) {
			break
		}
	}

	return s.es.GetSync().setWireValueUpdatedRemoteToLocal(ctx, time.UnixMicro(wireValueUpdated.GetUpdated()))
}

func (s *NodeService) syncWireValueLocalToRemote(ctx context.Context) error {
	wireValueUpdated, err := s.es.GetSync().getWireValueUpdated(ctx)
	if err != nil {
		return err
	}

	wireValueUpdated2, err := s.es.GetSync().getWireValueUpdatedLocalToRemote(ctx)
	if err != nil {
		return err
	}

	if wireValueUpdated.UnixMicro() <= wireValueUpdated2.UnixMicro() {
		return nil
	}

	after := wireValueUpdated2.UnixMicro()
	limit := uint32(100)

PULL:
	for {
		locals, err := s.es.GetWire().PullValue(ctx, &edges.PullWireValueRequest{After: after, Limit: limit})
		if err != nil {
			return err
		}

		for _, local := range locals.GetWire() {
			if local.GetUpdated() > wireValueUpdated.UnixMicro() {
				break PULL
			}

			_, err = s.WireServiceClient().SyncValue(ctx,
				&pb.WireValue{Id: local.GetId(), Value: local.GetValue(), Updated: local.GetUpdated()})
			if err != nil {
				s.es.Logger().Sugar().Errorf("SyncValue: %v", err)
				return err
			}

			after = local.GetUpdated()
		}

		if len(locals.GetWire()) < int(limit) {
			break
		}
	}

	return s.es.GetSync().setWireValueUpdatedLocalToRemote(ctx, wireValueUpdated)
}
