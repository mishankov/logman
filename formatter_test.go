package logman_test

import (
	"testing"
	"time"

	"github.com/mishankov/logman"
)

func TestDefaultFormatter(t *testing.T) {
	formatter := logman.NewDefaultFormatter("<_logLevel_> <_callLocation_> <_dateTime_>: _message_", logman.DefaultTimeFormat)

	tm, _ := time.Parse("2006-01-02 15:04:05 GMT-0700", "2006-01-02 15:04:05 GMT-0700")
	got := formatter.Format(logman.Debug, tm, "fake call location", "debug message")

	assertEqual(t, got, "<Debug> <fake call location> <2006-01-02 15:04:05 GMT-0700>: debug message")
}
