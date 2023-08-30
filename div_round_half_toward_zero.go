package div

// DivRoundHalfTowardZero computes division followed by rounding for two integers accurately.
// This function panics if b == 0.
func DivRoundHalfTowardZero(a, b int64) int64 {
	if b == 0 {
		panic(`cannot divide by 0`)
	}
	if (a < 0) != (b < 0) {
		// For negative result
		return DivRoundHalfUp(a, b)
	}
	// For positive result
	return DivRoundHalfDown(a, b)
}
