package datastore

import (
	"os"
	"path"
	"todo/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigration() {

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	sourceDir := "file://" + path.Join(dir, "/datastore/migrations/")

	migration, err := migrate.New(sourceDir, *config.MysqlConnectionString())
	if err != nil {
		panic(err)
	}
	defer migration.Close()

	err = migration.Up()
	if err != nil {
		if err != migrate.ErrNoChange {
			panic(err)
		}
	}

}
