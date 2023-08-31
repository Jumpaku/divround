package div

import "math"

// DivRoundHalfTowardZero computes division followed by rounding for two integers accurately.
func DivRoundHalfTowardZero(a, b int64) (int64, error) {
	if b == 0 {
		return 0, NewZeroDivisionError("DivRoundHalfTowardZero", a, b)
	}
	if a == math.MinInt64 && b == -1 {
		return 0, NewOverflowError("DivRoundHalfTowardZero", a, b)
	}
	if a == 0 {
		return 0, nil
	}
	if (a < 0) != (b < 0) {
		// For negative result
		return DivRoundHalfUp(a, b)
	}
	// For positive result
	return DivRoundHalfDown(a, b)
}
