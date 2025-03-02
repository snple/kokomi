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
	"github.com/uptrace/bun"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SlotService struct {
	cs *CoreService

	cores.UnimplementedSlotServiceServer
}

func newSlotService(cs *CoreService) *SlotService {
	return &SlotService{
		cs: cs,
	}
}

func (s *SlotService) Create(ctx context.Context, in *pb.Slot) (*pb.Slot, error) {
	var output pb.Slot
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetNodeId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Slot.NodeID")
		}

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Slot.Name")
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
			return &output, status.Error(codes.InvalidArgument, "Slot.Name min 2 character")
		}

		err = s.cs.GetDB().NewSelect().Model(&model.Slot{}).Where("node_id = ?", in.GetNodeId()).Where("name = ?", in.GetName()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			return &output, status.Error(codes.AlreadyExists, "Slot.Name must be unique")
		}
	}

	item := model.Slot{
		ID:      in.GetId(),
		NodeID:  in.GetNodeId(),
		Name:    in.GetName(),
		Desc:    in.GetDesc(),
		Tags:    in.GetTags(),
		Secret:  in.GetSecret(),
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

func (s *SlotService) Update(ctx context.Context, in *pb.Slot) (*pb.Slot, error) {
	var output pb.Slot
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Slot.ID")
		}

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Slot.Name")
		}
	}

	item, err := s.ViewByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "Slot.Name min 2 character")
		}

		modelItem := model.Slot{}
		err = s.cs.GetDB().NewSelect().Model(&modelItem).Where("node_id = ?", item.NodeID).Where("name = ?", in.GetName()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			if modelItem.ID != item.ID {
				return &output, status.Error(codes.AlreadyExists, "Slot.Name must be unique")
			}
		}
	}

	item.Name = in.GetName()
	item.Desc = in.GetDesc()
	item.Tags = in.GetTags()
	item.Secret = in.GetSecret()
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

func (s *SlotService) View(ctx context.Context, in *pb.Id) (*pb.Slot, error) {
	var output pb.Slot
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Slot.ID")
		}
	}

	item, err := s.ViewByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *SlotService) Name(ctx context.Context, in *cores.SlotNameRequest) (*pb.Slot, error) {
	var output pb.Slot
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Slot.Name")
		}
	}

	item, err := s.ViewByNodeIDAndName(ctx, in.GetNodeId(), in.GetName())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *SlotService) NameFull(ctx context.Context, in *pb.Name) (*pb.Slot, error) {
	var output pb.Slot
	// var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Slot.Name")
		}
	}

	nodeName := consts.DEFAULT_NODE
	itemName := in.GetName()

	if strings.Contains(itemName, ".") {
		splits := strings.Split(itemName, ".")
		if len(splits) != 2 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Slot.Name")
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

func (s *SlotService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Slot.ID")
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

func (s *SlotService) List(ctx context.Context, in *cores.SlotListRequest) (*cores.SlotListResponse, error) {
	var err error
	var output cores.SlotListResponse

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

	var items []model.Slot

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
		item := pb.Slot{}

		s.copyModelToOutput(&item, &items[i])

		output.Slot = append(output.Slot, &item)
	}

	return &output, nil
}

func (s *SlotService) Link(ctx context.Context, in *cores.SlotLinkRequest) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Slot.ID")
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

func (s *SlotService) Clone(ctx context.Context, in *cores.SlotCloneRequest) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Slot.ID")
		}
	}

	err = s.cs.getClone().slot(ctx, s.cs.GetDB(), in.GetId(), in.GetNodeId())
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *SlotService) ViewByID(ctx context.Context, id string) (model.Slot, error) {
	item := model.Slot{
		ID: id,
	}

	err := s.cs.GetDB().NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, Slot.ID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *SlotService) ViewByNodeIDAndName(ctx context.Context, nodeID, name string) (model.Slot, error) {
	item := model.Slot{}

	err := s.cs.GetDB().NewSelect().Model(&item).Where("node_id = ?", nodeID).Where("name = ?", name).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, NodeID: %v, Slot.Name: %v", err, nodeID, name)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *SlotService) copyModelToOutput(output *pb.Slot, item *model.Slot) {
	output.Id = item.ID
	output.NodeId = item.NodeID
	output.Name = item.Name
	output.Desc = item.Desc
	output.Tags = item.Tags
	output.Secret = item.Secret
	output.Config = item.Config
	output.Link = s.cs.GetStatus().GetLink(item.ID)
	output.Status = item.Status
	output.Created = item.Created.UnixMicro()
	output.Updated = item.Updated.UnixMicro()
	output.Deleted = item.Deleted.UnixMicro()
}

func (s *SlotService) afterUpdate(ctx context.Context, item *model.Slot) error {
	var err error

	err = s.cs.GetSync().setNodeUpdated(ctx, s.cs.GetDB(), item.NodeID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Sync.setNodeUpdated: %v", err)
	}

	err = s.cs.GetSyncGlobal().setUpdated(ctx, s.cs.GetDB(), model.SYNC_GLOBAL_SLOT, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "SyncGlobal.setUpdated: %v", err)
	}

	return nil
}

func (s *SlotService) afterDelete(ctx context.Context, item *model.Slot) error {
	var err error

	err = s.cs.GetSync().setNodeUpdated(ctx, s.cs.GetDB(), item.NodeID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Sync.setNodeUpdated: %v", err)
	}

	err = s.cs.GetSyncGlobal().setUpdated(ctx, s.cs.GetDB(), model.SYNC_GLOBAL_SLOT, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "SyncGlobal.setUpdated: %v", err)
	}

	return nil
}

func (s *SlotService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Slot, error) {
	var output pb.Slot
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Slot.ID")
		}
	}

	item, err := s.viewWithDeleted(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *SlotService) viewWithDeleted(ctx context.Context, id string) (model.Slot, error) {
	item := model.Slot{
		ID: id,
	}

	err := s.cs.GetDB().NewSelect().Model(&item).WherePK().WhereAllWithDeleted().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, Slot.ID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *SlotService) Pull(ctx context.Context, in *cores.SlotPullRequest) (*cores.SlotPullResponse, error) {
	var err error
	var output cores.SlotPullResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	var items []model.Slot

	query := s.cs.GetDB().NewSelect().Model(&items)

	if in.GetNodeId() != "" {
		query.Where("node_id = ?", in.GetNodeId())
	}

	err = query.Where("updated > ?", time.UnixMicro(in.GetAfter())).WhereAllWithDeleted().Order("updated ASC").Limit(int(in.GetLimit())).Scan(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Query: %v", err)
	}

	for i := 0; i < len(items); i++ {
		item := pb.Slot{}

		s.copyModelToOutput(&item, &items[i])

		output.Slot = append(output.Slot, &item)
	}

	return &output, nil
}

func (s *SlotService) Sync(ctx context.Context, in *pb.Slot) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid slot_id")
		}

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Slot.Name")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Slot.Updated")
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
				return &output, status.Error(codes.InvalidArgument, "Slot.Name min 2 character")
			}

			err = s.cs.GetDB().NewSelect().Model(&model.Slot{}).Where("node_id = ?", in.GetNodeId()).Where("name = ?", in.GetName()).Scan(ctx)
			if err != nil {
				if err != sql.ErrNoRows {
					return &output, status.Errorf(codes.Internal, "Query: %v", err)
				}
			} else {
				return &output, status.Error(codes.AlreadyExists, "Slot.Name must be unique")
			}
		}

		item := model.Slot{
			ID:      in.GetId(),
			NodeID:  in.GetNodeId(),
			Name:    in.GetName(),
			Desc:    in.GetDesc(),
			Tags:    in.GetTags(),
			Secret:  in.GetSecret(),
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
				return &output, status.Error(codes.InvalidArgument, "Slot.Name min 2 character")
			}

			modelItem := model.Slot{}
			err = s.cs.GetDB().NewSelect().Model(&modelItem).Where("node_id = ?", item.NodeID).Where("name = ?", in.GetName()).Scan(ctx)
			if err != nil {
				if err != sql.ErrNoRows {
					return &output, status.Errorf(codes.Internal, "Query: %v", err)
				}
			} else {
				if modelItem.ID != item.ID {
					return &output, status.Error(codes.AlreadyExists, "Slot.Name must be unique")
				}
			}
		}

		item.Name = in.GetName()
		item.Desc = in.GetDesc()
		item.Tags = in.GetTags()
		item.Secret = in.GetSecret()
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
