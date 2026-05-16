package main

/*
Exercise 020:

Define a "generator" that iterates over the numbers divisible by 7 between
0 and n (exclusive). Python uses `yield`; in Go we expose two equivalent
flavours:

  - Ex020:        returns the full slice (easy to test).
  - Ex020Stream:  returns a receive-only channel (the closest analogue to a
                  Python generator), so you can range over it lazily.

Example:
  Ex020(20) -> [0, 7, 14]
*/

// Ex020 returns all non-negative numbers strictly less than n that are
// divisible by 7.
func Ex020(n int) []int {
	if n <= 0 {
		return nil
	}

	out := []int{}
	for i := 0; i < n; i++ {
		if i%7 == 0 {
			out = append(out, i)
		}
	}
	return out
}

// Ex020Stream yields the same numbers as Ex020 over a channel.
func Ex020Stream(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			if i%7 == 0 {
				ch <- i
			}
		}
	}()
	return ch
}
