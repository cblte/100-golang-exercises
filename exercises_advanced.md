# Exercises for Advanced

> Carsten Brueggenolte &lt;mail@cbrueggenolte.de&gt;
> v1.0.0, March 2021

This document contains all the advanced exercises. If you find an error or a bug or something that could be written better, please get in contact with me. Thanks.

## Table of Contents

1. [Exercise 018: Validate passwords](#exercise-018-validate-passwords)
2. [Exercise 019: Sort (name, age, score) records](#exercise-019-sort-name-age-score-records)
3. [Exercise 020: Generator for numbers divisible by 7](#exercise-020-generator-for-numbers-divisible-by-7)

---

## Exercise 018: Validate passwords

A website requires users to register with a username and password. Write a program to check the validity of each password against the following criteria:

1. At least one letter `[a-z]`
2. At least one letter `[A-Z]`
3. At least one digit `[0-9]`
4. At least one character from `[$#@]`
5. Length between 6 and 12 (inclusive)
6. No whitespace

The program receives a comma-separated list of candidate passwords and returns the valid ones in the same order, comma-separated.

**Example**

Input:

```
ABd1234@1,a F1#,2w3E*,2We3345
```

Output:

```
ABd1234@1
```

### Hints

- The `unicode` package has helpers like `IsLower`, `IsUpper`, `IsDigit`, `IsSpace`.
- A single pass over the string is enough to compute every flag.

### Solution

- Reference: [018/exercise018.go](018/exercise018.go)
- Practice here: [018/practice/practice.go](018/practice/practice.go)

---

## Exercise 019: Sort (name, age, score) records

Sort a list of `name,age,score` lines in ascending order, first by name, then by age, then by score. (Compare each field as a string, matching the reference Python solution.)

**Example**

Input:

```
Tom,19,80
John,20,90
Jony,17,91
Jony,17,93
Json,21,85
```

Output:

```
John,20,90
Jony,17,91
Jony,17,93
Json,21,85
Tom,19,80
```

### Hints

- Use `sort.SliceStable` with a comparator that splits each entry with `strings.Split`.
- Walk through the fields and break on the first one that differs.

### Solution

- Reference: [019/exercise019.go](019/exercise019.go)
- Practice here: [019/practice/practice.go](019/practice/practice.go)

---

## Exercise 020: Generator for numbers divisible by 7

Define a generator that iterates over the numbers divisible by 7 between `0` and `n` (exclusive). The original Python challenge uses `yield`; in Go the closest analogue is sending values over a channel from a goroutine.

We expose two flavours:

- `Ex020(n)` returns the full slice — easy to test.
- `Ex020Stream(n)` returns a receive-only channel that yields the same values lazily and closes when done.

**Example**

```
Ex020(20)        -> [0, 7, 14]
Ex020Stream(50)  -> 0, 7, 14, 21, 28, 35, 42, 49
```

### Hints

- For `Ex020Stream`, start a goroutine that sends on the channel and `defer close(ch)` so callers can `range` over it.

### Solution

- Reference: [020/exercise020.go](020/exercise020.go)
- Practice here: [020/practice/practice.go](020/practice/practice.go)
