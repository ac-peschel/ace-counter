package db

import (
	"database/sql"
	"time"
)

func CreateIncrement(db *sql.DB, counterId int) (int, error) {
	_, err := db.Exec(`
		INSERT INTO increments (counter_id, created_at)
		VALUES (?, ?)
	`, counterId, time.Now().Unix())
	if err != nil {
		return 0, err
	}

	return GetIncrements(db, counterId)
}

func GetIncrements(db *sql.DB, counterId int) (int, error) {
	var sum int
	err := db.QueryRow(`
		SELECT COUNT(*)
		FROM increments
		WHERE counter_id = ?
	`, counterId).Scan(&sum)
	if err != nil {
		return 0, err
	}

	return sum, nil
}
