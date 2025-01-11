package formatters_test

import (
	"strings"
	"testing"
	"time"

	"github.com/mishankov/logman"
	"github.com/mishankov/logman/formatters"
	"github.com/mishankov/logman/internal/testutils"
)

func TestJSONFormatter(t *testing.T) {
	formatter := formatters.NewJSONFormatter()

	tm, _ := time.Parse("2006-01-02 15:04:05 GMT-0700", "2006-01-02 15:04:05 GMT-0700")
	got := formatter.Format(logman.Debug, tm, "fake/call/location:44", "some message")

	if !(strings.HasPrefix(got, "{") && strings.HasSuffix(got, "}")) {
		t.Errorf("%q is expected to be JSON", got)
	}

	// Keys of JSON are not ordered, so check individual keys
	testutils.AssertContains(t, got, `"log_level":"Debug"`)
	testutils.AssertContains(t, got, `"call_location":"fake/call/location:44"`)
	testutils.AssertContains(t, got, `"message":"some message"`)
	testutils.AssertContains(t, got, `"date_time":"2006-01-02 15:04:05 GMT-0700"`)
}
