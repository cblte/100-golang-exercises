package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Exercise 001")
	var numbers []string
	for i := 2000; i <= 3200; i++ {
		if i%7 == 0 && i%5 != 0 {
			numbers = append(numbers, strconv.Itoa(i))
		}
	}
	fmt.Println(strings.Join(numbers, ","))
}
