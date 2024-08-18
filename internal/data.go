package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

var (
	ErrFileNotFound      = fmt.Errorf("file not found")
	ErrInvalidCronFormat = fmt.Errorf("invalid cron format")
)

// Task represents a task
// Recurrence is a cron format but only supports Dom, Month, and Dow
// Dom: Day of the month: 1-31, "* / , - ?"
// Month: Month of the year: 1-12 or JAN-DEC, "* / , -"
// Dow: Day of the week: 0-6 or SUN-SAT, "* / , - ?"
type Task struct {
	Title       string `json:"title"`
	Recurrence  string `json:"recurrence"`
	Description string `json:"description"`
}

// Next returns the next time of the task
func (t *Task) Next(tt time.Time) (time.Time, error) {
	parser := cron.NewParser(cron.Dom | cron.Month | cron.Dow)
	schedule, err := parser.Parse(t.Recurrence)
	if err != nil {
		return time.Time{}, ErrInvalidCronFormat
	}
	return schedule.Next(tt), nil
}

// In returns true if the task is in the given time (only check the day)
func (t *Task) In(tt time.Time) (bool, error) {
	next, err := t.Next(tt.Add(-time.Hour * 24))
	if err != nil {
		return false, err
	}
	return next.Year() == tt.Year() && next.Month() == tt.Month() && next.Day() == tt.Day(), nil
}

func (t *Task) Verify() error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	schedule := cron.NewParser(cron.Dom | cron.Month | cron.Dow)
	_, err := schedule.Parse(t.Recurrence)
	if err != nil {
		return errors.Join(err, ErrInvalidCronFormat)
	}
	return nil
}

func LoadTasks(taskFile string) ([]Task, error) {
	file, err := os.Open(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFileNotFound
		}
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var tasks []Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, fmt.Errorf("failed to decode file: %w", err)
	}

	return tasks, nil
}

func SaveTasks(taskFile string, tasks []Task) error {
	file, err := os.Create(taskFile)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	if err := enc.Encode(tasks); err != nil {
		return fmt.Errorf("failed to encode file: %w", err)
	}

	return nil
}
