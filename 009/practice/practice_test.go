package practice

import (
	"reflect"
	"testing"
)

func TestEx009(t *testing.T) {
	input := []string{"Hello world", "Practice makes perfect"}
	want := []string{"HELLO WORLD", "PRACTICE MAKES PERFECT"}
	got := Ex009(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Ex009() = %v, want %v", got, want)
	}
}
