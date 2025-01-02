package logman

import (
	"bytes"
	"testing"
)

func TestDefaultFormatter(t *testing.T) {
	buffer := &bytes.Buffer{}
	timer := &FakeTimeProvider{}
	formatter := NewDefaultFormatter("<_logLevel_> <_dateTime_>: _message_")

	logger := NewLogger(buffer, timer, formatter)
	logger.Debug("debug message")

	AssertEqual(t, buffer.String(), "<DEBUG> <2006-01-02 15:04:05 GMT-0700>: debug message")
}
