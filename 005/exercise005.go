package main

/*
Define a module which has at least two methods:

- ReadString: to get a string from console input
- PrintString: to print the string in upper case.

Also please include simple test function to test the class methods.
*/

import (
	"bufio"
	"os"
	"strings"
)

var input string = ""

// ReadString reads user input from command line, terminated via 'enter'
func ReadString() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
}

// PrintString prints the string entered via GetString
func PrintString() string {
	return strings.ToUpper(input)
}
