package div

import "math"

// DivRoundHalfDown computes division followed by rounding for two integers accurately.
// This function panics if b == 0.
func DivRoundHalfDown(a, b int64) int64 {
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
			if -a > -(b - a) {
				return 1
			}
			return 0
		}
		return DivRoundHalfDown(-(a-b), -b) + 1
	case a < 0:
		if a+b > 0 {
			if -a >= a+b {
				return -1
			}
			return 0
		}
		return -DivRoundHalfUp(-(a+b), b) - 1
	case b < 0:
		if a+b < 0 {
			if a >= -(a + b) {
				return -1
			}
			return 0
		}
		return -DivRoundHalfUp(-(a+b), b) - 1
	default:
		q, r := a/b, a%b
		if r > b-r {
			return q + 1
		}
		return q
	}
}
