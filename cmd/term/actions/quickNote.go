package actions

import (
	"database/sql"
	"fmt"

	"github.com/jacobkania/devnotes/db/model"
)

func QuickNote(d *sql.DB, text string) error {
	newNote := model.Note{
		Contents: text,
	}

	err := newNote.Save(d)
	if err != nil {
		return err
	}

	fmt.Printf("DevNotes: Added!\n")
	return nil
}
