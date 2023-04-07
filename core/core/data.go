package core

import (
	"bytes"
	"context"
	"encoding/json"
	"text/template"
	"time"

	"github.com/danclive/nson-go"
	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/core/model"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
	"github.com/snple/kokomi/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DataService struct {
	cs *CoreService

	cores.UnimplementedDataServiceServer
}

func newDateService(cs *CoreService) *DataService {
	return &DataService{
		cs: cs,
	}
}

func (s *DataService) Upload(ctx context.Context, in *cores.DataUploadRequest) (*cores.DataUploadResponse, error) {
	var err error
	var output cores.DataUploadResponse
	output.Message = "Failed"

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetContent()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid content")
		}

		if len(in.GetDeviceId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid device id")
		}
	}

	output.Id = in.GetId()

	device, err := s.cs.GetDevice().view(ctx, in.GetDeviceId())
	if err != nil {
		return &output, err
	}

	if device.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Device Status != ON")
	}

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
	s.cs.GetTag().SetTagValue(tag.ID, value)
}

func (s *DataService) Compile(ctx context.Context, in *cores.DataQueryRequest) (*pb.Message, error) {
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

	influx := s.cs.GetInfluxDB()
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

func (s *DataService) Query(in *cores.DataQueryRequest, stream cores.DataService_QueryServer) error {
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

	influx := s.cs.GetInfluxDB()
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

func (s *DataService) QueryTag(in *cores.DataQueryTagRequest, stream cores.DataService_QueryTagServer) error {
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

	influx := s.cs.GetInfluxDB()
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
