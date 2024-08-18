package internal

import (
	"fmt"
	"testing"
	"time"
)

func tdate(y int, m time.Month, d int) time.Time {
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

func TestTask_In(t *testing.T) {
	cases := []struct {
		name string
		task Task
		day  time.Time
		want bool
	}{
		{
			name: "same day @daily",
			task: Task{
				Recurrence: "* * *",
			},
			day:  tdate(2024, time.August, 18),
			want: true,
		},
		{
			name: "not same day @weekly (every Monday)",
			task: Task{
				Recurrence: fmt.Sprintf("* * %d", time.Monday),
			},
			day:  tdate(2024, time.August, 18), // Sunday
			want: false,
		},
		{
			name: "same day @weekly (every Monday)",
			task: Task{
				Recurrence: fmt.Sprintf("* * %d", time.Monday),
			},
			day:  tdate(2024, time.August, 19), // Monday
			want: true,
		},
		{
			name: "same day @monthly (every 1th)",
			task: Task{
				Recurrence: "1 * *",
			},
			day:  tdate(2024, time.August, 1),
			want: true,
		},
		{
			name: "not same day @monthly (every 1th)",
			task: Task{
				Recurrence: "1 * *",
			},
			day:  tdate(2024, time.August, 2),
			want: false,
		},
		{
			name: "same day every 6 month",
			task: Task{
				Recurrence: "1 1,7 *",
			},
			day:  tdate(2024, time.January, 1),
			want: true,
		},
		{
			name: "not same day every 6 month",
			task: Task{
				Recurrence: "1 1,7 *",
			},
			day:  tdate(2024, time.August, 1),
			want: false,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.task.In(tt.day)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
