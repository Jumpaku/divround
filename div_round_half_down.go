package div

// DivRoundHalfDown computes division followed by rounding for two integers accurately.
// This function panics if b == 0.
func DivRoundHalfDown[T SignedInteger](a, b T) T {
	if b == 0 {
		panic(`cannot divide by 0`)
	}
	// ceil(a/b - 1/2)
	return DivCeil[T](a*2-b, 2*b)
}
