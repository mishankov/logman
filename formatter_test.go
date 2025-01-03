package logman_test

import (
	"testing"

	"github.com/mishankov/logman"
)

func TestDefaultFormatter(t *testing.T) {
	formatter := logman.NewDefaultFormatter("<_logLevel_> <_callLocation_> <_dateTime_>: _message_")

	got := formatter.Format(logman.Debug, "fake date", "fake call location", "debug message")

	assertEqual(t, got, "<Debug> <fake call location> <fake date>: debug message")
}
