package core

import (
	"context"
	"database/sql"
	"time"

	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/core/model"
	"github.com/snple/kokomi/util"
	"github.com/uptrace/bun"
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

	{
		seedDevice := func() error {
			device := model.Device{
				ID:      util.RandomID(),
				Name:    consts.DEFAULT_DEVICE,
				Status:  consts.ON,
				Created: time.Now(),
				Updated: time.Now(),
			}

			_, err = db.NewInsert().Model(&device).Exec(ctx)
			if err != nil {
				return err
			}

			return nil
		}

		err = db.NewSelect().Model(&model.Device{}).Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				if err = seedDevice(); err != nil {
					return err
				}
			}

			return err
		}
	}

	return nil
}
