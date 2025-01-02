package logman_test

import (
	"bytes"
	"testing"

	"github.com/mishankov/logman"
)

func TestLogger_Debug(t *testing.T) {
	buffer := &bytes.Buffer{}
	timer := &FakeTimeProvider{}
	formatter := logman.NewDefaultFormatter(logman.DefaultFormat)

	logger := logman.NewLogger(buffer, timer, formatter)
	logger.Debug("message")

	AssertEqual(t, buffer.String(), "[2006-01-02 15:04:05 GMT-0700] [Debug] - message")
}

func TestLogger_Info(t *testing.T) {
	buffer := &bytes.Buffer{}
	timer := &FakeTimeProvider{}
	formatter := logman.NewDefaultFormatter(logman.DefaultFormat)

	logger := logman.NewLogger(buffer, timer, formatter)
	logger.Info("message")

	AssertEqual(t, buffer.String(), "[2006-01-02 15:04:05 GMT-0700] [Info] - message")
}

func TestLogger_Warn(t *testing.T) {
	buffer := &bytes.Buffer{}
	timer := &FakeTimeProvider{}
	formatter := logman.NewDefaultFormatter(logman.DefaultFormat)

	logger := logman.NewLogger(buffer, timer, formatter)
	logger.Warn("message")

	AssertEqual(t, buffer.String(), "[2006-01-02 15:04:05 GMT-0700] [Warn] - message")
}

func TestLogger_Error(t *testing.T) {
	buffer := &bytes.Buffer{}
	timer := &FakeTimeProvider{}
	formatter := logman.NewDefaultFormatter(logman.DefaultFormat)

	logger := logman.NewLogger(buffer, timer, formatter)
	logger.Error("message")

	AssertEqual(t, buffer.String(), "[2006-01-02 15:04:05 GMT-0700] [Error] - message")
}

func TestLogger_Fatal(t *testing.T) {
	buffer := &bytes.Buffer{}
	timer := &FakeTimeProvider{}
	formatter := logman.NewDefaultFormatter(logman.DefaultFormat)

	logger := logman.NewLogger(buffer, timer, formatter)
	logger.Fatal("message")

	AssertEqual(t, buffer.String(), "[2006-01-02 15:04:05 GMT-0700] [Fatal] - message")
}

func ExampleLogger_Debug() {
	logger := logman.NewDefaultLogger()
	// Using fake time provider for test to pass. Remove it in your code
	logger.Timer = &FakeTimeProvider{}

	logger.Debug("message")

	// Output: [2006-01-02 15:04:05 GMT-0700] [Debug] - message
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
