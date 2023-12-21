package date

import "time"

func FormatAsISO8601(t time.Time) string {
	return t.Format("2006-01-02T15:04:05Z07:00")
}