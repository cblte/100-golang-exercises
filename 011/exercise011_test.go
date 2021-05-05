package main

import (
	"reflect"
	"testing"
)

// Write a program which accepts a sequence of comma separated 4 digit binary
// numbers as its input and then check whether they are divisible by 5 or not.
// The numbers that are divisible by 5 are to be printed in a comma
// separated sequence.

func TestEx011(t *testing.T) {

	input := "0100,0011,1010,1001"
	want := "1010"

	got := Ex011(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Ex011() = %v, want %v", got, want)
	}

}
