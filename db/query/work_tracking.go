package query

import (
	"database/sql"

	"github.com/jacobkania/devnotes/db/model"
)

const (
	qFindUnfinishedWorkTracking = `
		SELECT
			id,
			start_time,
			end_time,
			created_at,
			updated_at
		FROM work_tracking
		WHERE start_time IS NOT NULL
		AND end_time IS NULL
		ORDER BY start_time
		LIMIT 1
	`
)

// FindByCreatedDate retrieves all notes created on the specified date
func FindUnfinishedWorkTracking(db *sql.DB) (*model.WorkTracking, error) {
	rows, err := db.Query(qFindUnfinishedWorkTracking)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	work_tracking := &model.WorkTracking{}

	if err := rows.Scan(&work_tracking.ID, &work_tracking.StartTime, &work_tracking.EndTime, &work_tracking.CreatedAt, &work_tracking.UpdatedAt); err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return work_tracking, nil
}
