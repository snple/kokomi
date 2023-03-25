package edge

import (
	"context"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"snple.com/kokomi/edge/model"
	"snple.com/kokomi/pb"
	"snple.com/kokomi/pb/edges"
	"snple.com/kokomi/pb/nodes"
)

func (s *NodeService) syncRemoteToLocal(ctx context.Context) error {
	deviceUpdated, err := s.SyncServiceClient().GetDeviceUpdated(ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	remoteDeviceUpdated, err := s.es.GetSync().getRemoteDeviceUpdated(ctx)
	if err != nil {
		return err
	}

	if deviceUpdated.GetUpdated() <= remoteDeviceUpdated.UnixMilli() {
		return nil
	}

	// device
	{
		remote, err := s.DeviceServiceClient().View(ctx, &pb.MyEmpty{})
		if err != nil {
			return err
		}

		local, err := s.es.GetDevice().View(ctx, &pb.MyEmpty{})
		if err != nil {
			if code, ok := status.FromError(err); ok {
				if code.Code() == codes.NotFound {
					_, err = s.es.GetDevice().Create(ctx, remote)
					if err != nil {
						return err
					}
				} else {
					return err
				}
			} else {
				return err
			}
		} else if remote.GetUpdated() > local.GetUpdated() {
			_, err = s.es.GetDevice().Update(ctx, remote)
			if err != nil {
				return err
			}
		}
	}

	// slot
	{
		after := remoteDeviceUpdated.UnixMilli()
		limit := uint32(10)

		for {
			remotes, err := s.SlotServiceClient().Pull(ctx, &nodes.PullSlotRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetSlot() {
				local, err := s.es.GetSlot().ViewWithDeleted(ctx, &pb.Id{Id: remote.GetId()})
				if err != nil {
					if code, ok := status.FromError(err); ok {
						if code.Code() == codes.NotFound {
							_, err = s.es.GetSlot().Create(ctx, remote)
							if err != nil {
								return err
							}
						} else {
							return err
						}
					} else {
						return err
					}
				} else if remote.GetUpdated() > local.GetUpdated() {
					_, err = s.es.GetSlot().Update(ctx, remote)
					if err != nil {
						return err
					}
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
		after := remoteDeviceUpdated.UnixMilli()
		limit := uint32(10)

		for {
			remotes, err := s.OptionServiceClient().Pull(ctx, &nodes.PullOptionRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetOption() {
				if !strings.HasPrefix(remote.GetName(), model.OPTION_PRIVATE_PREFIX) {
					local, err := s.es.GetOption().ViewWithDeleted(ctx, &pb.Id{Id: remote.GetId()})
					if err != nil {
						if code, ok := status.FromError(err); ok {
							if code.Code() == codes.NotFound {
								_, err = s.es.GetOption().Create(ctx, remote)
								if err != nil {
									return err
								}
							} else {
								return err
							}
						} else {
							return err
						}
					} else if remote.GetUpdated() > local.GetUpdated() {
						_, err = s.es.GetOption().Update(ctx, remote)
						if err != nil {
							return err
						}
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
		after := remoteDeviceUpdated.UnixMilli()
		limit := uint32(10)

		for {
			remotes, err := s.PortServiceClient().Pull(ctx, &nodes.PullPortRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetPort() {
				local, err := s.es.GetPort().ViewWithDeleted(ctx, &pb.Id{Id: remote.GetId()})
				if err != nil {
					if code, ok := status.FromError(err); ok {
						if code.Code() == codes.NotFound {
							_, err = s.es.GetPort().Create(ctx, remote)
							if err != nil {
								return err
							}
						} else {
							return err
						}
					} else {
						return err
					}
				} else if remote.GetUpdated() > local.GetUpdated() {
					_, err = s.es.GetPort().Update(ctx, remote)
					if err != nil {
						return err
					}
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
		after := remoteDeviceUpdated.UnixMilli()
		limit := uint32(10)

		for {
			remotes, err := s.ProxyServiceClient().Pull(ctx, &nodes.PullProxyRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetProxy() {
				local, err := s.es.GetProxy().ViewWithDeleted(ctx, &pb.Id{Id: remote.GetId()})
				if err != nil {
					if code, ok := status.FromError(err); ok {
						if code.Code() == codes.NotFound {
							_, err = s.es.GetProxy().Create(ctx, remote)
							if err != nil {
								return err
							}
						} else {
							return err
						}
					} else {
						return err
					}
				} else if remote.GetUpdated() > local.GetUpdated() {
					_, err = s.es.GetProxy().Update(ctx, remote)
					if err != nil {
						return err
					}
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
		after := remoteDeviceUpdated.UnixMilli()
		limit := uint32(10)

		for {
			remotes, err := s.SourceServiceClient().Pull(ctx, &nodes.PullSourceRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetSource() {
				local, err := s.es.GetSource().ViewWithDeleted(ctx, &pb.Id{Id: remote.GetId()})
				if err != nil {
					if code, ok := status.FromError(err); ok {
						if code.Code() == codes.NotFound {
							_, err = s.es.GetSource().Create(ctx, remote)
							if err != nil {
								return err
							}
						} else {
							return err
						}
					} else {
						return err
					}
				} else if remote.GetUpdated() > local.GetUpdated() {
					_, err = s.es.GetSource().Update(ctx, remote)
					if err != nil {
						return err
					}
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
		after := remoteDeviceUpdated.UnixMilli()
		limit := uint32(10)

		for {
			remotes, err := s.TagServiceClient().Pull(ctx, &nodes.PullTagRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetTag() {
				local, err := s.es.GetTag().ViewWithDeleted(ctx, &pb.Id{Id: remote.GetId()})
				if err != nil {
					if code, ok := status.FromError(err); ok {
						if code.Code() == codes.NotFound {
							_, err = s.es.GetTag().Create(ctx, remote)
							if err != nil {
								return err
							}
						} else {
							return err
						}
					} else {
						return err
					}
				} else if remote.GetUpdated() > local.GetUpdated() {
					_, err = s.es.GetTag().Update(ctx, remote)
					if err != nil {
						return err
					}
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetTag()) < int(limit) {
				break
			}
		}
	}

	// var
	{
		after := remoteDeviceUpdated.UnixMilli()
		limit := uint32(10)

		for {
			remotes, err := s.VarServiceClient().Pull(ctx, &nodes.PullVarRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetVar() {
				local, err := s.es.GetVar().ViewWithDeleted(ctx, &pb.Id{Id: remote.GetId()})
				if err != nil {
					if code, ok := status.FromError(err); ok {
						if code.Code() == codes.NotFound {
							_, err = s.es.GetVar().Create(ctx, remote)
							if err != nil {
								return err
							}
						} else {
							return err
						}
					} else {
						return err
					}
				} else if remote.GetUpdated() > local.GetUpdated() {
					_, err = s.es.GetVar().Update(ctx, remote)
					if err != nil {
						return err
					}
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetVar()) < int(limit) {
				break
			}
		}
	}

	// cable
	{
		after := remoteDeviceUpdated.UnixMilli()
		limit := uint32(10)

		for {
			remotes, err := s.CableServiceClient().Pull(ctx, &nodes.PullCableRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetCable() {
				local, err := s.es.GetCable().ViewWithDeleted(ctx, &pb.Id{Id: remote.GetId()})
				if err != nil {
					if code, ok := status.FromError(err); ok {
						if code.Code() == codes.NotFound {
							_, err = s.es.GetCable().Create(ctx, remote)
							if err != nil {
								return err
							}
						} else {
							return err
						}
					} else {
						return err
					}
				} else if remote.GetUpdated() > local.GetUpdated() {
					_, err = s.es.GetCable().Update(ctx, remote)
					if err != nil {
						return err
					}
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
		after := remoteDeviceUpdated.UnixMilli()
		limit := uint32(10)

		for {
			remotes, err := s.WireServiceClient().Pull(ctx, &nodes.PullWireRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetWire() {
				local, err := s.es.GetWire().ViewWithDeleted(ctx, &pb.Id{Id: remote.GetId()})
				if err != nil {
					if code, ok := status.FromError(err); ok {
						if code.Code() == codes.NotFound {
							_, err = s.es.GetWire().Create(ctx, remote)
							if err != nil {
								return err
							}
						} else {
							return err
						}
					} else {
						return err
					}
				} else if remote.GetUpdated() > local.GetUpdated() {
					_, err = s.es.GetWire().Update(ctx, remote)
					if err != nil {
						return err
					}
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetWire()) < int(limit) {
				break
			}
		}
	}

	return s.es.GetSync().setRemoteDeviceUpdated(ctx, time.UnixMilli(deviceUpdated.GetUpdated()))
}

func (s *NodeService) syncLocalToRemote(ctx context.Context) error {
	deviceUpdated, err := s.es.GetSync().getDeviceUpdated(ctx)
	if err != nil {
		return err
	}

	localDeviceUpdated, err := s.es.GetSync().getLocalDeviceUpdated(ctx)
	if err != nil {
		return err
	}

	if deviceUpdated.UnixMilli() <= localDeviceUpdated.UnixMilli() {
		return nil
	}

	// device
	{
		local, err := s.es.GetDevice().View(ctx, &pb.MyEmpty{})
		if err != nil {
			return err
		}

		remote, err := s.DeviceServiceClient().View(ctx, &pb.MyEmpty{})
		if err != nil {
			return err
		}

		if local.GetUpdated() > remote.GetUpdated() {
			_, err := s.DeviceServiceClient().Update(ctx, local)
			if err != nil {
				return err
			}
		}
	}

	// slot
	{
		after := localDeviceUpdated.UnixMilli()
		limit := uint32(10)

		for {
			locals, err := s.es.GetSlot().Pull(ctx, &edges.PullSlotRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetSlot() {
				remote, err := s.SlotServiceClient().ViewWithDeleted(ctx, &pb.Id{Id: local.GetId()})
				if err != nil {
					if code, ok := status.FromError(err); ok {
						if code.Code() == codes.NotFound {
							_, err = s.SlotServiceClient().Create(ctx, local)
							if err != nil {
								return err
							}
						} else {
							return err
						}
					} else {
						return err
					}
				} else if local.GetUpdated() > remote.GetUpdated() {
					_, err = s.SlotServiceClient().Update(ctx, local)
					if err != nil {
						return err
					}
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
		after := localDeviceUpdated.UnixMilli()
		limit := uint32(10)

		for {
			locals, err := s.es.GetOption().Pull(ctx, &edges.PullOptionRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetOption() {
				if !strings.HasPrefix(local.GetName(), model.OPTION_PRIVATE_PREFIX) {
					remote, err := s.OptionServiceClient().ViewWithDeleted(ctx, &pb.Id{Id: local.GetId()})
					if err != nil {
						if code, ok := status.FromError(err); ok {
							if code.Code() == codes.NotFound {
								_, err = s.OptionServiceClient().Create(ctx, local)
								if err != nil {
									return err
								}
							} else {
								return err
							}
						} else {
							return err
						}
					} else if local.GetUpdated() > remote.GetUpdated() {
						_, err = s.OptionServiceClient().Update(ctx, local)
						if err != nil {
							return err
						}
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
		after := localDeviceUpdated.UnixMilli()
		limit := uint32(10)

		for {
			locals, err := s.es.GetPort().Pull(ctx, &edges.PullPortRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetPort() {
				remote, err := s.PortServiceClient().ViewWithDeleted(ctx, &pb.Id{Id: local.GetId()})
				if err != nil {
					if code, ok := status.FromError(err); ok {
						if code.Code() == codes.NotFound {
							_, err = s.PortServiceClient().Create(ctx, local)
							if err != nil {
								return err
							}
						} else {
							return err
						}
					} else {
						return err
					}
				} else if local.GetUpdated() > remote.GetUpdated() {
					_, err = s.PortServiceClient().Update(ctx, local)
					if err != nil {
						return err
					}
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
		after := localDeviceUpdated.UnixMilli()
		limit := uint32(10)

		for {
			locals, err := s.es.GetSource().Pull(ctx, &edges.PullSourceRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetSource() {
				remote, err := s.SourceServiceClient().ViewWithDeleted(ctx, &pb.Id{Id: local.GetId()})
				if err != nil {
					if code, ok := status.FromError(err); ok {
						if code.Code() == codes.NotFound {
							_, err = s.SourceServiceClient().Create(ctx, local)
							if err != nil {
								return err
							}
						} else {
							return err
						}
					} else {
						return err
					}
				} else if local.GetUpdated() > remote.GetUpdated() {
					_, err = s.SourceServiceClient().Update(ctx, local)
					if err != nil {
						return err
					}
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
		after := localDeviceUpdated.UnixMilli()
		limit := uint32(10)

		for {
			locals, err := s.es.GetTag().Pull(ctx, &edges.PullTagRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetTag() {
				remote, err := s.TagServiceClient().ViewWithDeleted(ctx, &pb.Id{Id: local.GetId()})
				if err != nil {
					if code, ok := status.FromError(err); ok {
						if code.Code() == codes.NotFound {
							_, err = s.TagServiceClient().Create(ctx, local)
							if err != nil {
								return err
							}
						} else {
							return err
						}
					} else {
						return err
					}
				} else if local.GetUpdated() > remote.GetUpdated() {
					_, err = s.TagServiceClient().Update(ctx, local)
					if err != nil {
						return err
					}
				}

				after = local.GetUpdated()
			}

			if len(locals.GetTag()) < int(limit) {
				break
			}
		}
	}

	// var
	{
		after := localDeviceUpdated.UnixMilli()
		limit := uint32(10)

		for {
			locals, err := s.es.GetVar().Pull(ctx, &edges.PullVarRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetVar() {
				remote, err := s.VarServiceClient().ViewWithDeleted(ctx, &pb.Id{Id: local.GetId()})
				if err != nil {
					if code, ok := status.FromError(err); ok {
						if code.Code() == codes.NotFound {
							_, err = s.VarServiceClient().Create(ctx, local)
							if err != nil {
								return err
							}
						} else {
							return err
						}
					} else {
						return err
					}
				} else if local.GetUpdated() > remote.GetUpdated() {
					_, err = s.VarServiceClient().Update(ctx, local)
					if err != nil {
						return err
					}
				}

				after = local.GetUpdated()
			}

			if len(locals.GetVar()) < int(limit) {
				break
			}
		}
	}

	// cable
	{
		after := localDeviceUpdated.UnixMilli()
		limit := uint32(10)

		for {
			locals, err := s.es.GetCable().Pull(ctx, &edges.PullCableRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetCable() {
				remote, err := s.CableServiceClient().ViewWithDeleted(ctx, &pb.Id{Id: local.GetId()})
				if err != nil {
					if code, ok := status.FromError(err); ok {
						if code.Code() == codes.NotFound {
							_, err = s.CableServiceClient().Create(ctx, local)
							if err != nil {
								return err
							}
						} else {
							return err
						}
					} else {
						return err
					}
				} else if local.GetUpdated() > remote.GetUpdated() {
					_, err = s.CableServiceClient().Update(ctx, local)
					if err != nil {
						return err
					}
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
		after := localDeviceUpdated.UnixMilli()
		limit := uint32(10)

		for {
			locals, err := s.es.GetWire().Pull(ctx, &edges.PullWireRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetWire() {
				remote, err := s.WireServiceClient().ViewWithDeleted(ctx, &pb.Id{Id: local.GetId()})
				if err != nil {
					if code, ok := status.FromError(err); ok {
						if code.Code() == codes.NotFound {
							_, err = s.WireServiceClient().Create(ctx, local)
							if err != nil {
								return err
							}
						} else {
							return err
						}
					} else {
						return err
					}
				} else if local.GetUpdated() > remote.GetUpdated() {
					_, err = s.WireServiceClient().Update(ctx, local)
					if err != nil {
						return err
					}
				}

				after = local.GetUpdated()
			}

			if len(locals.GetWire()) < int(limit) {
				break
			}
		}
	}

	return s.es.GetSync().setLocalDeviceUpdated(ctx, deviceUpdated)
}

func (s *NodeService) syncRemoteToLocalTagValue(ctx context.Context) error {
	tagValueUpdated, err := s.SyncServiceClient().GetTagValueUpdated(ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	remoteTagValueUpdated, err := s.es.GetSync().getRemoteTagValueUpdated(ctx)
	if err != nil {
		return err
	}

	if tagValueUpdated.GetUpdated() <= remoteTagValueUpdated.UnixMilli() {
		return nil
	}

	after := remoteTagValueUpdated.UnixMilli()
	limit := uint32(10)

	for {
		remoteValues, err := s.TagServiceClient().PullValue(ctx, &nodes.PullTagValueRequest{After: after, Limit: limit})
		if err != nil {
			return err
		}

		for _, value := range remoteValues.GetTag() {
			_, err = s.es.GetTag().SetValue(ctx, &pb.TagValue{Id: value.GetId(), Value: value.GetValue()})
			if err != nil {
				s.es.Logger().Sugar().Errorf("SetValue: %v", err)
				return err
			}

			after = value.GetUpdated()
		}

		if len(remoteValues.GetTag()) < int(limit) {
			break
		}
	}

	return s.es.GetSync().setRemoteTagValueUpdated(ctx, time.UnixMilli(tagValueUpdated.GetUpdated()))
}

func (s *NodeService) syncRemoteToLocalWireValue(ctx context.Context) error {
	wireValueUpdated, err := s.SyncServiceClient().GetWireValueUpdated(ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	remoteWireValueUpdated, err := s.es.GetSync().getRemoteWireValueUpdated(ctx)
	if err != nil {
		return err
	}

	if wireValueUpdated.GetUpdated() <= remoteWireValueUpdated.UnixMilli() {
		return nil
	}

	after := remoteWireValueUpdated.UnixMilli()
	limit := uint32(10)

	for {
		remoteValues, err := s.WireServiceClient().PullValue(ctx, &nodes.PullWireValueRequest{After: after, Limit: limit})
		if err != nil {
			return err
		}

		for _, remote := range remoteValues.GetWire() {
			_, err = s.es.GetWire().SetValueUnchecked(ctx,
				&pb.WireValue{Id: remote.GetId(), Value: remote.GetValue(), Updated: remote.Updated})
			if err != nil {
				s.es.Logger().Sugar().Errorf("SetWireValue: %v", err)
				return err
			}

			after = remote.GetUpdated()
		}

		if len(remoteValues.GetWire()) < int(limit) {
			break
		}
	}

	return s.es.GetSync().setRemoteWireValueUpdated(ctx, time.UnixMilli(wireValueUpdated.GetUpdated()))
}

func (s *NodeService) syncLocalToRemoteWireValue(ctx context.Context) error {
	wireValueUpdated, err := s.es.GetSync().getWireValueUpdated(ctx)
	if err != nil {
		return err
	}

	localWireValueUpdated, err := s.es.GetSync().getLocalWireValueUpdated(ctx)
	if err != nil {
		return err
	}

	if wireValueUpdated.UnixMilli() <= localWireValueUpdated.UnixMilli() {
		return nil
	}

	after := localWireValueUpdated.UnixMilli()
	limit := uint32(10)

	for {
		locals, err := s.es.GetWire().PullValue(ctx, &edges.PullWireValueRequest{After: after, Limit: limit})
		if err != nil {
			return err
		}

		for _, local := range locals.GetWire() {
			_, err = s.WireServiceClient().SetValueUnchecked(ctx,
				&pb.WireValue{Id: local.GetId(), Value: local.GetValue(), Updated: local.GetUpdated()})
			if err != nil {
				s.es.Logger().Sugar().Errorf("SetWireValue: %v", err)
				return err
			}

			after = local.GetUpdated()
		}

		if len(locals.GetWire()) < int(limit) {
			break
		}
	}

	return s.es.GetSync().setLocalWireValueUpdated(ctx, wireValueUpdated)
}
