package formatters_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/mishankov/logman"
	"github.com/mishankov/logman/formatters"
	"github.com/mishankov/logman/internal/testutils"
)

func TestDefaultContextFormatter(t *testing.T) {
	formatter := formatters.NewDefaultContextFormatter(formatters.DefaultTimeLayout, []fmt.Stringer{ContextKey1, ContextKey2})
	tm, _ := time.Parse("2006-01-02 15:04:05 GMT-0700", "2006-01-02 15:04:05 GMT-0700")
	got := formatter.Format(testContext(), logman.Debug, tm, "fake call location", "debug message")

	testutils.AssertEqual(t, got, `time="2006-01-02 15:04:05 GMT-0700" level=Debug msg="debug message" key1=3 key2="some value"`)
}

func TestDefaultContextFormatterWithParams(t *testing.T) {
	formatter := formatters.NewDefaultContextFormatter(formatters.DefaultTimeLayout, []fmt.Stringer{ContextKey1, ContextKey2})
	tm, _ := time.Parse("2006-01-02 15:04:05 GMT-0700", "2006-01-02 15:04:05 GMT-0700")
	got := formatter.Format(testContext(), logman.Debug, tm, "fake call location", "debug message", "param1", 1, "param2", "value2")

	testutils.AssertEqual(t, got, `time="2006-01-02 15:04:05 GMT-0700" level=Debug msg="debug message" key1=3 key2="some value" param1=1 param2=value2`)
}

type ContextKey string

const (
	ContextKey1 ContextKey = "key1"
	ContextKey2 ContextKey = "key2"
)

func (k ContextKey) String() string { return string(k) }

func testContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, ContextKey1, 3)
	ctx = context.WithValue(ctx, ContextKey2, "some value")
	return ctx
}
