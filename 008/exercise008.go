package main

import "strings"

// Ex008 takes two numbers and generate a slice
// Input: without,hello,bag,world
// Output: bag,hello,without,world
func Ex008(input string) string {

	output := ""

	slice := strings.Split(input, ",")
	// sorting
	sorted := true

	for sorted {
		sorted = false
		for l := len(slice) - 1; l > 0; l-- {
			if slice[l] < slice[l-1] {
				slice[l], slice[l-1] = slice[l-1], slice[l]
				sorted = true
			}
		}
	}

	output = strings.Join(slice, ",")

	return output
}
