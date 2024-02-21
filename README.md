# terminal calculator

## how to use

```bash
$ go run main.go
```

## example

```console
$ go run main.go
input formula > 1 + 1 - 2 * 3 + 4 ^ 2
You entered:  1 + 1 - 2 * 3 + 4 ^ 2
RPN:  1 1 + 2 3 * - 4 2 ^ +
Result:  12

input formula > 4 ! + 5 C 2 + 4 P 2
You entered:  4 ! + 5 C 2 + 4 P 2
RPN:  4 ! 5 2 C + 4 2 P +
Result:  46

input formula > 1 + 2 * 3 / 4 - 5
You entered:  1 + 2 * 3 / 4 - 5
RPN:  1 2 3 * 4 / + 5 -
Result:  -3
```

- \+ is addition
- \- is subtraction
- \* is multiplication
- / is division
- C is combination and P is permutation
- ! is factorial
- ^ is power

/ is defined as integer division. For example, 3 / 2 = 1.
