package div

// DivCeil computes division followed by ceiling for two integers accurately.
// This function panic if b == 0.
func DivCeil[T SignedInteger](a, b T) T {
	if b == 0 {
		panic(`cannot divide by 0`)
	}
	switch {
	case a < 0 && b < 0:
		return DivCeil(-a, -b)
	case a < 0:
		return -DivFloor(-a, b)
	case b < 0:
		return -DivFloor(a, -b)
	default:
		return T((uint64(a) + uint64(b) - 1) / uint64(b))
	}
}
