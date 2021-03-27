package main

import (
	"reflect"
	"testing"
)

func TestEx004(t *testing.T) {

	input := "12, 23, 34, 45, 56, 67, 78, 89, 90"
	want := []int{12, 23, 34, 45, 56, 67, 78, 89, 90}
	got := Ex004(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Ex004() = %v, want %v", got, want)
	}

}
