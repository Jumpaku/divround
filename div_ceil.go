package divround

import "math"

// DivCeil computes division followed by ceiling (rounding toward positive infinity) for two integers.
func DivCeil(a, b int64) (int64, error) {
	if b == 0 {
		return 0, NewZeroDivisionError("DivCeil", a, b)
	}
	if a == math.MinInt64 && b == -1 {
		return 0, NewOverflowError("DivCeil", a, b)
	}
	if a == 0 {
		return 0, nil
	}

	switch {
	case a < 0 && b < 0:
		if b < a {
			return 1, nil
		}
		return ignoreErr(DivCeil(-(a-b), -b)) + 1, nil
	case a < 0:
		if a+b > 0 {
			return 0, nil
		}
		return -ignoreErr(DivFloor(-(a+b), b)) - 1, nil
	case b < 0:
		if a+b < 0 {
			return 0, nil
		}
		return -ignoreErr(DivFloor(-(a+b), b)) - 1, nil
	default:
		q, r := a/b, a%b
		if r > 0 {
			return q + 1, nil
		}
		return q, nil
	}
}
