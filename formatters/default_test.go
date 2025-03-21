package formatters_test

import (
	"context"
	"testing"
	"time"

	"github.com/mishankov/testman/assert"

	"github.com/mishankov/logman"
	"github.com/mishankov/logman/formatters"
)

func TestDefaultFormatter(t *testing.T) {
	formatter := formatters.NewDefaultFormatter("<_logLevel_> <_callLocation_> <_dateTime_>: _message_", formatters.DefaultTimeLayout)

	tm, _ := time.Parse("2006-01-02 15:04:05 GMT-0700", "2006-01-02 15:04:05 GMT-0700")
	got := formatter.Format(context.TODO(), logman.Debug, tm, "fake call location", "debug message")

	assert.Equal(t, got, "<Debug> <fake call location> <2006-01-02 15:04:05 GMT-0700>: debug message")
}

func TestPartialFields(t *testing.T) {
	formatter := formatters.NewDefaultFormatter("<_logLevel_> <_dateTime_>: _message_", formatters.DefaultTimeLayout)

	tm, _ := time.Parse("2006-01-02 15:04:05 GMT-0700", "2006-01-02 15:04:05 GMT-0700")
	got := formatter.Format(context.TODO(), logman.Debug, tm, "fake call location", "debug message")

	assert.Equal(t, got, "<Debug> <2006-01-02 15:04:05 GMT-0700>: debug message")
}

func TestStructuredParamsDefault(t *testing.T) {
	formatter := formatters.NewDefaultFormatter("<_logLevel_> <_callLocation_> <_dateTime_>: _message_ _params_", formatters.DefaultTimeLayout)

	tm, _ := time.Parse("2006-01-02 15:04:05 GMT-0700", "2006-01-02 15:04:05 GMT-0700")
	got := formatter.Format(context.TODO(), logman.Debug, tm, "fake call location", "debug message", "key", "someValue", "key2", 3)

	assert.Equal(t, got, "<Debug> <fake call location> <2006-01-02 15:04:05 GMT-0700>: debug message key=someValue key2=3")
}

func BenchmarkFormatter(b *testing.B) {
	formatter := formatters.NewDefaultFormatter(formatters.DefaultFormat, formatters.DefaultTimeLayout)
	tm, _ := time.Parse("2006-01-02 15:04:05 GMT-0700", "2006-01-02 15:04:05 GMT-0700")

	b.ResetTimer()
	for range b.N {
		formatter.Format(context.TODO(), logman.Debug, tm, "fake call location", "debug message")
	}
}
