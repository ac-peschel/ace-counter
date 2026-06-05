package db

import (
	"database/sql"

	"github.com/ac-peschel/ace-counter/internal/model"
)

func DeleteCounter(db *sql.DB, id int) error {
	_, err := db.Exec(`
		DELETE FROM counters
		WHERE id = ?
	`, id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCounter(db *sql.DB, id int, name, owner string) error {
	_, err := db.Exec(`
		UPDATE counters
		SET name = ?, owner = ?
		WHERE id = ?
	`, name, owner, id)
	if err != nil {
		return err
	}

	return nil
}

func CreateCounter(db *sql.DB) (*model.Counter, error) {
	var c model.Counter
	err := db.QueryRow(`
		INSERT INTO counters (name, owner)
		VALUES (?, ?)
		RETURNING id, name, owner
	`, "Neuer Counter", "Celina").Scan(
		&c.Id,
		&c.Name,
		&c.Owner,
	)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func GetCounterByID(db *sql.DB, id int) (*model.Counter, error) {
	row := db.QueryRow(`
		SELECT id, name, owner
		FROM counters
		WHERE id = ?
	`, id)

	var c model.Counter

	err := row.Scan(
		&c.Id,
		&c.Name,
		&c.Owner,
	)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func GetAllCounter(db *sql.DB) ([]model.Counter, error) {
	rows, err := db.Query(`
		SELECT id, name, owner
		FROM counters
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var counter []model.Counter
	for rows.Next() {
		var c model.Counter
		err := rows.Scan(
			&c.Id,
			&c.Name,
			&c.Owner,
		)
		if err != nil {
			return nil, err
		}
		counter = append(counter, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return counter, nil
}
