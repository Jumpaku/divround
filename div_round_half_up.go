package div

// DivRoundHalfUp computes division followed by rounding for two integers accurately.
// This function panics if b == 0.
func DivRoundHalfUp(a, b int64) int64 {
	if b == 0 {
		panic(`cannot divide by 0`)
	}
	// floor(a/b + 1/2) == floor((2*a + b)/(2*b))
	return DivFloor(a*2+b, 2*b)
}
