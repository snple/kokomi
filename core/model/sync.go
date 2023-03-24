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
	SYNC_DEVICE_SUFFIX     = ""
	SYNC_TAG_VALUE_SUFFIX  = "_tgv"
	SYNC_WIRE_VALUE_SUFFIX = "_wev"
)

type SyncGlobal struct {
	bun.BaseModel `bun:"sync_global"`
	Key           string    `bun:"type:TEXT,pk" json:"key"`
	Updated       time.Time `bun:"updated" json:"updated"`
}

const (
	SYNC_USER = "user"
)
