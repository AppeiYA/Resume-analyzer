package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	*sqlx.DB
}

func ConnectDB(databaseUrl string) (*DB, error) {
	db, err := sqlx.Connect("postgres", databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
    }

	if pingerr := db.Ping(); pingerr != nil {
		return nil, fmt.Errorf("failed to ping db: %w", pingerr)
	}

	db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)

	return &DB{db}, nil
}