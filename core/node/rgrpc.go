package node

import (
	"github.com/snple/rgrpc"
	"snple.com/kokomi/pb/edges"
)

type RrpcService struct {
	ns *NodeService

	rgrpc.UnimplementedRgrpcServiceServer
}

var _ rgrpc.RgrpcServiceServer = (*RrpcService)(nil)

func newRrpcService(ns *NodeService) *RrpcService {
	return &RrpcService{ns: ns}
}

func (s *RrpcService) OpenRgrpc(stream rgrpc.RgrpcService_OpenRgrpcServer) error {
	channel := rgrpc.NewClient(stream)
	defer channel.Close()

	deviceID, err := validateToken(stream.Context())
	if err != nil {
		return err
	}

	s.ns.Logger().Sugar().Infof("rgrpc connect success, id: %v", deviceID)

	client := edges.NewControlServiceClient(channel)

	s.ns.cs.GetControl().AddClient(deviceID, client)
	defer s.ns.cs.GetControl().RemoveClient(deviceID, client)

	defer s.ns.Logger().Sugar().Infof("rgrpc break connect, id: %v", deviceID)

	<-channel.Done()
	return channel.Err()
}
