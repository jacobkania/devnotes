package work_tracking

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jacobkania/devnotes/db/model"
	"github.com/jacobkania/devnotes/db/util"
)

func ManageOverlap(d *sql.DB, unfinished *model.WorkTracking, newWorkTracking *model.WorkTracking) error {
	if unfinished == nil || unfinished.EndTime.Valid {
		return nil
	}

	startTimeAgo := time.Since(unfinished.StartTime.Time).Round(time.Second)

	fmt.Printf("You have an unfinished work tracking session! ")
	fmt.Printf(
		"Started: %s (%s ago)\n",
		unfinished.StartTime.Time.Format(util.DATE_TIME),
		startTimeAgo,
	)

	r, err := util.QuestionC(
		"What would you like to do with the previous session?",
		"Terminate it now",
		"Terminate it 1 hour ago",
		"Terminate it at a specific time",
		"Discard",
	)
	if err != nil {
		return err
	}

	switch r {
	case 0:
		unfinished.EndTime = newWorkTracking.StartTime
		err := unfinished.Save(d)
		if err != nil {
			return err
		}

	case 1:
		oneHourAgo := time.Now().Add(-1 * time.Hour)

		if oneHourAgo.Before(unfinished.StartTime.Time) {
			unfinished.EndTime = unfinished.StartTime
		} else {
			unfinished.EndTime = sql.NullTime{Time: oneHourAgo, Valid: true}
		}

		err := unfinished.Save(d)
		if err != nil {
			return err
		}

	case 2:
		fmt.Println("Unimplemented")
		return errors.New("Fuck")

	case 3:
		err := unfinished.Destroy(d)
		if err != nil {
			return err
		}
	}

	return nil
}
