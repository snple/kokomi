package edge

import (
	"bytes"
	"context"
	"encoding/json"
	"text/template"
	"time"

	"github.com/danclive/nson-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"snple.com/kokomi/edge/model"
	"snple.com/kokomi/pb"
	"snple.com/kokomi/pb/edges"
	"snple.com/kokomi/util"
)

type DataService struct {
	es *EdgeService

	edges.UnimplementedDataServiceServer
}

func newDataService(es *EdgeService) *DataService {
	return &DataService{
		es: es,
	}
}

func (s *DataService) Upload(ctx context.Context, in *edges.DataUploadRequest) (*edges.DataUploadResponse, error) {
	var err error
	var output edges.DataUploadResponse
	output.Message = "Failed"

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetContent()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid content")
		}
	}

	output.Id = in.GetId()

	switch in.GetContentType() {
	case 1:
		err = s.uploadContentType1(ctx, in, &output)
	default:
		return &output, status.Error(codes.InvalidArgument, "Please supply supported content type")
	}

	if err != nil {
		return &output, err
	}

	output.Message = "Success"

	return &output, nil
}

func (s *DataService) CacheTagValue(tag *model.Tag, value nson.Value) {
	s.es.GetTag().SetTagValue(tag.ID, value)
}

func (s *DataService) Compile(ctx context.Context, in *edges.DataQueryRequest) (*pb.Message, error) {
	var err error
	var output pb.Message

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetFlux()) == 0 {
			return &output, status.Errorf(codes.InvalidArgument, "Please supply valid flux")
		}
	}

	influx := s.es.GetInfluxDB()
	if influx.IsNone() {
		return &output, status.Error(codes.FailedPrecondition, "Influxdb is not enable")
	}
	influxdb := influx.Unwrap()

	vars := map[string]string{
		"start":  "-1h",
		"stop":   "now()",
		"window": "10m",
		"fn":     "mean",
	}

	vars2 := in.GetVars()
	if vars2 != nil {
		for key, val := range vars2 {
			vars[key] = util.TryConvertTimeZone(val)
		}
	}

	vars["bucket"] = influxdb.Bucket()

	tmpl, err := template.New("").Parse(in.GetFlux())
	if err != nil {
		return &output, status.Errorf(codes.FailedPrecondition, "Template parse: %v", err)
	}

	buffer := new(bytes.Buffer)

	err = tmpl.Execute(buffer, vars)
	if err != nil {
		return &output, status.Errorf(codes.FailedPrecondition, "Template exec: %v", err)
	}

	output.Message = buffer.String()

	return &output, nil
}

func (s *DataService) Query(in *edges.DataQueryRequest, stream edges.DataService_QueryServer) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetFlux()) == 0 {
			return status.Errorf(codes.InvalidArgument, "Please supply valid flux")
		}
	}

	influx := s.es.GetInfluxDB()
	if influx.IsNone() {
		return status.Error(codes.FailedPrecondition, "Influxdb is not enable")
	}
	influxdb := influx.Unwrap()

	vars := map[string]string{
		"start":  "-1h",
		"stop":   "now()",
		"window": "10m",
		"fn":     "mean",
	}

	vars2 := in.GetVars()
	if vars2 != nil {
		for key, val := range vars2 {
			vars[key] = util.TryConvertTimeZone(val)
		}
	}

	vars["bucket"] = influxdb.Bucket()

	tmpl, err := template.New("").Parse(in.GetFlux())
	if err != nil {
		return status.Errorf(codes.FailedPrecondition, "Template parse: %v", err)
	}

	buffer := new(bytes.Buffer)

	err = tmpl.Execute(buffer, vars)
	if err != nil {
		return status.Errorf(codes.FailedPrecondition, "Template exec: %v", err)
	}

	result, err := influxdb.QueryAPI().Query(stream.Context(), buffer.String())
	if err != nil {
		return status.Errorf(codes.FailedPrecondition, "InfluxDB Query: %v", err)
	}

	for result.Next() {
		values := result.Record().Values()

		for key, val := range values {
			if t, ok := val.(time.Time); ok {
				values[key] = util.TimeFormat(t)
			}
		}

		bytes, err := json.Marshal(values)
		if err != nil {
			return status.Errorf(codes.Internal, "Json Marshal: %v", err)
		}

		stream.Send(&pb.Message{Message: string(bytes)})
	}

	if err = result.Err(); err != nil {
		return status.Errorf(codes.Internal, "InfluxDB Query: %v", err)
	}

	return nil
}

func (s *DataService) QueryTag(in *edges.DataQueryTagRequest, stream edges.DataService_QueryTagServer) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return status.Errorf(codes.InvalidArgument, "Please supply valid tag id")
		}
	}

	influx := s.es.GetInfluxDB()
	if influx.IsNone() {
		return status.Error(codes.FailedPrecondition, "Influxdb is not enable")
	}
	influxdb := influx.Unwrap()

	vars := map[string]string{
		"start":  "-1h",
		"stop":   "now()",
		"window": "10m",
		"fn":     "mean",
	}

	vars2 := in.GetVars()
	if vars2 != nil {
		for key, val := range vars2 {
			vars[key] = util.TryConvertTimeZone(val)
		}
	}

	vars["tag"] = in.GetId()
	vars["bucket"] = influxdb.Bucket()

	script := `
	from(bucket: "{{.bucket}}")
	    |> range(start: {{.start}}, stop: {{.stop}})
	    |> filter(fn: (r) => r["_field"] == "{{.tag}}")
	    |> aggregateWindow(every: {{.window}}, fn: {{.fn}}, createEmpty: false)
		|> drop(columns: ["_measurement", "_field", "_start", "_stop", "source"])
	`

	tmpl, err := template.New("").Parse(script)
	if err != nil {
		return status.Errorf(codes.FailedPrecondition, "Template parse: %v", err)
	}

	buffer := new(bytes.Buffer)

	err = tmpl.Execute(buffer, vars)
	if err != nil {
		return status.Errorf(codes.FailedPrecondition, "Template exec: %v", err)
	}

	result, err := influxdb.QueryAPI().Query(stream.Context(), buffer.String())
	if err != nil {
		return status.Errorf(codes.FailedPrecondition, "InfluxDB Query: %v", err)
	}

	for result.Next() {
		values := result.Record().Values()

		for key, val := range values {
			if t, ok := val.(time.Time); ok {
				values[key] = util.TimeFormat(t)
			}
		}

		bytes, err := json.Marshal(values)
		if err != nil {
			return status.Errorf(codes.Internal, "Json Marshal: %v", err)
		}

		stream.Send(&pb.Message{Message: string(bytes)})
	}

	if err = result.Err(); err != nil {
		return status.Errorf(codes.Internal, "InfluxDB Query: %v", err)
	}

	return nil
}
