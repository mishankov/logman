package logman_test

import (
	"bytes"
	"testing"

	"github.com/mishankov/logman"
)

func TestDefaultFormatter(t *testing.T) {
	buffer := &bytes.Buffer{}
	timer := &FakeTimeProvider{}
	formatter := logman.NewDefaultFormatter("<_logLevel_> <_dateTime_>: _message_")

	logger := logman.NewLogger(buffer, timer, formatter)
	logger.Debug("debug message")

	AssertEqual(t, buffer.String(), "<Debug> <2006-01-02 15:04:05 GMT-0700>: debug message\n")
}
