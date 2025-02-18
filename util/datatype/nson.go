package datatype

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/danclive/nson-go"
)

func NsonValueToFloat64(value nson.Value) (float64, bool) {
	switch value.Tag() {
	case nson.TAG_I32:
		return float64(value.(nson.I32)), true
	case nson.TAG_U32:
		return float64(value.(nson.U32)), true
	case nson.TAG_I64:
		return float64(value.(nson.I64)), true
	case nson.TAG_U64:
		return float64(value.(nson.U64)), true
	case nson.TAG_F32:
		return float64(value.(nson.F32)), true
	case nson.TAG_F64:
		return float64(value.(nson.F64)), true
	case nson.TAG_BOOL:
		if value.(nson.Bool) {
			return 1, true
		}

		return 0, true
	}

	return 0, false
}

func EncodeNsonValue(value nson.Value) (string, error) {
	v := ""

	if value == nil {
		return v, errors.New("value is nil")
	}

	switch value.Tag() {
	case nson.TAG_I32:
		v = fmt.Sprintf("%d", int32(value.(nson.I32)))
	case nson.TAG_I64:
		v = fmt.Sprintf("%d", int64(value.(nson.I64)))
	case nson.TAG_U32:
		v = fmt.Sprintf("%d", uint32(value.(nson.U32)))
	case nson.TAG_U64:
		v = fmt.Sprintf("%d", uint64(value.(nson.U64)))
	case nson.TAG_F32:
		v = fmt.Sprintf("%f", float32(value.(nson.F32)))
	case nson.TAG_F64:
		v = fmt.Sprintf("%f", float64(value.(nson.F64)))
	case nson.TAG_BOOL:
		v = fmt.Sprintf("%t", bool(value.(nson.Bool)))
	case nson.TAG_STRING:
		v = string(value.(nson.String))
	case nson.TAG_NULL:
		v = ""
	case nson.TAG_BINARY:
		v = value.(nson.Binary).Hex()
	case nson.TAG_TIMESTAMP:
		v = fmt.Sprintf("%d", value.(nson.Timestamp))
	case nson.TAG_ID:
		v = value.(nson.Id).Hex()
	default:
		return v, fmt.Errorf("unsupported value type: %v", value.Tag())
	}

	return v, nil
}

func DecodeNsonValue(value string, tag uint8) (nson.Value, error) {
	var nsonValue nson.Value

	switch tag {
	case nson.TAG_I32:
		value, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return nsonValue, err
		}

		nsonValue = nson.I32(value)
	case nson.TAG_I64:
		value, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nsonValue, err
		}

		nsonValue = nson.I64(value)
	case nson.TAG_U32:
		value, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return nsonValue, err
		}

		nsonValue = nson.U32(value)
	case nson.TAG_U64:
		value, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return nsonValue, err
		}

		nsonValue = nson.U64(value)
	case nson.TAG_F32:
		value, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return nsonValue, err
		}

		nsonValue = nson.F32(value)
	case nson.TAG_F64:
		value, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nsonValue, err
		}

		nsonValue = nson.F64(value)
	case nson.TAG_BOOL:
		if value == "true" || value == "1" {
			nsonValue = nson.Bool(true)
		} else {
			nsonValue = nson.Bool(false)
		}
	case nson.TAG_STRING:
		nsonValue = nson.String(value)
	case nson.TAG_NULL:
		nsonValue = nson.Null{}
	case nson.TAG_BINARY:
		binary, err := nson.BinaryFromHex(value)
		if err != nil {
			return nsonValue, err
		}
		nsonValue = nson.Binary(binary)
	case nson.TAG_TIMESTAMP:
		value, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nsonValue, err
		}
		nsonValue = nson.Timestamp(value)
	case nson.TAG_ID:
		id, err := nson.IdFromHex(value)
		if err != nil {
			return nsonValue, err
		}
		nsonValue = nson.Id(id)
	default:
		return nsonValue, fmt.Errorf("unsupported value type: %v", tag)
	}

	return nsonValue, nil
}
