package datatype

import "github.com/danclive/nson-go"

type DataType string

const (
	DataTypeBool   DataType = "BOOL"
	DataTypeF32    DataType = "F32"
	DataTypeF64    DataType = "F64"
	DataTypeI8     DataType = "I8"
	DataTypeU8     DataType = "U8"
	DataTypeI16    DataType = "I16"
	DataTypeU16    DataType = "U16"
	DataTypeI32    DataType = "I32"
	DataTypeI64    DataType = "I64"
	DataTypeU32    DataType = "U32"
	DataTypeU64    DataType = "U64"
	DataTypeString DataType = "STRING"
)

func (t DataType) Size() int {
	switch t {
	case DataTypeBool, DataTypeI8, DataTypeU8:
		return 1
	case DataTypeI16, DataTypeU16:
		return 2
	case DataTypeI32, DataTypeU32, DataTypeF32:
		return 4
	case DataTypeI64, DataTypeU64, DataTypeF64:
		return 8
	}

	return 0
}

func (t DataType) IsNumber() bool {
	switch t {
	case DataTypeI8, DataTypeI16, DataTypeI32,
		DataTypeU8, DataTypeU16, DataTypeU32,
		DataTypeI64, DataTypeU64,
		DataTypeF32, DataTypeF64:
		return true
	}

	return false
}

func (t DataType) DefaultValue() nson.Value {
	switch t {
	case DataTypeI8, DataTypeI16, DataTypeI32:
		return nson.I32(0)
	case DataTypeU8, DataTypeU16, DataTypeU32:
		return nson.U32(0)
	case DataTypeI64:
		return nson.I64(0)
	case DataTypeU64:
		return nson.U64(0)
	case DataTypeF32:
		return nson.F32(0)
	case DataTypeF64:
		return nson.F64(0)
	case DataTypeBool:
		return nson.Bool(false)
	case DataTypeString:
		return nson.String("")
	default:
		return nson.Null{}
	}
}

func (t DataType) Tag() uint8 {
	switch t {
	case DataTypeI8, DataTypeI16, DataTypeI32:
		return nson.TAG_I32
	case DataTypeU8, DataTypeU16, DataTypeU32:
		return nson.TAG_U32
	case DataTypeI64:
		return nson.TAG_I64
	case DataTypeU64:
		return nson.TAG_U64
	case DataTypeF32:
		return nson.TAG_F32
	case DataTypeF64:
		return nson.TAG_F64
	case DataTypeBool:
		return nson.TAG_BOOL
	case DataTypeString:
		return nson.TAG_STRING
	default:
		return nson.TAG_NULL
	}
}
