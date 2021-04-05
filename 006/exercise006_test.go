package main

import (
	"reflect"
	"testing"
)

// TestEx006 test if Ex006 does correct calculation of
// Q = Square root of [(2 * C * D)/H]
func TestEx006(t *testing.T) {

	input := []int{100, 150, 180}
	want := []int{18, 22, 24}
	got := Ex006(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Ex006() = %v, want %v", got, want)
	}
}
