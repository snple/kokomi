package core

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/snple/kokomi/core/model"
	"github.com/snple/kokomi/util"
	"github.com/uptrace/bun"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type cloneService struct {
	cs *CoreService
}

func newCloneService(cs *CoreService) *cloneService {
	return &cloneService{
		cs: cs,
	}
}

func (s *cloneService) device(ctx context.Context, db bun.IDB, deviceID string) error {
	var err error

	device := model.Device{
		ID: deviceID,
	}

	err = db.NewSelect().Model(&device).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	device.ID = util.RandomID()
	device.Name = fmt.Sprintf("%v_clone_%v", device.Name, randNameSuffix())

	device.Created = time.Now()
	device.Updated = time.Now()

	_, err = db.NewInsert().Model(&device).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	// slot
	{
		var slots []model.Slot

		err = db.NewSelect().Model(&slots).Where("device_id = ?", deviceID).Order("id ASC").Scan(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		for _, slot := range slots {
			slot.ID = util.RandomID()
			slot.DeviceID = device.ID

			slot.Created = time.Now()
			slot.Updated = time.Now()

			_, err = db.NewInsert().Model(&slot).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}
		}
	}

	// option
	{
		var options []model.Option

		err = db.NewSelect().Model(&options).Where("device_id = ?", deviceID).Order("id ASC").Scan(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		for _, option := range options {
			option.ID = util.RandomID()
			option.DeviceID = device.ID

			option.Created = time.Now()
			option.Updated = time.Now()

			_, err = db.NewInsert().Model(&option).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}
		}
	}

	// port
	{
		var slots []model.Port

		err = db.NewSelect().Model(&slots).Where("device_id = ?", deviceID).Order("id ASC").Scan(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		for _, slot := range slots {
			slot.ID = util.RandomID()
			slot.DeviceID = device.ID

			slot.Created = time.Now()
			slot.Updated = time.Now()

			_, err = db.NewInsert().Model(&slot).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}
		}
	}

	// proxy
	{
		var slots []model.Proxy

		err = db.NewSelect().Model(&slots).Where("device_id = ?", deviceID).Order("id ASC").Scan(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		for _, slot := range slots {
			slot.ID = util.RandomID()
			slot.DeviceID = device.ID

			slot.Created = time.Now()
			slot.Updated = time.Now()

			_, err = db.NewInsert().Model(&slot).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}
		}
	}

	tagIDMap := make(map[string]string, 0)

	// source
	{
		var sources []model.Source

		err = db.NewSelect().Model(&sources).Where("device_id = ?", deviceID).Order("id ASC").Scan(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		for _, source := range sources {
			oldSourceId := source.ID

			source.ID = util.RandomID()
			source.DeviceID = device.ID

			source.Created = time.Now()
			source.Updated = time.Now()

			_, err = db.NewInsert().Model(&source).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}

			// tag
			{
				var tags []model.Tag

				err = db.NewSelect().Model(&tags).Where("source_id = ?", oldSourceId).Order("id ASC").Scan(ctx)
				if err != nil {
					return status.Errorf(codes.Internal, "Query: %v", err)
				}

				for _, tag := range tags {
					newId := util.RandomID()
					tagIDMap[tag.ID] = newId

					tag.ID = newId
					tag.SourceID = source.ID
					tag.DeviceID = source.DeviceID

					tag.Created = time.Now()
					tag.Updated = time.Now()

					_, err = db.NewInsert().Model(&tag).Exec(ctx)
					if err != nil {
						return status.Errorf(codes.Internal, "Insert: %v", err)
					}
				}
			}
		}
	}

	// const
	{
		var constants []model.Const

		err = db.NewSelect().Model(&constants).Where("device_id = ?", deviceID).Order("id ASC").Scan(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		for _, constant := range constants {
			constant.ID = util.RandomID()
			constant.DeviceID = device.ID

			constant.Created = time.Now()
			constant.Updated = time.Now()

			_, err = db.NewInsert().Model(&constant).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}
		}
	}

	err = s.cs.GetSync().setDeviceUpdated(ctx, db, device.ID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) slot(ctx context.Context, db bun.IDB, slotID, deviceID string) error {
	var err error

	item := model.Slot{
		ID: slotID,
	}

	err = db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	// device validation
	if len(deviceID) > 0 {
		device := model.Device{
			ID: deviceID,
		}

		err = db.NewSelect().Model(&device).WherePK().Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return status.Error(codes.InvalidArgument, "Please supply valid DeviceID")
			}

			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		item.DeviceID = device.ID
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	err = s.cs.GetSync().setDeviceUpdated(ctx, db, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) option(ctx context.Context, db bun.IDB, optionID, deviceID string) error {
	var err error

	item := model.Option{
		ID: optionID,
	}

	err = db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	// device validation
	if len(deviceID) > 0 {
		device := model.Device{
			ID: deviceID,
		}

		err = db.NewSelect().Model(&device).WherePK().Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return status.Error(codes.InvalidArgument, "Please supply valid DeviceID")
			}

			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		item.DeviceID = device.ID
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	err = s.cs.GetSync().setDeviceUpdated(ctx, db, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) port(ctx context.Context, db bun.IDB, portID, deviceID string) error {
	var err error

	item := model.Port{
		ID: portID,
	}

	err = db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	// device validation
	if len(deviceID) > 0 {
		device := model.Device{
			ID: deviceID,
		}

		err = db.NewSelect().Model(&device).WherePK().Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return status.Error(codes.InvalidArgument, "Please supply valid DeviceID")
			}

			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		item.DeviceID = device.ID
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	err = s.cs.GetSync().setDeviceUpdated(ctx, db, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) proxy(ctx context.Context, db bun.IDB, proxyID, deviceID string) error {
	var err error

	item := model.Proxy{
		ID: proxyID,
	}

	err = db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	// device validation
	if len(deviceID) > 0 {
		device := model.Device{
			ID: deviceID,
		}

		err = db.NewSelect().Model(&device).WherePK().Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return status.Error(codes.InvalidArgument, "Please supply valid DeviceID")
			}

			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		item.DeviceID = device.ID
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	err = s.cs.GetSync().setDeviceUpdated(ctx, db, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) source(ctx context.Context, db bun.IDB, sourceID, deviceID string) error {
	var err error

	item := model.Source{
		ID: sourceID,
	}

	err = db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	// device validation
	if len(deviceID) > 0 {
		device := model.Device{
			ID: deviceID,
		}

		err = db.NewSelect().Model(&device).WherePK().Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return status.Error(codes.InvalidArgument, "Please supply valid DeviceID")
			}

			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		item.DeviceID = device.ID
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	// clone tags
	{
		var tags []model.Tag

		err = db.NewSelect().Model(&tags).Where("source_id = ?", sourceID).Order("id ASC").Scan(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		for _, tag := range tags {
			tag.ID = util.RandomID()
			tag.SourceID = item.ID
			tag.DeviceID = item.DeviceID

			tag.Created = time.Now()
			tag.Updated = time.Now()

			_, err = db.NewInsert().Model(&tag).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}
		}
	}

	err = s.cs.GetSync().setDeviceUpdated(ctx, db, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) tag(ctx context.Context, db bun.IDB, tagID, sourceID string) error {
	var err error

	item := model.Tag{
		ID: tagID,
	}

	err = db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	// source validation
	if len(sourceID) > 0 {
		source := model.Source{
			ID: sourceID,
		}

		err = db.NewSelect().Model(&source).WherePK().Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return status.Error(codes.InvalidArgument, "Please supply valid Source.ID")
			}

			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		item.SourceID = source.ID
		item.DeviceID = source.DeviceID
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	err = s.cs.GetSync().setDeviceUpdated(ctx, db, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) const_(ctx context.Context, db bun.IDB, constID, deviceID string) error {
	var err error

	item := model.Const{
		ID: constID,
	}

	err = db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	// device validation
	if len(deviceID) > 0 {
		device := model.Device{
			ID: deviceID,
		}

		err = db.NewSelect().Model(&device).WherePK().Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return status.Error(codes.InvalidArgument, "Please supply valid DeviceID")
			}

			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		item.DeviceID = device.ID
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	err = s.cs.GetSync().setDeviceUpdated(ctx, db, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func randNameSuffix() string {
	buf := new(bytes.Buffer)

	random := rand.Uint32()
	binary.Write(buf, binary.BigEndian, random)

	return hex.EncodeToString(buf.Bytes())
}
