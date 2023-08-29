package div

import "math"

// DivFloor computes division followed by flooring for two integers accurately.
// This function panics if b == 0.
func DivFloor(a, b int64) int64 {
	if b == 0 {
		panic(`cannot divide by 0`)
	}
	if a == math.MinInt64 && b == -1 {
		panic("overflow")
	}
	if a == 0 {
		return 0
	}
	switch {
	case a < 0 && b < 0:
		if a > b {
			return 0
		}
		return DivFloor(-(a-b), -b) + 1
	case a < 0:
		if a+b > 0 {
			return -1
		}
		return -DivCeil(-(a+b), b) - 1
	case b < 0:
		if a+b < 0 {
			return -1
		}
		return -DivCeil(a, -b)
	default:
		return a / b
	}
}
