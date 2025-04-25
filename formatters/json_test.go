package formatters_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/mishankov/testman/assert"

	"github.com/mishankov/logman"
	"github.com/mishankov/logman/formatters"
)

func TestJSONFormatter(t *testing.T) {
	formatter := formatters.NewJSONFormatter()

	tm, _ := time.Parse("2006-01-02 15:04:05 GMT-0700", "2006-01-02 15:04:05 GMT-0700")
	got := formatter.Format(context.TODO(), logman.Debug, tm, "fake/call/location:44", "some message")

	assert.True(t, strings.HasPrefix(got, "{") && strings.HasSuffix(got, "}"))

	// Keys of JSON are not ordered, so check individual keys
	assert.Contains(t, got, `"logLevel":"Debug"`)
	assert.Contains(t, got, `"callLocation":"fake/call/location:44"`)
	assert.Contains(t, got, `"message":"some message"`)
	assert.Contains(t, got, `"time":"2006-01-02 15:04:05 GMT-0700"`)
}

func TestStructuredParamsJSON(t *testing.T) {
	formatter := formatters.NewJSONFormatter()

	tm, _ := time.Parse("2006-01-02 15:04:05 GMT-0700", "2006-01-02 15:04:05 GMT-0700")
	got := formatter.Format(context.TODO(), logman.Debug, tm, "fake/call/location:44", "some message", "key", "someValue", "key2", 3)

	if !strings.HasPrefix(got, "{") || !strings.HasSuffix(got, "}") {
		t.Errorf("%q is expected to be JSON", got)
	}

	assert.Contains(t, got, `"key":"someValue"`)
	assert.Contains(t, got, `"key2":3`)
}
