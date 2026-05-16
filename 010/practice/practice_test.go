package practice

import "testing"

func TestEx010(t *testing.T) {
	input := "hello world and practice makes perfect and hello world again"
	want := "again and hello makes perfect practice world"
	got := Ex010(input)

	if got != want {
		t.Errorf("Ex010() = %v, want %v", got, want)
	}
}
