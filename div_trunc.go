package div

// DivTrunc computes division followed by rounding toward zero for two integers accurately.
// This function panics if b == 0.
func DivTrunc[T SignedInteger](a, b T) T {
	if b == 0 {
		panic(`cannot divide by 0`)
	}
	if (a < 0) != (b < 0) {
		// For negative result
		return DivCeil[T](a, b)
	}
	// For positive result
	return DivFloor[T](a, b)
}
