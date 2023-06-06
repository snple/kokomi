package core

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/danclive/nson-go"
	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/core/model"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
	"github.com/snple/kokomi/util"
	"github.com/snple/kokomi/util/datatype"
	"github.com/snple/types"
	"github.com/snple/types/cache"
	"github.com/uptrace/bun"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AttrService struct {
	cs *CoreService

	valueCache *cache.Cache[nson.Value]

	cores.UnimplementedAttrServiceServer
}

func newAttrService(cs *CoreService) *AttrService {
	return &AttrService{
		cs:         cs,
		valueCache: cache.NewCache[nson.Value](nil),
	}
}

func (s *AttrService) Create(ctx context.Context, in *pb.Attr) (*pb.Attr, error) {
	var output pb.Attr
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetClassId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.ClassID")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.Name")
		}
	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "Attr.Name min 2 character")
		}

		err = s.cs.GetDB().NewSelect().Model(&model.Attr{}).Where("name = ?", in.GetName()).Where("class_id = ?", in.GetClassId()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			return &output, status.Error(codes.AlreadyExists, "Attr.Name must be unique")
		}
	}

	item := model.Attr{
		ID:       in.GetId(),
		ClassID:  in.GetClassId(),
		Name:     in.GetName(),
		Desc:     in.GetDesc(),
		Type:     in.GetType(),
		Tags:     in.GetTags(),
		DataType: in.GetDataType(),
		HValue:   in.GetHValue(),
		LValue:   in.GetLValue(),
		TagID:    in.GetTagId(),
		Config:   in.GetConfig(),
		Status:   in.GetStatus(),
		Access:   in.GetAccess(),
		Save:     in.GetSave(),
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	// class validation
	{
		class, err := s.cs.GetClass().view(ctx, in.GetClassId())
		if err != nil {
			return &output, err
		}

		item.DeviceID = class.DeviceID
	}

	// tag validation
	if in.GetTagId() != "" {
		tag, err := s.cs.GetTag().view(ctx, in.GetTagId())
		if err != nil {
			return &output, err
		}

		if tag.DeviceID != item.DeviceID {
			return &output, status.Error(codes.NotFound, "Query: tag.DeviceID != item.deviceID")
		}
	}

	if len(item.ID) == 0 {
		item.ID = util.RandomID()
	}

	_, err = s.cs.GetDB().NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Insert: %v", err)
	}

	if err = s.afterUpdate(ctx, &item); err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	output.Value, err = s.getAttrValue(ctx, &item)
	if err != nil {
		return &output, err
	}

	return &output, nil
}

func (s *AttrService) Update(ctx context.Context, in *pb.Attr) (*pb.Attr, error) {
	var output pb.Attr
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.ID")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.Name")
		}
	}

	item, err := s.view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "Attr.Name min 2 character")
		}

		modelItem := model.Attr{}
		err = s.cs.GetDB().NewSelect().Model(&modelItem).Where("class_id = ?", item.ClassID).Where("name = ?", in.GetName()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			if modelItem.ID != item.ID {
				return &output, status.Error(codes.AlreadyExists, "Attr.Name must be unique")
			}
		}
	}

	// tag validation
	if in.GetTagId() != "" {
		tag, err := s.cs.GetTag().view(ctx, in.GetTagId())
		if err != nil {
			return &output, err
		}

		if tag.DeviceID != item.DeviceID {
			return &output, status.Error(codes.NotFound, "Query: tag.DeviceID != item.deviceID")
		}
	}

	item.Name = in.GetName()
	item.Desc = in.GetDesc()
	item.Tags = in.GetTags()
	item.Type = in.GetType()
	item.DataType = in.GetDataType()
	item.HValue = in.GetHValue()
	item.LValue = in.GetLValue()
	item.TagID = in.GetTagId()
	item.Config = in.GetConfig()
	item.Status = in.GetStatus()
	item.Access = in.GetAccess()
	item.Save = in.GetSave()
	item.Updated = time.Now()

	_, err = s.cs.GetDB().NewUpdate().Model(&item).WherePK().Exec(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Update: %v", err)
	}

	if err = s.afterUpdate(ctx, &item); err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	output.Value, err = s.getAttrValue(ctx, &item)
	if err != nil {
		return &output, err
	}

	return &output, nil
}

func (s *AttrService) View(ctx context.Context, in *pb.Id) (*pb.Attr, error) {
	var output pb.Attr
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.ID")
		}
	}

	item, err := s.view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	output.Value, err = s.getAttrValue(ctx, &item)
	if err != nil {
		return &output, err
	}

	return &output, nil
}

func (s *AttrService) ViewByName(ctx context.Context, in *cores.ViewAttrByNameRequest) (*pb.Attr, error) {
	var output pb.Attr
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetDeviceId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid DeviceID")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.Name")
		}
	}

	item, err := s.ViewByDeviceIDAndName(ctx, in.GetDeviceId(), in.GetName())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	output.Name = in.GetName()

	output.Value, err = s.getAttrValue(ctx, &item)
	if err != nil {
		return &output, err
	}

	return &output, nil
}

func (s *AttrService) ViewByNameFull(ctx context.Context, in *pb.Name) (*pb.Attr, error) {
	var output pb.Attr
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.Name")
		}
	}

	deviceName := consts.DEFAULT_DEVICE
	className := consts.DEFAULT_CLASS
	itemName := in.GetName()

	if strings.Contains(itemName, ".") {
		splits := strings.Split(itemName, ".")

		switch len(splits) {
		case 2:
			className = splits[0]
			itemName = splits[1]
		case 3:
			deviceName = splits[0]
			className = splits[1]
			itemName = splits[2]
		default:
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.Name")
		}
	}

	device, err := s.cs.GetDevice().viewByName(ctx, deviceName)
	if err != nil {
		return &output, err
	}

	class, err := s.cs.GetClass().ViewByDeviceIDAndName(ctx, device.ID, className)
	if err != nil {
		return &output, err
	}

	item, err := s.ViewByClassIDAndName(ctx, class.ID, itemName)
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	output.Name = in.GetName()

	output.Value, err = s.getAttrValue(ctx, &item)
	if err != nil {
		return &output, err
	}

	return &output, nil
}

func (s *AttrService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.ID")
		}
	}

	item, err := s.view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	item.Updated = time.Now()
	item.Deleted = time.Now()

	_, err = s.cs.GetDB().NewUpdate().Model(&item).Column("updated", "deleted").WherePK().Exec(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Delete: %v", err)
	}

	if err = s.afterDelete(ctx, &item); err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *AttrService) List(ctx context.Context, in *cores.ListAttrRequest) (*cores.ListAttrResponse, error) {
	var err error
	var output cores.ListAttrResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	defaultPage := pb.Page{
		Limit:  10,
		Offset: 0,
	}

	if in.GetPage() == nil {
		in.Page = &defaultPage
	}

	output.Page = in.GetPage()

	var items []model.Attr

	query := s.cs.GetDB().NewSelect().Model(&items)

	if len(in.GetDeviceId()) > 0 {
		query.Where("device_id = ?", in.GetDeviceId())
	}

	if len(in.GetClassId()) > 0 {
		query.Where("class_id = ?", in.GetClassId())
	}

	if len(in.GetPage().GetSearch()) > 0 {
		search := fmt.Sprintf("%%%v%%", in.GetPage().GetSearch())

		query.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			q = q.Where(`"name" LIKE ?`, search).
				WhereOr(`"desc" LIKE ?`, search).
				WhereOr(`"address" LIKE ?`, search)

			return q
		})
	}

	if len(in.GetTags()) > 0 {
		tagsSplit := strings.Split(in.GetTags(), ",")

		if len(tagsSplit) == 1 {
			search := fmt.Sprintf("%%%v%%", tagsSplit[0])

			query = query.Where(`"tags" LIKE ?`, search)
		} else {
			query = query.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
				for i := 0; i < len(tagsSplit); i++ {
					search := fmt.Sprintf("%%%v%%", tagsSplit[i])

					q = q.WhereOr(`"tags" LIKE ?`, search)
				}

				return q
			})
		}
	}

	if len(in.GetType()) > 0 {
		query = query.Where(`type = ?`, in.GetType())
	}

	if len(in.GetPage().GetOrderBy()) > 0 && (in.GetPage().GetOrderBy() == "id" || in.GetPage().GetOrderBy() == "name" ||
		in.GetPage().GetOrderBy() == "created" || in.GetPage().GetOrderBy() == "updated") {
		query.Order(in.GetPage().GetOrderBy() + " " + in.GetPage().GetSort().String())
	} else {
		query.Order("id ASC")
	}

	count, err := query.Offset(int(in.GetPage().GetOffset())).Limit(int(in.GetPage().GetLimit())).ScanAndCount(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Query: %v", err)
	}

	output.Count = uint32(count)

	for i := 0; i < len(items); i++ {
		item := pb.Attr{}

		s.copyModelToOutput(&item, &items[i])

		item.Value, err = s.getAttrValue(ctx, &items[i])
		if err != nil {
			return &output, err
		}

		output.Attr = append(output.Attr, &item)
	}

	return &output, nil
}

func (s *AttrService) Clone(ctx context.Context, in *cores.CloneAttrRequest) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.ID")
		}
	}

	err = s.cs.getClone().cloneAttr(ctx, s.cs.GetDB(), in.GetId(), in.GetClassId())
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *AttrService) GetValue(ctx context.Context, in *pb.Id) (*pb.AttrValue, error) {
	var err error
	var output pb.AttrValue

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.ID")
		}
	}

	item, err := s.view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	output.Id = in.GetId()

	if item.TagID != "" {
		reply, err := s.cs.GetTag().GetValue(ctx, &pb.Id{Id: item.TagID})
		if err != nil {
			return nil, err
		}

		output.Value = reply.GetValue()
		output.Updated = reply.GetUpdated()
	} else {
		var value nson.Value = nson.Null{}
		if v := s.getAttrValueValue(in.GetId()); v.IsSome() {
			cv := v.Unwrap()
			value = cv.Data
			output.Updated = cv.Updated.UnixMicro()
		}

		output.Value, err = datatype.EncodeNsonValue(value)
		if err != nil {
			return &output, status.Errorf(codes.InvalidArgument, "EncodeValue: %v", err)
		}
	}

	return &output, nil
}

func (s *AttrService) SetValue(ctx context.Context, in *pb.AttrValue) (*pb.MyBool, error) {
	return s.setValue(ctx, in, true)
}

func (s *AttrService) SetValueUnchecked(ctx context.Context, in *pb.AttrValue) (*pb.MyBool, error) {
	return s.setValue(ctx, in, false)
}

func (s *AttrService) setValue(ctx context.Context, in *pb.AttrValue, check bool) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.ID")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.Value")
		}
	}

	// attr
	item, err := s.view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	if item.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Attr.Status != ON")
	}

	if check {
		if item.Access != consts.ON {
			return &output, status.Errorf(codes.FailedPrecondition, "Attr.Access != ON")
		}
	}

	nsonValue, err := datatype.DecodeNsonValue(in.GetValue(), item.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	// validation device and class
	{
		// device
		{
			device, err := s.cs.GetDevice().view(ctx, item.DeviceID)
			if err != nil {
				return &output, err
			}

			if device.Status != consts.ON {
				return &output, status.Errorf(codes.FailedPrecondition, "Device.Status != ON")
			}
		}

		// class
		{
			class, err := s.cs.GetClass().view(ctx, item.ClassID)
			if err != nil {
				return &output, err
			}

			if class.Status != consts.ON {
				return &output, status.Errorf(codes.FailedPrecondition, "Class.Status != ON")
			}
		}
	}

	if item.TagID == "" {
		if check {
			return s.cs.GetTag().SetValue(ctx, &pb.TagValue{Id: item.TagID, Value: in.GetValue()})
		}

		return s.cs.GetTag().SetValueUnchecked(ctx, &pb.TagValue{Id: item.TagID, Value: in.GetValue()})
	}

	s.setAttrValue(item.ID, nsonValue)

	output.Bool = true

	return &output, nil
}

func (s *AttrService) GetValueByName(ctx context.Context, in *cores.GetAttrValueByNameRequest) (*cores.AttrNameValue, error) {
	var err error
	var output cores.AttrNameValue

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetDeviceId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid DeviceID")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.Name")
		}
	}

	item, err := s.ViewByDeviceIDAndName(ctx, in.GetDeviceId(), in.GetName())
	if err != nil {
		return &output, err
	}

	output.DeviceId = in.GetDeviceId()
	output.Name = in.GetName()

	if item.TagID == "" {
		reply, err := s.cs.GetTag().GetValue(ctx, &pb.Id{Id: item.TagID})
		if err != nil {
			return nil, err
		}

		output.Value = reply.GetValue()
		output.Updated = reply.GetUpdated()
	} else {
		var value nson.Value = nson.Null{}
		if v := s.getAttrValueValue(item.ID); v.IsSome() {
			cv := v.Unwrap()
			value = cv.Data
			output.Updated = cv.Updated.UnixMicro()
		}

		output.Value, err = datatype.EncodeNsonValue(value)
		if err != nil {
			return &output, status.Errorf(codes.InvalidArgument, "EncodeValue: %v", err)
		}
	}

	return &output, nil
}

func (s *AttrService) SetValueByName(ctx context.Context, in *cores.AttrNameValue) (*pb.MyBool, error) {
	return s.setValueByName(ctx, in, true)
}

func (s *AttrService) SetValueByNameUnchecked(ctx context.Context, in *cores.AttrNameValue) (*pb.MyBool, error) {
	return s.setValueByName(ctx, in, false)
}

func (s *AttrService) setValueByName(ctx context.Context, in *cores.AttrNameValue, check bool) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetDeviceId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid DeviceID")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.Name")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.Value")
		}
	}

	// device
	device, err := s.cs.GetDevice().view(ctx, in.GetDeviceId())
	if err != nil {
		return &output, err
	}

	if device.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Device.Status != ON")
	}

	// name
	className := consts.DEFAULT_CLASS
	itemName := in.GetName()

	if strings.Contains(itemName, ".") {
		splits := strings.Split(itemName, ".")
		if len(splits) != 2 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.Name")
		}

		className = splits[0]
		itemName = splits[1]
	}

	// class
	class, err := s.cs.GetClass().ViewByDeviceIDAndName(ctx, device.ID, className)
	if err != nil {
		return &output, err
	}

	if class.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Class.Status != ON")
	}

	// attr
	item, err := s.ViewByClassIDAndName(ctx, class.ID, itemName)
	if err != nil {
		return &output, err
	}

	if item.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Attr.Status != ON")
	}

	if check {
		if item.Access != consts.ON {
			return &output, status.Errorf(codes.FailedPrecondition, "Attr.Access != ON")
		}
	}

	nsonValue, err := datatype.DecodeNsonValue(in.GetValue(), item.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	if item.TagID == "" {
		if check {
			return s.cs.GetTag().SetValue(ctx, &pb.TagValue{Id: item.TagID, Value: in.GetValue()})
		}

		return s.cs.GetTag().SetValueUnchecked(ctx, &pb.TagValue{Id: item.TagID, Value: in.GetValue()})
	}

	s.setAttrValue(item.ID, nsonValue)

	output.Bool = true

	return &output, nil
}

func (s *AttrService) view(ctx context.Context, id string) (model.Attr, error) {
	item := model.Attr{
		ID: id,
	}

	err := s.cs.GetDB().NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, Attr.ID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *AttrService) ViewByDeviceIDAndName(ctx context.Context, deviceID, name string) (model.Attr, error) {
	item := model.Attr{}

	className := consts.DEFAULT_CLASS
	itemName := name

	if strings.Contains(itemName, ".") {
		splits := strings.Split(itemName, ".")
		if len(splits) != 2 {
			return item, status.Error(codes.InvalidArgument, "Please supply valid Attr.Name")
		}

		className = splits[0]
		itemName = splits[1]
	}

	class, err := s.cs.GetClass().ViewByDeviceIDAndName(ctx, deviceID, className)
	if err != nil {
		return item, err
	}

	return s.ViewByClassIDAndName(ctx, class.ID, itemName)
}

func (s *AttrService) ViewByClassIDAndName(ctx context.Context, classID, name string) (model.Attr, error) {
	item := model.Attr{}

	err := s.cs.GetDB().NewSelect().Model(&item).Where("class_id = ?", classID).Where("name = ?", name).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, ClassID: %v, Attr.Name: %v", err, classID, name)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *AttrService) copyModelToOutput(output *pb.Attr, item *model.Attr) {
	output.Id = item.ID
	output.DeviceId = item.DeviceID
	output.ClassId = item.ClassID
	output.Name = item.Name
	output.Desc = item.Desc
	output.Tags = item.Tags
	output.Type = item.Type
	output.DataType = item.DataType
	output.HValue = item.HValue
	output.LValue = item.LValue
	output.TagId = item.TagID
	output.Config = item.Config
	output.Status = item.Status
	output.Access = item.Access
	output.Save = item.Save
	output.Created = item.Created.UnixMicro()
	output.Updated = item.Updated.UnixMicro()
	output.Deleted = item.Deleted.UnixMicro()
}

func (s *AttrService) getAttrValue(ctx context.Context, item *model.Attr) (string, error) {
	if len(item.TagID) > 0 {
		return s.cs.GetTag().getTagValue(ctx, item.TagID)
	}

	var value nson.Value = nson.Null{}
	if v := s.valueCache.Get(item.ID); v.IsSome() {
		value = v.Unwrap()
	}

	return datatype.EncodeNsonValue(value)
}

func (s *AttrService) getAttrValueValue(id string) types.Option[cache.Value[nson.Value]] {
	return s.valueCache.GetValue(id)
}

func (s *AttrService) setAttrValue(id string, value nson.Value) {
	s.valueCache.Set(id, value, 0)
}

func (s *AttrService) afterUpdate(ctx context.Context, item *model.Attr) error {
	var err error

	err = s.cs.GetSync().setDeviceUpdated(ctx, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *AttrService) afterDelete(ctx context.Context, item *model.Attr) error {
	var err error

	err = s.cs.GetSync().setDeviceUpdated(ctx, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *AttrService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Attr, error) {
	var output pb.Attr
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.ID")
		}
	}

	item, err := s.viewWithDeleted(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *AttrService) viewWithDeleted(ctx context.Context, id string) (model.Attr, error) {
	item := model.Attr{
		ID: id,
	}

	err := s.cs.GetDB().NewSelect().Model(&item).WherePK().WhereAllWithDeleted().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, Attr.ID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *AttrService) Pull(ctx context.Context, in *cores.PullAttrRequest) (*cores.PullAttrResponse, error) {
	var err error
	var output cores.PullAttrResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	var items []model.Attr

	query := s.cs.GetDB().NewSelect().Model(&items)

	if in.GetDeviceId() != "" {
		query.Where("device_id = ?", in.GetDeviceId())
	}

	if in.GetClassId() != "" {
		query.Where("class_id = ?", in.GetClassId())
	}

	if in.GetType() != "" {
		query.Where(`type = ?`, in.GetType())
	}

	err = query.Where("updated > ?", time.UnixMicro(in.GetAfter())).WhereAllWithDeleted().Order("updated ASC").Limit(int(in.GetLimit())).Scan(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Query: %v", err)
	}

	for i := 0; i < len(items); i++ {
		item := pb.Attr{}

		s.copyModelToOutput(&item, &items[i])

		output.Attr = append(output.Attr, &item)
	}

	return &output, nil
}

func (s *AttrService) Sync(ctx context.Context, in *pb.Attr) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.ID")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.Name")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Attr.Updated")
		}
	}

	insert := false
	update := false

	item, err := s.viewWithDeleted(ctx, in.GetId())
	if err != nil {
		if code, ok := status.FromError(err); ok {
			if code.Code() == codes.NotFound {
				insert = true
				goto SKIP
			}
		}

		return &output, err
	}

	update = true

SKIP:

	// insert
	if insert {
		// device validation
		{
			_, err = s.cs.GetDevice().viewWithDeleted(ctx, in.GetDeviceId())
			if err != nil {
				return &output, err
			}
		}

		// class validation
		{
			class, err := s.cs.GetClass().viewWithDeleted(ctx, in.GetClassId())
			if err != nil {
				return &output, err
			}

			if class.DeviceID != in.GetDeviceId() {
				return &output, status.Error(codes.NotFound, "Query: class.DeviceID != in.GetDeviceId()")
			}
		}

		// name validation
		{
			if len(in.GetName()) < 2 {
				return &output, status.Error(codes.InvalidArgument, "Attr.Name min 2 character")
			}

			err = s.cs.GetDB().NewSelect().Model(&model.Attr{}).Where("name = ?", in.GetName()).Where("class_id = ?", in.GetClassId()).Scan(ctx)
			if err != nil {
				if err != sql.ErrNoRows {
					return &output, status.Errorf(codes.Internal, "Query: %v", err)
				}
			} else {
				return &output, status.Error(codes.AlreadyExists, "Attr.Name must be unique")
			}
		}

		// tag validation
		if in.GetTagId() != "" {
			tag, err := s.cs.GetTag().viewWithDeleted(ctx, in.GetTagId())
			if err != nil {
				return &output, err
			}

			if tag.DeviceID != in.GetDeviceId() {
				return &output, status.Error(codes.NotFound, "Query: tag.DeviceID != in.GetDeviceId()")
			}
		}

		item := model.Attr{
			ID:       in.GetId(),
			DeviceID: in.GetDeviceId(),
			ClassID:  in.GetClassId(),
			Name:     in.GetName(),
			Desc:     in.GetDesc(),
			Type:     in.GetType(),
			Tags:     in.GetTags(),
			DataType: in.GetDataType(),
			HValue:   in.GetHValue(),
			LValue:   in.GetLValue(),
			TagID:    in.GetTagId(),
			Config:   in.GetConfig(),
			Status:   in.GetStatus(),
			Access:   in.GetAccess(),
			Save:     in.GetSave(),
			Created:  time.UnixMicro(in.GetCreated()),
			Updated:  time.UnixMicro(in.GetUpdated()),
		}

		_, err = s.cs.GetDB().NewInsert().Model(&item).Exec(ctx)
		if err != nil {
			return &output, status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	// update
	if update {
		if in.GetDeviceId() != item.DeviceID {
			return &output, status.Error(codes.NotFound, "Query: in.GetDeviceId() != item.DeviceID")
		}

		if in.GetUpdated() <= item.Updated.UnixMicro() {
			return &output, nil
		}

		// name validation
		{
			if len(in.GetName()) < 2 {
				return &output, status.Error(codes.InvalidArgument, "Attr.Name min 2 character")
			}

			modelItem := model.Attr{}
			err = s.cs.GetDB().NewSelect().Model(&modelItem).Where("class_id = ?", item.ClassID).Where("name = ?", in.GetName()).Scan(ctx)
			if err != nil {
				if err != sql.ErrNoRows {
					return &output, status.Errorf(codes.Internal, "Query: %v", err)
				}
			} else {
				if modelItem.ID != item.ID {
					return &output, status.Error(codes.AlreadyExists, "Attr.Name must be unique")
				}
			}
		}

		// tag validation
		if in.GetTagId() != "" {
			tag, err := s.cs.GetTag().viewWithDeleted(ctx, in.GetTagId())
			if err != nil {
				return &output, err
			}

			if tag.DeviceID != item.DeviceID {
				return &output, status.Error(codes.NotFound, "Query: tag.DeviceID != item.deviceID")
			}
		}

		item.Name = in.GetName()
		item.Desc = in.GetDesc()
		item.Tags = in.GetTags()
		item.Type = in.GetType()
		item.DataType = in.GetDataType()
		item.HValue = in.GetHValue()
		item.LValue = in.GetLValue()
		item.TagID = in.GetTagId()
		item.Config = in.GetConfig()
		item.Status = in.GetStatus()
		item.Access = in.GetAccess()
		item.Save = in.GetSave()
		item.Updated = time.UnixMicro(in.GetUpdated())
		item.Deleted = time.UnixMicro(in.GetDeleted())

		_, err = s.cs.GetDB().NewUpdate().Model(&item).WherePK().WhereAllWithDeleted().Exec(ctx)
		if err != nil {
			return &output, status.Errorf(codes.Internal, "Update: %v", err)
		}
	}

	if err = s.afterUpdate(ctx, &item); err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}
