package practice

import "testing"

func TestEx008(t *testing.T) {
	input := "without,hello,bag,world"
	want := "bag,hello,without,world"
	got := Ex008(input)

	if got != want {
		t.Errorf("Ex008() = %v, want %v", got, want)
	}
}
