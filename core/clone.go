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

	"github.com/snple/beacon/core/model"
	"github.com/snple/beacon/util"
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

func (s *cloneService) node(ctx context.Context, db bun.IDB, nodeID string) error {
	var err error

	node := model.Node{
		ID: nodeID,
	}

	err = db.NewSelect().Model(&node).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	node.ID = util.RandomID()
	node.Name = fmt.Sprintf("%v_clone_%v", node.Name, randNameSuffix())

	node.Created = time.Now()
	node.Updated = time.Now()

	_, err = db.NewInsert().Model(&node).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	// slot
	{
		var slots []model.Slot

		err = db.NewSelect().Model(&slots).Where("node_id = ?", nodeID).Order("id ASC").Scan(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		for _, slot := range slots {
			slot.ID = util.RandomID()
			slot.NodeID = node.ID

			slot.Created = time.Now()
			slot.Updated = time.Now()

			_, err = db.NewInsert().Model(&slot).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}
		}
	}

	// source
	{
		var sources []model.Source

		err = db.NewSelect().Model(&sources).Where("node_id = ?", nodeID).Order("id ASC").Scan(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		for _, source := range sources {
			oldSourceId := source.ID

			source.ID = util.RandomID()
			source.NodeID = node.ID

			source.Created = time.Now()
			source.Updated = time.Now()

			_, err = db.NewInsert().Model(&source).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}

			// pin
			{
				var pins []model.Pin

				err = db.NewSelect().Model(&pins).Where("source_id = ?", oldSourceId).Order("id ASC").Scan(ctx)
				if err != nil {
					return status.Errorf(codes.Internal, "Query: %v", err)
				}

				for _, pin := range pins {
					pin.ID = util.RandomID()
					pin.SourceID = source.ID
					pin.NodeID = source.NodeID

					pin.Created = time.Now()
					pin.Updated = time.Now()

					_, err = db.NewInsert().Model(&pin).Exec(ctx)
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

		err = db.NewSelect().Model(&constants).Where("node_id = ?", nodeID).Order("id ASC").Scan(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		for _, constant := range constants {
			constant.ID = util.RandomID()
			constant.NodeID = node.ID

			constant.Created = time.Now()
			constant.Updated = time.Now()

			_, err = db.NewInsert().Model(&constant).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}
		}
	}

	err = s.cs.GetSync().setNodeUpdated(ctx, db, node.ID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) slot(ctx context.Context, db bun.IDB, slotID, nodeID string) error {
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

	// node validation
	if nodeID != "" {
		node := model.Node{
			ID: nodeID,
		}

		err = db.NewSelect().Model(&node).WherePK().Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return status.Error(codes.InvalidArgument, "Please supply valid NodeID")
			}

			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		item.NodeID = node.ID
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	err = s.cs.GetSync().setNodeUpdated(ctx, db, item.NodeID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) source(ctx context.Context, db bun.IDB, sourceID, nodeID string) error {
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

	// node validation
	if nodeID != "" {
		node := model.Node{
			ID: nodeID,
		}

		err = db.NewSelect().Model(&node).WherePK().Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return status.Error(codes.InvalidArgument, "Please supply valid NodeID")
			}

			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		item.NodeID = node.ID
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	// clone pins
	{
		var pins []model.Pin

		err = db.NewSelect().Model(&pins).Where("source_id = ?", sourceID).Order("id ASC").Scan(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		for _, pin := range pins {
			pin.ID = util.RandomID()
			pin.SourceID = item.ID
			pin.NodeID = item.NodeID

			pin.Created = time.Now()
			pin.Updated = time.Now()

			_, err = db.NewInsert().Model(&pin).Exec(ctx)
			if err != nil {
				return status.Errorf(codes.Internal, "Insert: %v", err)
			}
		}
	}

	err = s.cs.GetSync().setNodeUpdated(ctx, db, item.NodeID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) pin(ctx context.Context, db bun.IDB, pinID, sourceID string) error {
	var err error

	item := model.Pin{
		ID: pinID,
	}

	err = db.NewSelect().Model(&item).WherePK().Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.NotFound, "Query: %v", err)
		}

		return status.Errorf(codes.Internal, "Query: %v", err)
	}

	// source validation
	if sourceID != "" {
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
		item.NodeID = source.NodeID
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	err = s.cs.GetSync().setNodeUpdated(ctx, db, item.NodeID, time.Now())
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	return nil
}

func (s *cloneService) const_(ctx context.Context, db bun.IDB, constID, nodeID string) error {
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

	// node validation
	if nodeID != "" {
		node := model.Node{
			ID: nodeID,
		}

		err = db.NewSelect().Model(&node).WherePK().Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return status.Error(codes.InvalidArgument, "Please supply valid NodeID")
			}

			return status.Errorf(codes.Internal, "Query: %v", err)
		}

		item.NodeID = node.ID
	}

	item.ID = util.RandomID()
	item.Name = fmt.Sprintf("%v_clone_%v", item.Name, randNameSuffix())

	item.Created = time.Now()
	item.Updated = time.Now()

	_, err = db.NewInsert().Model(&item).Exec(ctx)
	if err != nil {
		return status.Errorf(codes.Internal, "Insert: %v", err)
	}

	err = s.cs.GetSync().setNodeUpdated(ctx, db, item.NodeID, time.Now())
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
