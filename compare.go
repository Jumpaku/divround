package div

import "fmt"

// CompareResult represents a comparison result of two numbers.
type CompareResult int

const (
	// CompareResultEqual represents x == y is true
	CompareResultEqual CompareResult = 0
	// CompareResultEqual represents x < y is true
	CompareResultLess CompareResult = -1
	// CompareResultEqual represents x > y is true
	CompareResultGreater CompareResult = 1
)

func (cmp CompareResult) String() string {
	switch cmp {
	default:
		panic(fmt.Sprintf(`invalid value for CompareResult: %v`, int(cmp)))
	case CompareResultEqual:
		return `CompareResultEqual`
	case CompareResultLess:
		return `CompareResultLess`
	case CompareResultGreater:
		return `CompareResultGreater`
	}
}

// Compare accurately compares a/b and c/d.
// This function panic when b or d is 0.
func Compare[T SignedInteger](a, b, c, d T) CompareResult {
	if b == 0 || d == 0 {
		panic(`cannot divide by 0`)
	}

	if b < 0 {
		a = -a
		b = -b
	}

	if d < 0 {
		c = -c
		d = -d
	}

	cmp := int64(a)*int64(d) - int64(c)*int64(b)

	switch {
	case cmp < 0:
		return CompareResultLess
	case cmp > 0:
		return CompareResultGreater
	default:
		return CompareResultEqual
	}
}
