package main

/*
Exercise 017:

Compute the net amount of a bank account from a transaction log. Each entry
is a string of the form "<OP> <AMOUNT>" where OP is:
  - "D" for deposit
  - "W" for withdrawal

Example:
  Ex017([]string{"D 300", "D 300", "W 200", "D 100"}) -> 500
*/

import (
	"strconv"
	"strings"
)

// Ex017 returns the net account balance after applying every transaction.
// Malformed or empty lines are ignored.
func Ex017(transactions []string) int {
	net := 0

	for _, t := range transactions {
		fields := strings.Fields(t)
		if len(fields) != 2 {
			continue
		}

		amount, err := strconv.Atoi(fields[1])
		if err != nil {
			continue
		}

		switch fields[0] {
		case "D":
			net += amount
		case "W":
			net -= amount
		}
	}

	return net
}
