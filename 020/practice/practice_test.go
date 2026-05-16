package practice

import (
	"reflect"
	"testing"
)

func TestEx020(t *testing.T) {
	cases := []struct {
		in   int
		want []int
	}{
		{20, []int{0, 7, 14}},
		{50, []int{0, 7, 14, 21, 28, 35, 42, 49}},
		{0, nil},
		{-5, nil},
	}

	for _, c := range cases {
		got := Ex020(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Ex020(%d) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestEx020Stream(t *testing.T) {
	var got []int
	for v := range Ex020Stream(50) {
		got = append(got, v)
	}

	want := []int{0, 7, 14, 21, 28, 35, 42, 49}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Ex020Stream(50) = %v, want %v", got, want)
	}
}
