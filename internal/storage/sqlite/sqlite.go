package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

var (
	migrationsPath = "file://migrations"
)

func New(storagePath string) *sql.DB {

	if _, err := os.Stat(filepath.Dir(storagePath)); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(storagePath), 0755); err != nil {
			panic("failed to create storage directory: " + err.Error())
		}
	}

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		panic("failed to open db: " + err.Error())
	}

	err = migrateDB(db)
	if err != nil {
		panic("migrations failed " + err.Error())
	}
	return db
}

func migrateDB(db *sql.DB) error {
	op := "sqlite.migrate"
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		return fmt.Errorf("%s: sqlite.WithInstance: %w", op, err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"sqlite3", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
