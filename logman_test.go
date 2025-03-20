package logman_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/mishankov/testman/assert"

	"github.com/mishankov/logman"
	"github.com/mishankov/logman/formatters"
)

func TestCustomLogLevel(t *testing.T) {
	ll := logman.LogLevel(99)
	assert.Equal(t, ll.String(), "99")
}

func TestLogger(t *testing.T) {
	logger, buffer := testLoggerAndBuffer()

	message := "message"

	for _, logFunction := range loggerFunctions(logger) {
		logFunction(message)
		assert.Contains(t, buffer.String(), message)
		buffer.Reset()
	}
}

func TestCompositeMessage(t *testing.T) {
	logger, buffer := testLoggerAndBuffer()

	message := []string{"composite", "message"}

	for _, logFunction := range loggerFunctions(logger) {
		logFunction(message[0], message[1])
		assert.Contains(t, buffer.String(), strings.Join(message, " "))
		buffer.Reset()
	}
}

func TestFormattedMessages(t *testing.T) {
	logger, buffer := testLoggerAndBuffer()

	message := "my %v message %v"
	formats := []string{"awesome", "here"}

	for _, logFunction := range formatLoggerFunctions(logger) {
		logFunction(message, formats[0], formats[1])
		assert.Contains(t, buffer.String(), fmt.Sprintf(message, formats[0], formats[1]))
		buffer.Reset()
	}
}

func TestStructuredMessages(t *testing.T) {
	logger, buffer := testLoggerAndBuffer()

	message := "my message "
	formats := []string{"key", "value"}

	for _, logFunction := range structLoggerFunctions(logger) {
		logFunction(message, formats[0], formats[1])
		assert.Contains(t, buffer.String(), fmt.Sprintf("%v %v=%v", message, formats[0], formats[1]))
		buffer.Reset()
	}
}

type ContextValueKey string

func (cvk ContextValueKey) String() string {
	return string(cvk)
}

const (
	ContextValueKey1 ContextValueKey = "key1"
	ContextValueKey2 ContextValueKey = "key2"
)

func TestMessagesWithContext(t *testing.T) {
	buffer := &bytes.Buffer{}
	logger := logman.NewLogger(buffer, formatters.NewDefaultContextFormatter(formatters.DefaultTimeLayout, []fmt.Stringer{ContextValueKey1, ContextValueKey2}), &FakeFilter{true})

	message := "my message"
	ctx := context.WithValue(context.Background(), ContextValueKey1, "val1")
	ctx = context.WithValue(ctx, ContextValueKey2, "val2")

	for _, logFunction := range contextLoggerFunctions(logger) {
		logFunction(ctx, message)
		assert.Contains(t, buffer.String(), fmt.Sprintf("%v=%v", ContextValueKey1, "val1"))
		assert.Contains(t, buffer.String(), fmt.Sprintf("%v=%v", ContextValueKey2, "val2"))
		buffer.Reset()
	}

	t.Run("LogsCtx", func(t *testing.T) {
		logger.LogsCtx(ctx, logman.Debug, message)
		assert.Contains(t, buffer.String(), fmt.Sprintf("%v=%v", ContextValueKey1, "val1"))
		assert.Contains(t, buffer.String(), fmt.Sprintf("%v=%v", ContextValueKey2, "val2"))
		buffer.Reset()
	})
}

var errTest = errors.New("some error")

func TestErrorsAsMessages(t *testing.T) {
	logger, buffer := testLoggerAndBuffer()

	for _, logFunction := range loggerFunctions(logger) {
		logFunction(errTest)
		assert.Contains(t, buffer.String(), errTest.Error())
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
				assert.Contains(t, got, s)
			}

			buffer.Reset()
		}
	})

	t.Run("format functions", func(t *testing.T) {
		for _, logFunction := range formatLoggerFunctions(logger) {
			logFunction("some log")

			got := buffer.String()
			for _, s := range want {
				assert.Contains(t, got, s)
			}

			buffer.Reset()
		}
	})

	t.Run("structured functions", func(t *testing.T) {
		for _, logFunction := range structLoggerFunctions(logger) {
			logFunction("some log")

			got := buffer.String()
			for _, s := range want {
				assert.Contains(t, got, s)
			}

			buffer.Reset()
		}
	})

	t.Run("Log", func(t *testing.T) {
		logger.Log(logman.Debug, "some log")

		got := buffer.String()
		for _, s := range want {
			assert.Contains(t, got, s)
		}

		buffer.Reset()
	})

	t.Run("Logf", func(t *testing.T) {
		logger.Logf(logman.Debug, "some log")

		got := buffer.String()
		for _, s := range want {
			assert.Contains(t, got, s)
		}

		buffer.Reset()
	})

	t.Run("Logs", func(t *testing.T) {
		logger.Logs(logman.Debug, "some log")

		got := buffer.String()
		for _, s := range want {
			assert.Contains(t, got, s)
		}

		buffer.Reset()
	})

}

func TestFilter(t *testing.T) {
	logger, buffer := testLoggerAndBuffer()
	message := "some message"

	t.Run("no filter should always log", func(t *testing.T) {
		logger.Log(logman.Debug, message)
		assert.Contains(t, buffer.String(), message)
		buffer.Reset()
	})

	t.Run("no log if filter returns false", func(t *testing.T) {
		logger.Filter = &FakeFilter{false}
		logger.Log(logman.Debug, message)
		logger.Logf(logman.Debug, "%s", message)
		assert.Equal(t, buffer.Len(), 0)
		buffer.Reset()
	})

	t.Run("log if filter returns true", func(t *testing.T) {
		logger.Filter = &FakeFilter{true}
		logger.Log(logman.Debug, message)
		assert.Contains(t, buffer.String(), message)
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

		assert.True(t, strings.HasSuffix(got, "\n"))
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

func contextLoggerFunctions(logger *logman.Logger) []func(context.Context, string, ...any) {
	return []func(context.Context, string, ...any){
		logger.DebugsCtx, logger.InfosCtx, logger.WarnsCtx, logger.ErrorsCtx, logger.FatalsCtx,
	}
}
