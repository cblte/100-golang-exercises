package main

import (
	"sort"
	"strings"
)

// Ex010 accepts a String separated by whitespace.
// input gets sorted and duplicates are removed
func Ex010(input string) string {

	// make a map and split the input string
	m := make(map[string]bool)
	temp := strings.Split(input, " ")

	// add values to map as keys (making them unique)
	for _, v := range temp {
		m[v] = true
	}

	// convert map back into slice
	var result []string
	for key := range m {
		result = append(result, key)
	}

	// sort slice and join back toghether
	sort.Strings(result)
	output := strings.Join(result, " ")

	return output
}
