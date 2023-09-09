# Divround: Exact Division-and-Rounding Operation in Go

**Divround** is a Go library designed to facilitate exact division-and-rounding operations for integers.
Division-and-rounding is an operation defined as `Round(numerator / denominator)`, where `numerator` and `denominator` are integers and `Round` is a function that rounds the quotient into an integer.
This operation demands careful handling to ensure precision although it seems simple.

## Features

- **Exact Division-and-Rounding Operations for Integers** 
- **Overflow Check**
- **Zero Division Check**
- **Multiple Kind of Rounding Methods**

## Usage

To install **divround**, use go get:

```sh
go get github.com/Jumpaku/divround
```

Importing **divround** enables to call functions to perform the division-and-rounding operation, an example is shown as follows:

```go
import (
    "fmt"
    "log"
    . "github.com/Jumpaku/divround"
)

func main() {
    var (
        attack        = 65
        effectPercent = 25
        defense       = 80
    )

    // damage = attack * (1 + effectRate / 100) - defense
    damage, err := DivRound(attack * (100 + effectRate) - 100 * defense, 100)
    if err != nil {
        log.Panicf(`fail to compute damage: %v`, err)
    }

    fmt.Printf(`Damage: %d`, damage)
}
```

## Background and Motivation

Arithmetic operations enclosed in integers can be transformed into the form of division-and-rounding operation if the arithmetic operations consist of addition (`+`), subtraction (`-`), multiplication (`*`), and division (`/`).
Although there are some approaches to compute the division-and-rounding operation, using floating point number type (or decimal number type) should be avoided and using rational number type and primitive integer type is recommended, as described below:

* *floating point number type*: Do not use floating point number type because it includes essential inaccuracy and is not robust for rounding operation. The inaccuracy comes from inability to precisely represent decimal places or the limited number of significant digits which causes the calculation results that may include rounding errors. It is too difficult to understand and control the inaccuracy.
* *decimal number type*: Avoid to use decimal number type. Decimal number type, which may be supported by standard libraries, can represent decimal fractions exactly, guarantee precision, and control rounding. However the result of division may include rounding error and succeeding rounding obtains wrong result.
* *rational number type*: It is one of the recommended methods because rational number type can exactly represent rational numbers. Note that internal GCD operations may impact on performance.
* *primitive integer type*: It is one of the recommended methods because integer type can compute the division-and-rounding operation exactly and is natively supported by almost languages. Note that integer arithmetic operation for large numbers may cause overflow.

Focusing to Go, to perform the division-and-rounding operation using integer type demands careful and complex implementation although rational number type is supported by the standard library..
Therefore *divround* provides the functions to perform the operation for integer with error handling and multiple kind of rounding methods.

