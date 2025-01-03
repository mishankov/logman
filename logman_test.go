package logman_test

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/mishankov/logman"
)

func TestLogger(t *testing.T) {
	buffer := &bytes.Buffer{}
	timeFomatter := &FakeTimeFormatter{}
	formatter := logman.NewDefaultFormatter(logman.DefaultFormat)

	logger := logman.NewLogger(buffer, timeFomatter, formatter)

	tt := []struct {
		logFunction func(...any)
	}{
		{
			logFunction: logger.Debug,
		},
		{
			logFunction: logger.Info,
		},
		{
			logFunction: logger.Warn,
		},
		{
			logFunction: logger.Error,
		},
		{
			logFunction: logger.Fatal,
		},
	}

	message := "message"

	for _, test := range tt {
		test.logFunction(message)
		AssertContains(t, buffer.String(), message)
		buffer.Reset()
	}
}

func TestCompositeMessage(t *testing.T) {
	buffer := &bytes.Buffer{}
	timeFomatter := &FakeTimeFormatter{}
	formatter := logman.NewDefaultFormatter(logman.DefaultFormat)

	logger := logman.NewLogger(buffer, timeFomatter, formatter)

	tt := []struct {
		logFunction func(...any)
	}{
		{
			logFunction: logger.Debug,
		},
		{
			logFunction: logger.Info,
		},
		{
			logFunction: logger.Warn,
		},
		{
			logFunction: logger.Error,
		},
		{
			logFunction: logger.Fatal,
		},
	}

	message := []string{"composite", "message"}

	for _, test := range tt {
		test.logFunction(message[0], message[1])
		AssertContains(t, buffer.String(), strings.Join(message, " "))
		buffer.Reset()
	}
}

func TestFormatedMessages(t *testing.T) {
	buffer := &bytes.Buffer{}
	timeFomatter := &FakeTimeFormatter{}
	formatter := logman.NewDefaultFormatter(logman.DefaultFormat)

	logger := logman.NewLogger(buffer, timeFomatter, formatter)

	tt := []struct {
		logFunction func(string, ...any)
		want        string
	}{
		{
			logFunction: logger.Debugf,
		},
		{
			logFunction: logger.Infof,
		},
		{
			logFunction: logger.Warnf,
		},
		{
			logFunction: logger.Errorf,
		},
		{
			logFunction: logger.Fatalf,
		},
	}

	message := "my %v message %v"
	formats := []string{"awesome", "here"}

	for _, test := range tt {
		test.logFunction(message, formats[0], formats[1])
		AssertContains(t, buffer.String(), fmt.Sprintf(message, formats[0], formats[1]))
		buffer.Reset()
	}
}

func TestErrorsAsMessages(t *testing.T) {
	buffer := &bytes.Buffer{}
	timeFomatter := &FakeTimeFormatter{}
	formatter := logman.NewDefaultFormatter(logman.DefaultFormat)

	logger := logman.NewLogger(buffer, timeFomatter, formatter)

	tt := []struct {
		logFunction func(...any)
		want        string
	}{
		{
			logFunction: logger.Debug,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Debug] - some error\n",
		},
		{
			logFunction: logger.Info,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Info] - some error\n",
		},
		{
			logFunction: logger.Warn,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Warn] - some error\n",
		},
		{
			logFunction: logger.Error,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Error] - some error\n",
		},
		{
			logFunction: logger.Fatal,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Fatal] - some error\n",
		},
	}

	message := "some error"
	err := errors.New(message)

	for _, test := range tt {
		test.logFunction(err)
		AssertContains(t, buffer.String(), message)
		buffer.Reset()
	}
}

func TestCallLocation(t *testing.T) {
	buffer := &bytes.Buffer{}
	timeFomatter := &FakeTimeFormatter{}
	formatter := logman.NewDefaultFormatter(logman.DefaultFormat)

	logger := logman.NewLogger(buffer, timeFomatter, formatter)

	tt := []struct {
		logFunction func(...any)
	}{
		{
			logFunction: logger.Debug,
		},
		{
			logFunction: logger.Info,
		},
		{
			logFunction: logger.Warn,
		},
		{
			logFunction: logger.Error,
		},
		{
			logFunction: logger.Fatal,
		},
	}

	// Check module and function names
	want := []string{"logman_test", "TestCallLocation"}

	for _, test := range tt {
		test.logFunction("some log")
		got := buffer.String()

		for _, s := range want {
			AssertContains(t, got, s)
		}

		buffer.Reset()
	}
}

func ExampleLogger_Debug() {
	logger := logman.NewDefaultLogger()
	// Using fake time provider for test to pass. Remove it in your code
	logger.TimeFormatter = &FakeTimeFormatter{}

	logger.Debug("message")

	// Output: [2006-01-02 15:04:05 GMT-0700] [github.com/mishankov/logman_test.ExampleLogger_Debug:214] [Debug] - message
}

// Mocks

// FakeTimeFormatter implements TimeProvider interface for tests
type FakeTimeFormatter struct{}

func (ft *FakeTimeFormatter) Format(_ time.Time) string {
	return "2006-01-02 15:04:05 GMT-0700"
}

// Asserts

func AssertEqual(t *testing.T, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func AssertContains(t *testing.T, str, substr string) {
	t.Helper()
	if !strings.Contains(str, substr) {
		t.Errorf("expected %q to contain %q", str, substr)
	}
}
