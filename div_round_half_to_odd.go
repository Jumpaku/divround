package div

// DivRoundHalfToOdd computes division followed by rounding for two integers accurately.
// This function panics if b == 0.
func DivRoundHalfToOdd[T SignedInteger](a, b T) T {
	if b == 0 {
		panic(`cannot divide by 0`)
	}
	up := DivRoundHalfUp[T](a, b)
	if (up & 1) == 1 {
		return up
	}
	return DivRoundHalfDown[T](a, b)
}
