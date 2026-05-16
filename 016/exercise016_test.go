package main

import "testing"

func TestEx016(t *testing.T) {
	input := "1,2,3,4,5,6,7,8,9"
	want := "1,3,5,7,9"

	if got := Ex016(input); got != want {
		t.Errorf("Ex016() = %v, want %v", got, want)
	}
}
