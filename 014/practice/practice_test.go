package practice

import "testing"

func TestEx014(t *testing.T) {
	input := "Hello world!"
	want := "UPPER CASE 1, LOWER CASE 9"

	if got := Ex014(input); got != want {
		t.Errorf("Ex014() = %v, want %v", got, want)
	}
}
