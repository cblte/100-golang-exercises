# Exercises for Intermediate

> Carsten Brueggenolte &lt;mail@cbrueggenolte.de&gt;
> v1.0.0, March 2021

This document contains all the intermediate exercises. If you find an error or a bug or something that could be written better, please get in contact with me. Thanks.

## Table of Contents

1. [Exercise 006: Compute Q from a formula](#exercise-006-compute-q-from-a-formula)
2. [Exercise 007: 2-dimensional array of products](#exercise-007-2-dimensional-array-of-products)
3. [Exercise 008: Sort comma-separated words](#exercise-008-sort-comma-separated-words)
4. [Exercise 009: Capitalize every line](#exercise-009-capitalize-every-line)
5. [Exercise 010: Deduplicate and sort whitespace-separated words](#exercise-010-deduplicate-and-sort-whitespace-separated-words)
6. [Exercise 011: 4-digit binary numbers divisible by 5](#exercise-011-4-digit-binary-numbers-divisible-by-5)
7. [Exercise 012: Numbers between 100 and 300 whose digits are all even](#exercise-012-numbers-between-100-and-300-whose-digits-are-all-even)
8. [Exercise 013: Count letters and digits in a sentence](#exercise-013-count-letters-and-digits-in-a-sentence)
9. [Exercise 014: Count upper- and lower-case letters](#exercise-014-count-upper--and-lower-case-letters)
10. [Exercise 015: Compute a + aa + aaa + aaaa](#exercise-015-compute-a--aa--aaa--aaaa)
11. [Exercise 016: Filter odd numbers from a comma-separated list](#exercise-016-filter-odd-numbers-from-a-comma-separated-list)
12. [Exercise 017: Bank account net amount](#exercise-017-bank-account-net-amount)

---

## Exercise 006: Compute Q from a formula

Write a program that calculates and prints the value according to the given formula:

```
Q = Square root of [(2 * C * D) / H]
```

Following are the fixed values of C and H:

- `C` is 50. `H` is 30.
- `D` is the variable whose values should be input to your program in a comma-separated sequence.

**Example**

Let us assume the following comma separated input sequence is given to the program: `100,150,180`

The output of the program should be: `18,22,24`

### Hints

- Use `math.Sqrt` for the square root.
- Use `math.Round` for rounding a float operation.

### Solution

- Reference: [006/exercise006.go](006/exercise006.go)
- Practice here: [006/practice/practice.go](006/practice/practice.go)

---

## Exercise 007: 2-dimensional array of products

Write a program which takes 2 digits, `X`, `Y` as input and generates a 2-dimensional array. The element value in the i-th row and j-th column of the array should be `i*j`.

> Note: `i=0,1,..,X-1`; `j=0,1,..,Y-1`.

**Example**

Suppose the following inputs are given to the program: `3,5`

Then, the output of the program should be: `[[0, 0, 0, 0, 0], [0, 1, 2, 3, 4], [0, 2, 4, 6, 8]]`

### Hints

In case of input data being supplied to the question, it should be assumed to be a console input in a comma-separated form.

### Solution

- Reference: [007/exercise007.go](007/exercise007.go)
- Practice here: [007/practice/practice.go](007/practice/practice.go)

---

## Exercise 008: Sort comma-separated words

Write a program that accepts a comma separated sequence of words as input and prints the words in a comma-separated sequence after sorting them alphabetically.

Suppose the following input is supplied to the program:

```
without,hello,bag,world
```

Then, the output should be:

```
bag,hello,without,world
```

### Hints

In case of input data being supplied to the question, it should be assumed to be a console input.

### Solution

- Reference: [008/exercise008.go](008/exercise008.go)
- Practice here: [008/practice/practice.go](008/practice/practice.go)

---

## Exercise 009: Capitalize every line

Write a program that accepts sequence of lines as input and prints the lines after making all characters in the sentence capitalized.

Suppose the following input is supplied to the program:

```
Hello world
Practice makes perfect
```

Then, the output should be:

```
HELLO WORLD
PRACTICE MAKES PERFECT
```

### Hints

In case of input data being supplied to the question, it should be assumed to be a console input.

### Solution

- Reference: [009/exercise009.go](009/exercise009.go)
- Practice here: [009/practice/practice.go](009/practice/practice.go)

---

## Exercise 010: Deduplicate and sort whitespace-separated words

Write a program that accepts a sequence of whitespace separated words as input and prints the words after removing all duplicate words and sorting them alphanumerically.

Suppose the following input is supplied to the program:

```
hello world and practice makes perfect and hello world again
```

Then, the output should be:

```
again and hello makes perfect practice world
```

### Hints

In case of input data being supplied to the question, it should be assumed to be a console input. Use a set-like container to remove duplicated data automatically and then sort the data.

### Solution

- Reference: [010/exercise010.go](010/exercise010.go)
- Practice here: [010/practice/practice.go](010/practice/practice.go)

---

## Exercise 011: 4-digit binary numbers divisible by 5

Write a program which accepts a sequence of comma separated 4 digit binary numbers as its input and then check whether they are divisible by 5 or not. The numbers that are divisible by 5 are to be printed in a comma separated sequence.

**Example**

```
0100,0011,1010,1001
```

Then the output should be:

```
1010
```

Notes: Assume the data is input by console.

### Hints

In case of input data being supplied to the question, it should be assumed to be a console input.

### Solution

- Reference: [011/exercise011.go](011/exercise011.go)
- Practice here: [011/practice/practice.go](011/practice/practice.go)

---

## Exercise 012: Numbers between 100 and 300 whose digits are all even

Write a program, which will find all such numbers between 100 and 300 (both included) such that each digit of the number is an even number. The numbers obtained should be printed in a comma-separated sequence on a single line.

### Hints

In case of input data being supplied to the question, it should be assumed to be a console input.

### Solution

- Reference: [012/exercise012.go](012/exercise012.go)
- Practice here: [012/practice/practice.go](012/practice/practice.go)

---

## Exercise 013: Count letters and digits in a sentence

Write a program that accepts a sentence and calculate the number of letters and digits.

Suppose the following input is supplied to the program:

```
hello world! 123
```

Then, the output should be:

```
LETTERS 10
DIGITS 3
```

### Hints

In case of input data being supplied to the question, it should be assumed to be a console input.

### Solution

- Reference: [013/exercise013.go](013/exercise013.go)
- Practice here: [013/practice/practice.go](013/practice/practice.go)

---

## Exercise 014: Count upper- and lower-case letters

Write a program that accepts a sentence and calculates the number of upper-case letters and lower-case letters.

Suppose the following input is supplied to the program:

```
Hello world!
```

Then, the output should be:

```
UPPER CASE 1, LOWER CASE 9
```

### Hints

Use `unicode.IsUpper` and `unicode.IsLower` while iterating over the runes of the string.

### Solution

- Reference: [014/exercise014.go](014/exercise014.go)
- Practice here: [014/practice/practice.go](014/practice/practice.go)

---

## Exercise 015: Compute a + aa + aaa + aaaa

Given a single digit `a`, compute the sum `a + aa + aaa + aaaa`.

Suppose the following input is supplied to the program: `9`

Then, the output should be: `11106` (because `9 + 99 + 999 + 9999`).

### Hints

Build the repeated string with `strconv.Itoa` and `strconv.Atoi`, or notice that the answer is simply `a * 1234`.

### Solution

- Reference: [015/exercise015.go](015/exercise015.go)
- Practice here: [015/practice/practice.go](015/practice/practice.go)

---

## Exercise 016: Filter odd numbers from a comma-separated list

Accept a comma-separated sequence of integers and return only the odd ones, in the original order, comma-separated.

Suppose the following input is supplied to the program: `1,2,3,4,5,6,7,8,9`

Then, the output should be: `1,3,5,7,9`

### Hints

Use `strings.Split`, `strconv.Atoi` and the `%` operator.

### Solution

- Reference: [016/exercise016.go](016/exercise016.go)
- Practice here: [016/practice/practice.go](016/practice/practice.go)

---

## Exercise 017: Bank account net amount

Compute the net amount of a bank account from a transaction log. Each entry is of the form `<OP> <AMOUNT>` where `OP` is `D` for deposit or `W` for withdrawal.

Suppose the following entries are supplied to the program:

```
D 300
D 300
W 200
D 100
```

Then, the output should be: `500`

### Hints

- Split each line with `strings.Fields`.
- Skip empty or malformed lines instead of crashing.

### Solution

- Reference: [017/exercise017.go](017/exercise017.go)
- Practice here: [017/practice/practice.go](017/practice/practice.go)
