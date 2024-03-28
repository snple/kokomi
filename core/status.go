package core

import (
	"github.com/snple/kokomi/consts"
	"github.com/snple/types"
	"github.com/snple/types/cache"
)

type StatusService struct {
	cs *CoreService

	link *cache.Cache[int32]
}

func newStatusService(cs *CoreService) *StatusService {
	return &StatusService{
		cs:   cs,
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
	s.link.Set(key, status, s.cs.dopts.linkTTL)
}
