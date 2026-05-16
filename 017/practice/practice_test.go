package practice

import "testing"

func TestEx017(t *testing.T) {
	input := []string{"D 300", "D 300", "W 200", "D 100"}
	want := 500

	if got := Ex017(input); got != want {
		t.Errorf("Ex017() = %v, want %v", got, want)
	}
}

func TestEx017_IgnoresGarbage(t *testing.T) {
	input := []string{"", "D 1000", "X 999", "W 250", "bogus"}
	want := 750

	if got := Ex017(input); got != want {
		t.Errorf("Ex017() = %v, want %v", got, want)
	}
}
