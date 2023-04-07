package core

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/core/model"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
	"github.com/snple/kokomi/util"
	"github.com/snple/kokomi/util/metadata"
	"github.com/uptrace/bun"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OptionService struct {
	cs *CoreService

	cores.UnimplementedOptionServiceServer
}

func newOptionService(cs *CoreService) *OptionService {
	return &OptionService{
		cs: cs,
	}
}

func (s *OptionService) Create(ctx context.Context, in *pb.Option) (*pb.Option, error) {
	var output pb.Option
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid option name")
		}
	}

	item := model.Option{
		ID:       in.GetId(),
		DeviceID: in.GetDeviceId(),
		Name:     in.GetName(),
		Desc:     in.GetDesc(),
		Tags:     in.GetTags(),
		Type:     in.GetType(),
		Value:    in.GetValue(),
		Status:   in.GetStatus(),
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	// device validation
	{
		_, err = s.cs.GetDevice().view(ctx, in.GetDeviceId())
		if err != nil {
			return &output, err
		}
	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "option name min 2 character")
		}

		err = s.cs.GetDB().NewSelect().Model(&model.Option{}).Where("device_id = ?", in.GetDeviceId()).Where("name = ?", in.GetName()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			return &output, status.Error(codes.AlreadyExists, "option name must be unique")
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

	return &output, nil
}

func (s *OptionService) Update(ctx context.Context, in *pb.Option) (*pb.Option, error) {
	var output pb.Option
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid option id")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid option name")
		}
	}

	isSync := metadata.IsSync(ctx)

	var item model.Option

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
			return &output, status.Error(codes.InvalidArgument, "option name min 2 character")
		}

		modelOption := model.Option{}
		err = s.cs.GetDB().NewSelect().Model(&modelOption).Where("device_id = ?", item.DeviceID).Where("name = ?", in.GetName()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			if modelOption.ID != item.ID {
				return &output, status.Error(codes.AlreadyExists, "option name must be unique")
			}
		}
	}

	item.Name = in.GetName()
	item.Desc = in.GetDesc()
	item.Tags = in.GetTags()
	item.Type = in.GetType()
	item.Value = in.GetValue()
	item.Status = in.GetStatus()
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

	return &output, nil
}

func (s *OptionService) View(ctx context.Context, in *pb.Id) (*pb.Option, error) {
	var output pb.Option
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid option id")
		}
	}

	item, err := s.view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *OptionService) ViewByName(ctx context.Context, in *cores.ViewOptionByNameRequest) (*pb.Option, error) {
	var output pb.Option
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
			return &output, status.Error(codes.InvalidArgument, "Please supply valid option name")
		}
	}

	item, err := s.viewByDeviceIDAndName(ctx, in.GetDeviceId(), in.GetName())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *OptionService) ViewByNameFull(ctx context.Context, in *pb.Name) (*pb.Option, error) {
	var output pb.Option
	// var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid option name")
		}
	}

	deviceName := consts.DEFAULT_DEVICE
	itemName := in.GetName()

	if strings.Contains(itemName, ".") {
		splits := strings.Split(itemName, ".")
		if len(splits) != 2 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid option name")
		}

		deviceName = splits[0]
		itemName = splits[1]
	}

	device, err := s.cs.GetDevice().viewByName(ctx, deviceName)
	if err != nil {
		return &output, err
	}

	item, err := s.viewByDeviceIDAndName(ctx, device.ID, itemName)
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	output.Name = in.GetName()

	return &output, nil
}

func (s *OptionService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid option id")
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

func (s *OptionService) List(ctx context.Context, in *cores.ListOptionRequest) (*cores.ListOptionResponse, error) {
	var err error
	var output cores.ListOptionResponse

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

	var items []model.Option

	query := s.cs.GetDB().NewSelect().Model(&items)

	if len(in.GetDeviceId()) > 0 {
		query.Where("device_id = ?", in.GetDeviceId())
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
		item := pb.Option{}

		s.copyModelToOutput(&item, &items[i])

		output.Option = append(output.Option, &item)
	}

	return &output, nil
}

func (s *OptionService) Clone(ctx context.Context, in *cores.CloneOptionRequest) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid option id")
		}
	}

	err = s.cs.getClone().cloneOption(ctx, s.cs.GetDB(), in.GetId(), in.GetDeviceId())
	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *OptionService) view(ctx context.Context, id string) (model.Option, error) {
	item := model.Option{
		ID: id,
	}

	err := s.cs.GetDB().NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, OptionID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *OptionService) viewByDeviceIDAndName(ctx context.Context, deviceID, name string) (model.Option, error) {
	item := model.Option{}

	err := s.cs.GetDB().NewSelect().Model(&item).Where("device_id = ?", deviceID).Where("name = ?", name).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, Option DeviceID: %v, Name: %v", err, deviceID, name)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *OptionService) copyModelToOutput(output *pb.Option, item *model.Option) {
	output.Id = item.ID
	output.DeviceId = item.DeviceID
	output.Name = item.Name
	output.Desc = item.Desc
	output.Tags = item.Tags
	output.Type = item.Type
	output.Value = item.Value
	output.Status = item.Status
	output.Created = item.Created.UnixMilli()
	output.Updated = item.Updated.UnixMilli()
	output.Deleted = item.Deleted.UnixMilli()
}

func (s *OptionService) afterUpdate(ctx context.Context, item *model.Option) error {
	var err error

	err = s.cs.GetSync().setDeviceUpdated(ctx, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *OptionService) afterDelete(ctx context.Context, item *model.Option) error {
	var err error

	err = s.cs.GetSync().setDeviceUpdated(ctx, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *OptionService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Option, error) {
	var output pb.Option
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid option id")
		}
	}

	item, err := s.viewWithDeleted(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *OptionService) viewWithDeleted(ctx context.Context, id string) (model.Option, error) {
	item := model.Option{
		ID: id,
	}

	err := s.cs.GetDB().NewSelect().Model(&item).WherePK().WhereAllWithDeleted().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, OptionID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *OptionService) Pull(ctx context.Context, in *cores.PullOptionRequest) (*cores.PullOptionResponse, error) {
	var err error
	var output cores.PullOptionResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	var items []model.Option

	query := s.cs.GetDB().NewSelect().Model(&items)

	if len(in.GetDeviceId()) > 0 {
		query.Where("device_id = ?", in.GetDeviceId())
	}

	err = query.Where("updated > ?", time.UnixMilli(in.GetAfter())).WhereAllWithDeleted().Order("updated ASC").Limit(int(in.GetLimit())).Scan(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Query: %v", err)
	}

	for i := 0; i < len(items); i++ {
		item := pb.Option{}

		s.copyModelToOutput(&item, &items[i])

		output.Option = append(output.Option, &item)
	}

	return &output, nil
}
