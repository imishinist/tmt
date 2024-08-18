package internal

import (
	"time"

	"github.com/robfig/cron/v3"
)

func ParseRecurrence(recurrence string) (cron.Schedule, error) {
	parser := cron.NewParser(cron.Dom | cron.Month | cron.Dow)
	schedule, err := parser.Parse(recurrence)
	if err != nil {
		return nil, ErrInvalidCronFormat
	}
	return schedule, nil
}

func Match(schedule cron.Schedule, tt time.Time) bool {
	next := schedule.Next(tt.Add(-time.Hour * 24))
	return SameDay(next, tt)
}

func SameDay(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
}
