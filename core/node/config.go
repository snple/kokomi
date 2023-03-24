package node

import (
	"crypto/tls"
	"time"

	"github.com/quic-go/quic-go"
)

type nodeOptions struct {
	quicOptions      *quicOptions
	quicPingInterval time.Duration
}

type quicOptions struct {
	Addr       string
	TLSConfig  *tls.Config
	QUICConfig *quic.Config
}

func defaultNodeOptions() nodeOptions {
	return nodeOptions{
		quicPingInterval: 60 * time.Second,
	}
}

type NodeOption interface {
	apply(*nodeOptions)
}

var extraNodeOptions []NodeOption

type funcNodeOption struct {
	f func(*nodeOptions)
}

func (fdo *funcNodeOption) apply(do *nodeOptions) {
	fdo.f(do)
}

func newFuncNodeOption(f func(*nodeOptions)) *funcNodeOption {
	return &funcNodeOption{
		f: f,
	}
}

func WithQuic(addr string, tlsConfig *tls.Config, quicConfig *quic.Config) NodeOption {
	return newFuncNodeOption(func(o *nodeOptions) {
		tlsConfig2 := tlsConfig.Clone()
		if len(tlsConfig2.NextProtos) == 0 {
			tlsConfig2.NextProtos = []string{"kokomi"}
		}

		quicConfig2 := quicConfig.Clone()
		quicConfig2.EnableDatagrams = true

		o.quicOptions = &quicOptions{addr, tlsConfig2, quicConfig2}
	})
}

func WithQuicPingInterval(d time.Duration) NodeOption {
	return newFuncNodeOption(func(o *nodeOptions) {
		o.quicPingInterval = d
	})
}
