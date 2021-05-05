package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Ex012 run over two numbers
func Ex012(start, end int) string {

	values := []string{}

	// loop from start to end
	for num := int64(start); num < int64(end); num++ {
		temp := strconv.FormatInt(num, 10)

		a := int(temp[0])
		b := int(temp[1])
		c := int(temp[2])

		if a%2 == 0 && b%2 == 0 && c%2 == 0 {
			values = append(values, fmt.Sprint(num))
		}

	}
	return strings.Join(values, ",")
}
