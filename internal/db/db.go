package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func Open(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		PRAGMA journal_mode = WAL;
		PRAGMA synchronous = NORMAL;
		PRAGMA foreign_keys = ON;
	`)
	if err != nil {
		return nil, fmt.Errorf("pragma failed: %w", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS counters (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			owner TEXT NOT NULL
		);
	`)
	if err != nil {
		return nil, fmt.Errorf("schema init failed: %w", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS increments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			counter_id INTEGER NOT NULL,
			created_at INTEGER NOT NULL
		);
	`)
	if err != nil {
		return nil, fmt.Errorf("schema init failed: %w", err)
	}

	return db, nil
}
