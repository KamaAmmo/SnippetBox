package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, value, expectedValue T) {
	t.Helper()

	if value != expectedValue {
		t.Errorf("got: %v; want: %v", value, expectedValue)
	}
}

func StringContains(t *testing.T, actual, expectedSubstring string) {
	t.Helper()

	if !strings.Contains(actual, expectedSubstring) {
		t.Errorf("got: %q; expected to contain: %q", actual, expectedSubstring)
	}
}
