package practice

import "testing"

func TestEx011(t *testing.T) {
	input := "0100,0011,1010,1001"
	want := "1010"
	got := Ex011(input)

	if got != want {
		t.Errorf("Ex011() = %v, want %v", got, want)
	}
}
