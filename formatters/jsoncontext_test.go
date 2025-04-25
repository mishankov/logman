package formatters_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/mishankov/testman/assert"

	"github.com/mishankov/logman"
	"github.com/mishankov/logman/formatters"
)

func TestJSONContextFormatter(t *testing.T) {
	formatter := formatters.NewJSONContextFormatter(formatters.DefaultTimeLayout, []fmt.Stringer{ContextKey1, ContextKey2})

	tm, _ := time.Parse("2006-01-02 15:04:05 GMT-0700", "2006-01-02 15:04:05 GMT-0700")
	got := formatter.Format(testContext(), logman.Debug, tm, "fake/call/location:44", "some message")

	assert.True(t, strings.HasPrefix(got, "{") && strings.HasSuffix(got, "}"))

	// Keys of JSON are not ordered, so check individual keys
	assert.Contains(t, got, `"level":"Debug"`)
	assert.Contains(t, got, `"location":"fake/call/location:44"`)
	assert.Contains(t, got, `"time":"2006-01-02 15:04:05 GMT-0700"`)
	assert.Contains(t, got, `"msg":"some message"`)
	assert.Contains(t, got, `"key1":3`)
	assert.Contains(t, got, `"key2":"some value"`)
}

func TestJSONContextFormatterWithParams(t *testing.T) {
	formatter := formatters.NewJSONContextFormatter(formatters.DefaultTimeLayout, []fmt.Stringer{ContextKey1, ContextKey2})

	tm, _ := time.Parse("2006-01-02 15:04:05 GMT-0700", "2006-01-02 15:04:05 GMT-0700")
	got := formatter.Format(testContext(), logman.Debug, tm, "fake/call/location:44", "some message", "param1", "someValue", "param2", 3)

	if !strings.HasPrefix(got, "{") || strings.HasSuffix(got, "}") {
		t.Errorf("%q is expected to be JSON", got)
	}

	assert.Contains(t, got, `"param1":"someValue"`)
	assert.Contains(t, got, `"param2":3`)
}
