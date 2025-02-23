package emu

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/snple/beacon/consts"
	"github.com/snple/beacon/edge"
	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/edges"
	"github.com/snple/beacon/util"
	"github.com/snple/beacon/util/datatype"
	"go.uber.org/zap"
)

type EmuSlot struct {
	es *edge.EdgeService

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup

	dopts emuOptions
}

func Emu(es *edge.EdgeService, opts ...EmuOption) (*EmuSlot, error) {
	ctx, cancel := context.WithCancel(es.Context())

	gs := &EmuSlot{
		es:     es,
		ctx:    ctx,
		cancel: cancel,
		dopts:  defaultEmuOptions(),
	}

	for _, opt := range extraEmuOptions {
		opt.apply(&gs.dopts)
	}

	for _, opt := range opts {
		opt.apply(&gs.dopts)
	}

	return gs, nil
}

func (s *EmuSlot) Start() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	s.logger().Info("Emu slot started")

	ticker := time.NewTicker(s.dopts.tickerInterval)
	defer ticker.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			err := s.ticker()
			if err != nil {
				s.logger().Sugar().Errorf("Emu ticker: %v", err)
			}
		}
	}
}

func (s *EmuSlot) Stop() {
	s.cancel()
	s.closeWG.Wait()
}

func (s *EmuSlot) ticker() error {
	// t := time.Now().UnixMicro()

	offset := 0
	limit := 100

	for {
		request := edges.TagListRequest{
			Page: &pb.Page{
				Offset: uint32(offset),
				Limit:  uint32(limit),
			},
			Tags: "emu",
		}

		offset += limit

		reply, err := s.es.GetTag().List(s.ctx, &request)
		if err != nil {
			return err
		}

		{
			for _, tag := range reply.GetTag() {
				// if randomBool() {
				// 	continue
				// }

				if tag.GetStatus() != consts.ON {
					continue
				}

				if tag.GetAccess() == consts.ON {
					continue
				}

				if !datatype.DataType(tag.GetDataType()).IsNumber() {
					continue
				}

				min := float64(0)
				max := float64(100)

				if len(tag.LValue) > 0 {
					if lv, err := strconv.ParseFloat(tag.LValue, 64); err == nil {
						min = lv
					}
				}

				if len(tag.HValue) > 0 {
					if hv, err := strconv.ParseFloat(tag.HValue, 64); err == nil {
						max = hv
					}
				}

				value := util.RandNum(min, max)
				value = util.Round(value, 3)
				var value2 string

				switch datatype.DataType(tag.DataType) {
				case datatype.DataTypeI8, datatype.DataTypeI16, datatype.DataTypeI32:
					value2 = fmt.Sprintf("%v", int32(value))
				case datatype.DataTypeU8, datatype.DataTypeU16, datatype.DataTypeU32:
					value2 = fmt.Sprintf("%v", uint32(value))
				case datatype.DataTypeI64:
					value2 = fmt.Sprintf("%v", int64(value))
				case datatype.DataTypeU64:
					value2 = fmt.Sprintf("%v", uint64(value))
				case datatype.DataTypeF32:
					value2 = fmt.Sprintf("%v", float32(value))
				case datatype.DataTypeF64:
					value2 = fmt.Sprintf("%v", value)
				}

				if value2 != "" {
					s.logger().Sugar().Debugf("emu tag: %v, name: %v, value: %v, ", tag.GetId(), tag.GetName(), value2)
					_, err := s.es.GetTag().SyncValue(s.ctx, &pb.TagValue{Id: tag.GetId(), Value: value2, Updated: time.Now().UnixMicro()})
					if err != nil {
						s.logger().Sugar().Errorf("Emu SyncValue: %v", err)
					}
				}
			}
		}

		if len(reply.GetTag()) < limit {
			break
		}
	}

	return nil
}

func (s *EmuSlot) logger() *zap.Logger {
	return s.es.Logger()
}

type emuOptions struct {
	debug          bool
	tickerInterval time.Duration
}

func defaultEmuOptions() emuOptions {
	return emuOptions{
		debug:          false,
		tickerInterval: 60 * time.Second,
	}
}

type EmuOption interface {
	apply(*emuOptions)
}

var extraEmuOptions []EmuOption

type funcEmuOption struct {
	f func(*emuOptions)
}

func (fdo *funcEmuOption) apply(do *emuOptions) {
	fdo.f(do)
}

func newFuncEmuOption(f func(*emuOptions)) *funcEmuOption {
	return &funcEmuOption{
		f: f,
	}
}

func WithDebug(debug bool) EmuOption {
	return newFuncEmuOption(func(o *emuOptions) {
		o.debug = debug
	})
}

func WithTickerInterval(d time.Duration) EmuOption {
	return newFuncEmuOption(func(o *emuOptions) {
		o.tickerInterval = d
	})
}

func randomBool() bool {
	return rand.Intn(2) == 1 // 生成 0 或 1，然后做转换
}
