package edge

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/danclive/nson-go"
	"github.com/dgraph-io/badger/v4"
	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/edge/model"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/util"
	"github.com/snple/kokomi/util/datatype"
	"github.com/snple/types/cache"
	"github.com/uptrace/bun"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TagService struct {
	es *EdgeService

	cache *cache.Cache[model.Tag]

	edges.UnimplementedTagServiceServer
}

func newTagService(es *EdgeService) *TagService {
	return &TagService{
		es:    es,
		cache: cache.NewCache[model.Tag](nil),
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Source.ID")
		}

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Name")
		}
	}

	// source validation
	{
		_, err = s.es.GetSource().ViewByID(ctx, in.GetSourceId())
		if err != nil {
			return &output, err
		}
	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "Tag.Name min 2 character")
		}

		err = s.es.GetDB().NewSelect().Model(&model.Tag{}).Where("name = ?", in.GetName()).Where("source_id = ?", in.GetSourceId()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			return &output, status.Error(codes.AlreadyExists, "Tag.Name must be unique")
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
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	if item.ID == "" {
		item.ID = util.RandomID()
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

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Name")
		}
	}

	item, err := s.ViewByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "Tag.Name min 2 character")
		}

		modelItem := model.Tag{}
		err = s.es.GetDB().NewSelect().Model(&modelItem).Where("source_id = ?", item.SourceID).Where("name = ?", in.GetName()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			if modelItem.ID != item.ID {
				return &output, status.Error(codes.AlreadyExists, "Tag.Name must be unique")
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
	item.Updated = time.Now()

	_, err = s.es.GetDB().NewUpdate().Model(&item).WherePK().Exec(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Update: %v", err)
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

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}
	}

	item, err := s.ViewByID(ctx, in.GetId())
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

func (s *TagService) Name(ctx context.Context, in *pb.Name) (*pb.Tag, error) {
	var output pb.Tag
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Name")
		}
	}

	item, err := s.ViewByName(ctx, in.GetName())
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

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}
	}

	item, err := s.ViewByID(ctx, in.GetId())
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

func (s *TagService) List(ctx context.Context, in *edges.TagListRequest) (*edges.TagListResponse, error) {
	var err error
	var output edges.TagListResponse

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

	items := make([]model.Tag, 0, 10)

	query := s.es.GetDB().NewSelect().Model(&items)

	if len(in.GetSourceId()) > 0 {
		query.Where("source_id = ?", in.GetSourceId())
	}

	if in.GetPage().GetSearch() != "" {
		search := fmt.Sprintf("%%%v%%", in.GetPage().GetSearch())

		query.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			q = q.Where(`"name" LIKE ?`, search).
				WhereOr(`"desc" LIKE ?`, search).
				WhereOr(`"address" LIKE ?`, search)

			return q
		})
	}

	if in.GetTags() != "" {
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

	if in.GetType() != "" {
		query = query.Where(`type = ?`, in.GetType())
	}

	if in.GetPage().GetOrderBy() != "" && (in.GetPage().GetOrderBy() == "id" || in.GetPage().GetOrderBy() == "name" ||
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

func (s *TagService) Clone(ctx context.Context, in *edges.TagCloneRequest) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}
	}

	err = s.es.getClone().tag(ctx, s.es.GetDB(), in.GetId(), in.GetSourceId())
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *TagService) ViewByID(ctx context.Context, id string) (model.Tag, error) {
	item := model.Tag{
		ID: id,
	}

	err := s.es.GetDB().NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, Tag.ID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *TagService) ViewByName(ctx context.Context, name string) (model.Tag, error) {
	item := model.Tag{}

	sourceName := consts.DEFAULT_SOURCE
	itemName := name

	if strings.Contains(itemName, ".") {
		splits := strings.Split(itemName, ".")
		if len(splits) != 2 {
			return item, status.Error(codes.InvalidArgument, "Please supply valid Tag.Name")
		}

		sourceName = splits[0]
		itemName = splits[1]
	}

	source, err := s.es.GetSource().ViewByName(ctx, sourceName)
	if err != nil {
		return item, err
	}

	return s.ViewBySourceIDAndName(ctx, source.ID, itemName)
}

func (s *TagService) ViewBySourceIDAndName(ctx context.Context, sourceID, name string) (model.Tag, error) {
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

func (s *TagService) ViewBySourceIDAndAddress(ctx context.Context, sourceID, address string) (model.Tag, error) {
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
	output.Created = item.Created.UnixMicro()
	output.Updated = item.Updated.UnixMicro()
	output.Deleted = item.Deleted.UnixMicro()
}

func (s *TagService) afterUpdate(ctx context.Context, _ *model.Tag) error {
	var err error

	err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Sync.setDeviceUpdated: %v", err)
	}

	err = s.es.GetSync().setTagUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Sync.setTagUpdated: %v", err)
	}

	return nil
}

func (s *TagService) afterDelete(ctx context.Context, _ *model.Tag) error {
	var err error

	err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Sync.setDeviceUpdated: %v", err)
	}

	err = s.es.GetSync().setTagUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Sync.setTagUpdated: %v", err)
	}

	return nil
}

// sync

func (s *TagService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Tag, error) {
	var output pb.Tag
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
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
			return item, status.Errorf(codes.NotFound, "Query: %v, Tag.ID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *TagService) Pull(ctx context.Context, in *edges.TagPullRequest) (*edges.TagPullResponse, error) {
	var err error
	var output edges.TagPullResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	items := make([]model.Tag, 0, 10)

	query := s.es.GetDB().NewSelect().Model(&items)

	if in.GetSourceId() != "" {
		query.Where("source_id = ?", in.GetSourceId())
	}

	if in.GetType() != "" {
		query.Where(`type = ?`, in.GetType())
	}

	err = query.Where("updated > ?", time.UnixMicro(in.GetAfter())).WhereAllWithDeleted().Order("updated ASC").Limit(int(in.GetLimit())).Scan(ctx)
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

func (s *TagService) Sync(ctx context.Context, in *pb.Tag) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Name")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Updated")
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
		// name validation
		{
			if len(in.GetName()) < 2 {
				return &output, status.Error(codes.InvalidArgument, "Tag.Name min 2 character")
			}

			err = s.es.GetDB().NewSelect().Model(&model.Tag{}).Where("name = ?", in.GetName()).Where("source_id = ?", in.GetSourceId()).Scan(ctx)
			if err != nil {
				if err != sql.ErrNoRows {
					return &output, status.Errorf(codes.Internal, "Query: %v", err)
				}
			} else {
				return &output, status.Error(codes.AlreadyExists, "Tag.Name must be unique")
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
			Created:  time.UnixMicro(in.GetCreated()),
			Updated:  time.UnixMicro(in.GetUpdated()),
			Deleted:  time.UnixMicro(in.GetDeleted()),
		}

		// source validation
		{
			_, err = s.es.GetSource().viewWithDeleted(ctx, in.GetSourceId())
			if err != nil {
				return &output, err
			}
		}

		_, err = s.es.GetDB().NewInsert().Model(&item).Exec(ctx)
		if err != nil {
			return &output, status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	// update
	if update {
		if in.GetUpdated() <= item.Updated.UnixMicro() {
			return &output, nil
		}

		// name validation
		{
			if len(in.GetName()) < 2 {
				return &output, status.Error(codes.InvalidArgument, "Tag.Name min 2 character")
			}

			modelItem := model.Tag{}
			err = s.es.GetDB().NewSelect().Model(&modelItem).Where("source_id = ?", item.SourceID).Where("name = ?", in.GetName()).Scan(ctx)
			if err != nil {
				if err != sql.ErrNoRows {
					return &output, status.Errorf(codes.Internal, "Query: %v", err)
				}
			} else {
				if modelItem.ID != item.ID {
					return &output, status.Error(codes.AlreadyExists, "Tag.Name must be unique")
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
		item.Updated = time.UnixMicro(in.GetUpdated())
		item.Deleted = time.UnixMicro(in.GetDeleted())

		_, err = s.es.GetDB().NewUpdate().Model(&item).WherePK().WhereAllWithDeleted().Exec(ctx)
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

// cache

func (s *TagService) GC() {
	s.cache.GC()
}

func (s *TagService) ViewFromCacheByID(ctx context.Context, id string) (model.Tag, error) {
	if !s.es.dopts.cache {
		return s.ViewByID(ctx, id)
	}

	if option := s.cache.Get(id); option.IsSome() {
		return option.Unwrap(), nil
	}

	item, err := s.ViewByID(ctx, id)
	if err != nil {
		return item, err
	}

	s.cache.Set(id, item, s.es.dopts.cacheTTL)

	return item, nil
}

func (s *TagService) ViewFromCacheByName(ctx context.Context, name string) (model.Tag, error) {
	if !s.es.dopts.cache {
		return s.ViewByName(ctx, name)
	}

	if option := s.cache.Get(name); option.IsSome() {
		return option.Unwrap(), nil
	}

	item, err := s.ViewByName(ctx, name)
	if err != nil {
		return item, err
	}

	s.cache.Set(name, item, s.es.dopts.cacheTTL)

	return item, nil
}

func (s *TagService) ViewFromCacheBySourceIDAndName(ctx context.Context, sourceID, name string) (model.Tag, error) {
	if !s.es.dopts.cache {
		return s.ViewBySourceIDAndName(ctx, sourceID, name)
	}

	id := sourceID + name

	if option := s.cache.Get(id); option.IsSome() {
		return option.Unwrap(), nil
	}

	item, err := s.ViewBySourceIDAndName(ctx, sourceID, name)
	if err != nil {
		return item, err
	}

	s.cache.Set(id, item, s.es.dopts.cacheTTL)

	return item, nil
}

func (s *TagService) ViewFromCacheBySourceIDAndAddress(ctx context.Context, sourceID, address string) (model.Tag, error) {
	if !s.es.dopts.cache {
		return s.ViewBySourceIDAndAddress(ctx, sourceID, address)
	}

	id := sourceID + address

	if option := s.cache.Get(id); option.IsSome() {
		return option.Unwrap(), nil
	}

	item, err := s.ViewBySourceIDAndAddress(ctx, sourceID, address)
	if err != nil {
		return item, err
	}

	s.cache.Set(id, item, s.es.dopts.cacheTTL)

	return item, nil
}

// value

func (s *TagService) GetValue(ctx context.Context, in *pb.Id) (*pb.TagValue, error) {
	var err error
	var output pb.TagValue

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}
	}

	output.Id = in.GetId()

	item2, err := s.getTagValueUpdated(ctx, in.GetId())
	if err != nil {
		if code, ok := status.FromError(err); ok {
			if code.Code() == codes.NotFound {
				return &output, nil
			}
		}

		return &output, err
	}

	output.Value = item2.Value
	output.Updated = item2.Updated.UnixMicro()

	return &output, nil
}

func (s *TagService) SetValue(ctx context.Context, in *pb.TagValue) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Value")
		}
	}

	// tag
	item, err := s.ViewFromCacheByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	if item.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag.Status != ON")
	}

	_, err = datatype.DecodeNsonValue(in.GetValue(), item.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	// validation device and source
	{
		// source
		{
			source, err := s.es.GetSource().ViewFromCacheByID(ctx, item.SourceID)
			if err != nil {
				return &output, err
			}

			if source.Status != consts.ON {
				return &output, status.Errorf(codes.FailedPrecondition, "Source.Status != ON")
			}
		}
	}

	if err = s.setTagValueUpdated(ctx, &item, in.GetValue(), time.Now()); err != nil {
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

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Name")
		}
	}

	item, err := s.ViewFromCacheByName(ctx, in.GetName())
	if err != nil {
		return &output, err
	}

	output.Id = item.ID
	output.Name = in.GetName()

	item2, err := s.getTagValueUpdated(ctx, item.ID)
	if err != nil {
		if code, ok := status.FromError(err); ok {
			if code.Code() == codes.NotFound {
				return &output, nil
			}
		}

		return &output, err
	}

	output.Value = item2.Value
	output.Updated = item2.Updated.UnixMicro()

	return &output, nil
}

func (s *TagService) SetValueByName(ctx context.Context, in *pb.TagNameValue) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Name")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Value")
		}
	}

	// name
	sourceName := consts.DEFAULT_SOURCE
	itemName := in.GetName()

	if strings.Contains(itemName, ".") {
		splits := strings.Split(itemName, ".")
		if len(splits) != 2 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Name")
		}

		sourceName = splits[0]
		itemName = splits[1]
	}

	// source
	source, err := s.es.GetSource().ViewFromCacheByName(ctx, sourceName)
	if err != nil {
		return &output, err
	}

	if source.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Source.Status != ON")
	}

	// tag
	item, err := s.ViewFromCacheBySourceIDAndName(ctx, source.ID, itemName)
	if err != nil {
		return &output, err
	}

	if item.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag.Status != ON")
	}

	_, err = datatype.DecodeNsonValue(in.GetValue(), item.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	if err = s.setTagValueUpdated(ctx, &item, in.GetValue(), time.Now()); err != nil {
		return &output, err
	}

	if err = s.afterUpdateValue(ctx, &item, in.GetValue()); err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *TagService) getTagValue(ctx context.Context, id string) (string, error) {
	item2, err := s.getTagValueUpdated(ctx, id)
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

func (s *TagService) afterUpdateValue(ctx context.Context, _ *model.Tag, _ string) error {
	var err error

	err = s.es.GetSync().setTagValueUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Sync.setTagValueUpdated: %v", err)
	}

	return nil
}

// sync value

func (s *TagService) ViewValue(ctx context.Context, in *pb.Id) (*pb.TagValueUpdated, error) {
	var output pb.TagValueUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}
	}

	item, err := s.getTagValueUpdated(ctx, in.GetId())
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

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}
	}

	item, err := s.getTagValueUpdated(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	idb, err := nson.IdFromHex(item.ID)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "IdFromHex: %v", err)
	}

	ts := uint64(time.Now().UnixMicro())

	txn := s.es.GetBadgerDB().NewTransactionAt(ts, true)
	defer txn.Discard()

	err = txn.Delete(append([]byte(model.TAG_VALUE_PREFIX), idb...))
	if err != nil {
		return &output, status.Errorf(codes.Internal, "BadgerDB Delete: %v", err)
	}

	err = txn.CommitAt(ts, nil)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "BadgerDB CommitAt: %v", err)
	}

	output.Bool = true

	return &output, nil
}

func (s *TagService) PullValue(ctx context.Context, in *edges.TagPullValueRequest) (*edges.TagPullValueResponse, error) {
	var err error
	var output edges.TagPullValueResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	items := make([]model.TagValue, 0, 10)

	{
		after := time.UnixMicro(in.GetAfter())

		txn := s.es.GetBadgerDB().NewTransactionAt(uint64(time.Now().UnixMicro()), false)
		defer txn.Discard()

		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		opts.SinceTs = uint64(in.GetAfter())

		it := txn.NewIterator(opts)
		defer it.Close()

		prefix := []byte(model.TAG_VALUE_PREFIX)

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			dbitem := it.Item()

			item := model.TagValue{}
			err = dbitem.Value(func(val []byte) error {
				return json.Unmarshal(val, &item)
			})
			if err != nil {
				return &output, status.Errorf(codes.Internal, "BadgerDB view value: %v", err)
			}

			if !item.Updated.After(after) {
				continue
			}

			if in.GetSourceId() != "" && in.GetSourceId() != item.SourceID {
				continue
			}

			items = append(items, item)
		}

		sort.Sort(model.SortTagValue(items))

		if len(items) > int(in.GetLimit()) {
			items = items[0:in.GetLimit()]
		}
	}

	for i := 0; i < len(items); i++ {
		item := pb.TagValueUpdated{}

		s.copyModelToOutputTagValue(&item, &items[i])

		output.Tag = append(output.Tag, &item)
	}

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

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Value")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Value.Updated")
		}
	}

	// tag
	item, err := s.ViewByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	_, err = datatype.DecodeNsonValue(in.GetValue(), item.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	value, err := s.getTagValueUpdated(ctx, in.GetId())
	if err != nil {
		if code, ok := status.FromError(err); ok {
			if code.Code() == codes.NotFound {
				goto UPDATED
			}
		}

		return &output, err
	}

	if in.GetUpdated() <= value.Updated.UnixMicro() {
		return &output, nil
	}

UPDATED:
	if err = s.setTagValueUpdated(ctx, &item, in.GetValue(), time.UnixMicro(in.GetUpdated())); err != nil {
		return &output, err
	}

	if err = s.afterUpdateValue(ctx, &item, in.GetValue()); err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *TagService) setTagValueUpdated(_ context.Context, item *model.Tag, value string, updated time.Time) error {
	item2 := model.TagValue{
		ID:       item.ID,
		SourceID: item.SourceID,
		Value:    value,
		Updated:  updated,
	}

	idb, err := nson.IdFromHex(item.ID)
	if err != nil {
		return status.Errorf(codes.Internal, "IdFromHex: %v", err)
	}

	data, err := json.Marshal(item2)
	if err != nil {
		return status.Errorf(codes.Internal, "json.Marshal: %v", err)
	}

	{
		ts := uint64(updated.UnixMicro())

		txn := s.es.GetBadgerDB().NewTransactionAt(ts, true)
		defer txn.Discard()

		err = txn.Set(append([]byte(model.TAG_VALUE_PREFIX), idb...), data)
		if err != nil {
			return status.Errorf(codes.Internal, "BadgerDB Set: %v", err)
		}

		err = txn.CommitAt(ts, nil)
		if err != nil {
			return status.Errorf(codes.Internal, "BadgerDB CommitAt: %v", err)
		}
	}

	return nil
}

func (s *TagService) getTagValueUpdated(_ context.Context, id string) (model.TagValue, error) {
	item := model.TagValue{
		ID: id,
	}

	idb, err := nson.IdFromHex(item.ID)
	if err != nil {
		return item, status.Errorf(codes.Internal, "IdFromHex: %v", err)
	}

	txn := s.es.GetBadgerDB().NewTransactionAt(uint64(time.Now().UnixMicro()), false)
	defer txn.Discard()

	dbitem, err := txn.Get(append([]byte(model.TAG_VALUE_PREFIX), idb...))
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return item, status.Errorf(codes.NotFound, "Tag.ID: %v", item.ID)
		}
		return item, status.Errorf(codes.Internal, "BadgerDB Get: %v", err)
	}

	err = dbitem.Value(func(val []byte) error {
		return json.Unmarshal(val, &item)
	})
	if err != nil {
		return item, status.Errorf(codes.Internal, "BadgerDB Get Value: %v", err)
	}

	return item, nil
}

func (s *TagService) copyModelToOutputTagValue(output *pb.TagValueUpdated, item *model.TagValue) {
	output.Id = item.ID
	output.SourceId = item.SourceID
	output.Value = item.Value
	output.Updated = item.Updated.UnixMicro()
}

// write

func (s *TagService) GetWrite(ctx context.Context, in *pb.Id) (*pb.TagValue, error) {
	var err error
	var output pb.TagValue

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}
	}

	output.Id = in.GetId()

	item2, err := s.getTagWriteUpdated(ctx, in.GetId())
	if err != nil {
		if code, ok := status.FromError(err); ok {
			if code.Code() == codes.NotFound {
				return &output, nil
			}
		}

		return &output, err
	}

	output.Value = item2.Value
	output.Updated = item2.Updated.UnixMicro()

	return &output, nil
}

func (s *TagService) SetWrite(ctx context.Context, in *pb.TagValue) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Value")
		}
	}

	// tag
	item, err := s.ViewFromCacheByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	if item.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag.Status != ON")
	}

	if item.Access != consts.WRITE {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag.Access != WRITE")
	}

	_, err = datatype.DecodeNsonValue(in.GetValue(), item.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	// validation device and source
	{
		// source
		{
			source, err := s.es.GetSource().ViewFromCacheByID(ctx, item.SourceID)
			if err != nil {
				return &output, err
			}

			if source.Status != consts.ON {
				return &output, status.Errorf(codes.FailedPrecondition, "Source.Status != ON")
			}
		}
	}

	if err = s.setTagWriteUpdated(ctx, &item, in.GetValue(), time.Now()); err != nil {
		return &output, err
	}

	if err = s.afterUpdateWrite(ctx, &item, in.GetValue()); err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *TagService) GetWriteByName(ctx context.Context, in *pb.Name) (*pb.TagNameValue, error) {
	var err error
	var output pb.TagNameValue

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Name")
		}
	}

	item, err := s.ViewFromCacheByName(ctx, in.GetName())
	if err != nil {
		return &output, err
	}

	output.Id = item.ID
	output.Name = in.GetName()

	item2, err := s.getTagWriteUpdated(ctx, item.ID)
	if err != nil {
		if code, ok := status.FromError(err); ok {
			if code.Code() == codes.NotFound {
				return &output, nil
			}
		}

		return &output, err
	}

	output.Value = item2.Value
	output.Updated = item2.Updated.UnixMicro()

	return &output, nil
}

func (s *TagService) SetWriteByName(ctx context.Context, in *pb.TagNameValue) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Name")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Value")
		}
	}

	// name
	sourceName := consts.DEFAULT_SOURCE
	itemName := in.GetName()

	if strings.Contains(itemName, ".") {
		splits := strings.Split(itemName, ".")
		if len(splits) != 2 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Name")
		}

		sourceName = splits[0]
		itemName = splits[1]
	}

	// source
	source, err := s.es.GetSource().ViewFromCacheByName(ctx, sourceName)
	if err != nil {
		return &output, err
	}

	if source.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Source.Status != ON")
	}

	// tag
	item, err := s.ViewFromCacheBySourceIDAndName(ctx, source.ID, itemName)
	if err != nil {
		return &output, err
	}

	if item.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag.Status != ON")
	}

	if item.Access != consts.WRITE {
		return &output, status.Errorf(codes.FailedPrecondition, "Tag.Access != WRITE")
	}

	_, err = datatype.DecodeNsonValue(in.GetValue(), item.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	if err = s.setTagWriteUpdated(ctx, &item, in.GetValue(), time.Now()); err != nil {
		return &output, err
	}

	if err = s.afterUpdateWrite(ctx, &item, in.GetValue()); err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *TagService) getTagWrite(ctx context.Context, id string) (string, error) {
	item2, err := s.getTagWriteUpdated(ctx, id)
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

func (s *TagService) afterUpdateWrite(ctx context.Context, _ *model.Tag, _ string) error {
	var err error

	err = s.es.GetSync().setTagWriteUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Sync.setTagWriteUpdated: %v", err)
	}

	return nil
}

// sync value

func (s *TagService) ViewWrite(ctx context.Context, in *pb.Id) (*pb.TagValueUpdated, error) {
	var output pb.TagValueUpdated
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}
	}

	item, err := s.getTagWriteUpdated(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutputTagValue(&output, &item)

	return &output, nil
}

func (s *TagService) DeleteWrite(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}
	}

	item, err := s.getTagWriteUpdated(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	idb, err := nson.IdFromHex(item.ID)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "IdFromHex: %v", err)
	}

	ts := uint64(time.Now().UnixMicro())

	txn := s.es.GetBadgerDB().NewTransactionAt(ts, true)
	defer txn.Discard()

	err = txn.Delete(append([]byte(model.TAG_VALUE_PREFIX), idb...))
	if err != nil {
		return &output, status.Errorf(codes.Internal, "BadgerDB Delete: %v", err)
	}

	err = txn.CommitAt(ts, nil)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "BadgerDB CommitAt: %v", err)
	}

	output.Bool = true

	return &output, nil
}

func (s *TagService) PullWrite(ctx context.Context, in *edges.TagPullValueRequest) (*edges.TagPullValueResponse, error) {
	var err error
	var output edges.TagPullValueResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	items := make([]model.TagValue, 0, 10)

	{
		after := time.UnixMicro(in.GetAfter())

		txn := s.es.GetBadgerDB().NewTransactionAt(uint64(time.Now().UnixMicro()), false)
		defer txn.Discard()

		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		opts.SinceTs = uint64(in.GetAfter())

		it := txn.NewIterator(opts)
		defer it.Close()

		prefix := []byte(model.TAG_WRITE_PREFIX)

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			dbitem := it.Item()

			item := model.TagValue{}
			err = dbitem.Value(func(val []byte) error {
				return json.Unmarshal(val, &item)
			})
			if err != nil {
				return &output, status.Errorf(codes.Internal, "BadgerDB view value: %v", err)
			}

			if !item.Updated.After(after) {
				continue
			}

			if in.GetSourceId() != "" && in.GetSourceId() != item.SourceID {
				continue
			}

			items = append(items, item)
		}

		sort.Sort(model.SortTagValue(items))

		if len(items) > int(in.GetLimit()) {
			items = items[0:in.GetLimit()]
		}
	}

	for i := 0; i < len(items); i++ {
		item := pb.TagValueUpdated{}

		s.copyModelToOutputTagValue(&item, &items[i])

		output.Tag = append(output.Tag, &item)
	}

	return &output, nil
}

func (s *TagService) SyncWrite(ctx context.Context, in *pb.TagValue) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.ID")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Value")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Tag.Value.Updated")
		}
	}

	// tag
	item, err := s.ViewByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	_, err = datatype.DecodeNsonValue(in.GetValue(), item.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	value, err := s.getTagWriteUpdated(ctx, in.GetId())
	if err != nil {
		if code, ok := status.FromError(err); ok {
			if code.Code() == codes.NotFound {
				goto UPDATED
			}
		}

		return &output, err
	}

	if in.GetUpdated() <= value.Updated.UnixMicro() {
		return &output, nil
	}

UPDATED:
	if err = s.setTagWriteUpdated(ctx, &item, in.GetValue(), time.UnixMicro(in.GetUpdated())); err != nil {
		return &output, err
	}

	if err = s.afterUpdateWrite(ctx, &item, in.GetValue()); err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *TagService) setTagWriteUpdated(_ context.Context, item *model.Tag, value string, updated time.Time) error {
	item2 := model.TagValue{
		ID:       item.ID,
		SourceID: item.SourceID,
		Value:    value,
		Updated:  updated,
	}

	idb, err := nson.IdFromHex(item.ID)
	if err != nil {
		return status.Errorf(codes.Internal, "IdFromHex: %v", err)
	}

	data, err := json.Marshal(item2)
	if err != nil {
		return status.Errorf(codes.Internal, "json.Marshal: %v", err)
	}

	{
		ts := uint64(updated.UnixMicro())

		txn := s.es.GetBadgerDB().NewTransactionAt(ts, true)
		defer txn.Discard()

		err = txn.Set(append([]byte(model.TAG_WRITE_PREFIX), idb...), data)
		if err != nil {
			return status.Errorf(codes.Internal, "BadgerDB Set: %v", err)
		}

		err = txn.CommitAt(ts, nil)
		if err != nil {
			return status.Errorf(codes.Internal, "BadgerDB CommitAt: %v", err)
		}
	}

	return nil
}

func (s *TagService) getTagWriteUpdated(_ context.Context, id string) (model.TagValue, error) {
	item := model.TagValue{
		ID: id,
	}

	idb, err := nson.IdFromHex(item.ID)
	if err != nil {
		return item, status.Errorf(codes.Internal, "IdFromHex: %v", err)
	}

	txn := s.es.GetBadgerDB().NewTransactionAt(uint64(time.Now().UnixMicro()), false)
	defer txn.Discard()

	dbitem, err := txn.Get(append([]byte(model.TAG_WRITE_PREFIX), idb...))
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return item, status.Errorf(codes.NotFound, "Tag.ID: %v", item.ID)
		}
		return item, status.Errorf(codes.Internal, "BadgerDB Get: %v", err)
	}

	err = dbitem.Value(func(val []byte) error {
		return json.Unmarshal(val, &item)
	})
	if err != nil {
		return item, status.Errorf(codes.Internal, "BadgerDB Get Value: %v", err)
	}

	return item, nil
}
