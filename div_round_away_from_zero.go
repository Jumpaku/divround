package div

// DivRoundAwayFromZero computes division followed by rounding away from zero for two integers accurately.
// This function panics if b == 0.
func DivRoundAwayFromZero(a, b int64) int64 {
	if b == 0 {
		panic(`cannot divide by 0`)
	}
	if (a < 0) != (b < 0) {
		// For negative result
		return DivFloor(a, b)
	}
	// For positive result
	return DivCeil(a, b)
}
