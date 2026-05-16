package main

import "testing"

func TestEx018(t *testing.T) {
	input := "ABd1234@1,a F1#,2w3E*,2We3345"
	want := "ABd1234@1"

	if got := Ex018(input); got != want {
		t.Errorf("Ex018() = %q, want %q", got, want)
	}
}

func TestEx018_MultipleValid(t *testing.T) {
	input := "Abcd1@xy,Abcd1@xy,short1@,NoDigit@x,NOLOWER1@"
	want := "Abcd1@xy,Abcd1@xy"

	if got := Ex018(input); got != want {
		t.Errorf("Ex018() = %q, want %q", got, want)
	}
}
