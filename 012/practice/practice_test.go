package practice

import "testing"

func TestEx012(t *testing.T) {
	want := "200,202,204,206,208,220,222,224,226,228,240,242,244,246,248,260,262,264,266,268,280,282,284,286,288"
	got := Ex012(100, 300)

	if got != want {
		t.Errorf("\nEx012() = %v\nwant    = %v", got, want)
	}
}
