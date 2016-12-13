package util

import "testing"

func TestRandString(t *testing.T) {
	s := RandString(10)
	if len(s) != 10 {
		t.Fatalf("Generated 10 characters string, but got \"%s\"(%d charactes)", s, len(s))
	}
}
