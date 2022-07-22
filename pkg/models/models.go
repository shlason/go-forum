package models

import (
	"database/sql"

	"github.com/shlason/go-forum/pkg/configs"
)

var db *sql.DB

func init() {
	d, err := sql.Open(configs.DbCfg.Dialect, configs.DbCfg.Dsn)
	if err != nil {
		panic(err)
	}
	db = d
}
