package edge

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/danclive/nson-go"
	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/edge/model"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/util/datatype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	UPLOAD_PREFIX    = "__"
	UPLOAD_SOURCE    = "__src"
	UPLOAD_TIMESTAMP = "__ts"

	UPLOAD_DEFAULT_SOURCE = consts.DEFAULT_SOURCE
)

func (s *DataService) uploadContentType1(ctx context.Context, in *edges.DataUploadRequest, output *edges.DataUploadResponse) error {
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

	validation := newValidationSource2(s)

	i := 0

	if array[0].Tag() == nson.TAG_U32 {
		if array[1].Tag() != nson.TAG_U32 {
			return status.Error(codes.InvalidArgument, "Please supply valid content format")
		}

		i = 2
	}

	for ; i < len(array); i += 2 {
		if array[i].Tag() != nson.TAG_MESSAGE_ID {
			return status.Error(codes.InvalidArgument, "tagID type != nson.TAG_MESSAGE_ID")
		}

		tagID := array[i].(nson.MessageId).Hex()
		value := array[i+1]

		tag := model.Tag{
			ID: tagID,
		}

		tag, err := s.es.GetTag().view(ctx, tagID)
		if err != nil {
			return err
		}

		if tag.Status != consts.ON || tag.Upload != consts.ON {
			continue
		}

		// validation device and source
		valid, err := validation.validation(ctx, s.es, tag.SourceID)
		if err != nil {
			return err
		}

		if !valid {
			continue
		}

		// validation data type
		err = s.validationDataTypeNson(&tag, value)
		if err != nil {
			return err
		}

		// cache
		if in.GetCache() {
			s.CacheTagValue(&tag, value)
		}
	}

	return nil
}

func (s *DataService) validationDataTypeNson(tag *model.Tag, value nson.Value) error {
	switch datatype.DataType(tag.DataType) {
	case datatype.DataTypeI8, datatype.DataTypeI16, datatype.DataTypeI32:
		if value.Tag() != nson.TAG_I32 {
			return status.Errorf(codes.InvalidArgument, "Tag %v,%v value.Tag() != nson.TAG_I32", tag.ID, tag.Name)
		}
	case datatype.DataTypeU8, datatype.DataTypeU16, datatype.DataTypeU32:
		if value.Tag() != nson.TAG_U32 {
			return status.Errorf(codes.InvalidArgument, "Tag %v,%v value.Tag() != nson.TAG_U32", tag.ID, tag.Name)
		}
	case datatype.DataTypeI64:
		if value.Tag() != nson.TAG_I64 {
			return status.Errorf(codes.InvalidArgument, "Tag %v,%v value.Tag() != nson.TAG_I64", tag.ID, tag.Name)
		}
	case datatype.DataTypeU64:
		if value.Tag() != nson.TAG_U64 {
			return status.Errorf(codes.InvalidArgument, "Tag %v,%v value.Tag() != nson.TAG_U64", tag.ID, tag.Name)
		}
	case datatype.DataTypeF32:
		if value.Tag() != nson.TAG_F32 {
			return status.Errorf(codes.InvalidArgument, "Tag %v,%v value.Tag() != nson.TAG_F32", tag.ID, tag.Name)
		}
	case datatype.DataTypeF64:
		if value.Tag() != nson.TAG_F64 {
			return status.Errorf(codes.InvalidArgument, "Tag %v,%v value.Tag() != nson.TAG_F64", tag.ID, tag.Name)
		}
	case datatype.DataTypeBool:
		if value.Tag() != nson.TAG_BOOL {
			return status.Errorf(codes.InvalidArgument, "Tag %v,%v value.Tag() != nson.TAG_BOOL", tag.ID, tag.Name)
		}
	case datatype.DataTypeString:
		if value.Tag() != nson.TAG_STRING {
			return status.Errorf(codes.InvalidArgument, "Tag %v,%v value.Tag() != nson.TAG_STRING", tag.ID, tag.Name)
		}
	default:
		return status.Errorf(codes.InvalidArgument, "Tag %v,%v unsupported value.Tag()", tag.ID, tag.Name)
	}

	return nil
}

func (s *DataService) validationDataTypeJson(tag *model.Tag, value interface{}) (nson.Value, error) {
	var value2 nson.Value

	switch datatype.DataType(tag.DataType) {
	case datatype.DataTypeI8, datatype.DataTypeI16, datatype.DataTypeI32:
		if number, ok := value.(json.Number); ok {
			v, err := number.Int64()
			if err != nil {
				return value2, status.Errorf(codes.InvalidArgument, "Tag %v,%v DecodeValue: %v", tag.ID, tag.Name, err)
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

type validationSource1 struct {
	s            *DataService
	sources      map[string]string // name, id
	blackSources map[string]struct{}
}

func newValidationSource1(s *DataService) *validationSource1 {
	return &validationSource1{
		s:            s,
		sources:      make(map[string]string),
		blackSources: make(map[string]struct{}),
	}
}

func (v *validationSource1) validation(ctx context.Context, es *EdgeService, sourceName string) (sourceID string, valid bool, err error) {
	if id, ok := v.sources[sourceName]; ok {
		sourceID = id
		valid = true
		return
	}

	if _, ok := v.blackSources[sourceName]; ok {
		return
	}

	// source
	{
		var source model.Source
		source, err = es.GetSource().viewByName(ctx, sourceName)
		if err != nil {
			return
		}

		if source.Status != consts.ON {
			v.blackSources[sourceName] = struct{}{}
			return
		}

		v.sources[sourceName] = source.ID

		sourceID = source.ID
		valid = true
		return
	}
}

type validationSource2 struct {
	s            *DataService
	sources      map[string]struct{}
	blackSources map[string]struct{}
}

func newValidationSource2(s *DataService) *validationSource2 {
	return &validationSource2{
		s:            s,
		sources:      make(map[string]struct{}),
		blackSources: make(map[string]struct{}),
	}
}

func (v *validationSource2) validation(ctx context.Context, es *EdgeService, sourceID string) (valid bool, err error) {
	// black list

	if _, ok := v.blackSources[sourceID]; ok {
		return
	}

	// source
	if _, ok := v.sources[sourceID]; !ok {
		var source model.Source
		source, err = es.GetSource().view(ctx, sourceID)
		if err != nil {
			return
		}

		if source.Status != consts.ON {
			v.blackSources[sourceID] = struct{}{}
			return
		}

		v.sources[sourceID] = struct{}{}
	}

	valid = true
	return
}
