package model

import (
	"time"

	"github.com/danclive/nson-go"
	"github.com/uptrace/bun"
	"snple.com/kokomi/util/datatype"
)

type Source struct {
	bun.BaseModel `bun:"source"`
	ID            string    `bun:"type:TEXT,pk" json:"id"`
	Name          string    `bun:"name,type:TEXT" json:"name"`
	Desc          string    `bun:"desc,type:TEXT" json:"desc"`
	Tags          string    `bun:"tags,type:TEXT" json:"tags"`
	Type          string    `bun:"type,type:TEXT" json:"type"`
	Source        string    `bun:"source,type:TEXT" json:"source"`
	Params        string    `bun:"params,type:TEXT" json:"params"`
	Config        string    `bun:"config,type:TEXT" json:"config"`
	Status        int32     `bun:"status" json:"status"`
	Upload        int32     `bun:"upload" json:"upload"`
	Save          int32     `bun:"save" json:"save"`
	Deleted       time.Time `bun:"deleted,soft_delete" json:"-"`
	Created       time.Time `bun:"created" json:"created"`
	Updated       time.Time `bun:"updated" json:"updated"`
}

type Tag struct {
	bun.BaseModel `bun:"tag"`
	ID            string    `bun:"type:TEXT,pk" json:"id"`
	SourceID      string    `bun:"source_id,type:TEXT" json:"source_id"`
	Name          string    `bun:"name,type:TEXT" json:"name"`
	Desc          string    `bun:"desc,type:TEXT" json:"desc"`
	Type          string    `bun:"type,type:TEXT" json:"type"`
	Tags          string    `bun:"tags,type:TEXT" json:"tags"`
	DataType      string    `bun:"data_type,type:TEXT" json:"data_type"`
	Address       string    `bun:"address,type:TEXT" json:"address"`
	HValue        string    `bun:"h_value,type:TEXT" json:"h_value"`
	LValue        string    `bun:"l_value,type:TEXT" json:"l_value"`
	Config        string    `bun:"config,type:TEXT" json:"config"`
	Status        int32     `bun:"status" json:"status"`
	Access        int32     `bun:"access" json:"access"`
	Upload        int32     `bun:"upload" json:"upload"`
	Save          int32     `bun:"save" json:"save"`
	Deleted       time.Time `bun:"deleted,soft_delete" json:"-"`
	Created       time.Time `bun:"created" json:"created"`
	Updated       time.Time `bun:"updated" json:"updated"`
}

func (t *Tag) DefaultValue() nson.Value {
	return datatype.DataType(t.DataType).DefaultValue()
}

func (t *Tag) ValueTag() uint8 {
	return datatype.DataType(t.DataType).Tag()
}

type TagValue struct {
	bun.BaseModel `bun:"tag_value"`
	ID            string    `bun:"type:TEXT,pk" json:"id"`
	SourceID      string    `bun:"source_id,type:TEXT" json:"source_id"`
	Value         string    `bun:"value,type:TEXT" json:"value"`
	Deleted       time.Time `bun:"deleted,soft_delete" json:"-"`
	Updated       time.Time `bun:"updated" json:"updated"`
}
