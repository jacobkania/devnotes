package dbutil

import "time"

func CurrentTimeUTCZ() string {
	return TimeUTCZ(time.Now())
}

func TimeUTCZ(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}
