package div

// DivRoundHalfTowardZero computes division followed by rounding for two integers accurately.
// This function panics if b == 0.
func DivRoundHalfTowardZero[T SignedInteger](a, b T) T {
	if b == 0 {
		panic(`cannot divide by 0`)
	}
	switch {
	case a < 0 && b < 0:
		return DivRoundHalfDown(-a, -b)
	case a < 0:
		return -DivRoundHalfDown(-a, b)
	case b < 0:
		return -DivRoundHalfDown(a, -b)
	default:
		return DivRoundHalfDown(a, b)
	}
}
