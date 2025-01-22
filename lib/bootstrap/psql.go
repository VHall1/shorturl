package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/vhall1/shorturl/lib/config"
)

type psqlConf struct {
	DatabaseUrl string
}

func NewPostgres() (*sql.DB, error) {
	conf := psqlConf{}
	if err := config.Load(&conf); err != nil {
		return nil, fmt.Errorf("failed to load psql config: %v", err)
	}

	db, err := sql.Open("postgres", conf.DatabaseUrl)
	if err != nil {
		return nil, fmt.Errorf("psql failed to connect: %v", err)
	}

	return db, nil
}
