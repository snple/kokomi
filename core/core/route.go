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
	"github.com/uptrace/bun"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RouteService struct {
	cs *CoreService

	cores.UnimplementedRouteServiceServer
}

func newRouteService(cs *CoreService) *RouteService {
	return &RouteService{
		cs: cs,
	}
}

func (s *RouteService) Create(ctx context.Context, in *cores.Route) (*cores.Route, error) {
	var output cores.Route
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Route.Name")
		}

		if in.GetSrc() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Route.Src")
		}

		if in.GetDst() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Route.Dst")
		}

	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "Route.Name min 2 character")
		}

		err = s.cs.GetDB().NewSelect().Model(&model.Route{}).Where("name = ?", in.GetName()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			return &output, status.Error(codes.AlreadyExists, "Route.Name must be unique")
		}
	}

	// validation src and dst
	{
		_, err = s.cs.GetCable().ViewByID(ctx, in.GetSrc())
		if err != nil {
			return &output, err
		}

		_, err = s.cs.GetCable().ViewByID(ctx, in.GetDst())
		if err != nil {
			return &output, err
		}

		{
			modelItems := []model.Route{}
			err = s.cs.GetDB().NewSelect().Model(&modelItems).Where("src = ?", in.GetDst()).Scan(ctx)
			if err != nil {
				return &output, err
			}

			if len(modelItems) > 0 {
				if len(modelItems) == 1 {
					if modelItems[0].DST == in.GetSrc() {
						goto ALLOW1
					}
				}

				return &output, status.Error(codes.InvalidArgument, "Break the rules")
			}

		ALLOW1:
		}

		{
			modelItems := []model.Route{}
			err = s.cs.GetDB().NewSelect().Model(&modelItems).Where("dst = ?", in.GetSrc()).Scan(ctx)
			if err != nil {
				return &output, err
			}

			if len(modelItems) > 0 {
				if len(modelItems) == 1 {
					if modelItems[0].SRC == in.GetDst() {
						goto ALLOW2
					}
				}

				return &output, status.Error(codes.InvalidArgument, "Break the rules")
			}

		ALLOW2:
		}

		{
			err = s.cs.GetDB().NewSelect().Model(&model.Route{}).
				Where("src = ?", in.GetSrc()).Where("dst = ?", in.GetDst()).Scan(ctx)
			if err != nil {
				if err != sql.ErrNoRows {
					return &output, status.Errorf(codes.Internal, "Query: %v", err)
				}
			} else {
				return &output, status.Error(codes.AlreadyExists, "Route.Rrc->Route.Dst must be unique")
			}
		}
	}

	item := model.Route{
		ID:      in.GetId(),
		Name:    in.GetName(),
		Desc:    in.GetDesc(),
		Tags:    in.GetTags(),
		Type:    in.GetType(),
		SRC:     in.GetSrc(),
		DST:     in.GetDst(),
		Config:  in.GetConfig(),
		Status:  in.GetStatus(),
		Created: time.Now(),
		Updated: time.Now(),
	}

	if len(item.ID) == 0 {
		item.ID = util.RandomID()
	}

	_, err = s.cs.GetDB().NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Insert: %v", err)
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *RouteService) Update(ctx context.Context, in *cores.Route) (*cores.Route, error) {
	var output cores.Route
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Route.ID")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Route.Name")
		}
	}

	item, err := s.ViewByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "Route.Name min 2 character")
		}

		modelItem := model.Route{}
		err = s.cs.GetDB().NewSelect().Model(&modelItem).Where("name = ?", in.GetName()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			if modelItem.ID != item.ID {
				return &output, status.Error(codes.AlreadyExists, "Route.Name must be unique")
			}
		}
	}

	item.Name = in.GetName()
	item.Desc = in.GetDesc()
	item.Tags = in.GetTags()
	item.Type = in.GetType()
	// item.SRC = in.GetSrc() // Cannot be updated
	// item.DST = in.GetDst() // Cannot be updated
	item.Config = in.GetConfig()
	item.Status = in.GetStatus()
	item.Updated = time.Now()

	_, err = s.cs.GetDB().NewUpdate().Model(&item).WherePK().Exec(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Update: %v", err)
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *RouteService) View(ctx context.Context, in *pb.Id) (*cores.Route, error) {
	var output cores.Route
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Route.ID")
		}
	}

	item, err := s.ViewByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *RouteService) Name(ctx context.Context, in *pb.Name) (*cores.Route, error) {
	var output cores.Route
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Route.Name")
		}
	}

	item, err := s.viewByName(ctx, in.GetName())
	if err != nil {
		return nil, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *RouteService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Route.ID")
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

	output.Bool = true

	return &output, nil
}

func (s *RouteService) List(ctx context.Context, in *cores.RouteListRequest) (*cores.RouteListResponse, error) {
	var err error
	var output cores.RouteListResponse

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

	var items []model.Route

	query := s.cs.GetDB().NewSelect().Model(&items)

	if in.GetSrc() != "" {
		query.Where("src = ?", in.GetSrc())
	}

	if in.GetDst() != "" {
		query.Where("dst = ?", in.GetDst())
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
		item := cores.Route{}

		s.copyModelToOutput(&item, &items[i])

		output.Route = append(output.Route, &item)
	}

	return &output, nil
}

func (s *RouteService) ViewByID(ctx context.Context, id string) (model.Route, error) {
	item := model.Route{
		ID: id,
	}

	err := s.cs.GetDB().NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, Route.ID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *RouteService) viewByName(ctx context.Context, name string) (model.Route, error) {
	item := model.Route{}

	err := s.cs.GetDB().NewSelect().Model(&item).Where("name = ?", name).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, Route.Name: %v", err, name)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *RouteService) copyModelToOutput(output *cores.Route, item *model.Route) {
	output.Id = item.ID
	output.Name = item.Name
	output.Desc = item.Desc
	output.Tags = item.Tags
	output.Type = item.Type
	output.Src = item.SRC
	output.Dst = item.DST
	output.Config = item.Config
	output.Status = item.Status
	output.Created = item.Created.UnixMicro()
	output.Updated = item.Updated.UnixMicro()
}

func (s *RouteService) listBySrcAndStatusON(ctx context.Context, src string) ([]model.Route, error) {
	var items []model.Route

	err := s.cs.GetDB().NewSelect().Model(&items).
		Where("src = ?", src).Where("status = ?", consts.ON).Scan(ctx)
	if err != nil {
		return items, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return items, nil
}
