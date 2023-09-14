package divround

import "math"

// DivRound computes division followed by rounding away from zero for two integers accurately.
func DivRound(a, b int64) (int64, error) {
	if b == 0 {
		return 0, NewZeroDivisionError("DivRound", a, b)
	}
	if a == math.MinInt64 && b == -1 {
		return 0, NewOverflowError("DivRound", a, b)
	}
	if a == 0 {
		return 0, nil
	}
	if (a < 0) != (b < 0) {
		// For negative result
		return DivRoundHalfDown(a, b)
	}
	// For positive result
	return DivRoundHalfUp(a, b)
}
