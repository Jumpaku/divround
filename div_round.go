package div

// DivRound computes division followed by rounding away from zero for two integers accurately.
// This function panics if b == 0.
func DivRound(a, b int64) int64 {
	if b == 0 {
		panic(`cannot divide by 0`)
	}
	if (a < 0) != (b < 0) {
		// For negative result
		return DivRoundHalfDown(a, b)
	}
	// For positive result
	return DivRoundHalfUp(a, b)
}
