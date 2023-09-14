package main

import (
	"strconv"
	"strings"
)

// Ex011 takes comma separated string with 4 digits binary numbers
// returns a comma separated string which the numbers which can
// be divided by 5
func Ex011(input string) string {

	// split string
	numbers := strings.Split(input, ",")
	result := []string{}

	// loop over numbers
	for _, v := range numbers {
		// convert
		converted, _ := strconv.ParseInt(v, 2, 64)
		// if mod 5 == 0
		if converted%5 == 0 {
			// convert back to binary and add to result slice
			// s := strconv.FormatInt(converted, 2)
			// result = append(result, s)

			// just add v to the result
			result = append(result, v)
		}

	}
	// join output back together
	output := strings.Join(result, ",")

	return output
}
