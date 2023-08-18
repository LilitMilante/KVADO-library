package app

import (
	"database/sql"
	"errors"

	"KVADO-library/migrations"

	"github.com/pressly/goose/v3"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToMySQL(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = upMigrations(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func upMigrations(db *sql.DB) error {
	goose.SetBaseFS(migrations.FS)
	goose.SetLogger(goose.NopLogger())

	err := goose.SetDialect("postgres")
	if err != nil {
		return err
	}

	err = goose.Up(db, ".")
	if err != nil && !errors.Is(err, goose.ErrNoNextVersion) {
		return err
	}

	return nil
}
