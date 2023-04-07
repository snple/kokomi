package edge

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"
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

type WireService struct {
	es *EdgeService

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup

	edges.UnimplementedWireServiceServer
}

func newWireService(es *EdgeService) *WireService {
	ctx, cancel := context.WithCancel(es.Context())

	return &WireService{
		es:     es,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *WireService) Create(ctx context.Context, in *pb.Wire) (*pb.Wire, error) {
	var output pb.Wire
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetCableId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid cable_id")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid wire name")
		}
	}

	item := model.Wire{
		ID:       in.GetId(),
		CableID:  in.GetCableId(),
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

	// cable validation
	{
		_, err = s.es.GetCable().view(ctx, in.GetCableId())
		if err != nil {
			return &output, err
		}

	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "wire name min 2 character")
		}

		err = s.es.GetDB().NewSelect().Model(&model.Wire{}).Where("name = ?", in.GetName()).Where("cable_id = ?", in.GetCableId()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			return &output, status.Error(codes.AlreadyExists, "wire name must be unique")
		}
	}

	// tag validation
	if in.GetTagId() != "" {
		_, err = s.es.GetTag().view(ctx, in.GetTagId())
		if err != nil {
			return &output, err
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

	output.Value, err = s.getWireValue(ctx, &item)
	if err != nil {
		return &output, err
	}

	return &output, nil
}

func (s *WireService) Update(ctx context.Context, in *pb.Wire) (*pb.Wire, error) {
	var output pb.Wire
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid wire id")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid wire name")
		}
	}

	isSync := metadata.IsSync(ctx)

	var item model.Wire

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
			return &output, status.Error(codes.InvalidArgument, "wire name min 2 character")
		}

		modelItem := model.Wire{}
		err = s.es.GetDB().NewSelect().Model(&modelItem).Where("cable_id = ?", item.CableID).Where("name = ?", in.GetName()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			if modelItem.ID != item.ID {
				return &output, status.Error(codes.AlreadyExists, "wire name must be unique")
			}
		}
	}

	// tag validation
	if in.GetTagId() != "" {
		_, err = s.es.GetTag().view(ctx, in.GetTagId())
		if err != nil {
			return &output, err
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

	output.Value, err = s.getWireValue(ctx, &item)
	if err != nil {
		return &output, err
	}

	return &output, nil
}

func (s *WireService) View(ctx context.Context, in *pb.Id) (*pb.Wire, error) {
	var output pb.Wire
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid wire id")
		}
	}

	item, err := s.view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	output.Value, err = s.getWireValue(ctx, &item)
	if err != nil {
		return &output, err
	}

	return &output, nil
}

func (s *WireService) ViewByName(ctx context.Context, in *pb.Name) (*pb.Wire, error) {
	var output pb.Wire
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid wire name")
		}
	}

	item, err := s.viewByName(ctx, in.GetName())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	output.Name = in.GetName()

	output.Value, err = s.getWireValue(ctx, &item)
	if err != nil {
		return &output, err
	}

	return &output, nil
}

func (s *WireService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid wire id")
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

func (s *WireService) List(ctx context.Context, in *edges.ListWireRequest) (*edges.ListWireResponse, error) {
	var err error
	var output edges.ListWireResponse

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

	var items []model.Wire

	query := s.es.GetDB().NewSelect().Model(&items)

	if len(in.GetCableId()) > 0 {
		query.Where("cable_id = ?", in.GetCableId())
	}

	if len(in.GetPage().GetSearch()) > 0 {
		search := fmt.Sprintf("%%%v%%", in.GetPage().GetSearch())

		query.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			q = q.Where(`"name" LIKE ?`, search).
				WhereOr(`"desc" LIKE ?`, search)

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
		item := pb.Wire{}

		s.copyModelToOutput(&item, &items[i])

		item.Value, err = s.getWireValue(ctx, &items[i])
		if err != nil {
			return &output, err
		}

		output.Wire = append(output.Wire, &item)
	}

	return &output, nil
}

func (s *WireService) Clone(ctx context.Context, in *edges.CloneWireRequest) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid wire id")
		}
	}

	err = s.es.getClone().cloneWire(ctx, s.es.GetDB(), in.GetId(), in.GetCableId())
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *WireService) GetValue(ctx context.Context, in *pb.Id) (*pb.WireValue, error) {
	var err error
	var output pb.WireValue

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid wire id")
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

func (s *WireService) SetValue(ctx context.Context, in *pb.WireValue) (*pb.MyBool, error) {
	return s.setValue(ctx, in, true)
}

func (s *WireService) SetValueUnchecked(ctx context.Context, in *pb.WireValue) (*pb.MyBool, error) {
	return s.setValue(ctx, in, false)
}

func (s *WireService) setValue(ctx context.Context, in *pb.WireValue, check bool) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	isSync := false
	if !check {
		isSync = metadata.IsSync(ctx)
	}

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid wire id")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid value")
		}

		if isSync && in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid var value updated")
		}
	}

	// wire
	item, err := s.view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	t := time.Now()
	if isSync {
		t = time.UnixMilli(in.GetUpdated())
	} else {
		if item.Status != consts.ON {
			return &output, status.Errorf(codes.FailedPrecondition, "Wire Status != ON")
		}

		if check {
			if item.Access != consts.ON {
				return &output, status.Errorf(codes.FailedPrecondition, "Wire Access != ON")
			}
		}
	}

	_, err = datatype.DecodeNsonValue(in.GetValue(), item.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	// cable
	{
		cable, err := s.es.GetCable().view(ctx, item.CableID)
		if err != nil {
			return &output, err
		}

		if cable.Status != consts.ON {
			return &output, status.Errorf(codes.FailedPrecondition, "Cable Status != ON")
		}
	}

	if err = s.updateWireValue(ctx, &item, in.GetValue(), t); err != nil {
		return &output, err
	}

	if err = s.afterUpdateValue(ctx, &item, in.GetValue()); err != nil {
		return &output, err
	}

	if item.TagID != "" {
		_, err = s.es.GetTag().SetValue(ctx, &pb.TagValue{Id: item.TagID, Value: in.GetValue()})
		if err != nil {
			return &output, err
		}
	}

	output.Bool = true

	return &output, nil
}

func (s *WireService) GetValueByName(ctx context.Context, in *pb.Name) (*pb.WireNameValue, error) {
	var err error
	var output pb.WireNameValue

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid wire name")
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

func (s *WireService) SetValueByName(ctx context.Context, in *pb.WireNameValue) (*pb.MyBool, error) {
	return s.setValueByName(ctx, in, true)
}

func (s *WireService) SetValueByNameUnchecked(ctx context.Context, in *pb.WireNameValue) (*pb.MyBool, error) {
	return s.setValueByName(ctx, in, false)
}

func (s *WireService) setValueByName(ctx context.Context, in *pb.WireNameValue, check bool) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid wire name")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid value")
		}
	}

	// name
	cableName := consts.DEFAULT_CABLE
	itemName := in.GetName()

	if strings.Contains(itemName, ".") {
		splits := strings.Split(itemName, ".")
		if len(splits) != 2 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid wire name")
		}

		cableName = splits[0]
		itemName = splits[1]
	}

	// cable
	cable, err := s.es.GetCable().viewByName(ctx, cableName)
	if err != nil {
		return &output, err
	}

	if cable.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Cable Status != ON")
	}

	// wire
	item, err := s.viewByCableIDAndName(ctx, cable.ID, itemName)
	if err != nil {
		return &output, err
	}

	if item.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Wire Status != ON")
	}

	if check {
		if item.Access != consts.ON {
			return &output, status.Errorf(codes.FailedPrecondition, "Wire Access != ON")
		}
	}

	_, err = datatype.DecodeNsonValue(in.GetValue(), item.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	if err = s.updateWireValue(ctx, &item, in.GetValue(), time.Now()); err != nil {
		return &output, err
	}

	if err = s.afterUpdateValue(ctx, &item, in.GetValue()); err != nil {
		return &output, err
	}

	if item.TagID != "" {
		_, err = s.es.GetTag().SetValue(ctx, &pb.TagValue{Id: item.TagID, Value: in.GetValue()})
		if err != nil {
			return &output, err
		}
	}

	output.Bool = true

	return &output, nil
}

func (s *WireService) view(ctx context.Context, id string) (model.Wire, error) {
	item := model.Wire{
		ID: id,
	}

	err := s.es.GetDB().NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, WireID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *WireService) viewByName(ctx context.Context, name string) (model.Wire, error) {
	item := model.Wire{}

	cableName := consts.DEFAULT_CABLE
	itemName := name

	if strings.Contains(itemName, ".") {
		splits := strings.Split(itemName, ".")
		if len(splits) != 2 {
			return item, status.Error(codes.InvalidArgument, "Please supply valid wire name")
		}

		cableName = splits[0]
		itemName = splits[1]

	}

	cable, err := s.es.GetCable().viewByName(ctx, cableName)
	if err != nil {
		return item, err
	}

	return s.viewByCableIDAndName(ctx, cable.ID, itemName)
}

func (s *WireService) viewByCableIDAndName(ctx context.Context, cableID, name string) (model.Wire, error) {
	item := model.Wire{}

	err := s.es.GetDB().NewSelect().Model(&item).Where("cable_id = ?", cableID).Where("name = ?", name).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, Wire CableID: %v, Name: %v", err, cableID, name)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *WireService) copyModelToOutput(output *pb.Wire, item *model.Wire) {
	output.Id = item.ID
	output.CableId = item.CableID
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
	output.Created = item.Created.UnixMilli()
	output.Updated = item.Updated.UnixMilli()
	output.Deleted = item.Deleted.UnixMilli()
}

func (s *WireService) afterUpdate(ctx context.Context, item *model.Wire) error {
	var err error

	err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	err = s.es.GetSync().setWireUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *WireService) afterDelete(ctx context.Context, item *model.Wire) error {
	var err error

	err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	err = s.es.GetSync().setWireUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *WireService) getWireValue(ctx context.Context, item *model.Wire) (string, error) {
	item2, err := s.viewValueUpdated(ctx, item.ID)
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

func (s *WireService) updateWireValue(ctx context.Context, item *model.Wire, value string, updated time.Time) error {
	var err error

	item2 := model.WireValue{
		ID:      item.ID,
		CableID: item.CableID,
		Value:   value,
		Updated: updated,
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

func (s *WireService) afterUpdateValue(ctx context.Context, item *model.Wire, value string) error {
	var err error

	err = s.es.GetSync().setWireValueUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *WireService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Wire, error) {
	var output pb.Wire
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid wire id")
		}
	}

	item, err := s.viewWithDeleted(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *WireService) viewWithDeleted(ctx context.Context, id string) (model.Wire, error) {
	item := model.Wire{
		ID: id,
	}

	err := s.es.GetDB().NewSelect().Model(&item).WherePK().WhereAllWithDeleted().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, WireID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *WireService) Pull(ctx context.Context, in *edges.PullWireRequest) (*edges.PullWireResponse, error) {
	var err error
	var output edges.PullWireResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	var items []model.Wire

	query := s.es.GetDB().NewSelect().Model(&items)

	if len(in.GetCableId()) > 0 {
		query.Where("cable_id = ?", in.GetCableId())
	}

	err = query.Where("updated > ?", time.UnixMilli(in.GetAfter())).WhereAllWithDeleted().Order("updated ASC").Limit(int(in.GetLimit())).Scan(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Query: %v", err)
	}

	for i := 0; i < len(items); i++ {
		item := pb.Wire{}

		s.copyModelToOutput(&item, &items[i])

		output.Wire = append(output.Wire, &item)
	}

	return &output, nil
}

func (s *WireService) ViewValue(ctx context.Context, in *pb.Id) (*pb.WireValueUpdated, error) {
	var output pb.WireValueUpdated
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

	s.copyModelToOutputWireValue(&output, &item)

	return &output, nil
}

func (s *WireService) DeleteValue(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid wire id")
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

func (s *WireService) PullValue(ctx context.Context, in *edges.PullWireValueRequest) (*edges.PullWireValueResponse, error) {
	var err error
	var output edges.PullWireValueResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	var items []model.WireValue

	query := s.es.GetDB().NewSelect().Model(&items)

	err = query.Where("updated > ?", time.UnixMilli(in.GetAfter())).Order("updated ASC").Limit(int(in.GetLimit())).Scan(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Query: %v", err)
	}

	for i := 0; i < len(items); i++ {
		item := pb.WireValueUpdated{}

		s.copyModelToOutputWireValue(&item, &items[i])

		output.Wire = append(output.Wire, &item)
	}

	return &output, nil
}

func (s *WireService) viewValueUpdated(ctx context.Context, id string) (model.WireValue, error) {
	item := model.WireValue{
		ID: id,
	}

	err := s.es.GetDB().NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, WireID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *WireService) copyModelToOutputWireValue(output *pb.WireValueUpdated, item *model.WireValue) {
	output.Id = item.ID
	output.CableId = item.CableID
	output.Value = item.Value
	output.Updated = item.Updated.UnixMilli()
}

func (s *WireService) Start() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	ticker := time.NewTicker(s.es.dopts.syncWireValueFromTagValue)
	defer ticker.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			err := s.syncValueFromTagValue(s.ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("syncWireValueFromTag: %v", err)
			}
		}
	}
}

func (s *WireService) Stop() {
	s.cancel()
	s.closeWG.Wait()
}

func (s *WireService) syncValueFromTagValue(ctx context.Context) error {
	var err error
	var items []model.Wire

	err = s.es.GetDB().NewSelect().Model(&items).
		Where("tag_id <> ?", "").Where("status = ?", consts.ON).Scan(ctx)
	if err != nil {
		return err
	}

	for i := 0; i < len(items); i++ {
		if option := s.es.GetTag().GetTagValueValue(items[i].TagID); option.IsSome() {
			// cable
			{
				cable, err := s.es.GetCable().view(ctx, items[i].CableID)
				if err != nil {
					return err
				}

				if cable.Status != consts.ON {
					continue
				}
			}

			tagValue := option.Unwrap()

			value, err := datatype.EncodeNsonValue(tagValue.Data)
			if err != nil {
				return status.Errorf(codes.InvalidArgument, "EncodeValue: %v", err)
			}

			item2, err := s.viewValueUpdated(ctx, items[i].TagID)
			if err != nil {
				if code, ok := status.FromError(err); ok {
					if code.Code() == codes.NotFound {
						goto NOT_FOUND
					}
				}

				return err
			}

			if tagValue.Updated == item2.Updated {
				continue
			}

		NOT_FOUND:

			if err = s.updateWireValue(ctx, &items[i], value, tagValue.Updated); err != nil {
				return err
			}

			if err = s.afterUpdateValue(ctx, &items[i], value); err != nil {
				return err
			}
		}
	}

	return nil
}
