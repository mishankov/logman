package formatters_test

import (
	"testing"
	"time"

	"github.com/mishankov/logman"
	"github.com/mishankov/logman/formatters"
	"github.com/mishankov/logman/internal/testutils"
)

func TestDefaultFormatter(t *testing.T) {
	formatter := formatters.NewDefaultFormatter("<_logLevel_> <_callLocation_> <_dateTime_>: _message_", formatters.DefaultTimeFormat)

	tm, _ := time.Parse("2006-01-02 15:04:05 GMT-0700", "2006-01-02 15:04:05 GMT-0700")
	got := formatter.Format(logman.Debug, tm, "fake call location", "debug message")

	testutils.AssertEqual(t, got, "<Debug> <fake call location> <2006-01-02 15:04:05 GMT-0700>: debug message")
}
