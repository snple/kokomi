package core

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/snple/beacon/consts"
	"github.com/snple/beacon/core/model"
	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/cores"
	"github.com/snple/beacon/util"
	"github.com/snple/types/cache"
	"github.com/uptrace/bun"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SourceService struct {
	cs *CoreService

	cache *cache.Cache[model.Source]

	cores.UnimplementedSourceServiceServer
}

func newSourceService(cs *CoreService) *SourceService {
	return &SourceService{
		cs:    cs,
		cache: cache.NewCache[model.Source](nil),
	}
}

func (s *SourceService) Create(ctx context.Context, in *pb.Source) (*pb.Source, error) {
	var output pb.Source
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetNodeId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Source.NodeID")
		}

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Source.Name")
		}
	}

	// node validation
	{
		_, err = s.cs.GetNode().ViewByID(ctx, in.GetNodeId())
		if err != nil {
			return &output, err
		}
	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "Source.Name min 2 character")
		}

		err = s.cs.GetDB().NewSelect().Model(&model.Source{}).Where("node_id = ?", in.GetNodeId()).Where("name = ?", in.GetName()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			return &output, status.Error(codes.AlreadyExists, "Source.Name must be unique")
		}
	}

	item := model.Source{
		ID:      in.GetId(),
		NodeID:  in.GetNodeId(),
		Name:    in.GetName(),
		Desc:    in.GetDesc(),
		Tags:    in.GetTags(),
		Source:  in.GetSource(),
		Params:  in.GetParams(),
		Config:  in.GetConfig(),
		Status:  in.GetStatus(),
		Created: time.Now(),
		Updated: time.Now(),
	}

	if item.ID == "" {
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

	return &output, nil
}

func (s *SourceService) Update(ctx context.Context, in *pb.Source) (*pb.Source, error) {
	var output pb.Source
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Source.ID")
		}

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Source.Name")
		}
	}

	item, err := s.ViewByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "Source.Name min 2 character")
		}

		modelItem := model.Source{}
		err = s.cs.GetDB().NewSelect().Model(&modelItem).Where("node_id = ?", item.NodeID).Where("name = ?", in.GetName()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			if modelItem.ID != item.ID {
				return &output, status.Error(codes.AlreadyExists, "Source.Name must be unique")
			}
		}
	}

	item.Name = in.GetName()
	item.Desc = in.GetDesc()
	item.Tags = in.GetTags()
	item.Source = in.GetSource()
	item.Params = in.GetParams()
	item.Config = in.GetConfig()
	item.Status = in.GetStatus()
	item.Updated = time.Now()

	_, err = s.cs.GetDB().NewUpdate().Model(&item).WherePK().Exec(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Update: %v", err)
	}

	if err = s.afterUpdate(ctx, &item); err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *SourceService) View(ctx context.Context, in *pb.Id) (*pb.Source, error) {
	var output pb.Source
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Source.ID")
		}
	}

	item, err := s.ViewByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *SourceService) Name(ctx context.Context, in *cores.SourceNameRequest) (*pb.Source, error) {
	var output pb.Source
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetNodeId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid NodeID")
		}

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Source.Name")
		}
	}

	item, err := s.ViewByNodeIDAndName(ctx, in.GetNodeId(), in.GetName())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *SourceService) NameFull(ctx context.Context, in *pb.Name) (*pb.Source, error) {
	var output pb.Source
	// var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Source.Name")
		}
	}

	nodeName := consts.DEFAULT_NODE
	itemName := in.GetName()

	if strings.Contains(itemName, ".") {
		splits := strings.Split(itemName, ".")
		if len(splits) != 2 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Source.Name")
		}

		nodeName = splits[0]
		itemName = splits[1]
	}

	node, err := s.cs.GetNode().ViewByName(ctx, nodeName)
	if err != nil {
		return &output, err
	}

	item, err := s.ViewByNodeIDAndName(ctx, node.ID, itemName)
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	output.Name = in.GetName()

	return &output, nil
}

func (s *SourceService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Source.ID")
		}
	}

	item, err := s.ViewByID(ctx, in.GetId())
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

func (s *SourceService) List(ctx context.Context, in *cores.SourceListRequest) (*cores.SourceListResponse, error) {
	var err error
	var output cores.SourceListResponse

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

	var items []model.Source

	query := s.cs.GetDB().NewSelect().Model(&items)

	if in.GetNodeId() != "" {
		query.Where("node_id = ?", in.GetNodeId())
	}

	if in.GetPage().GetSearch() != "" {
		search := fmt.Sprintf("%%%v%%", in.GetPage().GetSearch())

		query.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			q = q.Where(`"name" LIKE ?`, search).
				WhereOr(`"desc" LIKE ?`, search)

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

	if len(in.GetSource()) > 0 {
		query = query.Where(`source = ?`, in.GetSource())
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
		item := pb.Source{}

		s.copyModelToOutput(&item, &items[i])

		output.Source = append(output.Source, &item)
	}

	return &output, nil
}

func (s *SourceService) Link(ctx context.Context, in *cores.SourceLinkRequest) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Source.ID")
		}
	}

	item, err := s.ViewByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.cs.GetStatus().SetLink(item.ID, in.GetStatus())

	output.Bool = true

	return &output, nil
}

func (s *SourceService) Clone(ctx context.Context, in *cores.SourceCloneRequest) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Source.ID")
		}
	}

	tx, err := s.cs.GetDB().BeginTx(ctx, nil)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "BeginTx: %v", err)
	}
	var done bool
	defer func() {
		if !done {
			_ = tx.Rollback()
		}
	}()

	err = s.cs.getClone().source(ctx, tx, in.GetId(), in.GetNodeId())
	if err != nil {
		return &output, err
	}

	done = true
	err = tx.Commit()
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Commit: %v", err)
	}

	output.Bool = true

	return &output, nil
}

func (s *SourceService) ViewByID(ctx context.Context, id string) (model.Source, error) {
	item := model.Source{
		ID: id,
	}

	err := s.cs.GetDB().NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, Source.ID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *SourceService) ViewByNodeIDAndName(ctx context.Context, nodeID, name string) (model.Source, error) {
	item := model.Source{}

	err := s.cs.GetDB().NewSelect().Model(&item).Where("node_id = ?", nodeID).Where("name = ?", name).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, NodeID: %v, Source.Name: %v", err, nodeID, name)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *SourceService) copyModelToOutput(output *pb.Source, item *model.Source) {
	output.Id = item.ID
	output.NodeId = item.NodeID
	output.Name = item.Name
	output.Desc = item.Desc
	output.Tags = item.Tags
	output.Source = item.Source
	output.Params = item.Params
	output.Config = item.Config
	output.Link = s.cs.GetStatus().GetLink(item.ID)
	output.Status = item.Status
	output.Created = item.Created.UnixMicro()
	output.Updated = item.Updated.UnixMicro()
	output.Deleted = item.Deleted.UnixMicro()
}

func (s *SourceService) afterUpdate(ctx context.Context, item *model.Source) error {
	var err error

	err = s.cs.GetSync().setNodeUpdated(ctx, s.cs.GetDB(), item.NodeID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Sync.setNodeUpdated: %v", err)
	}

	err = s.cs.GetSyncGlobal().setUpdated(ctx, s.cs.GetDB(), model.SYNC_GLOBAL_SOURCE, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "SyncGlobal.setUpdated: %v", err)
	}

	return nil
}

func (s *SourceService) afterDelete(ctx context.Context, item *model.Source) error {
	var err error

	err = s.cs.GetSync().setNodeUpdated(ctx, s.cs.GetDB(), item.NodeID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Sync.setNodeUpdated: %v", err)
	}

	err = s.cs.GetSyncGlobal().setUpdated(ctx, s.cs.GetDB(), model.SYNC_GLOBAL_SOURCE, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "SyncGlobal.setUpdated: %v", err)
	}

	return nil
}

// sync

func (s *SourceService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Source, error) {
	var output pb.Source
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Source.ID")
		}
	}

	item, err := s.viewWithDeleted(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *SourceService) viewWithDeleted(ctx context.Context, id string) (model.Source, error) {
	item := model.Source{
		ID: id,
	}

	err := s.cs.GetDB().NewSelect().Model(&item).WherePK().WhereAllWithDeleted().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, Source.ID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *SourceService) Pull(ctx context.Context, in *cores.SourcePullRequest) (*cores.SourcePullResponse, error) {
	var err error
	var output cores.SourcePullResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	var items []model.Source

	query := s.cs.GetDB().NewSelect().Model(&items)

	if in.GetNodeId() != "" {
		query.Where("node_id = ?", in.GetNodeId())
	}

	if in.GetSource() != "" {
		query.Where(`source = ?`, in.GetSource())
	}

	err = query.Where("updated > ?", time.UnixMicro(in.GetAfter())).WhereAllWithDeleted().Order("updated ASC").Limit(int(in.GetLimit())).Scan(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Query: %v", err)
	}

	for i := 0; i < len(items); i++ {
		item := pb.Source{}

		s.copyModelToOutput(&item, &items[i])

		output.Source = append(output.Source, &item)
	}

	return &output, nil
}

func (s *SourceService) Sync(ctx context.Context, in *pb.Source) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Source.ID")
		}

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Source.Name")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Source.Updated")
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
		// node validation
		{
			_, err = s.cs.GetNode().viewWithDeleted(ctx, in.GetNodeId())
			if err != nil {
				return &output, err
			}
		}

		// name validation
		{
			if len(in.GetName()) < 2 {
				return &output, status.Error(codes.InvalidArgument, "Source.Name min 2 character")
			}

			err = s.cs.GetDB().NewSelect().Model(&model.Source{}).Where("node_id = ?", in.GetNodeId()).Where("name = ?", in.GetName()).Scan(ctx)
			if err != nil {
				if err != sql.ErrNoRows {
					return &output, status.Errorf(codes.Internal, "Query: %v", err)
				}
			} else {
				return &output, status.Error(codes.AlreadyExists, "Source.Name must be unique")
			}
		}

		item := model.Source{
			ID:      in.GetId(),
			NodeID:  in.GetNodeId(),
			Name:    in.GetName(),
			Desc:    in.GetDesc(),
			Tags:    in.GetTags(),
			Source:  in.GetSource(),
			Params:  in.GetParams(),
			Config:  in.GetConfig(),
			Status:  in.GetStatus(),
			Created: time.UnixMicro(in.GetCreated()),
			Updated: time.UnixMicro(in.GetUpdated()),
			Deleted: time.UnixMicro(in.GetDeleted()),
		}

		_, err = s.cs.GetDB().NewInsert().Model(&item).Exec(ctx)
		if err != nil {
			return &output, status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	// update
	if update {
		if in.GetNodeId() != item.NodeID {
			return &output, status.Error(codes.NotFound, "Query: in.GetNodeId() != item.NodeID")
		}

		if in.GetUpdated() <= item.Updated.UnixMicro() {
			return &output, nil
		}

		// name validation
		{
			if len(in.GetName()) < 2 {
				return &output, status.Error(codes.InvalidArgument, "Source.Name min 2 character")
			}

			modelItem := model.Source{}
			err = s.cs.GetDB().NewSelect().Model(&modelItem).Where("node_id = ?", item.NodeID).Where("name = ?", in.GetName()).Scan(ctx)
			if err != nil {
				if err != sql.ErrNoRows {
					return &output, status.Errorf(codes.Internal, "Query: %v", err)
				}
			} else {
				if modelItem.ID != item.ID {
					return &output, status.Error(codes.AlreadyExists, "Source.Name must be unique")
				}
			}
		}

		item.Name = in.GetName()
		item.Desc = in.GetDesc()
		item.Tags = in.GetTags()
		item.Source = in.GetSource()
		item.Params = in.GetParams()
		item.Config = in.GetConfig()
		item.Status = in.GetStatus()
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

// cache

func (s *SourceService) GC() {
	s.cache.GC()
}

func (s *SourceService) ViewFromCacheByID(ctx context.Context, id string) (model.Source, error) {
	if !s.cs.dopts.cache {
		return s.ViewByID(ctx, id)
	}

	if option := s.cache.Get(id); option.IsSome() {
		return option.Unwrap(), nil
	}

	item, err := s.ViewByID(ctx, id)
	if err != nil {
		return item, err
	}

	s.cache.Set(id, item, s.cs.dopts.cacheTTL)

	return item, nil
}

func (s *SourceService) ViewFromCacheByNodeIDAndName(ctx context.Context, nodeID, name string) (model.Source, error) {
	if !s.cs.dopts.cache {
		return s.ViewByNodeIDAndName(ctx, nodeID, name)
	}

	id := nodeID + name

	if option := s.cache.Get(id); option.IsSome() {
		return option.Unwrap(), nil
	}

	item, err := s.ViewByNodeIDAndName(ctx, nodeID, name)
	if err != nil {
		return item, err
	}

	s.cache.Set(id, item, s.cs.dopts.cacheTTL)

	return item, nil
}
