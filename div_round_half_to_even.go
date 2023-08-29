package div

// DivRoundHalfToEven computes division followed by rounding for two integers accurately.
// This function panics if b == 0.
func DivRoundHalfToEven(a, b int64) int64 {
	if b == 0 {
		panic(`cannot divide by 0`)
	}
	up := DivRoundHalfUp(a, b)
	if (up & 1) != 1 {
		return up
	}
	return DivRoundHalfDown(a, b)
}
