package main

/*
Exercise 018:

Validate passwords. The criteria are:
  1. At least one letter [a-z]
  2. At least one number [0-9]
  3. At least one letter [A-Z]
  4. At least one character from [$#@]
  5. Length between 6 and 12 (inclusive)
  6. No whitespace

The program receives a comma-separated list of candidate passwords and
returns the valid ones in the same order, comma-separated.

Example:
  Ex018("ABd1234@1,a F1#,2w3E*,2We3345") -> "ABd1234@1"
*/

import (
	"strings"
	"unicode"
)

// Ex018 returns the valid passwords from the comma-separated input.
func Ex018(input string) string {
	candidates := strings.Split(input, ",")
	valid := []string{}

	for _, p := range candidates {
		if isValidPassword018(p) {
			valid = append(valid, p)
		}
	}

	return strings.Join(valid, ",")
}

func isValidPassword018(p string) bool {
	if len(p) < 6 || len(p) > 12 {
		return false
	}

	var hasLower, hasUpper, hasDigit, hasSpecial bool

	for _, r := range p {
		switch {
		case unicode.IsSpace(r):
			return false
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsDigit(r):
			hasDigit = true
		case r == '$' || r == '#' || r == '@':
			hasSpecial = true
		}
	}

	return hasLower && hasUpper && hasDigit && hasSpecial
}
