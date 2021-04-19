package main

import (
	"reflect"
	"testing"
)

// Write a program that accepts a comma separated sequence of words as input
// and prints the words in a comma-separated sequence after sorting them
// alphabetically.
func TestEx008(t *testing.T) {

	input := "without,hello,bag,world"
	want := "bag,hello,without,world"

	got := Ex008(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Ex008() = %v, want %v", got, want)
	}
}
