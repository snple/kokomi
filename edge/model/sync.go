package model

import "time"

type Sync struct {
	Key     string    `bun:"type:TEXT,pk" json:"key"`
	Updated time.Time `bun:"updated" json:"updated"`
}

const (
	SYNC_PREFIX    = "sync_"
	SYNC_NODE      = "sync_node"
	SYNC_SLOT      = "sync_slot"
	SYNC_SOURCE    = "sync_source"
	SYNC_TAG       = "sync_tag"
	SYNC_CONST     = "sync_const"
	SYNC_TAG_VALUE = "sync_tag_value"
	SYNC_TAG_WRITE = "sync_tag_write"

	SYNC_NODE_REMOTE_TO_LOCAL = "sync_node_rtl"
	SYNC_NODE_LOCAL_TO_REMOTE = "sync_node_ltr"

	SYNC_TAG_VALUE_REMOTE_TO_LOCAL = "sync_tgv_rtl"
	SYNC_TAG_VALUE_LOCAL_TO_REMOTE = "sync_tgv_ltr"

	SYNC_TAG_WRITE_REMOTE_TO_LOCAL = "sync_tgw_rtl"
	SYNC_TAG_WRITE_LOCAL_TO_REMOTE = "sync_tgw_ltr"
)
