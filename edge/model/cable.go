package model

import (
	"time"

	"github.com/danclive/nson-go"
	"github.com/snple/kokomi/util/datatype"
	"github.com/uptrace/bun"
)

type Cable struct {
	bun.BaseModel `bun:"cable"`
	ID            string    `bun:"type:TEXT,pk" json:"id"`
	Name          string    `bun:"name,type:TEXT" json:"name"`
	Desc          string    `bun:"desc,type:TEXT" json:"desc"`
	Tags          string    `bun:"tags,type:TEXT" json:"tags"`
	Type          string    `bun:"type,type:TEXT" json:"type"`
	Config        string    `bun:"config,type:TEXT" json:"config"`
	Status        int32     `bun:"status" json:"status"`
	Save          int32     `bun:"save" json:"save"`
	Deleted       time.Time `bun:"deleted,soft_delete" json:"-"`
	Created       time.Time `bun:"created" json:"created"`
	Updated       time.Time `bun:"updated" json:"updated"`
}

type Wire struct {
	bun.BaseModel `bun:"wire"`
	ID            string    `bun:"type:TEXT,pk" json:"id"`
	CableID       string    `bun:"cable_id,type:TEXT" json:"cable_id"`
	Name          string    `bun:"name,type:TEXT" json:"name"`
	Desc          string    `bun:"desc,type:TEXT" json:"desc"`
	Type          string    `bun:"type,type:TEXT" json:"type"`
	Tags          string    `bun:"tags,type:TEXT" json:"tags"`
	DataType      string    `bun:"data_type,type:TEXT" json:"data_type"`
	HValue        string    `bun:"h_value,type:TEXT" json:"h_value"`
	LValue        string    `bun:"l_value,type:TEXT" json:"l_value"`
	Config        string    `bun:"config,type:TEXT" json:"config"`
	Status        int32     `bun:"status" json:"status"`
	Access        int32     `bun:"access" json:"access"`
	Save          int32     `bun:"save" json:"save"`
	Deleted       time.Time `bun:"deleted,soft_delete" json:"-"`
	Created       time.Time `bun:"created" json:"created"`
	Updated       time.Time `bun:"updated" json:"updated"`
}

func (t *Wire) DefaultValue() nson.Value {
	return datatype.DataType(t.DataType).DefaultValue()
}

func (t *Wire) ValueTag() uint8 {
	return datatype.DataType(t.DataType).Tag()
}

type WireValue struct {
	ID      string    `bun:"type:TEXT,pk" json:"id"`
	CableID string    `bun:"cable_id,type:TEXT" json:"cable_id"`
	Value   string    `bun:"value,type:TEXT" json:"value"`
	Updated time.Time `bun:"updated" json:"updated"`
}

const (
	WIRE_VALUE_PREFIX = "wev_"
)
