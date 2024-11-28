package query

import (
	"database/sql"
	"time"

	"github.com/jacobkania/devnotes/db/model"
)

const (
	qFindNoteByCreatedDate = `SELECT id, contents, created_at, updated_at FROM notes WHERE DATE(created_at, 'localtime') = ?`
)

// FindByCreatedDate retrieves all notes created on the specified date
func FindNoteByCreatedDate(db *sql.DB, date time.Time) ([]model.Note, error) {
	rows, err := db.Query(qFindNoteByCreatedDate, date.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []model.Note
	for rows.Next() {
		var note model.Note
		if err := rows.Scan(&note.ID, &note.Contents, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}
