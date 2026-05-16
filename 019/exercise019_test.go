package main

import (
	"reflect"
	"testing"
)

func TestEx019(t *testing.T) {
	input := []string{
		"Tom,19,80",
		"John,20,90",
		"Jony,17,91",
		"Jony,17,93",
		"Json,21,85",
	}
	want := []string{
		"John,20,90",
		"Jony,17,91",
		"Jony,17,93",
		"Json,21,85",
		"Tom,19,80",
	}

	got := Ex019(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Ex019() = %v, want %v", got, want)
	}
}
