package main

import (
	"reflect"
	"testing"
)

// Write a program that accepts a sequence of whitespace separated words as
// input and prints the words after removing all duplicate words and sorting
// them alphanumerically.
func TestEx010(t *testing.T) {

	input := "hello world and practice makes perfect and hello world again"
	want := "again and hello makes perfect practice world"

	got := Ex010(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Ex010() = %v, want %v", got, want)
	}
}
