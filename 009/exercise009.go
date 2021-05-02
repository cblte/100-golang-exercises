package main

import "strings"

// Ex009 convert input to upper case
func Ex009(input []string) []string {

	output := []string{}

	for _, v := range input {
		output = append(output, strings.ToUpper(v))
	}

	return output
}
