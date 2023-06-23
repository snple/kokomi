package edge

import (
	"context"
	"database/sql"
	"time"

	"github.com/snple/kokomi/edge/model"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeviceService struct {
	es *EdgeService

	edges.UnimplementedDeviceServiceServer
}

func newDeviceService(es *EdgeService) *DeviceService {
	return &DeviceService{
		es: es,
	}
}

// func (s *DeviceService) Create(ctx context.Context, in *pb.Device) (*pb.Device, error) {
// 	var output pb.Device
// 	var err error

// 	// basic validation
// 	{
// 		if in == nil {
// 			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
// 		}

// 		if len(in.GetName()) == 0 {
// 			return &output, status.Error(codes.InvalidArgument, "Please supply valid Device.Name")
// 		}
// 	}

// 	// name validation
// 	{
// 		if len(in.GetName()) < 2 {
// 			return &output, status.Error(codes.InvalidArgument, "Device.Name min 2 character")
// 		}

// 		err = s.es.GetDB().NewSelect().Model(&model.Device{}).Where("name = ?", in.GetName()).Scan(ctx)
// 		if err != nil {
// 			if err != sql.ErrNoRows {
// 				return &output, status.Errorf(codes.Internal, "Query: %v", err)
// 			}
// 		} else {
// 			return &output, status.Error(codes.AlreadyExists, "Device.Name must be unique")
// 		}
// 	}

// 	item := model.Device{
// 		ID:       in.GetId(),
// 		Name:     in.GetName(),
// 		Desc:     in.GetDesc(),
// 		Tags:     in.GetTags(),
// 		Type:     in.GetType(),
// 		Location: in.GetLocation(),
// 		Config:   in.GetConfig(),
// 		Status:   in.GetStatus(),
// 		Created:  time.Now(),
// 		Updated:  time.Now(),
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

func (s *DeviceService) Update(ctx context.Context, in *pb.Device) (*pb.Device, error) {
	var output pb.Device
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Device.Name")
		}
	}

	// name validation
	{
		if len(in.GetName()) < 2 {
			return &output, status.Error(codes.InvalidArgument, "Device.Name min 2 character")
		}
	}

	item, err := s.view(ctx)
	if err != nil {
		return &output, err
	}

	item.Name = in.GetName()
	item.Desc = in.GetDesc()
	item.Tags = in.GetTags()
	item.Type = in.GetType()
	item.Location = in.GetLocation()
	item.Config = in.GetConfig()
	item.Status = in.GetStatus()
	item.Updated = time.Now()

	_, err = s.es.GetDB().NewUpdate().Model(&item).WherePK().Exec(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Update: %v", err)
	}

	// update device id
	if len(in.GetId()) > 0 && in.GetId() != item.ID {
		_, err = s.es.GetDB().NewUpdate().Model(&item).Set("id = ?", in.GetId()).WherePK().Exec(ctx)
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

func (s *DeviceService) View(ctx context.Context, in *pb.MyEmpty) (*pb.Device, error) {
	var output pb.Device
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	item, err := s.view(ctx)
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *DeviceService) Destory(ctx context.Context, in *pb.MyEmpty) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

	err = func() error {
		models := []interface{}{
			(*model.Slot)(nil),
			(*model.Option)(nil),
			(*model.Port)(nil),
			(*model.Proxy)(nil),
			(*model.Source)(nil),
			(*model.Tag)(nil),
			(*model.Const)(nil),
			(*model.Cable)(nil),
			(*model.Wire)(nil),
		}

		tx, err := s.es.GetDB().BeginTx(ctx, nil)
		if err != nil {
			return status.Errorf(codes.Internal, "BeginTx: %v", err)
		}
		var done bool
		defer func() {
			if !done {
				_ = tx.Rollback()
			}
		}()

		for _, model := range models {
			_, err = tx.NewDelete().Model(model).Where("1 = 1").ForceDelete().Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Delete: %v", err)
			}
		}

		done = true
		err = tx.Commit()
		if err != nil {
			return status.Errorf(codes.Internal, "Commit: %v", err)
		}

		return nil
	}()

	if err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}

func (s *DeviceService) view(ctx context.Context) (model.Device, error) {
	item := model.Device{}

	err := s.es.GetDB().NewSelect().Model(&item).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *DeviceService) copyModelToOutput(output *pb.Device, item *model.Device) {
	output.Id = item.ID
	output.Name = item.Name
	output.Desc = item.Desc
	output.Tags = item.Tags
	output.Type = item.Type
	output.Location = item.Location
	output.Config = item.Config
	output.Status = item.Status
	output.Link = s.es.GetStatus().GetDeviceLink()
	output.Created = item.Created.UnixMicro()
	output.Updated = item.Updated.UnixMicro()
	output.Deleted = item.Updated.UnixMicro()
}

func (s *DeviceService) afterUpdate(ctx context.Context, item *model.Device) error {
	var err error

	err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *DeviceService) ViewWithDeleted(ctx context.Context, in *pb.MyEmpty) (*pb.Device, error) {
	var output pb.Device
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	item, err := s.viewWithDeleted(ctx)
	if err != nil {
		return &output, err
	}

	s.copyModelToOutput(&output, &item)

	return &output, nil
}

func (s *DeviceService) viewWithDeleted(ctx context.Context) (model.Device, error) {
	item := model.Device{}

	err := s.es.GetDB().NewSelect().Model(&item).WhereAllWithDeleted().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, status.Errorf(codes.NotFound, "Query: %v, Device.ID: %v", err, item.ID)
		}

		return item, status.Errorf(codes.Internal, "Query: %v", err)
	}

	return item, nil
}

func (s *DeviceService) Sync(ctx context.Context, in *pb.Device) (*pb.MyBool, error) {
	var output pb.MyBool
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetId()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid DeviceID")
		}

		if len(in.GetName()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Device.Name")
		}

		if in.GetUpdated() == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid device updated")
		}
	}

	insert := false
	update := false

	item, err := s.viewWithDeleted(ctx)
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

	//	insert
	if insert {
		// name validation
		{
			if len(in.GetName()) < 2 {
				return &output, status.Error(codes.InvalidArgument, "Device.Name min 2 character")
			}

			err = s.es.GetDB().NewSelect().Model(&model.Device{}).Where("name = ?", in.GetName()).Scan(ctx)
			if err != nil {
				if err != sql.ErrNoRows {
					return &output, status.Errorf(codes.Internal, "Query: %v", err)
				}
			} else {
				return &output, status.Error(codes.AlreadyExists, "Device.Name must be unique")
			}
		}

		item := model.Device{
			ID:       in.GetId(),
			Name:     in.GetName(),
			Desc:     in.GetDesc(),
			Tags:     in.GetTags(),
			Type:     in.GetType(),
			Location: in.GetLocation(),
			Config:   in.GetConfig(),
			Status:   in.GetStatus(),
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
				return &output, status.Error(codes.InvalidArgument, "Device.Name min 2 character")
			}

			modelItem := model.Device{}
			err = s.es.GetDB().NewSelect().Model(&modelItem).Where("name = ?", in.GetName()).Scan(ctx)
			if err != nil {
				if err != sql.ErrNoRows {
					return &output, status.Errorf(codes.Internal, "Query: %v", err)
				}
			} else {
				if modelItem.ID != item.ID {
					return &output, status.Error(codes.AlreadyExists, "Device.Name must be unique")
				}
			}
		}

		item.Name = in.GetName()
		item.Desc = in.GetDesc()
		item.Tags = in.GetTags()
		item.Type = in.GetType()
		item.Location = in.GetLocation()
		item.Config = in.GetConfig()
		item.Status = in.GetStatus()
		item.Updated = time.UnixMicro(in.GetUpdated())
		item.Deleted = time.UnixMicro(in.GetDeleted())

		_, err = s.es.GetDB().NewUpdate().Model(&item).WherePK().WhereAllWithDeleted().Exec(ctx)
		if err != nil {
			return &output, status.Errorf(codes.Internal, "Update: %v", err)
		}

		// update device id
		if len(in.GetId()) > 0 && in.GetId() != item.ID {
			_, err = s.es.GetDB().NewUpdate().Model(&item).Set("id = ?", in.GetId()).WherePK().WhereAllWithDeleted().Exec(ctx)
			if err != nil {
				return &output, status.Errorf(codes.Internal, "Update: %v", err)
			}
		}
	}

	if err = s.afterUpdate(ctx, &item); err != nil {
		return &output, err
	}

	output.Bool = true

	return &output, nil
}
