package app

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"KVADO-library/migrations"

	"github.com/pressly/goose/v3"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToMySQL(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	// try to ping 3 seconds
	for i := 0; i < 6; i++ {
		err = db.Ping()
		if err == nil {
			break
		}

		time.Sleep(500 * time.Millisecond)
	}

	if err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	//applying migrations
	err = upMigrations(db)
	if err != nil {
		return nil, fmt.Errorf("up migrations: %w", err)
	}

	return db, nil
}

func upMigrations(db *sql.DB) error {
	goose.SetBaseFS(migrations.FS)
	goose.SetLogger(goose.NopLogger())

	err := goose.SetDialect("mysql")
	if err != nil {
		return fmt.Errorf("set dialect: %w", err)
	}

	err = goose.Up(db, ".")
	if err != nil && !errors.Is(err, goose.ErrNoNextVersion) {
		return fmt.Errorf("applying migrations: %w", err)
	}

	return nil
}
