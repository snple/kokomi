package zstd

import (
	"io"
	"sync"

	"github.com/klauspost/compress/zstd"
	"google.golang.org/grpc/encoding"
)

// Name is the name registered for the zstd compressor.
const Name = "zstd"

var eopts []zstd.EOption

func Register(opts ...zstd.EOption) {
	eopts = opts

	encoding.RegisterCompressor(&compressor{})
}

type writer struct {
	*zstd.Encoder
	pool *sync.Pool
}

func (c *compressor) Compress(w io.Writer) (io.WriteCloser, error) {
	z, inPool := c.poolCompressor.Get().(*writer)
	if !inPool {
		enc, err := zstd.NewWriter(io.Discard, eopts...)
		if err != nil {
			return nil, err
		}
		return &writer{Encoder: enc, pool: &c.poolCompressor}, nil
	}

	z.Encoder.Reset(w)
	return z, nil
}

func (z *writer) Close() error {
	defer z.pool.Put(z)
	return z.Encoder.Close()
}

type reader struct {
	*zstd.Decoder
	pool *sync.Pool
}

func (c *compressor) Decompress(r io.Reader) (io.Reader, error) {
	z, inPool := c.poolDecompressor.Get().(*reader)
	if !inPool {
		dec, err := zstd.NewReader(r)
		if err != nil {
			return nil, err
		}
		return &reader{Decoder: dec, pool: &c.poolDecompressor}, nil
	}
	if err := z.Reset(r); err != nil {
		c.poolDecompressor.Put(z)
		return nil, err
	}
	return z, nil
}

func (z *reader) Read(p []byte) (n int, err error) {
	n, err = z.Decoder.Read(p)
	if err == io.EOF {
		z.pool.Put(z)
	}
	return n, err
}

func (c *compressor) Name() string {
	return Name
}

type compressor struct {
	poolCompressor   sync.Pool
	poolDecompressor sync.Pool
}
