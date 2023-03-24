// go:build windows
//go:build windows
// +build windows

package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func dbOpen(file string) (*sql.DB, error) {
	return sql.Open("sqlite",
		fmt.Sprintf("file:%v?_pragma=busy_timeout(5000)&_pragma=synchronous(NORMAL)&_pragma=journal_mode(WAL)", file))
}
