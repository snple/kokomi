package edge

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/snple/kokomi/edge/model"
	"github.com/snple/kokomi/util"
	"github.com/uptrace/bun"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type cloneService struct {
	es *EdgeService
}

func newCloneService(es *EdgeService) *cloneService {
	return &cloneService{
		es: es,
	}
}

func (s *cloneService) slot(ctx context.Context, db bun.IDB, slotID string) error {
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

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	{
		err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}

		err = s.es.GetSync().setSlotUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func (s *cloneService) option(ctx context.Context, db bun.IDB, optionID string) error {
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

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	{
		err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}

		err = s.es.GetSync().setOptionUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func (s *cloneService) port(ctx context.Context, db bun.IDB, portID string) error {
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

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	{
		err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}

		err = s.es.GetSync().setPortUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func (s *cloneService) proxy(ctx context.Context, db bun.IDB, proxyID string) error {
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

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	{
		err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}

		err = s.es.GetSync().setProxyUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func (s *cloneService) source(ctx context.Context, db bun.IDB, sourceID string) error {
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

			tag.Created = time.Now()
			tag.Updated = time.Now()

			_, err = db.NewInsert().Model(&tag).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}
		}
	}

	{
		err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}

		err = s.es.GetSync().setSourceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
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
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	{
		err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}

		err = s.es.GetSync().setTagUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func (s *cloneService) const_(ctx context.Context, db bun.IDB, constID string) error {
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

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	{
		err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}

		err = s.es.GetSync().setConstUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func (s *cloneService) cable(ctx context.Context, db bun.IDB, cableID string) error {
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

			wire.Created = time.Now()
			wire.Updated = time.Now()

			_, err = db.NewInsert().Model(&wire).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}
		}
	}

	{
		err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}

		err = s.es.GetSync().setCableUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func (s *cloneService) wire(ctx context.Context, db bun.IDB, wireID, cableID string) error {
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
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	{
		err = s.es.GetSync().setDeviceUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}

		err = s.es.GetSync().setWireUpdated(ctx, time.Now())
		if err != nil {
			return status.Errorf(codes.Internal, "Insert: %v", err)
		}
	}

	return nil
}

func randNameSuffix() string {
	buf := new(bytes.Buffer)

	random := rand.Uint32()
	binary.Write(buf, binary.BigEndian, random)

	return hex.EncodeToString(buf.Bytes())
}
