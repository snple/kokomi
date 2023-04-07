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

type JumperService struct {
	cs *CoreService

	cores.UnimplementedJumperServiceServer
}

func newJumperService(cs *CoreService) *JumperService {
	return &JumperService{
		cs: cs,
	}
}

func (s *JumperService) Create(ctx context.Context, in *cores.Jumper) (*cores.Jumper, error) {
	var output cores.Jumper
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid jumper name")
		}

		if in.GetSrc() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid jumper src")
		}

		if in.GetDst() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid jumper dst")
		}

	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "jumper name min 2 character")
		}

		err = s.cs.GetDB().NewSelect().Model(&model.Jumper{}).Where("name = ?", in.GetName()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			return &output, status.Error(codes.AlreadyExists, "jumper name must be unique")
		}
	}

	// validation src and dst
	{
		_, err = s.cs.GetCable().view(ctx, in.GetSrc())
		if err != nil {
			return &output, err
		}

		_, err = s.cs.GetCable().view(ctx, in.GetDst())
		if err != nil {
			return &output, err
		}

		{
			modelItems := []model.Jumper{}
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
			modelItems := []model.Jumper{}
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
	}

	item := model.Jumper{
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

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *JumperService) Update(ctx context.Context, in *cores.Jumper) (*cores.Jumper, error) {
	var output cores.Jumper
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid jumper id")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid jumper name")
		}
	}

	item, err := s.view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "jumper name min 2 character")
		}

		modelDevice := model.Jumper{}
		err = s.cs.GetDB().NewSelect().Model(&modelDevice).Where("name = ?", in.GetName()).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return &output, status.Errorf(codes.Internal, "Query: %v", err)
			}
		} else {
			if modelDevice.ID != item.ID {
				return &output, status.Error(codes.AlreadyExists, "jumper name must be unique")
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

func (s *JumperService) View(ctx context.Context, in *pb.Id) (*cores.Jumper, error) {
	var output cores.Jumper
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid jumper id")
		}
	}

	item, err := s.view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *JumperService) ViewByName(ctx context.Context, in *pb.Name) (*cores.Jumper, error) {
	var output cores.Jumper
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid jumper name")
		}
	}

	item, err := s.viewByName(ctx, in.GetName())
	if err != nil {
		return nil, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *JumperService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid jumper id")
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

	output.Bool = true

	return &output, nil
}

func (s *JumperService) List(ctx context.Context, in *cores.ListJumperRequest) (*cores.ListJumperResponse, error) {
	var err error
	var output cores.ListJumperResponse

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

	var items []model.Jumper

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
		item := cores.Jumper{}

		s.copyModelToOutput(&item, &items[i])

		output.Jumper = append(output.Jumper, &item)
	}

	return &output, nil
}

func (s *JumperService) view(ctx context.Context, id string) (model.Jumper, error) {
	item := model.Jumper{
		ID: id,
	}

	err := s.cs.GetDB().NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, JumperID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *JumperService) viewByName(ctx context.Context, name string) (model.Jumper, error) {
	item := model.Jumper{}

	err := s.cs.GetDB().NewSelect().Model(&item).Where("name = ?", name).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, jumper Name: %v", err, name)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *JumperService) copyModelToOutput(output *cores.Jumper, item *model.Jumper) {
	output.Id = item.ID
	output.Name = item.Name
	output.Desc = item.Desc
	output.Tags = item.Tags
	output.Type = item.Type
	output.Src = item.SRC
	output.Dst = item.DST
	output.Config = item.Config
	output.Status = item.Status
	output.Created = item.Created.UnixMilli()
	output.Updated = item.Updated.UnixMilli()
}

func (s *JumperService) listBySrcAndStatusON(ctx context.Context, src string) ([]model.Jumper, error) {
	var items []model.Jumper

	err := s.cs.GetDB().NewSelect().Model(&items).
		Where("src = ?", src).Where("status = ?", consts.ON).Scan(ctx)
	if err != nil {
		return items, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return items, nil
}
