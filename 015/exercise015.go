package main

/*
Exercise 015:

Given a single digit a (1..9), compute a + aa + aaa + aaaa.

Example:
  Ex015(9) -> 9 + 99 + 999 + 9999 = 11106
*/

import "strconv"

// Ex015 returns a + aa + aaa + aaaa where a is the input digit.
func Ex015(a int) int {
	s := strconv.Itoa(a)

	total := 0
	repeated := ""
	for i := 0; i < 4; i++ {
		repeated += s
		n, _ := strconv.Atoi(repeated)
		total += n
	}

	return total
}
