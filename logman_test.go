package logman_test

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/mishankov/logman"
)

func TestLogger(t *testing.T) {
	logger, buffer := testLoggerAndBuffer()

	message := "message"

	for _, logFunction := range loggerFunctions(logger) {
		logFunction(message)
		assertContains(t, buffer.String(), message)
		buffer.Reset()
	}
}

func TestCompositeMessage(t *testing.T) {
	logger, buffer := testLoggerAndBuffer()

	message := []string{"composite", "message"}

	for _, logFunction := range loggerFunctions(logger) {
		logFunction(message[0], message[1])
		assertContains(t, buffer.String(), strings.Join(message, " "))
		buffer.Reset()
	}
}

func TestFormatedMessages(t *testing.T) {
	logger, buffer := testLoggerAndBuffer()

	message := "my %v message %v"
	formats := []string{"awesome", "here"}

	for _, logFunction := range formatLoggerFunctions(logger) {
		logFunction(message, formats[0], formats[1])
		assertContains(t, buffer.String(), fmt.Sprintf(message, formats[0], formats[1]))
		buffer.Reset()
	}
}

func TestErrorsAsMessages(t *testing.T) {
	logger, buffer := testLoggerAndBuffer()

	message := "some error"
	err := errors.New(message)

	for _, logFunction := range loggerFunctions(logger) {
		logFunction(err)
		assertContains(t, buffer.String(), message)
		buffer.Reset()
	}
}

func TestCallLocation(t *testing.T) {
	logger, buffer := testLoggerAndBuffer()

	// Check module and function names
	want := []string{"logman_test", "TestCallLocation"}

	for _, logFunction := range loggerFunctions(logger) {
		logFunction("some log")
		got := buffer.String()

		for _, s := range want {
			assertContains(t, got, s)
		}

		buffer.Reset()
	}
}

func TestFilter(t *testing.T) {
	logger, buffer := testLoggerAndBuffer()
	message := "some message"

	t.Run("no filter should always log", func(t *testing.T) {
		logger.Log(logman.Debug, message)
		assertContains(t, buffer.String(), message)
		buffer.Reset()
	})

	t.Run("no log if filter returns false", func(t *testing.T) {
		logger.Filter = &FakeFilter{false}
		logger.Log(logman.Debug, message)
		logger.Logf(logman.Debug, "%s", message)
		assertEqual(t, buffer.Len(), 0)
		buffer.Reset()
	})

	t.Run("log if filter returns true", func(t *testing.T) {
		logger.Filter = &FakeFilter{true}
		logger.Log(logman.Debug, message)
		assertContains(t, buffer.String(), message)
	})
}

// Mocks

// FakeFilter implements Filter interface for tests
type FakeFilter struct {
	bool
}

func (ff *FakeFilter) Filter(logLevel logman.LogLevel, callLocation string, message string) bool {
	return ff.bool
}

// Helpers

func testLoggerAndBuffer() (*logman.Logger, *bytes.Buffer) {
	buffer := &bytes.Buffer{}
	formatter := logman.NewDefaultFormatter(logman.DefaultFormat, logman.DefaultTimeFormat)
	filter := &FakeFilter{true}
	logger := logman.NewLogger(buffer, formatter, filter)

	return logger, buffer
}

func loggerFunctions(logger *logman.Logger) []func(...any) {
	return []func(...any){
		logger.Debug, logger.Info, logger.Warn, logger.Error, logger.Fatal,
	}
}

func formatLoggerFunctions(logger *logman.Logger) []func(string, ...any) {
	return []func(string, ...any){
		logger.Debugf, logger.Infof, logger.Warnf, logger.Errorf, logger.Fatalf,
	}
}

// Asserts

func assertEqual(t *testing.T, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertContains(t *testing.T, str, substr string) {
	t.Helper()
	if !strings.Contains(str, substr) {
		t.Errorf("expected %q to contain %q", str, substr)
	}
}
