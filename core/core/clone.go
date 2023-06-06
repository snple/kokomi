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

func (s *cloneService) cloneDevice(ctx context.Context, db bun.IDB, deviceID string) error {
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

	// var
	{
		var variables []model.Var

		err = db.NewSelect().Model(&variables).Where("device_id = ?", deviceID).Order("id ASC").Scan(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		for _, variable := range variables {
			variable.ID = util.RandomID()
			variable.DeviceID = device.ID

			variable.Created = time.Now()
			variable.Updated = time.Now()

			_, err = db.NewInsert().Model(&variable).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}
		}
	}

	// cable
	{
		var cables []model.Cable

		err = db.NewSelect().Model(&cables).Where("device_id = ?", deviceID).Order("id ASC").Scan(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		for _, cable := range cables {
			oldCableId := cable.ID

			cable.ID = util.RandomID()
			cable.DeviceID = device.ID

			cable.Created = time.Now()
			cable.Updated = time.Now()

			_, err = db.NewInsert().Model(&cable).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}

			// wire
			{
				var wires []model.Wire

				err = db.NewSelect().Model(&wires).Where("cable_id = ?", oldCableId).Order("id ASC").Scan(ctx)
				if err != nil {
					return status.Errorf(codes.Internal, "Query: %v", err)
				}

				for _, wire := range wires {
					wire.ID = util.RandomID()
					wire.CableID = cable.ID
					wire.DeviceID = cable.DeviceID

					wire.Created = time.Now()
					wire.Updated = time.Now()

					_, err = db.NewInsert().Model(&wire).Exec(ctx)
					if err != nil {
						return status.Errorf(codes.Internal, "Insert: %v", err)
					}
				}
			}
		}
	}

	// class
	{
		var classes []model.Class

		err = db.NewSelect().Model(&classes).Where("device_id = ?", deviceID).Order("id ASC").Scan(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		for _, class := range classes {
			oldClassId := class.ID

			class.ID = util.RandomID()
			class.DeviceID = device.ID

			class.Created = time.Now()
			class.Updated = time.Now()

			_, err = db.NewInsert().Model(&class).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}

			// attr
			{
				var attrs []model.Attr

				err = db.NewSelect().Model(&attrs).Where("class_id = ?", oldClassId).Order("id ASC").Scan(ctx)
				if err != nil {
					return status.Errorf(codes.Internal, "Query: %v", err)
				}

				for _, attr := range attrs {
					attr.ID = util.RandomID()
					attr.ClassID = class.ID
					attr.DeviceID = class.DeviceID

					if attr.TagID != "" {
						if tagId, ok := tagIDMap[attr.TagID]; ok {
							attr.TagID = tagId
						}
					}

					attr.Created = time.Now()
					attr.Updated = time.Now()

					_, err = db.NewInsert().Model(&attr).Exec(ctx)
					if err != nil {
						return status.Errorf(codes.Internal, "Insert: %v", err)
					}
				}
			}
		}
	}

	err = s.cs.GetSync().setDeviceUpdated(ctx, device.ID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) cloneSlot(ctx context.Context, db bun.IDB, slotID, deviceID string) error {
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

	err = s.cs.GetSync().setDeviceUpdated(ctx, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) cloneOption(ctx context.Context, db bun.IDB, optionID, deviceID string) error {
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

	err = s.cs.GetSync().setDeviceUpdated(ctx, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) clonePort(ctx context.Context, db bun.IDB, portID, deviceID string) error {
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

	err = s.cs.GetSync().setDeviceUpdated(ctx, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) cloneProxy(ctx context.Context, db bun.IDB, proxyID, deviceID string) error {
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

	err = s.cs.GetSync().setDeviceUpdated(ctx, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) cloneSource(ctx context.Context, db bun.IDB, sourceID, deviceID string) error {
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

	err = s.cs.GetSync().setDeviceUpdated(ctx, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) cloneTag(ctx context.Context, db bun.IDB, tagID, sourceID string) error {
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

	err = s.cs.GetSync().setDeviceUpdated(ctx, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) cloneVar(ctx context.Context, db bun.IDB, varID, deviceID string) error {
	var err error

	item := model.Var{
		ID: varID,
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

	err = s.cs.GetSync().setDeviceUpdated(ctx, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) cloneCable(ctx context.Context, db bun.IDB, cableID, deviceID string) error {
	var err error

	item := model.Cable{
		ID: cableID,
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

	// clone wires
	{
		var wires []model.Wire

		err = db.NewSelect().Model(&wires).Where("cable_id = ?", cableID).Order("id ASC").Scan(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		for _, wire := range wires {
			wire.ID = util.RandomID()
			wire.CableID = item.ID
			wire.DeviceID = item.DeviceID

			wire.Created = time.Now()
			wire.Updated = time.Now()

			_, err = db.NewInsert().Model(&wire).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}
		}
	}

	err = s.cs.GetSync().setDeviceUpdated(ctx, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) cloneWire(ctx context.Context, db bun.IDB, wireID, cableID string) error {
	var err error

	item := model.Wire{
		ID: wireID,
	}

	err = db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	// cable validation
	if len(cableID) > 0 {
		cable := model.Cable{
			ID: cableID,
		}

		err = db.NewSelect().Model(&cable).WherePK().Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return status.Error(codes.InvalidArgument, "Please supply valid Cable.ID")
			}

			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		item.CableID = cable.ID
		item.DeviceID = cable.DeviceID
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	err = s.cs.GetSync().setDeviceUpdated(ctx, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) cloneClass(ctx context.Context, db bun.IDB, classID, deviceID string) error {
	var err error

	item := model.Class{
		ID: classID,
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

	// clone attrs
	{
		var attrs []model.Attr

		err = db.NewSelect().Model(&attrs).Where("class_id = ?", classID).Order("id ASC").Scan(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		for _, attr := range attrs {
			attr.ID = util.RandomID()
			attr.ClassID = item.ID
			attr.DeviceID = item.DeviceID

			attr.Created = time.Now()
			attr.Updated = time.Now()

			_, err = db.NewInsert().Model(&attr).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}
		}
	}

	err = s.cs.GetSync().setDeviceUpdated(ctx, item.DeviceID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) cloneAttr(ctx context.Context, db bun.IDB, attrID, classID string) error {
	var err error

	item := model.Attr{
		ID: attrID,
	}

	err = db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	// class validation
	if len(classID) > 0 {
		class := model.Class{
			ID: classID,
		}

		err = db.NewSelect().Model(&class).WherePK().Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return status.Error(codes.InvalidArgument, "Please supply valid ClassID")
			}

			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		item.ClassID = class.ID
		item.DeviceID = class.DeviceID
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	err = s.cs.GetSync().setDeviceUpdated(ctx, item.DeviceID, time.Now())
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
