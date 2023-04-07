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
	"github.com/snple/kokomi/util/metadata"
	"github.com/snple/types"
	"github.com/snple/types/cache"
	"github.com/uptrace/bun"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TagService struct {
	cs *CoreService

	valueCache *cache.Cache[nson.Value]

	cores.UnimplementedTagServiceServer
}

func newTagService(cs *CoreService) *TagService {
	return &TagService{
		cs:         cs,
		valueCache: cache.NewCache[nson.Value](nil),
	}
}

func (s *TagService) Create(ctx context.Context, in *pb.Tag) (*pb.Tag, error) {
	var output pb.Tag
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetSourceId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid source_id")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag name")
		}
	}

	item := model.Tag{
		ID:       in.GetId(),
		SourceID: in.GetSourceId(),
		Name:     in.GetName(),
		Desc:     in.GetDesc(),
		Type:     in.GetType(),
		Tags:     in.GetTags(),
		DataType: in.GetDataType(),
		Address:  in.GetAddress(),
		HValue:   in.GetHValue(),
		LValue:   in.GetLValue(),
		Config:   in.GetConfig(),
		Status:   in.GetStatus(),
		Access:   in.GetAccess(),
		Save:     in.GetSave(),
		Upload:   in.GetUpload(),
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	// source validation
	{
		source, err := s.cs.GetSource().view(ctx, in.GetSourceId())
		if err != nil {
			return &output, err
		}

		item.DeviceID = source.DeviceID
	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "tag name min 2 character")
		}

		err = s.cs.GetDB().NewSelect().Model(&model.Tag{}).Where("name = ?", in.GetName()).Where("source_id = ?", in.GetSourceId()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			return &output, status.Error(codes.AlreadyExists, "tag name must be unique")
		}
	}

	isSync := metadata.IsSync(ctx)

	if len(item.ID) == 0 {
		item.ID = util.RandomID()
	}

	if isSync {
		item.Created = time.UnixMilli(in.GetCreated())
		item.Updated = time.UnixMilli(in.GetUpdated())
	}

	_, err = s.cs.GetDB().NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Insert: %v", err)
	}

	if err = s.afterUpdate(ctx, &item); err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	value := s.GetTagValue(item.ID)

	output.Value, err = datatype.EncodeNsonValue(value)
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "EncodeValue: %v", err)
	}

	return &output, nil
}

func (s *TagService) Update(ctx context.Context, in *pb.Tag) (*pb.Tag, error) {
	var output pb.Tag
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag id")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag name")
		}
	}

	isSync := metadata.IsSync(ctx)

	var item model.Tag

	if isSync {
		item, err = s.viewWithDeleted(ctx, in.GetId())
		if err != nil {
			return &output, err
		}
	} else {
		item, err = s.view(ctx, in.GetId())
		if err != nil {
			return &output, err
		}
	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "tag name min 2 character")
		}

		modelItem := model.Tag{}
		err = s.cs.GetDB().NewSelect().Model(&modelItem).Where("source_id = ?", item.SourceID).Where("name = ?", in.GetName()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			if modelItem.ID != item.ID {
				return &output, status.Error(codes.AlreadyExists, "tag name must be unique")
			}
		}
	}

	item.Name = in.GetName()
	item.Desc = in.GetDesc()
	item.Tags = in.GetTags()
	item.Type = in.GetType()
	item.DataType = in.GetDataType()
	item.Address = in.GetAddress()
	item.HValue = in.GetHValue()
	item.LValue = in.GetLValue()
	item.Config = in.GetConfig()
	item.Status = in.GetStatus()
	item.Access = in.GetAccess()
	item.Save = in.GetSave()
	item.Upload = in.GetUpload()
	item.Updated = time.Now()

	if isSync {
		item.Updated = time.UnixMilli(in.GetUpdated())
		item.Deleted = time.UnixMilli(in.GetDeleted())

		_, err = s.cs.GetDB().NewUpdate().Model(&item).WherePK().WhereAllWithDeleted().Exec(ctx)
		if err != nil {
			return &output, status.Errorf(codes.Internal, "Update: %v", err)
		}
	} else {
		_, err = s.cs.GetDB().NewUpdate().Model(&item).WherePK().Exec(ctx)
		if err != nil {
			return &output, status.Errorf(codes.Internal, "Update: %v", err)
		}
	}

	if err = s.afterUpdate(ctx, &item); err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	value := s.GetTagValue(item.ID)

	output.Value, err = datatype.EncodeNsonValue(value)
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "EncodeValue: %v", err)
	}

	return &output, nil
}

func (s *TagService) View(ctx context.Context, in *pb.Id) (*pb.Tag, error) {
	var output pb.Tag
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag id")
		}
	}

	item, err := s.view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	value := s.GetTagValue(item.ID)

	output.Value, err = datatype.EncodeNsonValue(value)
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "EncodeValue: %v", err)
	}

	return &output, nil
}

func (s *TagService) ViewByName(ctx context.Context, in *cores.ViewTagByNameRequest) (*pb.Tag, error) {
	var output pb.Tag
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetDeviceId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid device id")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag name")
		}
	}

	item, err := s.viewByDeviceIDAndName(ctx, in.GetDeviceId(), in.GetName())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	output.Name = in.GetName()

	value := s.GetTagValue(item.ID)

	output.Value, err = datatype.EncodeNsonValue(value)
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "EncodeValue: %v", err)
	}

	return &output, nil
}

func (s *TagService) ViewByNameFull(ctx context.Context, in *pb.Name) (*pb.Tag, error) {
	var output pb.Tag
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag name")
		}
	}

	deviceName := consts.DEFAULT_DEVICE
	sourceName := consts.DEFAULT_SOURCE
	itemName := in.GetName()

	if strings.Contains(itemName, ".") {
		splits := strings.Split(itemName, ".")

		switch len(splits) {
		case 2:
			sourceName = splits[0]
			itemName = splits[1]
		case 3:
			deviceName = splits[0]
			sourceName = splits[1]
			itemName = splits[2]
		default:
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag name")
		}
	}

	device, err := s.cs.GetDevice().viewByName(ctx, deviceName)
	if err != nil {
		return &output, err
	}

	source, err := s.cs.GetSource().viewByDeviceIDAndName(ctx, device.ID, sourceName)
	if err != nil {
		return &output, err
	}

	item, err := s.viewBySourceIDAndName(ctx, source.ID, itemName)
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	output.Name = in.GetName()

	value := s.GetTagValue(item.ID)

	output.Value, err = datatype.EncodeNsonValue(value)
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "EncodeValue: %v", err)
	}

	return &output, nil
}

func (s *TagService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag id")
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

func (s *TagService) List(ctx context.Context, in *cores.ListTagRequest) (*cores.ListTagResponse, error) {
	var err error
	var output cores.ListTagResponse

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

	var items []model.Tag

	query := s.cs.GetDB().NewSelect().Model(&items)

	if len(in.GetDeviceId()) > 0 {
		query.Where("device_id = ?", in.GetDeviceId())
	}

	if len(in.GetSourceId()) > 0 {
		query.Where("source_id = ?", in.GetSourceId())
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
		item := pb.Tag{}

		s.copyModelToOutput(&item, &items[i])

		value := s.GetTagValue(items[i].ID)

		item.Value, err = datatype.EncodeNsonValue(value)
		if err != nil {
			return &output, status.Errorf(codes.InvalidArgument, "EncodeValue: %v", err)
		}

		output.Tag = append(output.Tag, &item)
	}

	return &output, nil
}

func (s *TagService) Clone(ctx context.Context, in *cores.CloneTagRequest) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag id")
		}
	}

	err = s.cs.getClone().cloneTag(ctx, s.cs.GetDB(), in.GetId(), in.GetSourceId())
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *TagService) GetValue(ctx context.Context, in *pb.Id) (*pb.TagValue, error) {
	var err error
	var output pb.TagValue

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag id")
		}
	}

	output.Id = in.GetId()

	var value nson.Value = nson.Null{}
	if v := s.GetTagValueValue(in.GetId()); v.IsSome() {
		cv := v.Unwrap()
		value = cv.Data
		output.Updated = cv.Updated.UnixMilli()
	}

	output.Value, err = datatype.EncodeNsonValue(value)
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "EncodeValue: %v", err)
	}

	return &output, nil
}

func (s *TagService) SetValue(ctx context.Context, in *pb.TagValue) (*pb.MyBool, error) {
	return s.setValue(ctx, in, true)
}

func (s *TagService) SetValueUnchecked(ctx context.Context, in *pb.TagValue) (*pb.MyBool, error) {
	return s.setValue(ctx, in, false)
}

func (s *TagService) setValue(ctx context.Context, in *pb.TagValue, check bool) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag id")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid value")
		}
	}

	// tag
	item, err := s.view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	if item.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag Status != ON")
	}

	if check {
		if item.Access != consts.ON {
			return &output, status.Errorf(codes.FailedPrecondition, "Tag Access != ON")
		}
	}

	nsonValue, err := datatype.DecodeNsonValue(in.GetValue(), item.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	// validation device and source
	{
		// device
		{
			device, err := s.cs.GetDevice().view(ctx, item.DeviceID)
			if err != nil {
				return &output, err
			}

			if device.Status != consts.ON {
				return &output, status.Errorf(codes.FailedPrecondition, "Device Status != ON")
			}
		}

		// source
		{
			source, err := s.cs.GetSource().view(ctx, item.SourceID)
			if err != nil {
				return &output, err
			}

			if source.Status != consts.ON {
				return &output, status.Errorf(codes.FailedPrecondition, "Source Status != ON")
			}
		}
	}

	if check {
		if err = s.updateTagValue(ctx, &item, in.GetValue()); err != nil {
			return &output, err
		}

		if err = s.afterUpdateValue(ctx, &item, in.GetValue()); err != nil {
			return &output, err
		}
	} else {
		s.SetTagValue(item.ID, nsonValue)
	}

	output.Bool = true

	return &output, nil
}

func (s *TagService) GetValueByName(ctx context.Context, in *cores.GetTagValueByNameRequest) (*cores.TagNameValue, error) {
	var err error
	var output cores.TagNameValue

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetDeviceId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid device id")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag name")
		}
	}

	item, err := s.viewByDeviceIDAndName(ctx, in.GetDeviceId(), in.GetName())
	if err != nil {
		return &output, err
	}

	output.DeviceId = in.GetDeviceId()
	output.Name = in.GetName()

	var value nson.Value = nson.Null{}
	if v := s.GetTagValueValue(item.ID); v.IsSome() {
		cv := v.Unwrap()
		value = cv.Data
		output.Updated = cv.Updated.UnixMilli()
	}

	output.Value, err = datatype.EncodeNsonValue(value)
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "EncodeValue: %v", err)
	}

	return &output, nil
}

func (s *TagService) SetValueByName(ctx context.Context, in *cores.TagNameValue) (*pb.MyBool, error) {
	return s.setValueByName(ctx, in, true)
}

func (s *TagService) SetValueByNameUnchecked(ctx context.Context, in *cores.TagNameValue) (*pb.MyBool, error) {
	return s.setValueByName(ctx, in, false)
}

func (s *TagService) setValueByName(ctx context.Context, in *cores.TagNameValue, check bool) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetDeviceId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid device id")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag name")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid value")
		}
	}

	// device
	device, err := s.cs.GetDevice().view(ctx, in.GetDeviceId())
	if err != nil {
		return &output, err
	}

	if device.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Device Status != ON")
	}

	// name
	sourceName := consts.DEFAULT_SOURCE
	itemName := in.GetName()

	if strings.Contains(itemName, ".") {
		splits := strings.Split(itemName, ".")
		if len(splits) != 2 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag name")
		}

		sourceName = splits[0]
		itemName = splits[1]
	}

	// source
	source, err := s.cs.GetSource().viewByDeviceIDAndName(ctx, device.ID, sourceName)
	if err != nil {
		return &output, err
	}

	if source.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Source Status != ON")
	}

	// tag
	item, err := s.viewBySourceIDAndName(ctx, source.ID, itemName)
	if err != nil {
		return &output, err
	}

	if item.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag Status != ON")
	}

	if check {
		if item.Access != consts.ON {
			return &output, status.Errorf(codes.FailedPrecondition, "Tag Access != ON")
		}
	}

	nsonValue, err := datatype.DecodeNsonValue(in.GetValue(), item.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	if check {
		if err = s.updateTagValue(ctx, &item, in.GetValue()); err != nil {
			return &output, err
		}

		if err = s.afterUpdateValue(ctx, &item, in.GetValue()); err != nil {
			return &output, err
		}
	} else {
		s.SetTagValue(item.ID, nsonValue)
	}

	output.Bool = true

	return &output, nil
}

func (s *TagService) view(ctx context.Context, id string) (model.Tag, error) {
	item := model.Tag{
		ID: id,
	}

	err := s.cs.GetDB().NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, TagID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *TagService) viewByDeviceIDAndName(ctx context.Context, deviceID, name string) (model.Tag, error) {
	item := model.Tag{}

	sourceName := consts.DEFAULT_SOURCE
	itemName := name

	if strings.Contains(itemName, ".") {
		splits := strings.Split(itemName, ".")
		if len(splits) != 2 {
			return item, status.Error(codes.InvalidArgument, "Please supply valid tag name")
		}

		sourceName = splits[0]
		itemName = splits[1]
	}

	source, err := s.cs.GetSource().viewByDeviceIDAndName(ctx, deviceID, sourceName)
	if err != nil {
		return item, err
	}

	return s.viewBySourceIDAndName(ctx, source.ID, itemName)
}

func (s *TagService) viewBySourceIDAndName(ctx context.Context, sourceID, name string) (model.Tag, error) {
	item := model.Tag{}

	err := s.cs.GetDB().NewSelect().Model(&item).Where("source_id = ?", sourceID).Where("name = ?", name).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, Tag SourceID: %v, Name: %v", err, sourceID, name)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *TagService) viewBySourceIDAndAddress(ctx context.Context, sourceID, address string) (model.Tag, error) {
	item := model.Tag{}

	err := s.cs.GetDB().NewSelect().Model(&item).Where("source_id = ?", sourceID).Where("address = ?", address).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, Tag SourceID: %v, Address: %v", err, sourceID, address)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *TagService) copyModelToOutput(output *pb.Tag, item *model.Tag) {
	output.Id = item.ID
	output.DeviceId = item.DeviceID
	output.SourceId = item.SourceID
	output.Name = item.Name
	output.Desc = item.Desc
	output.Tags = item.Tags
	output.Type = item.Type
	output.DataType = item.DataType
	output.Address = item.Address
	output.HValue = item.HValue
	output.LValue = item.LValue
	output.Config = item.Config
	output.Status = item.Status
	output.Access = item.Access
	output.Save = item.Save
	output.Upload = item.Upload
	output.Created = item.Created.UnixMilli()
	output.Updated = item.Updated.UnixMilli()
	output.Deleted = item.Deleted.UnixMilli()
}

func (s *TagService) GetTagValue(id string) nson.Value {
	if v := s.valueCache.Get(id); v.IsSome() {
		return v.Unwrap()
	}

	return nson.Null{}
}

func (s *TagService) GetTagValueValue(id string) types.Option[cache.Value[nson.Value]] {
	return s.valueCache.GetValue(id)
}

func (s *TagService) SetTagValue(id string, value nson.Value) {
	s.valueCache.Set(id, value, s.cs.dopts.valueCacheTTL)
}

func (s *TagService) afterUpdate(ctx context.Context, item *model.Tag) error {
	var err error

	err = s.cs.GetSync().setDeviceUpdated(ctx, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *TagService) afterDelete(ctx context.Context, item *model.Tag) error {
	var err error

	err = s.cs.GetSync().setDeviceUpdated(ctx, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *TagService) updateTagValue(ctx context.Context, item *model.Tag, value string) error {
	var err error

	item2 := model.TagValue{
		ID:       item.ID,
		DeviceID: item.DeviceID,
		SourceID: item.SourceID,
		Value:    value,
		Updated:  time.Now(),
	}

	ret, err := s.cs.GetDB().NewUpdate().Model(&item2).WherePK().WhereAllWithDeleted().Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Update: %v", err)
	}

	n, err := ret.RowsAffected()
	if err != nil {
		return status.Errorf(codes.Internal, "RowsAffected: %v", err)
	}

	if n < 1 {
		_, err = s.cs.GetDB().NewInsert().Model(&item2).Exec(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func (s *TagService) afterUpdateValue(ctx context.Context, item *model.Tag, value string) error {
	var err error

	err = s.cs.GetSync().setTagValueUpdated(ctx, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *TagService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Tag, error) {
	var output pb.Tag
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag id")
		}
	}

	item, err := s.viewWithDeleted(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *TagService) viewWithDeleted(ctx context.Context, id string) (model.Tag, error) {
	item := model.Tag{
		ID: id,
	}

	err := s.cs.GetDB().NewSelect().Model(&item).WherePK().WhereAllWithDeleted().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, TagID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *TagService) Pull(ctx context.Context, in *cores.PullTagRequest) (*cores.PullTagResponse, error) {
	var err error
	var output cores.PullTagResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	var items []model.Tag

	query := s.cs.GetDB().NewSelect().Model(&items)

	if len(in.GetDeviceId()) > 0 {
		query.Where("device_id = ?", in.GetDeviceId())
	}

	if len(in.GetSourceId()) > 0 {
		query.Where("source_id = ?", in.GetSourceId())
	}

	err = query.Where("updated > ?", time.UnixMilli(in.GetAfter())).WhereAllWithDeleted().Order("updated ASC").Limit(int(in.GetLimit())).Scan(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Query: %v", err)
	}

	for i := 0; i < len(items); i++ {
		item := pb.Tag{}

		s.copyModelToOutput(&item, &items[i])

		output.Tag = append(output.Tag, &item)
	}

	return &output, nil
}

func (s *TagService) ViewValue(ctx context.Context, in *pb.Id) (*pb.TagValueUpdated, error) {
	var output pb.TagValueUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag id")
		}
	}

	item, err := s.viewValueUpdated(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutputTagValue(&output, &item)

	return &output, nil
}

func (s *TagService) DeleteValue(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag id")
		}
	}

	item, err := s.viewValueUpdated(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	_, err = s.cs.GetDB().NewDelete().Model(&item).WherePK().Exec(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Delete: %v", err)
	}

	output.Bool = true

	return &output, nil
}

func (s *TagService) PullValue(ctx context.Context, in *cores.PullTagValueRequest) (*cores.PullTagValueResponse, error) {
	var err error
	var output cores.PullTagValueResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	var items []model.TagValue

	query := s.cs.GetDB().NewSelect().Model(&items)

	if len(in.GetDeviceId()) > 0 {
		query.Where("device_id = ?", in.GetDeviceId())
	}

	if len(in.GetSourceId()) > 0 {
		query.Where("source_id = ?", in.GetSourceId())
	}

	err = query.Where("updated > ?", time.UnixMilli(in.GetAfter())).Order("updated ASC").Limit(int(in.GetLimit())).Scan(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Query: %v", err)
	}

	for i := 0; i < len(items); i++ {
		item := pb.TagValueUpdated{}

		s.copyModelToOutputTagValue(&item, &items[i])

		output.Tag = append(output.Tag, &item)
	}

	return &output, nil
}

func (s *TagService) viewValueUpdated(ctx context.Context, id string) (model.TagValue, error) {
	item := model.TagValue{
		ID: id,
	}

	err := s.cs.GetDB().NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, TagID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *TagService) copyModelToOutputTagValue(output *pb.TagValueUpdated, item *model.TagValue) {
	output.Id = item.ID
	output.DeviceId = item.DeviceID
	output.SourceId = item.SourceID
	output.Value = item.Value
	output.Updated = item.Updated.UnixMilli()
}
