package models

import "time"

type LeaderboardOccurrence struct {
	User        User
	Substance   Substance
	UserID      uint `json:"user_id"`
	SubstanceID uint `json:"substance_id"`
	Occurrence  int64
}

type TimeDuration string

const (
	TODAY     TimeDuration = "TODAY"
	YESTERDAY TimeDuration = "YESTERDAY"
	THIS_WEEK TimeDuration = "THIS_WEEK"

	WEEK TimeDuration = "WEEK"

	THIS_MONTH TimeDuration = "THIS_MONTH"

	MONTH TimeDuration = "MONTH"

	THIS_YEAR TimeDuration = "THIS_YEAR"

	YEAR TimeDuration = "YEAR"

	ALL TimeDuration = "ALL"
)

var Durations = []TimeDuration{
	TODAY, YESTERDAY, THIS_WEEK, WEEK,
	THIS_MONTH, MONTH, THIS_YEAR, YEAR, ALL,
}

func GetTimeDuration(durationStr string) TimeDuration {
	switch durationStr {
	case "TODAY":
		return TODAY
	case "YESTERDAY":
		return YESTERDAY
	case "THIS_WEEK":
		return THIS_WEEK
	case "WEEK":
		return WEEK
	case "THIS_MONTH":
		return THIS_MONTH
	case "MONTH":
		return MONTH
	case "THIS_YEAR":
		return THIS_YEAR
	case "YEAR":
		return YEAR
	case "ALL":
		return ALL
	default:
		return ""
	}
}

func GetTimeValue(duration TimeDuration) time.Time {
	now := time.Now().Local()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	switch duration {
	case TODAY:
		return midnight
	case YESTERDAY:
		return midnight.AddDate(0, 0, -1)
	case THIS_WEEK:
		weekday := time.Duration(midnight.Weekday())
		monday := midnight.Add(-weekday * 24 * time.Hour)
		return monday
	case WEEK:
		return midnight.AddDate(0, 0, -7)
	case THIS_MONTH:
		year, month, _ := midnight.Date()
		firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, midnight.Location())
		return firstOfMonth
	case MONTH:
		return midnight.AddDate(0, -1, 0)
	case YEAR:
		return midnight.AddDate(-1, 0, 0)

	case THIS_YEAR:
		year := time.Date(midnight.Year(), time.January, 1, 0, 0, 0, 0, midnight.Location())
		return year
	default:
		return time.Time{} // Invalid duration, return zero time
	}
}
