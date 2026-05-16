package practice

import (
	"reflect"
	"testing"
)

func TestEx006(t *testing.T) {
	input := []int{100, 150, 180}
	want := []int{18, 22, 24}
	got := Ex006(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Ex006() = %v, want %v", got, want)
	}
}
