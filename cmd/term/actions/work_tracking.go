package actions

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jacobkania/devnotes/cmd/term/actions/work_tracking"
	"github.com/jacobkania/devnotes/db/model"
	"github.com/jacobkania/devnotes/db/query"
	"github.com/jacobkania/devnotes/db/util"
)

// TodayNotes prints all notes from today
func WorkTrackingStart(d *sql.DB) error {
	startTime := sql.NullTime{Time: time.Now(), Valid: true}

	newWorkTracking := &model.WorkTracking{
		StartTime: startTime,
	}

	unfinished, err := query.FindUnfinishedWorkTracking(d)
	if err != nil {
		return err
	}

	err = work_tracking.ManageOverlap(d, unfinished, newWorkTracking)
	if err != nil {
		return err
	}

	err = newWorkTracking.Save(d)
	if err != nil {
		return err
	}

	fmt.Printf("Started tracking time: %s\n", startTime.Time.Format(util.DATE_TIME))
	return nil
}

func WorkTrackingEnd(d *sql.DB) error {
	endTime := sql.NullTime{Time: time.Now(), Valid: true}

	unfinished, err := query.FindUnfinishedWorkTracking(d)
	if err != nil {
		return err
	}

	if unfinished == nil {
		return errors.New("no open work tracking session found")
	}

	unfinished.EndTime = endTime
	err = unfinished.Save(d)
	if err != nil {
		return err
	}

	fmt.Printf("Ended tracking time: %s\n", endTime.Time.Format(util.DATE_TIME))
	return nil
}
