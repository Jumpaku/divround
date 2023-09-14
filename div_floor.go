package divround

import "math"

// DivFloor computes division followed by flooring for two integers accurately.
func DivFloor(a, b int64) (int64, error) {
	if b == 0 {
		return 0, NewZeroDivisionError("DivFloor", a, b)
	}
	if a == math.MinInt64 && b == -1 {
		return 0, NewOverflowError("DivFloor", a, b)
	}
	if a == 0 {
		return 0, nil
	}

	switch {
	case a < 0 && b < 0:
		if a > b {
			return 0, nil
		}
		return ignoreErr(DivFloor(-(a-b), -b)) + 1, nil
	case a < 0:
		if a+b > 0 {
			return -1, nil
		}
		return -ignoreErr(DivCeil(-(a+b), b)) - 1, nil
	case b < 0:
		if a+b < 0 {
			return -1, nil
		}
		return -ignoreErr(DivCeil(a, -b)), nil
	default:
		return a / b, nil
	}
}
