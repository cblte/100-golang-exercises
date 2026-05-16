package practice

/*
Exercise 020:

Iterate over the numbers divisible by 7 between 0 and n (exclusive).
Implement two flavours:

  - Ex020(n) returns them as a slice (easy to test).
  - Ex020Stream(n) yields them over a channel (Go's closest equivalent to a
    Python generator using `yield`).

Example:
  Ex020(20) -> [0, 7, 14]

Tip: run `go test ./...` from this folder.
*/

// Ex020 should return all non-negative integers < n that are divisible by 7.
func Ex020(n int) []int {
	// TODO: write your solution here.
	return nil
}

// Ex020Stream should send the same numbers as Ex020 over a channel, then close it.
func Ex020Stream(n int) <-chan int {
	ch := make(chan int)
	// TODO: send divisible-by-7 numbers from a goroutine and close(ch).
	close(ch)
	return ch
}
