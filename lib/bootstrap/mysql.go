package bootstrap

import (
	"database/sql"
	"fmt"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/vhall1/shorturl/lib/config"
)

type mysqlConf struct {
	DatabaseUrl string
}

func NewMysqlConn(addr string) (*sql.DB, error) {
	conf := mysqlConf{}
	if err := config.Load(&conf); err != nil {
		return nil, fmt.Errorf("failed to load mysql config: %v", err)
	}

	db, err := sql.Open("mysql", addr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
