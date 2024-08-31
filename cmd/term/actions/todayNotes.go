package actions

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jacobkania/devnotes/db/query"
)

// TodayNotes prints all notes from today
func TodayNotes(d *sql.DB) error {
	todayNotes, err := query.FindNoteByCreatedDate(d, time.Now())
	if err != nil {
		return err
	}

	// get number of digits in largest number of todayNotes to properly pad the output
	digits := len(fmt.Sprintf("%d", len(todayNotes)-1))

	for i, note := range todayNotes {
		fmt.Printf("%*d: %s\n", digits, i, note.Contents)
	}

	return nil
}
