package shiftime

import (
	"github.com/snple/kokomi/pb"
)

func Device(item *pb.Device) {
	if item != nil {
		item.Created = item.Created / 1000
		item.Updated = item.Updated / 1000
		item.Deleted = item.Deleted / 1000
	}
}

func Devices(items []*pb.Device) {
	for _, item := range items {
		Device(item)
	}
}

func Slot(item *pb.Slot) {
	if item != nil {
		item.Created = item.Created / 1000
		item.Updated = item.Updated / 1000
		item.Deleted = item.Deleted / 1000
	}
}

func Slots(items []*pb.Slot) {
	for _, item := range items {
		Slot(item)
	}
}

func Source(item *pb.Source) {
	if item != nil {
		item.Created = item.Created / 1000
		item.Updated = item.Updated / 1000
		item.Deleted = item.Deleted / 1000
	}
}

func Sources(items []*pb.Source) {
	for _, item := range items {
		Source(item)
	}
}

func Tag(item *pb.Tag) {
	if item != nil {
		item.Created = item.Created / 1000
		item.Updated = item.Updated / 1000
		item.Deleted = item.Deleted / 1000
	}
}

func Tags(items []*pb.Tag) {
	for _, item := range items {
		Tag(item)
	}
}

func TagValue(item *pb.TagValue) {
	if item != nil {
		item.Updated = item.Updated / 1000
	}
}

func TagNameValue(item *pb.TagNameValue) {
	if item != nil {
		item.Updated = item.Updated / 1000
	}
}

func TagNameValues(items []*pb.TagNameValue) {
	for _, item := range items {
		TagNameValue(item)
	}
}

func User(item *pb.User) {
	if item != nil {
		item.Created = item.Created / 1000
		item.Updated = item.Updated / 1000
		item.Deleted = item.Deleted / 1000
	}
}

func Users(items []*pb.User) {
	for _, item := range items {
		User(item)
	}
}

func Const(item *pb.Const) {
	if item != nil {
		item.Created = item.Created / 1000
		item.Updated = item.Updated / 1000
		item.Deleted = item.Deleted / 1000
	}
}

func Consts(items []*pb.Const) {
	for _, item := range items {
		Const(item)
	}
}

func ConstValue(item *pb.ConstValue) {
	if item != nil {
		item.Updated = item.Updated / 1000
	}
}

func ConstNameValue(item *pb.ConstNameValue) {
	if item != nil {
		item.Updated = item.Updated / 1000
	}
}

func ConstNameValues(items []*pb.ConstNameValue) {
	for _, item := range items {
		ConstNameValue(item)
	}
}
