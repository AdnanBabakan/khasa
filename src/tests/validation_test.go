package tests

import (
	"khasa/src/utils/strings"
	"testing"
)

func TestIsValidInt(t *testing.T) {
	got := strings.IsValidInt("123")
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = strings.IsValidInt("123.123")
	want = false

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = strings.IsValidInt("abc")
	want = false

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestIsValidFloat(t *testing.T) {
	got := strings.IsValidFloat("123.123")
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = strings.IsValidFloat("123")
	want = true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = strings.IsValidFloat("abc")
	want = false

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
