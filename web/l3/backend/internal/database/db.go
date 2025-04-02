package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("DB connection error: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Cannot connect to DB: %w", err)
	}
	return db, nil
}
