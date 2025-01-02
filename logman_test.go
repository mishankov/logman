package logman

import (
	"bytes"
	"testing"
)

func TestLogger_Debug(t *testing.T) {
	buffer := &bytes.Buffer{}
	timer := &FakeTimeProvider{}
	formatter := NewDefaultFormatter("<_logLevel_> <_dateTime_>: _message_")

	logger := NewLogger(buffer, timer, formatter)
	logger.Debug("debug message")

	AssertEqual(t, buffer.String(), "<DEBUG> <2006-01-02 15:04:05 GMT-0700>: debug message")
}

func ExampleLogger_Debug() {
	logger := NewDefaultLogger()
	// Using fake time provider for test to pass. Remove it in your code
	logger.timer = &FakeTimeProvider{}

	logger.Debug("debug message")

	// Output: [2006-01-02 15:04:05 GMT-0700] [DEBUG] - debug message
}

// Mocks

// FakeTimeProvider implements TimeProvider interface for tests
type FakeTimeProvider struct{}

func (ft *FakeTimeProvider) Time() string {
	return "2006-01-02 15:04:05 GMT-0700"
}

// Asserts

func AssertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}
