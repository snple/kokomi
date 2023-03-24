package core

import (
	"log"
	"time"

	"go.uber.org/zap"
	"snple.com/kokomi/db"
)

type coreOptions struct {
	influxdb      *db.InfluxDB
	linkStatusTTL time.Duration
	valueCacheTTL time.Duration
	logger        *zap.Logger
}

func defaultCoreOptions() coreOptions {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("zap.NewDevelopment(): %v", err)
	}

	return coreOptions{
		linkStatusTTL: 3 * time.Minute,
		valueCacheTTL: 3 * time.Minute,
		logger:        logger,
	}
}

type CoreOption interface {
	apply(*coreOptions)
}

var extraCoreOptions []CoreOption

type funcCoreOption struct {
	f func(*coreOptions)
}

func (fdo *funcCoreOption) apply(do *coreOptions) {
	fdo.f(do)
}

func newFuncCoreOption(f func(*coreOptions)) *funcCoreOption {
	return &funcCoreOption{
		f: f,
	}
}

func WithInfluxDB(influxdb *db.InfluxDB) CoreOption {
	return newFuncCoreOption(func(o *coreOptions) {
		o.influxdb = influxdb
	})
}

func WithLinkStatusTTL(d time.Duration) CoreOption {
	return newFuncCoreOption(func(o *coreOptions) {
		o.linkStatusTTL = d
	})
}

func WithValueCacheTTL(d time.Duration) CoreOption {
	return newFuncCoreOption(func(o *coreOptions) {
		o.valueCacheTTL = d
	})
}

func WithLogger(logger *zap.Logger) CoreOption {
	return newFuncCoreOption(func(o *coreOptions) {
		o.logger = logger
	})
}
