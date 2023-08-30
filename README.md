# go-div
Go library for precise integer division and rounding.

`go-div` is a Go library that focuses on providing precise integer division and rounding operations. It's designed to handle scenarios where precision is crucial and the limitations of floating-point arithmetic are a concern.


## Features

- **Precise Integer Division and Rounding Operations** 
- **Overflow Check**
- **Multiple Kind of Rounding Operations**


## Motivation

In many applications, precision is paramount. Floating-point numbers, while versatile, come with inherent limitations that can lead to unexpected behavior. `go-div` provides an alternative approach for situations where precision is non-negotiable.

In many applications, the requirement for precision is absolute. Floating-point numbers, while versatile, come with inherent limitations that can give rise to unexpected inaccuracies in crucial calculations. You should avoid using floating-point numbers if precision is required because they can result in cumulative rounding errors, loss of significant digits, and inexact comparisons.
