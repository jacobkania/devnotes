package model

import (
	"database/sql"
	"fmt"

	"github.com/jacobkania/devnotes/db/dbutil"
)

type WorkTracking struct {
	ID        int
	StartTime sql.NullTime
	EndTime   sql.NullTime
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

// Save inserts or updates the Note in the database
func (w *WorkTracking) Save(db *sql.DB) error {
	if w.ID == 0 {
		return w.insert(db)
	}
	return w.update(db)
}

func (w *WorkTracking) Destroy(db *sql.DB) error {
	stmt := `DELETE FROM work_tracking WHERE id = ?`
	result, err := db.Exec(stmt, w.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fmt.Errorf("issue deleting record, rows affected: %v", rowsAffected)
	}

	return nil
}

func (w *WorkTracking) insert(db *sql.DB) error {
	stmt := `INSERT INTO work_tracking (start_time, end_time, created_at) VALUES (?, ?, ?)`
	result, err := db.Exec(stmt, w.StartTime, w.EndTime, dbutil.CurrentTimeUTCZ())
	if err != nil {
		return err
	}

	lastInsId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	w.ID = int(lastInsId)
	return nil
}

func (w *WorkTracking) update(db *sql.DB) error {
	stmt := `UPDATE work_tracking SET start_time = ?, end_time = ?, updated_at = ? WHERE id = ?`
	_, err := db.Exec(stmt, w.StartTime, w.EndTime, dbutil.CurrentTimeUTCZ(), w.ID)
	return err
}

func (w *WorkTracking) Scan(row *sql.Row) error {
	return row.Scan(&w.ID, &w.StartTime, &w.EndTime, &w.CreatedAt, &w.UpdatedAt)
}
