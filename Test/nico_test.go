package main

import (
	"testing"
)

func TestAbs(t *testing.T) {
	got := Abs(-2)
	if got != 1 {
		t.Errorf("%d", got)
	}
}
