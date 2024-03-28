package slot

import (
	"github.com/snple/kokomi/pb/slots"
	"github.com/snple/kokomi/util/metadata"
	"github.com/snple/rgrpc"
)

type RgrpcService struct {
	ss *SlotService

	rgrpc.UnimplementedRgrpcServiceServer
}

func newRgrpcService(ss *SlotService) *RgrpcService {
	return &RgrpcService{
		ss: ss,
	}
}

var _ rgrpc.RgrpcServiceServer = (*RgrpcService)(nil)

func (s *RgrpcService) OpenRgrpc(stream rgrpc.RgrpcService_OpenRgrpcServer) error {
	channel := rgrpc.NewClient(stream)
	defer channel.Close()

	slotID, err := validateToken(stream.Context())
	if err != nil {
		return err
	}

	s.ss.Edge().Logger().Sugar().Infof("rgrpc connect success, id: %v, ip: %v",
		slotID, metadata.GetPeerAddr(stream.Context()))

	client := slots.NewControlServiceClient(channel)

	s.ss.Edge().GetControl().AddSlotClient(slotID, client)
	defer s.ss.Edge().GetControl().DeleteSlotClient(slotID, client)

	defer s.ss.Edge().Logger().Sugar().Infof("rgrpc break connect, id: %v, ip: %v",
		slotID, metadata.GetPeerAddr(stream.Context()))

	<-channel.Done()
	return channel.Err()
}
