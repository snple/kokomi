package edge

import (
	"github.com/snple/kokomi/consts"
	"github.com/snple/types"
	"github.com/snple/types/cache"
)

type StatusService struct {
	es *EdgeService

	link *cache.Cache[int32]
}

func newStatusService(es *EdgeService) *StatusService {
	return &StatusService{
		es:   es,
		link: cache.NewCache[int32](nil),
	}
}

func (s *StatusService) GetLink(key string) int32 {
	if v := s.link.Get(key); v.IsSome() {
		return v.Unwrap()
	}

	return consts.OFF
}

func (s *StatusService) GetLinkValue(key string) types.Option[cache.Value[int32]] {
	return s.link.GetValue(key)
}

func (s *StatusService) SetLink(key string, status int32) {
	s.link.Set(key, status, s.es.dopts.linkTTL)
}

func (s *StatusService) GetDeviceLink() int32 {
	if v := s.link.Get("device.ID"); v.IsSome() {
		return v.Unwrap()
	}

	return consts.OFF
}

func (s *StatusService) GetDeviceLinkValue() types.Option[cache.Value[int32]] {
	return s.link.GetValue("device.ID")
}

func (s *StatusService) SetDeviceLink(status int32) {
	s.link.Set("device.ID", status, 0)
}
