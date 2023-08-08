package tests

import (
	"khasa/src/utils/slices"
	"testing"
)

func TestContains(t *testing.T) {
	got := slices.Contains([]string{"a", "b", "c"}, "a")
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = slices.Contains([]int{1, 2, 3}, 1)
	want = true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
