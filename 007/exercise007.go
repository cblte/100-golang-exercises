package main

// Ex007 takes two numbers and generate a slice
// Then, the output of the program should be:
// [[0, 0, 0, 0, 0], [0, 1, 2, 3, 4], [0, 2, 4, 6, 8]]
func Ex007(x, y int) [][]int {

	slice := [][]int{}

	for i := 0; i < x; i++ {

		row := []int{}
		for j := 0; j < y; j++ {
			p := i * j
			row = append(row, p)
		}
		slice = append(slice, row)
	}

	return slice
}
