package main

import (
	"reflect"
	"testing"
)

// TestEx006 test if Ex006 does correct calculation of
// Q = Square root of [(2 * C * D)/H]
func TestEx007(t *testing.T) {

	a := 3
	b := 5
	want := [][]int{{0, 0, 0, 0, 0}, {0, 1, 2, 3, 4}, {0, 2, 4, 6, 8}}
	got := Ex007(a, b)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Ex007() = %v, want %v", got, want)
	}
}
