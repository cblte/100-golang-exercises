package practice

import "testing"

func TestEx015(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{9, 11106},
		{1, 1234},
		{2, 2468},
	}

	for _, c := range cases {
		if got := Ex015(c.in); got != c.want {
			t.Errorf("Ex015(%d) = %d, want %d", c.in, got, c.want)
		}
	}
}
