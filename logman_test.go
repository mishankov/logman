package logman_test

import (
	"bytes"
	"testing"

	"github.com/mishankov/logman"
)

func TestLogger(t *testing.T) {
	buffer := &bytes.Buffer{}
	timer := &FakeTimeProvider{}
	formatter := logman.NewDefaultFormatter(logman.DefaultFormat)

	logger := logman.NewLogger(buffer, timer, formatter)

	tt := []struct {
		logFunction func(...string)
		want        string
	}{
		{
			logFunction: logger.Debug,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Debug] - message\n",
		},
		{
			logFunction: logger.Info,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Info] - message\n",
		},
		{
			logFunction: logger.Warn,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Warn] - message\n",
		},
		{
			logFunction: logger.Error,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Error] - message\n",
		},
		{
			logFunction: logger.Fatal,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Fatal] - message\n",
		},
	}

	message := "message"

	for _, test := range tt {
		test.logFunction(message)
		AssertEqual(t, buffer.String(), test.want)
		buffer.Reset()
	}
}

func TestCompositeMessage(t *testing.T) {
	buffer := &bytes.Buffer{}
	timer := &FakeTimeProvider{}
	formatter := logman.NewDefaultFormatter(logman.DefaultFormat)

	logger := logman.NewLogger(buffer, timer, formatter)

	tt := []struct {
		logFunction func(...string)
		want        string
	}{
		{
			logFunction: logger.Debug,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Debug] - composite message\n",
		},
		{
			logFunction: logger.Info,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Info] - composite message\n",
		},
		{
			logFunction: logger.Warn,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Warn] - composite message\n",
		},
		{
			logFunction: logger.Error,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Error] - composite message\n",
		},
		{
			logFunction: logger.Fatal,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Fatal] - composite message\n",
		},
	}

	message := []string{"composite", "message"}

	for _, test := range tt {
		test.logFunction(message...)
		AssertEqual(t, buffer.String(), test.want)
		buffer.Reset()
	}
}

func TestFormatedMessages(t *testing.T) {
	buffer := &bytes.Buffer{}
	timer := &FakeTimeProvider{}
	formatter := logman.NewDefaultFormatter(logman.DefaultFormat)

	logger := logman.NewLogger(buffer, timer, formatter)

	tt := []struct {
		logFunction func(string, ...any)
		want        string
	}{
		{
			logFunction: logger.Debugf,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Debug] - my awesome message here\n",
		},
		{
			logFunction: logger.Infof,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Info] - my awesome message here\n",
		},
		{
			logFunction: logger.Warnf,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Warn] - my awesome message here\n",
		},
		{
			logFunction: logger.Errorf,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Error] - my awesome message here\n",
		},
		{
			logFunction: logger.Fatalf,
			want:        "[2006-01-02 15:04:05 GMT-0700] [Fatal] - my awesome message here\n",
		},
	}

	message := "my %v message %v"
	formats := []string{"awesome", "here"}

	for _, test := range tt {
		test.logFunction(message, formats[0], formats[1])
		AssertEqual(t, buffer.String(), test.want)
		buffer.Reset()
	}
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
