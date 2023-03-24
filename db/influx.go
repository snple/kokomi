package db

import (
	"context"
	"errors"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func ConnectInfluxDB(url, org, bucket, token string) (*InfluxDB, error) {
	client := influxdb2.NewClient(url, token)

	ready, err := client.Ready(context.Background())
	if err != nil {
		return nil, err
	}

	if *ready.Status != "ready" {
		client.Close()
		return nil, errors.New("client not ready")
	}

	reader := client.QueryAPI(org)
	query := fmt.Sprintf("from(bucket:\"%v\")|> range(start: -1h) |> filter(fn: (r) => r._measurement == \"123456\")", bucket)
	_, err = reader.Query(context.Background(), query)
	if err != nil {
		client.Close()
		return nil, err
	}

	return &InfluxDB{client, org, bucket}, nil
}

type InfluxDB struct {
	influxdb2.Client
	org    string
	bucket string
}

func (i *InfluxDB) WriteAPI() api.WriteAPI {
	return i.Client.WriteAPI(i.org, i.bucket)
}

func (i *InfluxDB) WriteAPIBlocking() api.WriteAPIBlocking {
	return i.Client.WriteAPIBlocking(i.org, i.bucket)
}

func (i *InfluxDB) Writer(batchSize int) *Writer {
	return newWriter(i.WriteAPIBlocking(), batchSize)
}

func (i *InfluxDB) QueryAPI() api.QueryAPI {
	return i.Client.QueryAPI(i.bucket)
}

func (i *InfluxDB) Bucket() string {
	return i.bucket
}

type Writer struct {
	batch     []*write.Point
	writeAPI  api.WriteAPIBlocking
	batchSize int
}

func (w *Writer) CurrentBatch() []*write.Point {
	return w.batch
}

func newWriter(writeAPI api.WriteAPIBlocking, batchSize int) *Writer {
	return &Writer{
		batch:     make([]*write.Point, 0, batchSize),
		writeAPI:  writeAPI,
		batchSize: batchSize,
	}
}

func (w *Writer) Write(ctx context.Context, p *write.Point) error {
	w.batch = append(w.batch, p)
	if len(w.batch) == w.batchSize {
		err := w.writeAPI.WritePoint(ctx, w.batch...)
		if err != nil {
			return err
		}
		w.batch = w.batch[:0]
	}
	return nil
}

func (w *Writer) Flush(ctx context.Context) error {
	if len(w.batch) > 0 {
		err := w.writeAPI.WritePoint(ctx, w.batch...)
		if err != nil {
			return err
		}
	}

	return nil
}
