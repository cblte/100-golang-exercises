package main

import "math"

const c = 50
const h = 30

// Ex006 calculates Q based on C=50 and H=50
// Formula used: Q = Square root of [(2 * C * D)/H]
func Ex006(d []int) []int {
	var s = make([]int, len(d))

	for i, num := range d {
		a := (2 * c * num) / h
		q := math.Sqrt(float64(a))

		var b int = int(math.Round(q))

		s[i] = b
	}

	return s
}
