package meetup

import (
	"time"
)

// WeekSchedule doc here
type WeekSchedule int

const (
	// Teenth doc here
	Teenth = iota
	// First doc here
	First
	// Second doc here
	Second
	// Third doc here
	Third
	// Fourth doc here
	Fourth
	// Last doc here
	Last
)

func findWeekDay(year int, month time.Month, weekday time.Weekday) time.Time {
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	delta := ((weekday - firstDay.Weekday()) + 7) % 7
	return firstDay.AddDate(0, 0, int(delta))
}

// Day doc here
func Day(schedule WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	t := findWeekDay(year, month, weekday)
	switch schedule {
	case Teenth:
		if t.Day() < 6 {
			return t.Day() + 14
		}
		return t.Day() + 7
	case First:
		return t.Day()
	case Second:
		return t.Day() + 7
	case Third:
		return t.Day() + 14
	case Fourth:
		return t.Day() + 21
	case Last:
		max := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
		return t.Day() + (max.Day()-t.Day())/7*7
	}
	return 0
}
