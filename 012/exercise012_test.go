package main

import (
	"reflect"
	"testing"
)

// Write a program, which will find all such numbers
// between 100 and 300 (both included) such that each digit of
// the number is an even number. The numbers obtained should be printed
// in a comma-separated sequence on a single line.

func TestEx012(t *testing.T) {

	inputA := 100
	inputB := 300
	want := "200,202,204,206,208,220,222,224,226,228,240,242,244,246,248,260,262,264,266,268,280,282,284,286,288"

	got := Ex012(inputA, inputB)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nEx012() = %v\nwant    = %v", got, want)
	}

}
