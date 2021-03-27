package main

/*
Exercise 003

With a given integral number n, write a program to generate a map that contains (i, i*i) such that is an integral number between 1 and n (both included), and then the program should print the map with representation of the value

Suppose the following input is supplied to the program: 8
Then, the output should be:
map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64]

*/

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Exercise 003")
	var n int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		log.Fatal("Error occured: ", err)
	}

	fmt.Printf("%v", Ex003(n))
}

// Ex003 returns a map with numbers are their squared values
func Ex003(n int) map[int]int {
	// create a map with the size of n
	var numbers = make(map[int]int, n)

	for i := 1; i <= n; i++ {
		numbers[i] = i * i
	}
	return numbers
}
