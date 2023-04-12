package model

import (
	"time"
)

type Sync struct {
	Key     string    `bun:"type:TEXT,pk" json:"key"`
	Updated time.Time `bun:"updated" json:"updated"`
}

const (
	SYNC_PREFIX     = "sync_"
	SYNC_DEVICE     = "sync_device"
	SYNC_SLOT       = "sync_slot"
	SYNC_OPTION     = "sync_option"
	SYNC_PORT       = "sync_port"
	SYNC_PROXY      = "sync_proxy"
	SYNC_SOURCE     = "sync_source"
	SYNC_TAG        = "sync_tag"
	SYNC_VAR        = "sync_var"
	SYNC_CABLE      = "sync_cable"
	SYNC_WIRE       = "sync_wire"
	SYNC_CLASS      = "sync_class"
	SYNC_ATTR       = "sync_attr"
	SYNC_TAG_VALUE  = "sync_tag_value"
	SYNC_WIRE_VALUE = "sync_wire_value"

	SYNC_LOCAL_DEVICE      = "sync_l_device"
	SYNC_REMOTE_DEVICE     = "sync_r_device"
	SYNC_LOCAL_TAG_VALUE   = "sync_l_tag_value"
	SYNC_REMOTE_TAG_VALUE  = "sync_r_tag_value"
	SYNC_LOCAL_WIRE_VALUE  = "sync_l_wire_value"
	SYNC_REMOTE_WIRE_VALUE = "sync_r_wire_value"
)
