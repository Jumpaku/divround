# Divround: Exact Division-and-Rounding Operation in Go

**Divround** is a Go library designed to facilitate exact division-and-rounding operations for integers.
Division-and-rounding is an operation defined as `Round(numerator/denominator)`, where the `numerator` and `denominator` are integers and `Round` is a function that rounds the quotient into an integer.
This operation demands careful handling to ensure precision although it seems simple.

## Features

- Exact Division-and-Rounding Operations for Integers
- Overflow Check
- Zero Division Check
- Multiple Kinds of Rounding Methods

## Usage

To install **divround**, use go get:

```sh
go get github.com/Jumpaku/divround
```

Importing **divround** allows to call functions to perform the division-and-rounding operation as shown in the example below:

```go
package main

import (
	"errors"
	"fmt"
	"log"

	. "github.com/Jumpaku/divround"
)

func main() {
	var (
		attack        int64 = 74
		effectPercent int64 = 25
		defense       int64 = 80
	)

	// damage = attack * (1 + effectPercent / 100) - defense
	damage, err := DivRound(attack*(100+effectPercent)-100*defense, 100)
	if err != nil {
		if errors.Is(err, ErrOverflow) {
			log.Panicf(`overflow occurred: %v`, err)
		}
		if errors.Is(err, ErrZeroDivision) {
			log.Panicf(`zero division occurred: %v`, err)
		}
	}

	fmt.Printf("Damage: %d\n", damage) // Damage: 13
}
```

## Background and Motivation

Arithmetic operations enclosed in integers can be transformed into the form of the division-and-rounding operation if the arithmetic operations consist of addition (`+`), subtraction (`-`), multiplication (`*`), and division (`/`).
Although there are some approaches to compute the division-and-rounding operation, using floating point number type (or decimal number type) should be avoided, and using rational number type and primitive integer type is recommended, as described below:

* *floating point number type*: Do not use the floating point number type because it includes essential inaccuracy and is not robust for rounding operations. The inaccuracy comes from the inability to precisely represent decimal places or the limited number of significant digits which causes the calculation results that may include rounding errors. It is too difficult to understand and control the inaccuracy.
* *decimal number type*: Avoid using decimal number type. Decimal number type, which may be supported by standard libraries, can represent decimal fractions exactly, guarantee precision, and control rounding. However, the result of division may include rounding errors, and succeeding rounding may obtain a wrong result.
* *rational number type*: It is one of the recommended methods because rational number type can exactly represent rational numbers. Note that internal GCD operations may impact upon performance.
* *primitive integer type*: It is one of the recommended methods because integer type can compute the division-and-rounding operation exactly and is natively supported by almost all languages. Note that integer arithmetic operations for large numbers may cause overflow.

Focusing on Go, performing the division-and-rounding operation using integer type demands careful and complex implementation although rational number type is supported by the standard library..
Therefore **divround** provides the functions to perform the operation for integers with error handling and multiple kinds of rounding methods.

## Detail Description

**Divround** provides the division-and-rounding operations as functions of type `func (int64, int64) (int64, error)`, which are named Div< _RoundingMethod_ >.

### Behaviors of Supported Rounding Methods

| a  /  b  | DivCeil | DivFloor | DivTrunc | DivRoundAwayFromZero | DivRound | DivRoundHalfZero | DivRoundHalfDown | DivRoundHalfUp | DivRoundHalfEven | DivRoundHalfOdd |
|:---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
|  55 / 10  |   6 |   5 |   5 |   6 |   6 |   5 |   5 |   6 |   6 |   5 |
|  41 / 10  |   5 |   4 |   4 |   5 |   4 |   4 |   4 |   4 |   4 |   4 |
|  40 / 10  |   4 |   4 |   4 |   4 |   4 |   4 |   4 |   4 |   4 |   4 |
|  39 / 10  |   4 |   3 |   3 |   4 |   4 |   4 |   4 |   4 |   4 |   4 |
|  25 / 10  |   3 |   2 |   2 |   3 |   3 |   2 |   2 |   3 |   2 |   3 |
| -25 / 10  |  -2 |  -3 |  -2 |  -3 |  -3 |  -2 |  -3 |  -2 |  -2 |  -3 |
| -39 / 10  |  -3 |  -4 |  -3 |  -4 |  -4 |  -4 |  -4 |  -4 |  -4 |  -4 |
| -40 / 10  |  -4 |  -4 |  -4 |  -4 |  -4 |  -4 |  -4 |  -4 |  -4 |  -4 |
| -41 / 10  |  -4 |  -5 |  -4 |  -5 |  -4 |  -4 |  -4 |  -4 |  -4 |  -4 |
| -55 / 10  |  -5 |  -6 |  -5 |  -6 |  -6 |  -5 |  -6 |  -5 |  -6 |  -5 |

### Errors

The functions provided by **divround** return the following errors:

* ErrOverflow: when `math.MinInt64 (-9223372036854775808)` is divided by `-1`.
* ErrDivisionByZero: when an integer is divided by `0`.
