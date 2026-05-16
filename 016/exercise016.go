package main

/*
Exercise 016:

Accept a comma-separated sequence of numbers and return, also comma-separated,
only the odd ones (preserving their original order).

Example:
  Ex016("1,2,3,4,5,6,7,8,9") -> "1,3,5,7,9"
*/

import (
	"strconv"
	"strings"
)

// Ex016 keeps only the odd numbers in the comma-separated input.
func Ex016(input string) string {
	parts := strings.Split(input, ",")
	out := []string{}

	for _, p := range parts {
		s := strings.TrimSpace(p)
		n, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		if n%2 != 0 {
			out = append(out, s)
		}
	}

	return strings.Join(out, ",")
}
