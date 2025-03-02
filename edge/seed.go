package edge

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/snple/beacon/edge/model"
	"github.com/snple/beacon/util"
	"github.com/uptrace/bun"
)

func Seed(db *bun.DB, deviceName string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if err = seed(tx, deviceName); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func seed(db bun.Tx, deviceName string) error {
	var err error
	ctx := context.Background()

	{
		seedDevice := func() error {
			device := model.Device{
				ID:      util.RandomID(),
				Name:    deviceName,
				Created: time.Now(),
				Updated: time.Now(),
			}

			_, err = db.NewInsert().Model(&device).Exec(ctx)
			if err != nil {
				return err
			}

			fmt.Printf("seed: the initial device created: %v\n", device.ID)

			return nil
		}

		// check if device already exists
		// if not, create it
		err = db.NewSelect().Model(&model.Device{}).Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				if err = seedDevice(); err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	return nil
}
