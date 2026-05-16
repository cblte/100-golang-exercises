package main

/*
Exercise 019:

Sort a list of "name,age,score" lines in ascending order by name, then age,
then score (all compared as strings, like the original Python challenge).

Example input:
  []string{
    "Tom,19,80",
    "John,20,90",
    "Jony,17,91",
    "Jony,17,93",
    "Json,21,85",
  }

Expected output (same shape, sorted):
  []string{
    "John,20,90",
    "Jony,17,91",
    "Jony,17,93",
    "Json,21,85",
    "Tom,19,80",
  }
*/

import (
	"sort"
	"strings"
)

// Ex019 returns a new slice sorted by name, age, score (string comparison on
// each field, matching the reference Python solution).
func Ex019(input []string) []string {
	out := make([]string, len(input))
	copy(out, input)

	sort.SliceStable(out, func(i, j int) bool {
		a := strings.Split(out[i], ",")
		b := strings.Split(out[j], ",")

		for k := 0; k < 3 && k < len(a) && k < len(b); k++ {
			if a[k] != b[k] {
				return a[k] < b[k]
			}
		}
		return false
	})

	return out
}
