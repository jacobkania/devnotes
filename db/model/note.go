package model

import (
	"database/sql"
	"time"
)

type Note struct {
	ID        int
	Contents  string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

func (n *Note) Save(db *sql.DB) error {
	if n.ID == 0 {
		return n.insert(db)
	}
	return n.update(db)
}

func (n *Note) insert(db *sql.DB) error {
	stmt := `INSERT INTO notes (contents, created_at) VALUES (?, ?)`
	result, err := db.Exec(stmt, n.Contents, time.Now())
	if err != nil {
		return err
	}

	lastInsId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	n.ID = int(lastInsId)
	return nil
}

func (n *Note) update(db *sql.DB) error {
	stmt := `UPDATE notes SET (contents, updated_at) VALUES (?, ?) WHERE id = ?`
	_, err := db.Exec(stmt, n.Contents, time.Now())
	return err
}
