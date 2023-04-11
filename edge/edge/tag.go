package edge

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/edge/model"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/util"
	"github.com/snple/kokomi/util/datatype"
	"github.com/snple/kokomi/util/metadata"
	"github.com/uptrace/bun"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TagService struct {
	es *EdgeService

	edges.UnimplementedTagServiceServer
}

func newTagService(es *EdgeService) *TagService {
	return &TagService{
		es: es,
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
		_, err = s.es.GetSource().view(ctx, in.GetSourceId())
		if err != nil {
			return &output, err
		}
	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "tag name min 2 character")
		}

		err = s.es.GetDB().NewSelect().Model(&model.Tag{}).Where("name = ?", in.GetName()).Where("source_id = ?", in.GetSourceId()).Scan(ctx)
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

	_, err = s.es.GetDB().NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Insert: %v", err)
	}

	if err = s.afterUpdate(ctx, &item); err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	output.Value, err = s.getTagValue(ctx, item.ID)
	if err != nil {
		return &output, err
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
		err = s.es.GetDB().NewSelect().Model(&modelItem).Where("source_id = ?", item.SourceID).Where("name = ?", in.GetName()).Scan(ctx)
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

		_, err = s.es.GetDB().NewUpdate().Model(&item).WherePK().WhereAllWithDeleted().Exec(ctx)
		if err != nil {
			return &output, status.Errorf(codes.Internal, "Update: %v", err)
		}
	} else {
		_, err = s.es.GetDB().NewUpdate().Model(&item).WherePK().Exec(ctx)
		if err != nil {
			return &output, status.Errorf(codes.Internal, "Update: %v", err)
		}
	}

	if err = s.afterUpdate(ctx, &item); err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	output.Value, err = s.getTagValue(ctx, item.ID)
	if err != nil {
		return &output, err
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

	output.Value, err = s.getTagValue(ctx, item.ID)
	if err != nil {
		return &output, err
	}

	return &output, nil
}

func (s *TagService) ViewByName(ctx context.Context, in *pb.Name) (*pb.Tag, error) {
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

	item, err := s.viewByName(ctx, in.GetName())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	output.Name = in.GetName()

	output.Value, err = s.getTagValue(ctx, item.ID)
	if err != nil {
		return &output, err
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

	_, err = s.es.GetDB().NewUpdate().Model(&item).Column("updated", "deleted").WherePK().Exec(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Delete: %v", err)
	}

	if err = s.afterDelete(ctx, &item); err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *TagService) List(ctx context.Context, in *edges.ListTagRequest) (*edges.ListTagResponse, error) {
	var err error
	var output edges.ListTagResponse

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

	query := s.es.GetDB().NewSelect().Model(&items)

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

		item.Value, err = s.getTagValue(ctx, items[i].ID)
		if err != nil {
			return &output, err
		}

		output.Tag = append(output.Tag, &item)
	}

	return &output, nil
}

func (s *TagService) Clone(ctx context.Context, in *edges.CloneTagRequest) (*pb.MyBool, error) {
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

	err = s.es.getClone().cloneTag(ctx, s.es.GetDB(), in.GetId(), in.GetSourceId())
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

	item2, err := s.viewValueUpdated(ctx, in.GetId())
	if err != nil {
		if code, ok := status.FromError(err); ok {
			if code.Code() == codes.NotFound {
				return &output, nil
			}
		}

		return &output, err
	}

	output.Value = item2.Value
	output.Updated = item2.Updated.UnixMilli()

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

	_, err = datatype.DecodeNsonValue(in.GetValue(), item.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	// validation device and source
	{
		// source
		{
			source, err := s.es.GetSource().view(ctx, item.SourceID)
			if err != nil {
				return &output, err
			}

			if source.Status != consts.ON {
				return &output, status.Errorf(codes.FailedPrecondition, "Source Status != ON")
			}
		}
	}

	if err = s.updateTagValue(ctx, &item, in.GetValue(), time.Now()); err != nil {
		return &output, err
	}

	if err = s.afterUpdateValue(ctx, &item, in.GetValue()); err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *TagService) SyncValue(ctx context.Context, in *pb.TagValue) (*pb.MyBool, error) {
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

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid var value updated")
		}
	}

	// tag
	item, err := s.view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	_, err = datatype.DecodeNsonValue(in.GetValue(), item.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	if err = s.updateTagValue(ctx, &item, in.GetValue(), time.UnixMilli(in.GetUpdated())); err != nil {
		return &output, err
	}

	if err = s.afterUpdateValue(ctx, &item, in.GetValue()); err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *TagService) GetValueByName(ctx context.Context, in *pb.Name) (*pb.TagNameValue, error) {
	var err error
	var output pb.TagNameValue

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag name")
		}
	}

	item, err := s.viewByName(ctx, in.GetName())
	if err != nil {
		return &output, err
	}

	output.Name = in.GetName()

	item2, err := s.viewValueUpdated(ctx, item.ID)
	if err != nil {
		if code, ok := status.FromError(err); ok {
			if code.Code() == codes.NotFound {
				return &output, nil
			}
		}

		return &output, err
	}

	output.Value = item2.Value
	output.Updated = item2.Updated.UnixMilli()

	return &output, nil
}

func (s *TagService) SetValueByName(ctx context.Context, in *pb.TagNameValue) (*pb.MyBool, error) {
	return s.setValueByName(ctx, in, true)
}

func (s *TagService) SetValueByNameUnchecked(ctx context.Context, in *pb.TagNameValue) (*pb.MyBool, error) {
	return s.setValueByName(ctx, in, false)
}

func (s *TagService) setValueByName(ctx context.Context, in *pb.TagNameValue, check bool) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid tag name")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid value")
		}
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
	source, err := s.es.GetSource().viewByName(ctx, sourceName)
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

	_, err = datatype.DecodeNsonValue(in.GetValue(), item.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	if err = s.updateTagValue(ctx, &item, in.GetValue(), time.Now()); err != nil {
		return &output, err
	}

	if err = s.afterUpdateValue(ctx, &item, in.GetValue()); err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *TagService) view(ctx context.Context, id string) (model.Tag, error) {
	item := model.Tag{
		ID: id,
	}

	err := s.es.GetDB().NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, TagID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *TagService) viewByName(ctx context.Context, name string) (model.Tag, error) {
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

	source, err := s.es.GetSource().viewByName(ctx, sourceName)
	if err != nil {
		return item, err
	}

	return s.viewBySourceIDAndName(ctx, source.ID, itemName)
}

func (s *TagService) viewBySourceIDAndName(ctx context.Context, sourceID, name string) (model.Tag, error) {
	item := model.Tag{}

	err := s.es.GetDB().NewSelect().Model(&item).Where("source_id = ?", sourceID).Where("name = ?", name).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, SourceID: %v, Name: %v", err, sourceID, name)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *TagService) viewBySourceIDAndAddress(ctx context.Context, sourceID, address string) (model.Tag, error) {
	item := model.Tag{}

	err := s.es.GetDB().NewSelect().Model(&item).Where("source_id = ?", sourceID).Where("address = ?", address).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, SourceID: %v, Address: %v", err, sourceID, address)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *TagService) copyModelToOutput(output *pb.Tag, item *model.Tag) {
	output.Id = item.ID
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

func (s *TagService) afterUpdate(ctx context.Context, item *model.Tag) error {
	var err error

	err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	err = s.es.GetSync().setTagUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *TagService) afterDelete(ctx context.Context, item *model.Tag) error {
	var err error

	err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	err = s.es.GetSync().setTagUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *TagService) getTagValue(ctx context.Context, id string) (string, error) {
	item2, err := s.viewValueUpdated(ctx, id)
	if err != nil {
		if code, ok := status.FromError(err); ok {
			if code.Code() == codes.NotFound {
				return "", nil
			}
		}

		return "", err
	}

	return item2.Value, nil
}

func (s *TagService) updateTagValue(ctx context.Context, item *model.Tag, value string, updated time.Time) error {
	var err error

	item2 := model.TagValue{
		ID:       item.ID,
		SourceID: item.SourceID,
		Value:    value,
		Updated:  updated,
	}

	ret, err := s.es.GetDB().NewUpdate().Model(&item2).WherePK().WhereAllWithDeleted().Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Update: %v", err)
	}

	n, err := ret.RowsAffected()
	if err != nil {
		return status.Errorf(codes.Internal, "RowsAffected: %v", err)
	}

	if n < 1 {
		_, err = s.es.GetDB().NewInsert().Model(&item2).Exec(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func (s *TagService) afterUpdateValue(ctx context.Context, item *model.Tag, value string) error {
	var err error

	err = s.es.GetSync().setTagValueUpdated(ctx, time.Now())
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

	err := s.es.GetDB().NewSelect().Model(&item).WherePK().WhereAllWithDeleted().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, TagID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *TagService) Pull(ctx context.Context, in *edges.PullTagRequest) (*edges.PullTagResponse, error) {
	var err error
	var output edges.PullTagResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	var items []model.Tag

	query := s.es.GetDB().NewSelect().Model(&items)

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

	_, err = s.es.GetDB().NewDelete().Model(&item).WherePK().Exec(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Delete: %v", err)
	}

	output.Bool = true

	return &output, nil
}

func (s *TagService) PullValue(ctx context.Context, in *edges.PullTagValueRequest) (*edges.PullTagValueResponse, error) {
	var err error
	var output edges.PullTagValueResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	var items []model.TagValue

	query := s.es.GetDB().NewSelect().Model(&items)

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

	err := s.es.GetDB().NewSelect().Model(&item).WherePK().Scan(ctx)
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
	output.SourceId = item.SourceID
	output.Value = item.Value
	output.Updated = item.Updated.UnixMilli()
}
