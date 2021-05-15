package main

import (
	"fmt"
	"unicode"
)

// Ex013 run over two numbers
func Ex013(input string) string {

	var digits int = 0
	var letters int = 0

	for _, char := range input {
		if unicode.IsDigit(char) {
			digits++
		} else if unicode.IsLetter(char) {
			letters++
		} else {
			continue
		}
	}
	return fmt.Sprintf("LETTERS %v, DIGITS %v", letters, digits)
}
