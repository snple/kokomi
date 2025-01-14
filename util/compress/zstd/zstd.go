package zstd

import (
	"fmt"
	"io"
	"sync"

	"github.com/klauspost/compress/zstd"
	"google.golang.org/grpc/encoding"
)

// Name is the name registered for the zstd compressor.
const Name = "zstd"

func init() {
	c := &compressor{}
	c.poolCompressor.New = func() any {
		enc, err := zstd.NewWriter(io.Discard)
		if err != nil {
			panic(err)
		}
		return &writer{Encoder: enc, pool: &c.poolCompressor}
	}
	encoding.RegisterCompressor(c)
}

type writer struct {
	*zstd.Encoder
	pool *sync.Pool
}

func SetLevel(level zstd.EncoderLevel) error {
	if level < zstd.SpeedFastest || level > zstd.SpeedBestCompression {
		return fmt.Errorf("grpc: invalid zstd compression level: %d", level)
	}
	c := encoding.GetCompressor(Name).(*compressor)
	c.poolCompressor.New = func() any {
		enc, err := zstd.NewWriter(io.Discard, zstd.WithEncoderLevel(level))
		if err != nil {
			panic(err)
		}
		return &writer{Encoder: enc, pool: &c.poolCompressor}
	}
	return nil
}

func (c *compressor) Compress(w io.Writer) (io.WriteCloser, error) {
	z := c.poolCompressor.Get().(*writer)
	z.Encoder.Reset(w)
	return z, nil
}

func (w *writer) Close() error {
	defer w.pool.Put(w)
	return w.Encoder.Close()
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
