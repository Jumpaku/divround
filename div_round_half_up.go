package divround

import "math"

// DivRoundHalfUp computes division followed by rounding for two integers accurately.
func DivRoundHalfUp(a, b int64) (int64, error) {
	if b == 0 {
		return 0, NewZeroDivisionError("DivRoundHalfUp", a, b)
	}
	if a == math.MinInt64 && b == -1 {
		return 0, NewOverflowError("DivRoundHalfUp", a, b)
	}
	if a == 0 {
		return 0, nil
	}
	switch {
	case a < 0 && b < 0:
		if a > b {
			if -a >= -(b - a) {
				return 1, nil
			}
			return 0, nil
		}
		return ignoreErr(DivRoundHalfUp(-(a-b), -b)) + 1, nil
	case a < 0:
		if a+b > 0 {
			if -a <= a+b {
				return 0, nil
			}
			return -1, nil
		}
		return -ignoreErr(DivRoundHalfDown(-(a+b), b)) - 1, nil
	case b < 0:
		if a+b < 0 {
			if a <= -(a + b) {
				return 0, nil
			}
			return -1, nil
		}
		return -ignoreErr(DivRoundHalfDown(-(a+b), b)) - 1, nil
	default:
		q, r := a/b, a%b
		if r >= b-r {
			return q + 1, nil
		}
		return q, nil
	}
}
