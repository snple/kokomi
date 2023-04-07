package status

import (
	"time"

	"github.com/danclive/nson-go"
	"github.com/snple/kokomi/consts"
	"github.com/snple/types"
	"github.com/snple/types/cache"
)

type StatusService struct {
	tag  *cache.Cache[nson.Value]
	link *cache.Cache[int32]
}

func NewStatusService() *StatusService {
	return &StatusService{
		tag:  cache.NewCache[nson.Value](nil),
		link: cache.NewCache[int32](nil),
	}
}

func (s *StatusService) TagGet(key string) types.Option[nson.Value] {
	return s.tag.Get(key)
}
func (s *StatusService) TagGetValue(key string) types.Option[cache.Value[nson.Value]] {
	return s.tag.GetValue(key)
}

func (s *StatusService) TagSet(key string, value nson.Value, ttl time.Duration) {
	s.tag.Set(key, value, ttl)
}

func (s *StatusService) LinkGet(key string) int32 {
	if v := s.link.Get(key); v.IsSome() {
		return v.Unwrap()
	}

	return consts.OFF
}

func (s *StatusService) LinkGetValue(key string) types.Option[cache.Value[int32]] {
	return s.link.GetValue(key)
}

func (s *StatusService) LinkSet(key string, status int32, ttl time.Duration) {
	s.link.Set(key, status, ttl)
}
