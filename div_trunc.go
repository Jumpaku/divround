package divround

import "math"

// DivTrunc computes division followed by truncation (rounding toward zero) for two integers.
func DivTrunc(a, b int64) (int64, error) {
	if b == 0 {
		return 0, NewZeroDivisionError("DivTrunc", a, b)
	}
	if a == math.MinInt64 && b == -1 {
		return 0, NewOverflowError("DivTrunc", a, b)
	}
	if a == 0 {
		return 0, nil
	}
	if (a < 0) != (b < 0) {
		// For negative result
		return DivCeil(a, b)
	}
	// For positive result
	return DivFloor(a, b)
}
