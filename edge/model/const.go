package model

import (
	"time"

	"github.com/danclive/nson-go"
	"github.com/snple/kokomi/util/datatype"
	"github.com/uptrace/bun"
)

type Const struct {
	bun.BaseModel `bun:"const"`
	ID            string    `bun:"type:TEXT,pk" json:"id"`
	Name          string    `bun:"name,type:TEXT" json:"name"`
	Desc          string    `bun:"desc,type:TEXT" json:"desc"`
	Tags          string    `bun:"tags,type:TEXT" json:"tags"`
	Type          string    `bun:"type,type:TEXT" json:"type"`
	DataType      string    `bun:"data_type,type:TEXT" json:"data_type"`
	Value         string    `bun:"value,type:TEXT" json:"value"`
	HValue        string    `bun:"h_value,type:TEXT" json:"h_value"`
	LValue        string    `bun:"l_value,type:TEXT" json:"l_value"`
	Config        string    `bun:"config,type:TEXT" json:"config"`
	Status        int32     `bun:"status" json:"status"`
	Deleted       time.Time `bun:"deleted,soft_delete" json:"-"`
	Created       time.Time `bun:"created" json:"created"`
	Updated       time.Time `bun:"updated" json:"updated"`
}

func (t *Const) DefaultValue() nson.Value {
	return datatype.DataType(t.DataType).DefaultValue()
}

func (t *Const) ValueTag() uint8 {
	return datatype.DataType(t.DataType).Tag()
}
