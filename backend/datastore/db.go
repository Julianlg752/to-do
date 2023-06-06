package datastore

import (
	"database/sql"
	"todo/config"

	"github.com/ansel1/merry/v2"
	"github.com/doug-martin/goqu/v9"
)

var (
	DB     *sql.DB
	GoquDB goqu.Database
)

func Connection() error {
	var err error
	DB, err = sql.Open("mysql", config.Dns())
	if err != nil {
		return merry.Wrap(err)
	}

	if err := DB.Ping(); err != nil {
		return merry.Wrap(err)

	}
	GoquDB = *goqu.Dialect("mysql").DB(DB)
	return nil
}
