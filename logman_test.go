package logman

import (
	"bytes"
	"testing"
)

func TestLogger_Debug(t *testing.T) {
	buffer := &bytes.Buffer{}
	timer := &FakeTimeProvider{}

	logger := NewLogger(buffer, timer)
	logger.Debug("debug message")

	AssertEqual(t, buffer.String(), "[2006-01-02 15:04:05 GMT-0700] [DEBUG] - debug message")
}

func ExampleLogger_Debug() {
	logger := NewDefaultLogger()
	// Using fake time provider to
	logger.timer = &FakeTimeProvider{}

	logger.Debug("debug message")

	// Output: [2006-01-02 15:04:05 GMT-0700] [DEBUG] - debug message
}

// Mocks

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
