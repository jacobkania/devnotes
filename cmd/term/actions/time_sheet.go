package actions

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/jacobkania/devnotes/db/query"
	"github.com/jacobkania/devnotes/db/util"
)

const (
	// possible timeRange format regexes
	// 2w - 2 weeks; 3d - 3 days; 4m - 4 months; 5y - 5 years
	timeRangeSimple = `^(\d+)([wdmy])$`
	// 2021-01-01 - 2021-12-31
	timeRangeFull = `^(\d{4}-\d{2}-\d{2}) - (\d{4}-\d{2}-\d{2})$`
	// 03/16 - 03/20
	timeRangeShort = `^(\d{2}/\d{2}) - (\d{2}/\d{2})$`
)

func TimeSheet(d *sql.DB, timeRange string) error {
	startTime, endTime, err := parseTime(timeRange)
	if err != nil {
		return err
	}

	results, err := query.FindWorkTrackingInTimeRange(d, startTime, endTime)
	if err != nil {
		return err
	}

	if len(results) == 0 {
		return errors.New("no work tracking found in the given time range")
	}

	for _, result := range results {
		timeLength := result.EndTime.Time.Sub(result.StartTime.Time).Round(time.Second)
		fmt.Printf("%+30s   -   %-30s |   %-15s \n", result.StartTime.Time.Format(util.DATE_TIME), result.EndTime.Time.Format(util.DATE_TIME), timeLength)
	}

	return nil
}

func parseTime(timeRange string) (startTime, endTime *time.Time, err error) {
	if len(timeRange) == 0 {
		return nil, nil, errors.New("time range is required")
	}

	if matches, _ := matchRegex(timeRangeSimple, timeRange); matches != nil {
		// simple time range
		numberStr := matches[1]
		unit := matches[2]
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			return nil, nil, errors.New("invalid number in time range")
		}
		endTime := time.Now()
		var startTime time.Time
		switch unit {
		case "w":
			startTime = endTime.AddDate(0, 0, -7*number)
		case "d":
			startTime = endTime.AddDate(0, 0, -number)
		case "m":
			startTime = endTime.AddDate(0, -number, 0)
		case "y":
			startTime = endTime.AddDate(-number, 0, 0)
		default:
			return nil, nil, errors.New("invalid unit in time range")
		}

		return &startTime, &endTime, nil
	} else if matches, _ := matchRegex(timeRangeFull, timeRange); matches != nil {
		// full time range
		startStr := matches[1]
		endStr := matches[2]
		layout := "2006-01-02"
		startTime, err := time.Parse(layout, startStr)
		if err != nil {
			return nil, nil, errors.New("invalid start date")
		}
		endTime, err := time.Parse(layout, endStr)
		if err != nil {
			return nil, nil, errors.New("invalid end date")
		}

		return &startTime, &endTime, nil
	} else if matches, _ := matchRegex(timeRangeShort, timeRange); matches != nil {
		// short time range
		startStr := matches[1]
		endStr := matches[2]
		year := time.Now().Year()
		layout := "01/02/2006"
		startTime, err := time.Parse(layout, startStr+"/"+strconv.Itoa(year))
		if err != nil {
			return nil, nil, errors.New("invalid start date")
		}
		endTime, err := time.Parse(layout, endStr+"/"+strconv.Itoa(year))
		if err != nil {
			return nil, nil, errors.New("invalid end date")
		}

		return &startTime, &endTime, nil
	} else {
		return nil, nil, errors.New("invalid time range format")
	}
}

func matchRegex(regex, str string) ([]string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}
	matches := r.FindStringSubmatch(str)
	if len(matches) == 0 {
		return nil, nil
	}
	return matches, nil
}
