package db

import (
	"database/sql"
	"fmt"
	"io/fs"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/jacobkania/devnotes/config"
	"github.com/uptrace/bun/driver/sqliteshim"

	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func Init(cfg *config.Config, migrateFS fs.FS) (db *sql.DB) {

	sqldb, err := sql.Open(sqliteshim.ShimName, cfg.DatabasePath)
	if err != nil {
		panic(err)
	}

	if err := runMigrations(sqldb, migrateFS); err != nil {
		panic(err)
	}

	return sqldb
}

func runMigrations(sqldb *sql.DB, migrateFS fs.FS) error {
	driver, err := sqlite3.WithInstance(sqldb, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("error initializing driver: %w", err)
	}

	// note: name "migrate" here must match with cmd/chat/main.go as embed source
	data, err := iofs.New(migrateFS, ".")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", data, "sqlite3", driver)
	if err != nil {
		return fmt.Errorf("error initializing migration: %w", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error running migration: %w", err)
	}
	return nil
}
