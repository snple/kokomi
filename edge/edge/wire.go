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
	"github.com/uptrace/bun"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WireService struct {
	es *EdgeService

	edges.UnimplementedWireServiceServer
}

func newWireService(es *EdgeService) *WireService {
	return &WireService{
		es: es,
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.CableID")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.Name")
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
			return &output, status.Error(codes.InvalidArgument, "Wire.Name min 2 character")
		}

		err = s.es.GetDB().NewSelect().Model(&model.Wire{}).Where("name = ?", in.GetName()).Where("cable_id = ?", in.GetCableId()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			return &output, status.Error(codes.AlreadyExists, "Wire.Name must be unique")
		}
	}

	if len(item.ID) == 0 {
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.ID")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.Name")
		}
	}

	item, err := s.view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "Wire.Name min 2 character")
		}

		modelItem := model.Wire{}
		err = s.es.GetDB().NewSelect().Model(&modelItem).Where("cable_id = ?", item.CableID).Where("name = ?", in.GetName()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			if modelItem.ID != item.ID {
				return &output, status.Error(codes.AlreadyExists, "Wire.Name must be unique")
			}
		}
	}

	item.Name = in.GetName()
	item.Desc = in.GetDesc()
	item.Tags = in.GetTags()
	item.Type = in.GetType()
	item.DataType = in.GetDataType()
	item.HValue = in.GetHValue()
	item.LValue = in.GetLValue()
	item.Config = in.GetConfig()
	item.Status = in.GetStatus()
	item.Access = in.GetAccess()
	item.Save = in.GetSave()
	item.Updated = time.Now()

	_, err = s.es.GetDB().NewUpdate().Model(&item).WherePK().Exec(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Update: %v", err)
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.ID")
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.Name")
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.ID")
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

	items := make([]model.Wire, 0, 10)

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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.ID")
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.ID")
		}
	}

	output.Id = in.GetId()

	item2, err := s.getWireValueUpdated(ctx, in.GetId())
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

func (s *WireService) SetValue(ctx context.Context, in *pb.WireValue) (*pb.MyBool, error) {
	return s.setValue(ctx, in, true)
}

func (s *WireService) SetValueUnchecked(ctx context.Context, in *pb.WireValue) (*pb.MyBool, error) {
	return s.setValue(ctx, in, false)
}

func (s *WireService) setValue(ctx context.Context, in *pb.WireValue, check bool) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.ID")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.Value")
		}
	}

	// wire
	item, err := s.view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	if item.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Wire.Status != ON")
	}

	if check {
		if item.Access != consts.ON {
			return &output, status.Errorf(codes.FailedPrecondition, "Wire.Access != ON")
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
			return &output, status.Errorf(codes.FailedPrecondition, "Cable.Status != ON")
		}
	}

	if err = s.setWireValueUpdated(ctx, &item, in.GetValue(), time.Now()); err != nil {
		return &output, err
	}

	if err = s.afterUpdateValue(ctx, &item, in.GetValue()); err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *WireService) SyncValue(ctx context.Context, in *pb.WireValue) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.ID")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.Value")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.Value.Updated")
		}
	}

	// wire
	item, err := s.view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	_, err = datatype.DecodeNsonValue(in.GetValue(), item.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	item2, err := s.getWireValueUpdated(ctx, in.GetId())
	if err != nil {
		if code, ok := status.FromError(err); ok {
			if code.Code() == codes.NotFound {
				goto UPDATED
			}
		}

		return &output, err
	}

	if in.GetUpdated() <= item2.Updated.UnixMicro() {
		return &output, nil
	}

UPDATED:
	if err = s.setWireValueUpdated(ctx, &item, in.GetValue(), time.UnixMicro(in.GetUpdated())); err != nil {
		return &output, err
	}

	if err = s.afterUpdateValue(ctx, &item, in.GetValue()); err != nil {
		return &output, err
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.Name")
		}
	}

	item, err := s.viewByName(ctx, in.GetName())
	if err != nil {
		return &output, err
	}

	output.Name = in.GetName()

	item2, err := s.getWireValueUpdated(ctx, item.ID)
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.Name")
		}

		if len(in.GetValue()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.Value")
		}
	}

	// name
	cableName := consts.DEFAULT_CABLE
	itemName := in.GetName()

	if strings.Contains(itemName, ".") {
		splits := strings.Split(itemName, ".")
		if len(splits) != 2 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.Name")
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
		return &output, status.Errorf(codes.FailedPrecondition, "Cable.Status != ON")
	}

	// wire
	item, err := s.ViewByCableIDAndName(ctx, cable.ID, itemName)
	if err != nil {
		return &output, err
	}

	if item.Status != consts.ON {
		return &output, status.Errorf(codes.FailedPrecondition, "Wire.Status != ON")
	}

	if check {
		if item.Access != consts.ON {
			return &output, status.Errorf(codes.FailedPrecondition, "Wire.Access != ON")
		}
	}

	_, err = datatype.DecodeNsonValue(in.GetValue(), item.ValueTag())
	if err != nil {
		return &output, status.Errorf(codes.InvalidArgument, "DecodeValue: %v", err)
	}

	if err = s.setWireValueUpdated(ctx, &item, in.GetValue(), time.Now()); err != nil {
		return &output, err
	}

	if err = s.afterUpdateValue(ctx, &item, in.GetValue()); err != nil {
		return &output, err
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
			return item, status.Errorf(codes.NotFound, "Query: %v, Wire.ID: %v", err, item.ID)
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
			return item, status.Error(codes.InvalidArgument, "Please supply valid Wire.Name")
		}

		cableName = splits[0]
		itemName = splits[1]

	}

	cable, err := s.es.GetCable().viewByName(ctx, cableName)
	if err != nil {
		return item, err
	}

	return s.ViewByCableIDAndName(ctx, cable.ID, itemName)
}

func (s *WireService) ViewByCableIDAndName(ctx context.Context, cableID, name string) (model.Wire, error) {
	item := model.Wire{}

	err := s.es.GetDB().NewSelect().Model(&item).Where("cable_id = ?", cableID).Where("name = ?", name).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, CableID: %v, Wire.Name: %v", err, cableID, name)
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
	output.Config = item.Config
	output.Status = item.Status
	output.Access = item.Access
	output.Save = item.Save
	output.Created = item.Created.UnixMicro()
	output.Updated = item.Updated.UnixMicro()
	output.Deleted = item.Deleted.UnixMicro()
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

func (s *WireService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Wire, error) {
	var output pb.Wire
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.ID")
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
			return item, status.Errorf(codes.NotFound, "Query: %v, Wire.ID: %v", err, item.ID)
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

	items := make([]model.Wire, 0, 10)

	query := s.es.GetDB().NewSelect().Model(&items)

	if in.GetCableId() != "" {
		query.Where("cable_id = ?", in.GetCableId())
	}

	if in.GetType() != "" {
		query.Where(`type = ?`, in.GetType())
	}

	err = query.Where("updated > ?", time.UnixMicro(in.GetAfter())).WhereAllWithDeleted().Order("updated ASC").Limit(int(in.GetLimit())).Scan(ctx)
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

func (s *WireService) Sync(ctx context.Context, in *pb.Wire) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid wire_id")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.Name")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.Updated")
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
		// cable validation
		{
			_, err = s.es.GetCable().viewWithDeleted(ctx, in.GetCableId())
			if err != nil {
				return &output, err
			}

		}

		// name validation
		{
			if len(in.GetName()) < 2 {
				return &output, status.Error(codes.InvalidArgument, "Wire.Name min 2 character")
			}

			err = s.es.GetDB().NewSelect().Model(&model.Wire{}).Where("name = ?", in.GetName()).Where("cable_id = ?", in.GetCableId()).Scan(ctx)
			if err != nil {
				if err != sql.ErrNoRows {
					return &output, status.Errorf(codes.Internal, "Query: %v", err)
				}
			} else {
				return &output, status.Error(codes.AlreadyExists, "Wire.Name must be unique")
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
			Config:   in.GetConfig(),
			Status:   in.GetStatus(),
			Access:   in.GetAccess(),
			Save:     in.GetSave(),
			Created:  time.UnixMicro(in.GetCreated()),
			Updated:  time.UnixMicro(in.GetUpdated()),
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
				return &output, status.Error(codes.InvalidArgument, "Wire.Name min 2 character")
			}

			modelItem := model.Wire{}
			err = s.es.GetDB().NewSelect().Model(&modelItem).Where("cable_id = ?", item.CableID).Where("name = ?", in.GetName()).Scan(ctx)
			if err != nil {
				if err != sql.ErrNoRows {
					return &output, status.Errorf(codes.Internal, "Query: %v", err)
				}
			} else {
				if modelItem.ID != item.ID {
					return &output, status.Error(codes.AlreadyExists, "Wire.Name must be unique")
				}
			}
		}

		item.Name = in.GetName()
		item.Desc = in.GetDesc()
		item.Tags = in.GetTags()
		item.Type = in.GetType()
		item.DataType = in.GetDataType()
		item.HValue = in.GetHValue()
		item.LValue = in.GetLValue()
		item.Config = in.GetConfig()
		item.Status = in.GetStatus()
		item.Access = in.GetAccess()
		item.Save = in.GetSave()
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

func (s *WireService) getWireValue(ctx context.Context, item *model.Wire) (string, error) {
	item2, err := s.getWireValueUpdated(ctx, item.ID)
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

func (s *WireService) afterUpdateValue(ctx context.Context, item *model.Wire, value string) error {
	var err error

	err = s.es.GetSync().setWireValueUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.ID")
		}
	}

	item, err := s.getWireValueUpdated(ctx, in.GetId())
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Wire.ID")
		}
	}

	item, err := s.getWireValueUpdated(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	idb, err := nson.MessageIdFromHex(item.ID)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "MessageIdFromHex: %v", err)
	}

	ts := uint64(time.Now().UnixMicro())

	txn := s.es.GetBadgerDB().NewTransactionAt(ts, true)
	defer txn.Discard()

	err = txn.Delete(append([]byte(model.WIRE_VALUE_PREFIX), idb...))
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

	items := make([]model.WireValue, 0, 10)

	{
		after := time.UnixMicro(in.GetAfter())

		txn := s.es.GetBadgerDB().NewTransactionAt(uint64(time.Now().UnixMicro()), false)
		defer txn.Discard()

		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		opts.SinceTs = uint64(in.GetAfter())

		it := txn.NewIterator(opts)
		defer it.Close()

		prefix := []byte(model.WIRE_VALUE_PREFIX)

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			dbitem := it.Item()

			item := model.WireValue{}
			err = dbitem.Value(func(val []byte) error {
				return json.Unmarshal(val, &item)
			})
			if err != nil {
				return &output, status.Errorf(codes.Internal, "BadgerDB view value: %v", err)
			}

			if !item.Updated.After(after) {
				continue
			}

			if in.GetCableId() != "" && in.GetCableId() != item.CableID {
				continue
			}

			items = append(items, item)
		}

		sort.Sort(sortWireValue(items))

		if len(items) > int(in.GetLimit()) {
			items = items[0:in.GetLimit()]
		}
	}

	for i := 0; i < len(items); i++ {
		item := pb.WireValueUpdated{}

		s.copyModelToOutputWireValue(&item, &items[i])

		output.Wire = append(output.Wire, &item)
	}

	return &output, nil
}

type sortWireValue []model.WireValue

func (a sortWireValue) Len() int           { return len(a) }
func (a sortWireValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortWireValue) Less(i, j int) bool { return a[i].Updated.Before(a[j].Updated) }

func (s *WireService) setWireValueUpdated(ctx context.Context, item *model.Wire, value string, updated time.Time) error {
	item2 := model.WireValue{
		ID:      item.ID,
		CableID: item.CableID,
		Value:   value,
		Updated: updated,
	}

	idb, err := nson.MessageIdFromHex(item.ID)
	if err != nil {
		return status.Errorf(codes.Internal, "MessageIdFromHex: %v", err)
	}

	data, err := json.Marshal(item2)
	if err != nil {
		return status.Errorf(codes.Internal, "json.Marshal: %v", err)
	}

	{
		ts := uint64(updated.UnixMicro())

		txn := s.es.GetBadgerDB().NewTransactionAt(ts, true)
		defer txn.Discard()

		err = txn.Set(append([]byte(model.WIRE_VALUE_PREFIX), idb...), data)
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

func (s *WireService) getWireValueUpdated(ctx context.Context, id string) (model.WireValue, error) {
	item := model.WireValue{
		ID: id,
	}

	idb, err := nson.MessageIdFromHex(item.ID)
	if err != nil {
		return item, status.Errorf(codes.Internal, "MessageIdFromHex: %v", err)
	}

	txn := s.es.GetBadgerDB().NewTransactionAt(uint64(time.Now().UnixMicro()), false)
	defer txn.Discard()

	dbitem, err := txn.Get(append([]byte(model.WIRE_VALUE_PREFIX), idb...))
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return item, status.Errorf(codes.NotFound, "Wire.ID: %v", item.ID)
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

func (s *WireService) copyModelToOutputWireValue(output *pb.WireValueUpdated, item *model.WireValue) {
	output.Id = item.ID
	output.CableId = item.CableID
	output.Value = item.Value
	output.Updated = item.Updated.UnixMicro()
}
