package db

import (
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func ConnectSqlite(file string, debug bool) (*bun.DB, error) {
	sqldb, err := dbOpen(file)
	if err != nil {
		return nil, err
	}

	// sqldb.SetMaxIdleConns(4)
	// sqldb.SetMaxOpenConns(20)
	sqldb.SetMaxOpenConns(1)

	db := bun.NewDB(sqldb, sqlitedialect.New())

	if debug {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	return db, nil
}
