package main

import (
	"reflect"
	"testing"
)

// Write a program that accepts a sentence and calculate the number of letters and digits.

func TestEx013(t *testing.T) {

	input := "hello world! 123"
	want := "LETTERS 10, DIGITS 3"

	got := Ex013(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nEx013() = %v\nwant    = %v", got, want)
	}

}
