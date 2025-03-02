package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Sync struct {
	bun.BaseModel `bun:"sync"`
	ID            string    `bun:"type:TEXT,pk" json:"id"`
	Updated       time.Time `bun:"updated" json:"updated"`
}

const (
	SYNC_NODE_SUFFIX      = ""
	SYNC_TAG_VALUE_SUFFIX = "_tv"
	SYNC_TAG_WRITE_SUFFIX = "_tw"
)

type SyncGlobal struct {
	bun.BaseModel `bun:"sync_global"`
	Name          string    `bun:"type:TEXT,pk" json:"name"`
	Updated       time.Time `bun:"updated" json:"updated"`
}

const (
	SYNC_GLOBAL_NODE   = "node"
	SYNC_GLOBAL_SLOT   = "slot"
	SYNC_GLOBAL_SOURCE = "source"
	SYNC_GLOBAL_TAG    = "tag"
	SYNC_GLOBAL_CONST  = "const"
	SYNC_GLOBAL_USER   = "user"
)
