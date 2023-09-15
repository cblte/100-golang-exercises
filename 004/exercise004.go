package main

import (
	"strconv"
	"strings"
)

// Ex004 takes a string of comma-separated numbers and returns a slice of int
func Ex004(input string) []int {
	// create a map with the size of n
	numbers := strings.Split(input, ",")

	length := len(numbers)
	var num = make([]int, length)

	for index, v := range numbers {
		s := strings.Trim(v, " ")
		num[index], _ = strconv.Atoi(s)
	}
	return num
}
