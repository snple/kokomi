//go:build linux || darwin
// +build linux darwin

package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func dbOpen(file string) (*sql.DB, error) {
	return sql.Open("sqlite3",
		fmt.Sprintf("file:%v?cache=shared&_journal_mode=WAL&_sync=1&_cache_size=16000", file))
}
