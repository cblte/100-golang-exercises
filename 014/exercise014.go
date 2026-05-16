package main

/*
Exercise 014:

Write a program that accepts a sentence and calculate the number of upper case
letters and lower case letters.

Suppose the following input is supplied to the program: "Hello world!"
Then, the output should be: "UPPER CASE 1, LOWER CASE 9"
*/

import (
	"fmt"
	"unicode"
)

// Ex014 counts upper- and lower-case letters in input and returns them in the
// format "UPPER CASE <u>, LOWER CASE <l>".
func Ex014(input string) string {
	var upper, lower int

	for _, r := range input {
		switch {
		case unicode.IsUpper(r):
			upper++
		case unicode.IsLower(r):
			lower++
		}
	}

	return fmt.Sprintf("UPPER CASE %d, LOWER CASE %d", upper, lower)
}
