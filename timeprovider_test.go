package logman_test

import (
	"testing"
	"time"

	"github.com/mishankov/logman"
)

func TestTime(t *testing.T) {
	timeFormat := "2006-01-02 15:04:05"
	timeProvider := logman.NewDefaultTimeProvider(timeFormat)

	got := timeProvider.Time()
	_, err := time.Parse(timeFormat, got)

	if err != nil {
		t.Fatal("Time() returned date time in invalid format:", err)
	}

}
