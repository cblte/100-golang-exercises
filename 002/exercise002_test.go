package main

import (
	"testing"
)

func TestEx002(t *testing.T) {

	// check for error
	want := 0
	got, err := Ex002(-10)
	if err == nil {
		t.Errorf("Ex002() = %v, want %v", got, err)
	}

	want = 1
	got, err = Ex002(0)
	if err == nil && got != uint64(want) {
		t.Errorf("Expected '%v' but got '%v'", 1, got)
	}

	want = 40320
	got, err = Ex002(8)
	if err == nil && got != uint64(want) {
		t.Errorf("Ex002() = %v, want %v", got, want)
	}

}
