package testutils

import (
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func Assert(t *testing.T, condition bool) {
	t.Helper()
	if !condition {
		t.Error("Condition expected to be true")
	}
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertDeepEqual[T any](t *testing.T, got, want T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertContains(t *testing.T, str, substr string) {
	t.Helper()
	if !strings.Contains(str, substr) {
		t.Errorf("expected %q to contain %q", str, substr)
	}
}

func AssertRegex(t *testing.T, got, wantRegex string) {
	t.Helper()
	r, err := regexp.Compile(wantRegex)
	if err != nil {
		t.Errorf("Regexp %q didn't compile: %v", wantRegex, err)
	}

	if !r.MatchString(got) {
		t.Errorf("%q didn't match regexp %q", got, wantRegex)
	}
}
