package div

// DivRound computes division followed by rounding away from zero for two integers accurately.
// This function panics if b == 0.
func DivRoundHalfAwayFromZero[T SignedInteger](a, b T) T {
	if b == 0 {
		panic(`cannot divide by 0`)
	}
	switch {
	case a < 0 && b < 0:
		return DivRoundHalfUp(-a, -b)
	case a < 0:
		return -DivRoundHalfUp(-a, b)
	case b < 0:
		return -DivRoundHalfUp(a, -b)
	default:
		return DivRoundHalfUp(a, b)
	}
}
