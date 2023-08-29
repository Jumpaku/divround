package div

// DivRoundHalfDown computes division followed by rounding for two integers accurately.
// This function panics if b == 0.
func DivRoundHalfDown(a, b int64) int64 {
	if b == 0 {
		panic(`cannot divide by 0`)
	}
	// ceil(a/b - 1/2)
	return DivCeil(a*2-b, 2*b)
}
