package div

// DivFloor computes division followed by flooring for two integers accurately.
// This function panics if b == 0.
func DivFloor[T SignedInteger](a, b T) T {
	if b == 0 {
		panic(`cannot divide by 0`)
	}
	switch {
	case a < 0 && b < 0:
		return DivFloor(-a, -b)
	case a < 0:
		return -DivCeil(-a, b)
	case b < 0:
		return -DivCeil(a, -b)
	default:
		return a / b
	}
}
