package main

import (
	"testing"
)

func TestEx001(t *testing.T) {

	want := "112,119,126,133,147,154,161,168,182,189,196"
	got := Ex001(100, 200)

	if got != want {
		t.Errorf("Ex001() = %v, want %v", got, want)
	}
}
