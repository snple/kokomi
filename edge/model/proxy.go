package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Proxy struct {
	bun.BaseModel `bun:"proxy"`
	ID            string    `bun:"type:TEXT,pk" json:"id"`
	Name          string    `bun:"name,type:TEXT" json:"name"`
	Desc          string    `bun:"desc,type:TEXT" json:"desc"`
	Tags          string    `bun:"tags,type:TEXT" json:"tags"`
	Type          string    `bun:"type,type:TEXT" json:"type"`
	Network       string    `bun:"network,type:TEXT" json:"network"`
	Address       string    `bun:"address,type:TEXT" json:"address"`
	Target        string    `bun:"target,type:TEXT" json:"target"`
	Config        string    `bun:"config,type:TEXT" json:"config"`
	Status        int32     `bun:"status" json:"status"`
	Deleted       time.Time `bun:"deleted,soft_delete" json:"-"`
	Created       time.Time `bun:"created" json:"created"`
	Updated       time.Time `bun:"updated" json:"updated"`
}
