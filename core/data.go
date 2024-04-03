package core

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"text/template"
	"time"

	"github.com/danclive/nson-go"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/core/model"
	"github.com/snple/kokomi/db"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
	"github.com/snple/kokomi/util"
	"github.com/snple/kokomi/util/datatype"
	"github.com/snple/kokomi/util/flux"
	"github.com/snple/types"
	"github.com/snple/types/cache"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DataService struct {
	cs *CoreService

	values *cache.Cache[nson.Value]

	cores.UnimplementedDataServiceServer
}

func newDateService(cs *CoreService) *DataService {
	return &DataService{
		cs:     cs,
		values: cache.NewCache[nson.Value](nil),
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid DeviceID")
		}
	}

	output.Id = in.GetId()

	device, err := s.cs.GetDevice().ViewByID(ctx, in.GetDeviceId())
	if err != nil {
		return &output, err
	}

	if device.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Device Status != ON")
	}

	switch in.GetContentType() {
	case 1:
		err = s.uploadContentType1(ctx, in, &output)
	case 2:
		err = s.uploadContentType2(ctx, in, &output)
	case 3:
		err = s.uploadContentType3(ctx, in, &output)
	case 12:
		err = s.uploadContentType12(ctx, in, &output)
	default:
		return &output, status.Error(codes.InvalidArgument, "Please supply supported content type")
	}

	if err != nil {
		return &output, err
	}

	output.Message = "Success"

	return &output, nil
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

func (s *DataService) QueryById(in *cores.DataQueryByIdRequest, stream cores.DataService_QueryByIdServer) error {
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
		"fill":   "none",
	}

	vars2 := in.GetVars()
	if vars2 != nil {
		for key, val := range vars2 {
			vars[key] = util.TryConvertTimeZone(val)
		}
	}

	vars["id"] = in.GetId()
	vars["bucket"] = influxdb.Bucket()

	script := flux.TemplateBasic

	if fill, ok := vars["fill"]; ok {
		switch fill {
		case "prev":
			script = flux.TemplateFill
		case "linear":
			script = flux.TemplateLinear
		}
	}

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

func (s *DataService) SetValue(tagID string, value nson.Value) {
	s.values.Set(tagID, value, 0)
}

func (s *DataService) GetValue(tagID string) types.Option[cache.Value[nson.Value]] {
	return s.values.GetValue(tagID)
}

func (s *DataService) UpdateValue(tag *model.Tag, value nson.Value) {
	s.updateValue(tag, value)
}

func (s *DataService) updateValue(tag *model.Tag, value nson.Value) {
	formatValue := func(src nson.Value) nson.Value {
		switch value.Tag() {
		case nson.TAG_F32:
			return nson.F32(util.Round(float64(src.(nson.F32)), 3))
		case nson.TAG_F64:
			return nson.F64(util.Round(float64(src.(nson.F64)), 3))
		default:
			return src
		}
	}

	updateValue := func(tag *model.Tag, value nson.Value) {
		valueString, err := datatype.EncodeNsonValue(value)
		if err != nil {
			return
		}

		ctx := context.Background()

		if err = s.cs.GetTag().setTagValueUpdated(ctx, tag, valueString, time.Now()); err != nil {
			return
		}

		if err = s.cs.GetTag().afterUpdateValue(ctx, tag, valueString); err != nil {
			return
		}
	}

	value = formatValue(value)

	if option := s.values.Get(tag.ID); option.IsSome() {
		oldvalue := option.Unwrap()

		oldvalue = formatValue(oldvalue)

		if oldvalue != value {
			updateValue(tag, value)
		}
	} else {
		updateValue(tag, value)
	}

	s.values.Set(tag.ID, value, 0)
}

func (s *DataService) uploadContentType1(ctx context.Context, in *cores.DataUploadRequest, _ *cores.DataUploadResponse) error {
	// [TagID, Value, TagID, Value, ...]

	reader := bytes.NewBuffer(in.GetContent())
	nsonValue, err := nson.Array{}.Decode(reader)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "nson::Array::Decode: %v", err)
	}

	array := nsonValue.(nson.Array)
	arrayLen := len(array)

	if arrayLen < 2 {
		return nil
	}

	if arrayLen%2 != 0 {
		return status.Error(codes.InvalidArgument, "Please supply valid content format")
	}

	// InfluxDB writer
	writer := types.None[*db.Writer]()
	option := s.cs.GetInfluxDB()
	if option.IsSome() {
		writer = types.Some(option.Unwrap().Writer(50))
	}

	timestamp := time.Now().Unix()
	i := 0

	if array[0].Tag() == nson.TAG_U32 {
		if array[1].Tag() != nson.TAG_U32 {
			return status.Error(codes.InvalidArgument, "Please supply valid content format")
		}

		i = 2

		// 时间戳
		ts := int64(array[0].(nson.U32))
		if ts > 0 {
			timestamp = ts
		}
	}

	for ; i < len(array); i += 2 {
		if array[i].Tag() != nson.TAG_MESSAGE_ID {
			return status.Error(codes.InvalidArgument, "tagID type != nson.TAG_MESSAGE_ID")
		}

		tagID := array[i].(nson.MessageId).Hex()
		value := array[i+1]

		tag, err := s.cs.GetTag().ViewByID(ctx, tagID)
		if err != nil {
			return err
		}

		if tag.DeviceID != in.GetDeviceId() {
			return status.Errorf(codes.InvalidArgument,
				"tag.DeviceID != DeviceID, TagID: %v, DeviceID: %v", tag.ID, in.GetDeviceId())
		}

		if tag.Status != consts.ON {
			continue
		}

		source, err := s.cs.GetSource().ViewFromCacheByID(ctx, tag.SourceID)
		if err != nil {
			return err
		}

		if source.Status != consts.ON {
			continue
		}

		// validation data type
		err = s.CheckDataTypeNson(&tag, value)
		if err != nil {
			return err
		}

		// realtime
		if in.GetRealtime() {
			s.updateValue(&tag, value)
		}

		// save
		if writer.IsSome() && in.GetSave() && source.Save == consts.ON && tag.Save == consts.ON {
			if value2, ok := datatype.NsonValueToFloat64(value); ok {

				point := s.NewPoint(&tag, value2, timestamp)

				err = writer.Unwrap().Write(ctx, point)
				if err != nil {
					return status.Errorf(codes.Internal, "Write: %v", err)
				}
			}
		}
	}

	if writer.IsSome() {
		err = writer.Unwrap().Flush(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Flush: %v", err)
		}
	}

	return nil
}

const (
	_UPLOAD_PREFIX              = "__"
	_UPLOAD_SOURCE              = "__src"
	_UPLOAD_TIMESTAMP           = "__ts"
	_UPLOAD_TYPE                = "__type"
	_UPLOAD_TYPE_NAME    uint32 = 0
	_UPLOAD_TYPE_ADDRESS uint32 = 1
	_UPLOAD_NAME                = "__name"
	_UPLOAD_VALUES              = "__vs"
)

func (s *DataService) uploadContentType2(ctx context.Context, in *cores.DataUploadRequest, _ *cores.DataUploadResponse) error {
	// {
	//  [TIMESTAMP: u32,]
	//  [SOURCE: String,]
	//  Type: u32, 0: Name, 1: Address, default 0
	//  Name|Address: Value, Name|Address: Value, ...
	// }

	// decode
	reader := bytes.NewBuffer(in.GetContent())
	nsonValue, err := nson.Message{}.Decode(reader)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "nson::Message::Decode: %v", err)
	}

	message := nsonValue.(nson.Message)

	// timestamp
	timestamp := time.Now().Unix()

	if ts, err := message.GetU32(_UPLOAD_TIMESTAMP); err == nil {
		if ts > 0 {
			timestamp = int64(ts)
		}
	}

	// source name
	sourceName := consts.DEFAULT_SOURCE

	if st, err := message.GetString(_UPLOAD_SOURCE); err == nil {
		sourceName = st
	}

	// upload type
	uploadType := _UPLOAD_TYPE_NAME

	if ty, err := message.GetU32(_UPLOAD_TYPE); err == nil {
		if ty == _UPLOAD_TYPE_NAME || ty == _UPLOAD_TYPE_ADDRESS {
			uploadType = ty
		}
	}

	// get and check source
	source, err := s.cs.GetSource().ViewFromCacheByDeviceIDAndName(ctx, in.GetDeviceId(), sourceName)
	if err != nil {
		return err
	}

	if source.Status != consts.ON {
		return nil
	}

	// InfluxDB writer
	writer := types.None[*db.Writer]()
	option := s.cs.GetInfluxDB()
	if option.IsSome() {
		writer = types.Some(option.Unwrap().Writer(50))
	}

	for name, value := range message {
		if strings.HasPrefix(name, _UPLOAD_PREFIX) {
			continue
		}

		// get and check tag
		var tag model.Tag
		if uploadType == _UPLOAD_TYPE_ADDRESS {
			tag, err = s.cs.GetTag().ViewBySourceIDAndAddress(ctx, source.ID, name)
		} else {
			tag, err = s.cs.GetTag().ViewBySourceIDAndName(ctx, source.ID, name)
		}
		if err != nil {
			return err
		}

		if tag.Status != consts.ON {
			continue
		}

		// validation data type
		err = s.CheckDataTypeNson(&tag, value)
		if err != nil {
			return err
		}

		// realtime
		if in.GetRealtime() {
			s.updateValue(&tag, value)
		}

		// save
		if writer.IsSome() && in.GetSave() && source.Save == consts.ON && tag.Save == consts.ON {
			if value2, ok := datatype.NsonValueToFloat64(value); ok {

				point := s.NewPoint(&tag, value2, timestamp)

				err = writer.Unwrap().Write(ctx, point)
				if err != nil {
					return status.Errorf(codes.Internal, "Write: %v", err)
				}
			}
		}
	}

	if writer.IsSome() {
		err = writer.Unwrap().Flush(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Flush: %v", err)
		}
	}

	return nil
}

func (s *DataService) uploadContentType3(ctx context.Context, in *cores.DataUploadRequest, _ *cores.DataUploadResponse) error {
	// {
	//  [SOURCE: String,]
	//  Type: u32, 0: Name, 1: Address, default 0
	//  Name: Name|Address,
	//  Series: Time: Value, Time: Value, ...
	// }

	// decode
	reader := bytes.NewBuffer(in.GetContent())
	nsonValue, err := nson.Array{}.Decode(reader)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "nson::Array::Decode: %v", err)
	}

	message := nsonValue.(nson.Message)

	// source name
	sourceName := consts.DEFAULT_SOURCE

	if st, err := message.GetString(_UPLOAD_SOURCE); err == nil {
		sourceName = st
	}

	// upload type
	uploadType := _UPLOAD_TYPE_NAME

	if ty, err := message.GetU32(_UPLOAD_TYPE); err == nil {
		if ty == _UPLOAD_TYPE_NAME || ty == _UPLOAD_TYPE_ADDRESS {
			uploadType = ty
		}
	}

	// name
	name, err := message.GetString(_UPLOAD_NAME)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Get 'Name' Field: %v", err)
	}

	// values
	values, err := message.GetArray(_UPLOAD_VALUES)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Get 'Values' Field: %v", err)
	}
	valuesLen := len(values)

	if valuesLen < 2 {
		return nil
	}

	if valuesLen%2 != 0 {
		return status.Error(codes.InvalidArgument, "Please supply valid 'Values' format")
	}

	// get and check source
	source, err := s.cs.GetSource().ViewFromCacheByDeviceIDAndName(ctx, in.GetDeviceId(), sourceName)
	if err != nil {
		return err
	}

	if source.Status != consts.ON {
		return nil
	}

	// get and check tag
	var tag model.Tag
	if uploadType == _UPLOAD_TYPE_ADDRESS {
		tag, err = s.cs.GetTag().ViewBySourceIDAndAddress(ctx, source.ID, name)
	} else {
		tag, err = s.cs.GetTag().ViewBySourceIDAndName(ctx, source.ID, name)
	}

	if err != nil {
		return err
	}

	if tag.Status != consts.ON {
		return nil
	}

	// InfluxDB writer
	writer := types.None[*db.Writer]()
	option := s.cs.GetInfluxDB()
	if option.IsSome() {
		writer = types.Some(option.Unwrap().Writer(50))
	}

	for n := 0; n < len(values); n += 2 {
		if values[n].Tag() != nson.TAG_U32 {
			return status.Error(codes.InvalidArgument, "timestamp type != nson.TAG_U32")
		}

		ts := values[n].(nson.U32)
		if ts == 0 {
			return status.Error(codes.InvalidArgument, "timestamp == 0")
		}

		value := values[n+1]

		// validation data type
		err = s.CheckDataTypeNson(&tag, value)
		if err != nil {
			return err
		}

		// save
		if writer.IsSome() && in.GetSave() && source.Save == consts.ON && tag.Save == consts.ON {
			if value2, ok := datatype.NsonValueToFloat64(value); ok {

				point := s.NewPoint(&tag, value2, int64(ts))

				err = writer.Unwrap().Write(ctx, point)
				if err != nil {
					return status.Errorf(codes.Internal, "Write: %v", err)
				}
			}
		}
	}

	if writer.IsSome() {
		err = writer.Unwrap().Flush(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Flush: %v", err)
		}
	}

	return nil
}

func (s *DataService) uploadContentType12(ctx context.Context, in *cores.DataUploadRequest, _ *cores.DataUploadResponse) error {
	// {
	//  [TIMESTAMP: u32,]
	//  [SOURCE: String,]
	//  Type: u32, 0: Name, 1: Address, default 0
	//  Name|Address: Value, Name|Address: Value, ...
	// }

	// decode
	decoder := json.NewDecoder(bytes.NewBuffer(in.GetContent()))
	decoder.UseNumber()

	message := map[string]any{}
	err := decoder.Decode(&message)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "json.Decode: %v", err)
	}

	// timestamp
	timestamp := time.Now().Unix()

	if v, ok := message[_UPLOAD_TIMESTAMP]; ok {
		if n, ok := v.(json.Number); ok {
			ts, e := n.Int64()
			if e != nil {
				return status.Errorf(codes.InvalidArgument, "Get Timestamp At json.Number.Int64(): %v", err)
			}
			if ts > 0 {
				timestamp = int64(ts)
			}
		}
	}

	// source name
	sourceName := consts.DEFAULT_SOURCE

	if v, ok := message[_UPLOAD_SOURCE]; ok {
		if sn, ok := v.(string); ok {
			sourceName = sn
		}
	}

	// upload type
	uploadType := _UPLOAD_TYPE_NAME

	if v, ok := message[_UPLOAD_TYPE]; ok {
		if n, ok := v.(json.Number); ok {
			tys, e := n.Int64()
			if e != nil {
				return status.Errorf(codes.InvalidArgument, "Get UploadType At json.Number.Int64(): %v", err)
			}
			ty := uint32(tys)
			if ty == _UPLOAD_TYPE_NAME || ty == _UPLOAD_TYPE_ADDRESS {
				uploadType = ty
			}
		}
	}

	// get and check source
	source, err := s.cs.GetSource().ViewFromCacheByDeviceIDAndName(ctx, in.GetDeviceId(), sourceName)
	if err != nil {
		return err
	}

	if source.Status != consts.ON {
		return nil
	}

	// InfluxDB writer
	writer := types.None[*db.Writer]()
	option := s.cs.GetInfluxDB()
	if option.IsSome() {
		writer = types.Some(option.Unwrap().Writer(50))
	}

	for name, value := range message {
		if strings.HasPrefix(name, _UPLOAD_PREFIX) {
			continue
		}

		// get and check tag
		var tag model.Tag
		if uploadType == _UPLOAD_TYPE_ADDRESS {
			tag, err = s.cs.GetTag().ViewBySourceIDAndAddress(ctx, source.ID, name)
		} else {
			tag, err = s.cs.GetTag().ViewBySourceIDAndName(ctx, source.ID, name)
		}
		if err != nil {
			return err
		}

		if tag.Status != consts.ON {
			continue
		}

		// validation data type
		value2, err := s.CheckDataTypeJson(&tag, value)
		if err != nil {
			return err
		}

		// realtime
		if in.GetRealtime() {
			s.updateValue(&tag, value2)
		}

		// save
		if writer.IsSome() && in.GetSave() && source.Save == consts.ON && tag.Save == consts.ON {
			if value2, ok := datatype.NsonValueToFloat64(value2); ok {

				point := s.NewPoint(&tag, value2, timestamp)

				err = writer.Unwrap().Write(ctx, point)
				if err != nil {
					return status.Errorf(codes.Internal, "Write: %v", err)
				}
			}
		}
	}

	if writer.IsSome() {
		err = writer.Unwrap().Flush(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Flush: %v", err)
		}
	}

	return nil
}

func (s *DataService) CheckDataTypeNson(tag *model.Tag, value nson.Value) error {
	switch datatype.DataType(tag.DataType) {
	case datatype.DataTypeI8, datatype.DataTypeI16, datatype.DataTypeI32:
		if value.Tag() != nson.TAG_I32 {
			return status.Errorf(codes.InvalidArgument, "Tag %v,%v, value.Tag() != nson.TAG_I32", tag.ID, tag.Name)
		}
	case datatype.DataTypeU8, datatype.DataTypeU16, datatype.DataTypeU32:
		if value.Tag() != nson.TAG_U32 {
			return status.Errorf(codes.InvalidArgument, "Tag %v,%v, value.Tag() != nson.TAG_U32", tag.ID, tag.Name)
		}
	case datatype.DataTypeI64:
		if value.Tag() != nson.TAG_I64 {
			return status.Errorf(codes.InvalidArgument, "Tag %v,%v, value.Tag() != nson.TAG_I64", tag.ID, tag.Name)
		}
	case datatype.DataTypeU64:
		if value.Tag() != nson.TAG_U64 {
			return status.Errorf(codes.InvalidArgument, "Tag %v,%v, value.Tag() != nson.TAG_U64", tag.ID, tag.Name)
		}
	case datatype.DataTypeF32:
		if value.Tag() != nson.TAG_F32 {
			return status.Errorf(codes.InvalidArgument, "Tag %v,%v, value.Tag() != nson.TAG_F32", tag.ID, tag.Name)
		}
	case datatype.DataTypeF64:
		if value.Tag() != nson.TAG_F64 {
			return status.Errorf(codes.InvalidArgument, "Tag %v,%v, value.Tag() != nson.TAG_F64", tag.ID, tag.Name)
		}
	case datatype.DataTypeBool:
		if value.Tag() != nson.TAG_BOOL {
			return status.Errorf(codes.InvalidArgument, "Tag %v,%v, value.Tag() != nson.TAG_BOOL", tag.ID, tag.Name)
		}
	case datatype.DataTypeString:
		if value.Tag() != nson.TAG_STRING {
			return status.Errorf(codes.InvalidArgument, "Tag %v,%v, value.Tag() != nson.TAG_STRING", tag.ID, tag.Name)
		}
	default:
		return status.Errorf(codes.InvalidArgument, "Tag %v,%v, unsupported value.Tag()", tag.ID, tag.Name)
	}

	return nil
}

func (s *DataService) CheckDataTypeJson(tag *model.Tag, value any) (nson.Value, error) {
	var value2 nson.Value

	switch datatype.DataType(tag.DataType) {
	case datatype.DataTypeI8, datatype.DataTypeI16, datatype.DataTypeI32:
		if number, ok := value.(json.Number); ok {
			v, err := number.Int64()
			if err != nil {
				return value2, status.Errorf(codes.InvalidArgument, "Tag %v,%v, DecodeValue: %v", tag.ID, tag.Name, err)
			}

			value2 = nson.I32(v)
		} else {
			return value2, status.Errorf(codes.InvalidArgument, "Tag %v,%v, value type != number", tag.ID, tag.Name)
		}
	case datatype.DataTypeU8, datatype.DataTypeU16, datatype.DataTypeU32:
		if number, ok := value.(json.Number); ok {
			v, err := number.Int64()
			if err != nil {
				return value2, status.Errorf(codes.InvalidArgument, "Tag %v,%v, DecodeValue: %v", tag.ID, tag.Name, err)
			}

			value2 = nson.U32(v)
		} else {
			return value2, status.Errorf(codes.InvalidArgument, "Tag %v,%v, value type != number", tag.ID, tag.Name)
		}
	case datatype.DataTypeI64:
		if number, ok := value.(json.Number); ok {
			v, err := number.Int64()
			if err != nil {
				return value2, status.Errorf(codes.InvalidArgument, "Tag %v,%v, DecodeValue: %v", tag.ID, tag.Name, err)
			}

			value2 = nson.I64(v)
		} else {
			return value2, status.Errorf(codes.InvalidArgument, "Tag %v,%v, value type != number", tag.ID, tag.Name)
		}
	case datatype.DataTypeU64:
		if number, ok := value.(json.Number); ok {
			v, err := number.Int64()
			if err != nil {
				return value2, status.Errorf(codes.InvalidArgument, "Tag %v,%v, DecodeValue: %v", tag.ID, tag.Name, err)
			}

			value2 = nson.U64(v)
		} else {
			return value2, status.Errorf(codes.InvalidArgument, "Tag %v,%v, value type != number", tag.ID, tag.Name)
		}
	case datatype.DataTypeF32:
		if number, ok := value.(json.Number); ok {
			v, err := number.Float64()
			if err != nil {
				return value2, status.Errorf(codes.InvalidArgument, "Tag %v,%v, DecodeValue: %v", tag.ID, tag.Name, err)
			}

			value2 = nson.F32(v)
		} else {
			return value2, status.Errorf(codes.InvalidArgument, "Tag %v,%v, value type != number", tag.ID, tag.Name)
		}
	case datatype.DataTypeF64:
		if number, ok := value.(json.Number); ok {
			v, err := number.Float64()
			if err != nil {
				return value2, status.Errorf(codes.InvalidArgument, "Tag %v,%v, DecodeValue: %v", tag.ID, tag.Name, err)
			}

			value2 = nson.F64(v)
		} else {
			return value2, status.Errorf(codes.InvalidArgument, "Tag %v,%v, value type != number", tag.ID, tag.Name)
		}
	case datatype.DataTypeBool:
		if number, ok := value.(json.Number); ok {
			v, err := number.Int64()
			if err != nil {
				return value2, status.Errorf(codes.InvalidArgument, "Tag %v,%v, DecodeValue: %v", tag.ID, tag.Name, err)
			}

			if v >= 1 {
				value2 = nson.Bool(true)
			} else {
				value2 = nson.Bool(false)
			}
		} else if v, ok := value.(bool); ok {
			value2 = nson.Bool(v)
		} else {
			return value2, status.Errorf(codes.InvalidArgument, "Please supply valid value type, Tag %v,%v, value type != number or bool", tag.ID, tag.Name)
		}
	case datatype.DataTypeString:
		if v, ok := value.(string); ok {
			value2 = nson.String(v)
		} else {
			return value2, status.Errorf(codes.InvalidArgument, "Please supply valid value type, Tag %v,%v, value type != string", tag.ID, tag.Name)
		}
	default:
		return value2, status.Errorf(codes.InvalidArgument, "Please supply valid value type, Tag %v,%v, unsupported value.Tag()", tag.ID, tag.Name)
	}

	return value2, nil
}

func (s *DataService) NewPoint(tag *model.Tag, value float64, timestamp int64) *write.Point {
	return write.NewPoint(
		tag.DeviceID,
		map[string]string{
			"source": tag.SourceID,
		},
		map[string]any{
			tag.ID: value,
		},
		time.Unix(timestamp, 0),
	)
}
