package edge

import (
	"github.com/snple/beacon/consts"
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

func (s *StatusService) GetNodeLink() int32 {
	if v := s.link.Get("node.ID"); v.IsSome() {
		return v.Unwrap()
	}

	return consts.OFF
}

func (s *StatusService) GetNodeLinkValue() types.Option[cache.Value[int32]] {
	return s.link.GetValue("node.ID")
}

func (s *StatusService) SetNodeLink(status int32) {
	s.link.Set("node.ID", status, 0)
}
