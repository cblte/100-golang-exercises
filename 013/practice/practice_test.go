package practice

import "testing"

func TestEx013(t *testing.T) {
	input := "hello world! 123"
	want := "LETTERS 10, DIGITS 3"
	got := Ex013(input)

	if got != want {
		t.Errorf("\nEx013() = %v\nwant    = %v", got, want)
	}
}
