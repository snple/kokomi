package model

import (
	"time"

	"github.com/snple/beacon/util/datatype"
	"github.com/uptrace/bun"
)

type Wire struct {
	bun.BaseModel `bun:"wire"`
	ID            string    `bun:"type:TEXT,pk" json:"id"`
	NodeID        string    `bun:"node_id,type:TEXT" json:"node_id"`
	Name          string    `bun:"name,type:TEXT" json:"name"`
	Desc          string    `bun:"desc,type:TEXT" json:"desc"`
	Tags          string    `bun:"tags,type:TEXT" json:"tags"`
	Source        string    `bun:"source,type:TEXT" json:"source"`
	Params        string    `bun:"params,type:TEXT" json:"params"`
	Config        string    `bun:"config,type:TEXT" json:"config"`
	Status        int32     `bun:"status" json:"status"`
	Deleted       time.Time `bun:"deleted,soft_delete" json:"-"`
	Created       time.Time `bun:"created" json:"created"`
	Updated       time.Time `bun:"updated" json:"updated"`
}

type Pin struct {
	bun.BaseModel `bun:"pin"`
	ID            string    `bun:"type:TEXT,pk" json:"id"`
	NodeID        string    `bun:"node_id,type:TEXT" json:"node_id"`
	WireID        string    `bun:"wire_id,type:TEXT" json:"wire_id"`
	Name          string    `bun:"name,type:TEXT" json:"name"`
	Desc          string    `bun:"desc,type:TEXT" json:"desc"`
	Tags          string    `bun:"tags,type:TEXT" json:"tags"`
	DataType      string    `bun:"data_type,type:TEXT" json:"data_type"`
	Address       string    `bun:"address,type:TEXT" json:"address"`
	Config        string    `bun:"config,type:TEXT" json:"config"`
	Status        int32     `bun:"status" json:"status"`
	Access        int32     `bun:"access" json:"access"`
	Deleted       time.Time `bun:"deleted,soft_delete" json:"-"`
	Created       time.Time `bun:"created" json:"created"`
	Updated       time.Time `bun:"updated" json:"updated"`
}

func (p *Pin) ValueTag() uint8 {
	return datatype.DataType(p.DataType).Tag()
}

type PinValue struct {
	bun.BaseModel `bun:"pin_value"`
	ID            string    `bun:"type:TEXT,pk" json:"id"`
	NodeID        string    `bun:"node_id,type:TEXT" json:"node_id"`
	WireID        string    `bun:"wire_id,type:TEXT" json:"wire_id"`
	Value         string    `bun:"value,type:TEXT" json:"value"`
	Deleted       time.Time `bun:"deleted,soft_delete" json:"-"`
	Updated       time.Time `bun:"updated" json:"updated"`
}

type PinWrite struct {
	bun.BaseModel `bun:"pin_write"`
	ID            string    `bun:"type:TEXT,pk" json:"id"`
	NodeID        string    `bun:"node_id,type:TEXT" json:"node_id"`
	WireID        string    `bun:"wire_id,type:TEXT" json:"wire_id"`
	Value         string    `bun:"value,type:TEXT" json:"value"`
	Deleted       time.Time `bun:"deleted,soft_delete" json:"-"`
	Updated       time.Time `bun:"updated" json:"updated"`
}
