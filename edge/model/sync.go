package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Sync struct {
	bun.BaseModel `bun:"sync"`
	Key           string    `bun:"type:TEXT,pk" json:"key"`
	Updated       time.Time `bun:"updated" json:"updated"`
}

const (
	SYNC_DEVICE     = "device"
	SYNC_SLOT       = "slot"
	SYNC_OPTION     = "option"
	SYNC_PORT       = "port"
	SYNC_PROXY      = "proxy"
	SYNC_SOURCE     = "source"
	SYNC_TAG        = "tag"
	SYNC_TAG_VALUE  = "tag_value"
	SYNC_VAR        = "var"
	SYNC_CABLE      = "cable"
	SYNC_WIRE       = "wire"
	SYNC_WIRE_VALUE = "wire_value"
)

const (
	SYNC_LOCAL_DEVICE      = "l_device"
	SYNC_REMOTE_DEVICE     = "r_device"
	SYNC_REMOTE_TAG_VALUE  = "r_tag_value"
	SYNC_LOCAL_WIRE_VALUE  = "l_wire_value"
	SYNC_REMOTE_WIRE_VALUE = "r_wire_value"
)
