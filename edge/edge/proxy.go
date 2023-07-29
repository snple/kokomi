package edge

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/snple/kokomi/edge/model"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/pb/nodes"
	"github.com/uptrace/bun"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProxyService struct {
	es *EdgeService

	edges.UnimplementedProxyServiceServer
}

func newProxyService(es *EdgeService) *ProxyService {
	return &ProxyService{
		es: es,
	}
}

// func (s *ProxyService) Create(ctx context.Context, in *pb.Proxy) (*pb.Proxy, error) {
// 	var output pb.Proxy
// 	var err error

// 	// basic validation
// 	{
// 		if in == nil {
// 			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
// 		}

// 		if len(in.GetName()) == 0 {
// 			return &output, status.Error(codes.InvalidArgument, "Please supply valid Proxy.Name")
// 		}
// 	}

// 	// name validation
// 	{
// 		if len(in.GetName()) < 2 {
// 			return &output, status.Error(codes.InvalidArgument, "Proxy.Name min 2 character")
// 		}

// 		err = s.es.GetDB().NewSelect().Model(&model.Proxy{}).Where("name = ?", in.GetName()).Scan(ctx)
// 		if err != nil {
// 			if err != sql.ErrNoRows {
// 				return &output, status.Errorf(codes.Internal, "Query: %v", err)
// 			}
// 		} else {
// 			return &output, status.Error(codes.AlreadyExists, "Proxy.Name must be unique")
// 		}
// 	}

// 	item := model.Proxy{
// 		ID:      in.GetId(),
// 		Name:    in.GetName(),
// 		Desc:    in.GetDesc(),
// 		Tags:    in.GetTags(),
// 		Type:    in.GetType(),
// 		Network: in.GetNetwork(),
// 		Address: in.GetAddress(),
// 		Target:  in.GetTarget(),
// 		Config:  in.GetConfig(),
// 		Status:  in.GetStatus(),
// 		Created: time.Now(),
// 		Updated: time.Now(),
// 	}

// 	if len(item.ID) == 0 {
// 		item.ID = util.RandomID()
// 	}

// 	_, err = s.es.GetDB().NewInsert().Model(&item).Exec(ctx)
// 	if err != nil {
// 		return &output, status.Errorf(codes.Internal, "Insert: %v", err)
// 	}

// 	if err = s.afterUpdate(ctx, &item); err != nil {
// 		return &output, err
// 	}

// 	s.copyModelToOutput(&output, &item)

// 	return &output, nil
// }

// func (s *ProxyService) Update(ctx context.Context, in *pb.Proxy) (*pb.Proxy, error) {
// 	var output pb.Proxy
// 	var err error

// 	// basic validation
// 	{
// 		if in == nil {
// 			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
// 		}

// 		if len(in.GetId()) == 0 {
// 			return &output, status.Error(codes.InvalidArgument, "Please supply valid Proxy.ID")
// 		}

// 		if len(in.GetName()) == 0 {
// 			return &output, status.Error(codes.InvalidArgument, "Please supply valid Proxy.Name")
// 		}
// 	}

// 	item, err := s.view(ctx, in.GetId())
// 	if err != nil {
// 		return &output, err
// 	}

// 	// name validation
// 	{
// 		if len(in.GetName()) < 2 {
// 			return &output, status.Error(codes.InvalidArgument, "Proxy.Name min 2 character")
// 		}

// 		modelItem := model.Proxy{}
// 		err = s.es.GetDB().NewSelect().Model(&modelItem).Where("name = ?", in.GetName()).Scan(ctx)
// 		if err != nil {
// 			if err != sql.ErrNoRows {
// 				return &output, status.Errorf(codes.Internal, "Query: %v", err)
// 			}
// 		} else {
// 			if modelItem.ID != item.ID {
// 				return &output, status.Error(codes.AlreadyExists, "Proxy.Name must be unique")
// 			}
// 		}
// 	}

// 	item.Name = in.GetName()
// 	item.Desc = in.GetDesc()
// 	item.Tags = in.GetTags()
// 	item.Type = in.GetType()
// 	item.Network = in.GetNetwork()
// 	item.Address = in.GetAddress()
// 	item.Target = in.GetTarget()
// 	item.Config = in.GetConfig()
// 	item.Status = in.GetStatus()
// 	item.Updated = time.Now()

// 	_, err = s.es.GetDB().NewUpdate().Model(&item).WherePK().Exec(ctx)
// 	if err != nil {
// 		return &output, status.Errorf(codes.Internal, "Update: %v", err)
// 	}

// 	if err = s.afterUpdate(ctx, &item); err != nil {
// 		return &output, err
// 	}

// 	s.copyModelToOutput(&output, &item)

// 	return &output, nil
// }

func (s *ProxyService) View(ctx context.Context, in *pb.Id) (*pb.Proxy, error) {
	var output pb.Proxy
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Proxy.ID")
		}
	}

	item, err := s.view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *ProxyService) Name(ctx context.Context, in *pb.Name) (*pb.Proxy, error) {
	var output pb.Proxy
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Proxy.Name")
		}
	}

	item, err := s.viewByName(ctx, in.GetName())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

// func (s *ProxyService) Delete(ctx context.Context, in *pb.Id) (*pb.MyBool, error) {
// 	var err error
// 	var output pb.MyBool

// 	// basic validation
// 	{
// 		if in == nil {
// 			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
// 		}

// 		if len(in.GetId()) == 0 {
// 			return &output, status.Error(codes.InvalidArgument, "Please supply valid Proxy.ID")
// 		}
// 	}

// 	item, err := s.view(ctx, in.GetId())
// 	if err != nil {
// 		return &output, err
// 	}

// 	item.Updated = time.Now()
// 	item.Deleted = time.Now()

// 	_, err = s.es.GetDB().NewUpdate().Model(&item).Column("updated", "deleted").WherePK().Exec(ctx)
// 	if err != nil {
// 		return &output, status.Errorf(codes.Internal, "Delete: %v", err)
// 	}

// 	if err = s.afterDelete(ctx, &item); err != nil {
// 		return &output, err
// 	}

// 	output.Bool = true

// 	return &output, nil
// }

func (s *ProxyService) List(ctx context.Context, in *edges.ProxyListRequest) (*edges.ProxyListResponse, error) {
	var err error
	var output edges.ProxyListResponse

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

	var items []model.Proxy

	query := s.es.GetDB().NewSelect().Model(&items)

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
		item := pb.Proxy{}

		s.copyModelToOutput(&item, &items[i])

		output.Proxy = append(output.Proxy, &item)
	}

	return &output, nil
}

func (s *ProxyService) Link(ctx context.Context, in *edges.ProxyLinkRequest) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	item, err := s.view(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.es.GetStatus().SetLink(item.ID, in.GetStatus())

	{
		if option := s.es.GetNode(); option.IsSome() {
			node := option.Unwrap()

			ctx := node.SetToken(context.Background())
			request := &nodes.ProxyLinkRequest{Id: in.GetId(), Status: in.GetStatus()}
			_, err := node.ProxyServiceClient().Link(ctx, request)
			if err != nil {
				return &output, err
			}
		}
	}

	output.Bool = true

	return &output, nil
}

func (s *ProxyService) Clone(ctx context.Context, in *edges.ProxyCloneRequest) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Proxy.ID")
		}
	}

	tx, err := s.es.GetDB().BeginTx(ctx, nil)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "BeginTx: %v", err)
	}
	var done bool
	defer func() {
		if !done {
			_ = tx.Rollback()
		}
	}()

	err = s.es.getClone().proxy(ctx, tx, in.GetId())
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

func (s *ProxyService) view(ctx context.Context, id string) (model.Proxy, error) {
	item := model.Proxy{
		ID: id,
	}

	err := s.es.GetDB().NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, Proxy.ID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *ProxyService) viewByName(ctx context.Context, name string) (model.Proxy, error) {
	item := model.Proxy{
		Name: name,
	}

	err := s.es.GetDB().NewSelect().Model(&item).Where("name = ?", name).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, Proxy.Name: %v", err, name)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *ProxyService) copyModelToOutput(output *pb.Proxy, item *model.Proxy) {
	output.Id = item.ID
	output.Name = item.Name
	output.Desc = item.Desc
	output.Tags = item.Tags
	output.Type = item.Type
	output.Network = item.Network
	output.Address = item.Address
	output.Target = item.Target
	output.Config = item.Config
	output.Link = s.es.GetStatus().GetLink(item.ID)
	output.Status = item.Status
	output.Created = item.Created.UnixMicro()
	output.Updated = item.Updated.UnixMicro()
	output.Deleted = item.Deleted.UnixMicro()
}

func (s *ProxyService) afterUpdate(ctx context.Context, item *model.Proxy) error {
	var err error

	err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	err = s.es.GetSync().setProxyUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *ProxyService) afterDelete(ctx context.Context, item *model.Proxy) error {
	var err error

	err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	err = s.es.GetSync().setProxyUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *ProxyService) ViewWithDeleted(ctx context.Context, in *pb.Id) (*pb.Proxy, error) {
	var output pb.Proxy
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Proxy.ID")
		}
	}

	item, err := s.viewWithDeleted(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *ProxyService) viewWithDeleted(ctx context.Context, id string) (model.Proxy, error) {
	item := model.Proxy{
		ID: id,
	}

	err := s.es.GetDB().NewSelect().Model(&item).WherePK().WhereAllWithDeleted().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, Proxy.ID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *ProxyService) Pull(ctx context.Context, in *edges.ProxyPullRequest) (*edges.ProxyPullResponse, error) {
	var err error
	var output edges.ProxyPullResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	output.After = in.GetAfter()
	output.Limit = in.GetLimit()

	var items []model.Proxy

	query := s.es.GetDB().NewSelect().Model(&items)

	if in.GetType() != "" {
		query.Where(`type = ?`, in.GetType())
	}

	err = query.Where("updated > ?", time.UnixMicro(in.GetAfter())).WhereAllWithDeleted().Order("updated ASC").Limit(int(in.GetLimit())).Scan(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Query: %v", err)
	}

	for i := 0; i < len(items); i++ {
		item := pb.Proxy{}

		s.copyModelToOutput(&item, &items[i])

		output.Proxy = append(output.Proxy, &item)
	}

	return &output, nil
}

func (s *ProxyService) Sync(ctx context.Context, in *pb.Proxy) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Proxy.ID")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Proxy.Name")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Proxy.Updated")
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
				return &output, status.Error(codes.InvalidArgument, "Proxy.Name min 2 character")
			}

			err = s.es.GetDB().NewSelect().Model(&model.Proxy{}).Where("name = ?", in.GetName()).Scan(ctx)
			if err != nil {
				if err != sql.ErrNoRows {
					return &output, status.Errorf(codes.Internal, "Query: %v", err)
				}
			} else {
				return &output, status.Error(codes.AlreadyExists, "Proxy.Name must be unique")
			}
		}

		item := model.Proxy{
			ID:      in.GetId(),
			Name:    in.GetName(),
			Desc:    in.GetDesc(),
			Tags:    in.GetTags(),
			Type:    in.GetType(),
			Network: in.GetNetwork(),
			Address: in.GetAddress(),
			Target:  in.GetTarget(),
			Config:  in.GetConfig(),
			Status:  in.GetStatus(),
			Created: time.UnixMicro(in.GetCreated()),
			Updated: time.UnixMicro(in.GetUpdated()),
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
				return &output, status.Error(codes.InvalidArgument, "Proxy.Name min 2 character")
			}

			modelItem := model.Proxy{}
			err = s.es.GetDB().NewSelect().Model(&modelItem).Where("name = ?", in.GetName()).Scan(ctx)
			if err != nil {
				if err != sql.ErrNoRows {
					return &output, status.Errorf(codes.Internal, "Query: %v", err)
				}
			} else {
				if modelItem.ID != item.ID {
					return &output, status.Error(codes.AlreadyExists, "Proxy.Name must be unique")
				}
			}
		}

		item.Name = in.GetName()
		item.Desc = in.GetDesc()
		item.Tags = in.GetTags()
		item.Type = in.GetType()
		item.Network = in.GetNetwork()
		item.Address = in.GetAddress()
		item.Target = in.GetTarget()
		item.Config = in.GetConfig()
		item.Status = in.GetStatus()
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
