# Exercises for Beginners

> Carsten Brueggenolte &lt;mail@cbrueggenolte.de&gt;
> v1.0.0, March 2021

This document contains all the beginner exercises. If you find an error or a bug or something that could be written better, please get in contact with me. Thanks.

## Table of Contents

1. [Exercise 001: Find numbers divisible by 7 but not by 5](#exercise-001-find-numbers-divisible-by-7-but-not-by-5)
2. [Exercise 002: Compute factorial](#exercise-002-compute-factorial)
3. [Exercise 003: Create a map with numbers squared](#exercise-003-create-a-map-with-numbers-squared)
4. [Exercise 004: Create a slice from comma-separated input string](#exercise-004-create-a-slice-from-comma-separated-input-string)
5. [Exercise 005: Define an "interface" with at least two methods](#exercise-005-define-an-interface-with-at-least-two-methods)

---

## Exercise 001: Find numbers divisible by 7 but not by 5

Write a program which will find all such numbers which are divisible by 7 but are not a multiple of 5, between 2000 and 3200 (both included). The numbers obtained should be printed in a comma-separated sequence on a single line.

### Hint

Consider using `strconv` and `strings.Join`.

### Solution

- Reference: [001/exercise001.go](001/exercise001.go)
- Practice here: [001/practice/practice.go](001/practice/practice.go)

---

## Exercise 002: Compute factorial

**Exercise Description**

Write a program which can compute the factorial of a given numbers. The results should be printed in a comma-separated sequence on a single line.

Suppose the following input is supplied to the program: `8`

Then, the output should be: `40320`

### Hints

In case of input data being supplied to the question, it should be assumed to be a console input.

### Solution

- Reference: [002/exercise002.go](002/exercise002.go)
- Practice here: [002/practice/practice.go](002/practice/practice.go)

---

## Exercise 003: Create a map with numbers squared

With a given integral number n, write a program to generate a map that contains `(i, i*i)` such that is an integral number between 1 and n (both included), and then the program should print the map with representation of the value.

Suppose the following input is supplied to the program: `8`

Then, the output should be: `map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64]`

### Hints

Use `make` for the map and `%v` verb for the output.

### Solution

- Reference: [003/exercise003.go](003/exercise003.go)
- Practice here: [003/practice/practice.go](003/practice/practice.go)

---

## Exercise 004: Create a slice from comma-separated input string

Write a program which accepts a sequence of comma-separated numbers from console and generate an slice out of them. Return the slice.

Suppose the following input is supplied to the program: `34, 67, 55, 33, 12, 98`.

Then, the output should be: `[34 67 55 33 12 98]`

### Hints

In case of input data being supplied to the question, it should be assumed to be a console input. The `strings` package has a `Split` method.

### Solution

- Reference: [004/exercise004.go](004/exercise004.go)
- Practice here: [004/practice/practice.go](004/practice/practice.go)

---

## Exercise 005: Define an "interface" with at least two methods

Create a separate file (module) which has at least two methods:

- `ReadString`: to read a string from console input.
- `PrintString`: to return the string in upper case.

Also create a `main.go` file that acts as calling class.

### Hints

- Use `bufio.NewScanner(os.Stdin)` to read in a full line of text.
- Use `go run .` to execute considering all files in the directory.

### Solution

- Reference: [005/exercise005.go](005/exercise005.go)
- Practice here: [005/practice/practice.go](005/practice/practice.go)
