package logman

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	timeFormat := "2006-01-02 15:04:05"
	timeProvider := NewDefaultTimeProvider(timeFormat)

	got := timeProvider.Time()
	_, err := time.Parse(timeFormat, got)

	if err != nil {
		t.Fatal("Time() returned date time in invalid format:", err)
	}

}
