package practice

import "testing"

func TestPrintString(t *testing.T) {
	SetInput("hello world!")

	want := "HELLO WORLD!"
	got := PrintString()

	if got != want {
		t.Errorf("PrintString() = %q, want %q", got, want)
	}
}
