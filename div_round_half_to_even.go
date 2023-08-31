package div

import "math"

// DivRoundHalfToEven computes division followed by rounding for two integers accurately.
func DivRoundHalfToEven(a, b int64) (int64, error) {
	if b == 0 {
		return 0, NewZeroDivisionError("DivRoundHalfToEven", a, b)
	}
	if a == math.MinInt64 && b == -1 {
		return 0, NewOverflowError("DivRoundHalfToEven", a, b)
	}
	if a == 0 {
		return 0, nil
	}
	up, _ := DivRoundHalfUp(a, b)
	if (up & 1) != 1 {
		return up, nil
	}
	return DivRoundHalfDown(a, b)
}
