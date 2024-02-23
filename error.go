package divround

import (
	"fmt"
)

var (
	ErrOverflow     error = overflowError{}
	ErrZeroDivision error = zeroDivisionError{}
)

func NewOverflowError(message string, numerator, denominator int64) error {
	return overflowError{message: message, numerator: numerator, denominator: denominator}
}
func NewZeroDivisionError(message string, numerator, denominator int64) error {
	return zeroDivisionError{message: message, numerator: numerator, denominator: denominator}
}

type overflowError struct {
	numerator   int64
	denominator int64
	message     string
}

func (err overflowError) Error() string {
	return fmt.Sprintf(`%s: overflow occurred at %d / %d`, err.message, err.numerator, err.denominator)
}

func (err overflowError) Is(target error) bool {
	_, ok := target.(overflowError)
	return ok
}

type zeroDivisionError struct {
	message     string
	numerator   int64
	denominator int64
}

func (err zeroDivisionError) Error() string {
	return fmt.Sprintf(`%s: division by zero occurred at %d / %d`, err.message, err.numerator, err.denominator)
}
func (err zeroDivisionError) Is(target error) bool {
	_, ok := target.(zeroDivisionError)
	return ok
}
