package common

import (
	"database/sql"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlConn(addr string) (*sql.DB, error) {
	db, err := sql.Open("mysql", addr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
