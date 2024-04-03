package core

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/core/model"
	"github.com/snple/kokomi/pb/cores"
	"github.com/snple/kokomi/util/datatype"
)

type SaveService struct {
	cs *CoreService

	tagValueUpdated int64
	lock            sync.RWMutex

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup
}

func newSaveService(cs *CoreService) *SaveService {
	ctx, cancel := context.WithCancel(cs.Context())

	s := &SaveService{
		cs:              cs,
		tagValueUpdated: time.Now().UnixMicro(),
		ctx:             ctx,
		cancel:          cancel,
	}

	return s
}

func (s *SaveService) start() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	if option := s.cs.GetInfluxDB(); option.IsNone() {
		return
	}

	s.cs.Logger().Info("start save service")

	ticker := time.NewTicker(s.cs.dopts.saveInterval)
	defer ticker.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			s.ticker()
		}
	}
}

func (s *SaveService) stop() {
	s.cancel()
	s.closeWG.Wait()
}

func (s *SaveService) ticker() {
	if err := s.checkTagValueUpdated(); err != nil {
		s.cs.Logger().Sugar().Errorf("checkTagValueUpdated: %v", err)
	}
}

func (s *SaveService) checkTagValueUpdated() error {
	valueUpdated := s.getTagValueUpdated()

	// InfluxDB writer
	option := s.cs.GetInfluxDB()
	if option.IsNone() {
		panic("influxdb not enable")
	}

	writer := option.Unwrap().Writer(50)

	after := valueUpdated
	limit := uint32(100)

	now := time.Now().UnixMicro()

PULL:
	for {
		remotes, err := s.cs.GetTag().PullValue(s.ctx,
			&cores.TagPullValueRequest{After: after, Limit: limit})
		if err != nil {
			return err
		}

		for _, remote := range remotes.GetTag() {
			if remote.GetUpdated() > now {
				break PULL
			}

			after = remote.GetUpdated()

			item, err := s.cs.GetTag().ViewByID(s.ctx, remote.GetId())
			if err != nil {
				return err
			}

			if item.Status != consts.ON || item.Save != consts.ON {
				continue
			}

			{
				device, err := s.cs.GetDevice().ViewFromCacheByID(s.ctx, remote.GetDeviceId())
				if err != nil {
					return err
				}

				if device.Status != consts.ON {
					continue
				}
			}

			{
				source, err := s.cs.GetSource().ViewFromCacheByID(s.ctx, remote.GetSourceId())
				if err != nil {
					return err
				}

				if source.Status != consts.ON || source.Save != consts.ON {
					continue
				}
			}

			value := float64(0)
			if datatype.DataType(item.DataType).IsNumber() {
				value, err = strconv.ParseFloat(remote.GetValue(), 64)
				if err != nil {
					return err
				}
			} else if datatype.DataType(item.DataType) == datatype.DataTypeBool {
				if remote.GetValue() == "true" || remote.GetValue() == "1" {
					value = 1
				}
			} else {
				continue
			}

			point := s.newTagPoint(&item, value, time.UnixMicro(remote.GetUpdated()).Unix())

			err = writer.Write(s.ctx, point)
			if err != nil {
				return err
			}
		}

		if len(remotes.GetTag()) < int(limit) {
			break
		}
	}

	err := writer.Flush(s.ctx)
	if err != nil {
		return err
	}

	s.setTagValueUpdated(after)

	return nil
}

func (s *SaveService) getTagValueUpdated() int64 {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.tagValueUpdated
}

func (s *SaveService) setTagValueUpdated(updated int64) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.tagValueUpdated = updated
}

func (s *SaveService) newTagPoint(tag *model.Tag, value float64, timestamp int64) *write.Point {
	return write.NewPoint(
		tag.DeviceID,
		map[string]string{
			"source": tag.SourceID,
		},
		map[string]interface{}{
			tag.ID: value,
		},
		time.Unix(timestamp, 0),
	)
}
