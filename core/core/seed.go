package core

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
	"snple.com/kokomi/consts"
	"snple.com/kokomi/core/model"
	"snple.com/kokomi/util"
)

func Seed(db *bun.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if err = seed(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func seed(db bun.Tx) error {
	var err error
	ctx := context.Background()

	hasDevice := false
	{
		device := model.Device{}

		err = db.NewSelect().Model(&device).Scan(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return err
			}
		} else {
			hasDevice = true
		}
	}

	if !hasDevice {
		deviceID := util.RandomID()

		device := model.Device{
			ID:      deviceID,
			Name:    consts.DEFAULT_DEVICE,
			Status:  consts.ON,
			Created: time.Now(),
			Updated: time.Now(),
		}

		_, err = db.NewInsert().Model(&device).Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
