package slot

import (
	"github.com/snple/rgrpc"
	"snple.com/kokomi/pb/slots"
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

	s.ss.es.Logger().Sugar().Infof("rgrpc connect success, id: %v", slotID)

	client := slots.NewControlServiceClient(channel)

	s.ss.es.GetControl().AddSlotClient(slotID, client)
	defer s.ss.es.GetControl().DeleteSlotClient(slotID, client)

	defer s.ss.es.Logger().Sugar().Infof("rgrpc break connect, id: %v", slotID)

	<-channel.Done()
	return channel.Err()
}
