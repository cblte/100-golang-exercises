package practice

import (
	"reflect"
	"testing"
)

func TestEx007(t *testing.T) {
	want := [][]int{{0, 0, 0, 0, 0}, {0, 1, 2, 3, 4}, {0, 2, 4, 6, 8}}
	got := Ex007(3, 5)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Ex007() = %v, want %v", got, want)
	}
}
