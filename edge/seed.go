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

func Seed(db *bun.DB, nodeName string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if err = seed(tx, nodeName); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func seed(db bun.Tx, nodeName string) error {
	var err error
	ctx := context.Background()

	{
		seedNode := func() error {
			node := model.Node{
				ID:      util.RandomID(),
				Name:    nodeName,
				Created: time.Now(),
				Updated: time.Now(),
			}

			_, err = db.NewInsert().Model(&node).Exec(ctx)
			if err != nil {
				return err
			}

			fmt.Printf("seed: the initial node created: %v\n", node.ID)

			return nil
		}

		// check if node already exists
		// if not, create it
		err = db.NewSelect().Model(&model.Node{}).Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				if err = seedNode(); err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	return nil
}
