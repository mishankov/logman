package logman_test

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/mishankov/logman"
	"github.com/mishankov/logman/formatters"
	"github.com/mishankov/logman/internal/testutils"
)

func TestCustomLogLevel(t *testing.T) {
	ll := logman.LogLevel(99)
	testutils.AssertEqual(t, ll.String(), "99")
}

func TestLogger(t *testing.T) {
	logger, buffer := testLoggerAndBuffer()

	message := "message"

	for _, logFunction := range loggerFunctions(logger) {
		logFunction(message)
		testutils.AssertContains(t, buffer.String(), message)
		buffer.Reset()
	}
}

func TestCompositeMessage(t *testing.T) {
	logger, buffer := testLoggerAndBuffer()

	message := []string{"composite", "message"}

	for _, logFunction := range loggerFunctions(logger) {
		logFunction(message[0], message[1])
		testutils.AssertContains(t, buffer.String(), strings.Join(message, " "))
		buffer.Reset()
	}
}

func TestFormattedMessages(t *testing.T) {
	logger, buffer := testLoggerAndBuffer()

	message := "my %v message %v"
	formats := []string{"awesome", "here"}

	for _, logFunction := range formatLoggerFunctions(logger) {
		logFunction(message, formats[0], formats[1])
		testutils.AssertContains(t, buffer.String(), fmt.Sprintf(message, formats[0], formats[1]))
		buffer.Reset()
	}
}

func TestStructuredMessages(t *testing.T) {
	logger, buffer := testLoggerAndBuffer()

	message := "my message "
	formats := []string{"key", "value"}

	for _, logFunction := range structLoggerFunctions(logger) {
		logFunction(message, formats[0], formats[1])
		testutils.AssertContains(t, buffer.String(), fmt.Sprintf("%v %v=%v", message, formats[0], formats[1]))
		buffer.Reset()
	}
}

var errTest = errors.New("some error")

func TestErrorsAsMessages(t *testing.T) {
	logger, buffer := testLoggerAndBuffer()

	for _, logFunction := range loggerFunctions(logger) {
		logFunction(errTest)
		testutils.AssertContains(t, buffer.String(), errTest.Error())
		buffer.Reset()
	}
}

func TestCallLocation(t *testing.T) {
	logger, buffer := testLoggerAndBuffer()

	// Check module and function names
	want := []string{"logman_test", "TestCallLocation"}

	t.Run("simple functions", func(t *testing.T) {
		for _, logFunction := range loggerFunctions(logger) {
			logFunction("some log")

			got := buffer.String()
			for _, s := range want {
				testutils.AssertContains(t, got, s)
			}

			buffer.Reset()
		}
	})

	t.Run("format functions", func(t *testing.T) {
		for _, logFunction := range formatLoggerFunctions(logger) {
			logFunction("some log")

			got := buffer.String()
			for _, s := range want {
				testutils.AssertContains(t, got, s)
			}

			buffer.Reset()
		}
	})

	t.Run("structured functions", func(t *testing.T) {
		for _, logFunction := range structLoggerFunctions(logger) {
			logFunction("some log")

			got := buffer.String()
			for _, s := range want {
				testutils.AssertContains(t, got, s)
			}

			buffer.Reset()
		}
	})

	t.Run("Log", func(t *testing.T) {
		logger.Log(logman.Debug, "some log")

		got := buffer.String()
		for _, s := range want {
			testutils.AssertContains(t, got, s)
		}

		buffer.Reset()
	})

	t.Run("Logf", func(t *testing.T) {
		logger.Logf(logman.Debug, "some log")

		got := buffer.String()
		for _, s := range want {
			testutils.AssertContains(t, got, s)
		}

		buffer.Reset()
	})

	t.Run("Logs", func(t *testing.T) {
		logger.Logs(logman.Debug, "some log")

		got := buffer.String()
		for _, s := range want {
			testutils.AssertContains(t, got, s)
		}

		buffer.Reset()
	})

}

func TestFilter(t *testing.T) {
	logger, buffer := testLoggerAndBuffer()
	message := "some message"

	t.Run("no filter should always log", func(t *testing.T) {
		logger.Log(logman.Debug, message)
		testutils.AssertContains(t, buffer.String(), message)
		buffer.Reset()
	})

	t.Run("no log if filter returns false", func(t *testing.T) {
		logger.Filter = &FakeFilter{false}
		logger.Log(logman.Debug, message)
		logger.Logf(logman.Debug, "%s", message)
		testutils.AssertEqual(t, buffer.Len(), 0)
		buffer.Reset()
	})

	t.Run("log if filter returns true", func(t *testing.T) {
		logger.Filter = &FakeFilter{true}
		logger.Log(logman.Debug, message)
		testutils.AssertContains(t, buffer.String(), message)
	})
}

func TestNewLine(t *testing.T) {
	testCases := []struct {
		f logman.Formatter
	}{
		{formatters.NewDefaultFormatter(formatters.DefaultFormat, formatters.DefaultTimeLayout)},
		{formatters.NewDefaultFormatter("_message_ [_dateTime_] [_logLevel_]", formatters.DefaultTimeLayout)},
		{formatters.NewJSONFormatter()},
	}

	for _, test := range testCases {
		logger, buffer := testLoggerAndBufferWithFormatter(test.f)

		logger.Debug("some message")

		got := buffer.String()
		if !strings.HasSuffix(got, "\n") {
			t.Errorf("Expected log %q line to end with new line", got)
		}
	}
}

// Mocks

// FakeFilter implements Filter interface for tests.
type FakeFilter struct {
	bool
}

func (ff *FakeFilter) Filter(_ logman.LogLevel, _, _ string) bool {
	return ff.bool
}

// Helpers

func testLoggerAndBuffer() (*logman.Logger, *bytes.Buffer) {
	buffer := &bytes.Buffer{}
	formatter := formatters.NewDefaultFormatter(formatters.DefaultFormat, formatters.DefaultTimeLayout)
	filter := &FakeFilter{true}
	logger := logman.NewLogger(buffer, formatter, filter)

	return logger, buffer
}

func testLoggerAndBufferWithFormatter(formatter logman.Formatter) (*logman.Logger, *bytes.Buffer) {
	buffer := &bytes.Buffer{}
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

func structLoggerFunctions(logger *logman.Logger) []func(string, ...any) {
	return []func(string, ...any){
		logger.Debugs, logger.Infos, logger.Warns, logger.Errors, logger.Fatals,
	}
}
