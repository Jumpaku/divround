package div_test

import (
	"fmt"
	"testing"

	"github.com/Jumpaku/go-div"
)

type testcaseCompare struct {
	a, b, c, d int64
	want       div.CompareResult
}

func TestCompare_Random(t *testing.T) {
	testCases := []testcaseCompare{
		{
			3, 5, 6, -4,
			div.CompareResultGreater,
		}, // 0.600000	 -1.500000
		{
			-8, 7, 6, 4,
			div.CompareResultLess,
		}, // -1.142857	 1.500000
		{
			-6, 1, -2, 5,
			div.CompareResultLess,
		}, // -6.000000	 -0.400000
		{
			3, 3, -8, 9,
			div.CompareResultGreater,
		}, // 1.000000	 -0.888889
		{
			7, -7, -2, -4,
			div.CompareResultLess,
		}, // -1.000000	 0.500000
		{
			-7, 8, 4, 9,
			div.CompareResultLess,
		}, // -0.875000	 0.444444
		{
			-7, 2, -9, 7,
			div.CompareResultLess,
		}, // -3.500000	 -1.285714
		{
			4, -9, -6, 9,
			div.CompareResultGreater,
		}, // -0.444444	 -0.666667
		{
			-8, 7, -5, 3,
			div.CompareResultGreater,
		}, // -1.142857	 -1.666667
		{
			9, -8, 9, 3,
			div.CompareResultLess,
		}, // -1.125000	 3.000000
		{
			-8, -2, -9, 1,
			div.CompareResultGreater,
		}, // 4.000000	 -9.000000
		{
			0, 6, -8, -9,
			div.CompareResultLess,
		}, // 0.000000	 0.888889
		{
			5, 9, 4, -5,
			div.CompareResultGreater,
		}, // 0.555556	 -0.800000
		{
			-4, 8, 0, -3,
			div.CompareResultLess,
		}, // -0.500000	 -0.000000
		{
			-9, 5, 2, 7,
			div.CompareResultLess,
		}, // -1.800000	 0.285714
		{
			-5, 5, 8, -1,
			div.CompareResultGreater,
		}, // -1.000000	 -8.000000
		{
			0, 8, -2, -6,
			div.CompareResultLess,
		}, // 0.000000	 0.333333
		{
			2, -8, -5, 8,
			div.CompareResultGreater,
		}, // -0.250000	 -0.625000
		{
			2, 7, -8, 8,
			div.CompareResultGreater,
		}, // 0.285714	 -1.000000
		{
			7, -3, 6, -1,
			div.CompareResultGreater,
		}, // -2.333333	 -6.000000
		{
			-9, 2, -5, 1,
			div.CompareResultGreater,
		}, // -4.500000	 -5.000000
		{
			4, -1, 6, -7,
			div.CompareResultLess,
		}, // -4.000000	 -0.857143
		{
			-4, 5, -3, 7,
			div.CompareResultLess,
		}, // -0.800000	 -0.428571
		{
			1, 6, 8, -3,
			div.CompareResultGreater,
		}, // 0.166667	 -2.666667
		{
			-7, -7, 6, -5,
			div.CompareResultGreater,
		}, // 1.000000	 -1.200000
		{
			-4, -5, -7, 2,
			div.CompareResultGreater,
		}, // 0.800000	 -3.500000
		{
			-2, 8, -1, -3,
			div.CompareResultLess,
		}, // -0.250000	 0.333333
		{
			-5, -5, 2, -9,
			div.CompareResultGreater,
		}, // 1.000000	 -0.222222
		{
			9, 4, 6, -2,
			div.CompareResultGreater,
		}, // 2.250000	 -3.000000
		{
			4, -3, -9, 3,
			div.CompareResultGreater,
		}, // -1.333333	 -3.000000
		{
			7, 4, 8, 8,
			div.CompareResultGreater,
		}, // 1.750000	 1.000000
		{
			6, -1, 1, 9,
			div.CompareResultLess,
		}, // -6.000000	 0.111111
		{
			7, -5, 1, -8,
			div.CompareResultLess,
		}, // -1.400000	 -0.125000
		{
			-5, -3, 0, 9,
			div.CompareResultGreater,
		}, // 1.666667	 0.000000
		{
			9, 5, -3, -7,
			div.CompareResultGreater,
		}, // 1.800000	 0.428571
		{
			7, -9, 4, -5,
			div.CompareResultGreater,
		}, // -0.777778	 -0.800000
		{
			-5, 8, -8, 5,
			div.CompareResultGreater,
		}, // -0.625000	 -1.600000
		{
			-7, -9, -8, 7,
			div.CompareResultGreater,
		}, // 0.777778	 -1.142857
		{
			-1, 4, 7, -1,
			div.CompareResultGreater,
		}, // -0.250000	 -7.000000
		{
			-2, -5, 3, -3,
			div.CompareResultGreater,
		}, // 0.400000	 -1.000000
		{
			7, 8, 4, -1,
			div.CompareResultGreater,
		}, // 0.875000	 -4.000000
		{
			3, 7, -3, 6,
			div.CompareResultGreater,
		}, // 0.428571	 -0.500000
		{
			-2, -2, -8, 5,
			div.CompareResultGreater,
		}, // 1.000000	 -1.600000
		{
			9, 1, -6, 6,
			div.CompareResultGreater,
		}, // 9.000000	 -1.000000
		{
			-8, -1, 0, 9,
			div.CompareResultGreater,
		}, // 8.000000	 0.000000
		{
			-2, 9, 4, 2,
			div.CompareResultLess,
		}, // -0.222222	 2.000000
		{
			0, 9, 3, 6,
			div.CompareResultLess,
		}, // 0.000000	 0.500000
		{
			4, -7, -6, -1,
			div.CompareResultLess,
		}, // -0.571429	 6.000000
		{
			8, 6, 4, 9,
			div.CompareResultGreater,
		}, // 1.333333	 0.444444
		{
			-2, 1, -2, 6,
			div.CompareResultLess,
		}, // -2.000000	 -0.333333
		{
			1, -8, 0, -1,
			div.CompareResultLess,
		}, // -0.125000	 -0.000000
		{
			-6, -9, -2, -3,
			div.CompareResultEqual,
		}, // 0.666667	 0.666667
		{
			-4, -3, 7, -8,
			div.CompareResultGreater,
		}, // 1.333333	 -0.875000
		{
			-9, 2, 1, 4,
			div.CompareResultLess,
		}, // -4.500000	 0.250000
		{
			-9, -1, -8, -9,
			div.CompareResultGreater,
		}, // 9.000000	 0.888889
		{
			-9, 4, 5, 1,
			div.CompareResultLess,
		}, // -2.250000	 5.000000
		{
			9, -1, -5, 8,
			div.CompareResultLess,
		}, // -9.000000	 -0.625000
		{
			-6, 4, 8, 7,
			div.CompareResultLess,
		}, // -1.500000	 1.142857
		{
			9, 3, 7, 5,
			div.CompareResultGreater,
		}, // 3.000000	 1.400000
		{
			-4, -8, -3, -1,
			div.CompareResultLess,
		}, // 0.500000	 3.000000
		{
			3, -3, -3, -7,
			div.CompareResultLess,
		}, // -1.000000	 0.428571
		{
			-9, 3, 4, -8,
			div.CompareResultLess,
		}, // -3.000000	 -0.500000
		{
			3, -3, 3, 3,
			div.CompareResultLess,
		}, // -1.000000	 1.000000
		{
			0, 8, -4, 4,
			div.CompareResultGreater,
		}, // 0.000000	 -1.000000
		{
			0, -6, -8, -5,
			div.CompareResultLess,
		}, // -0.000000	 1.600000
		{
			6, 1, 0, -5,
			div.CompareResultGreater,
		}, // 6.000000	 -0.000000
		{
			-4, 7, -4, -5,
			div.CompareResultLess,
		}, // -0.571429	 0.800000
		{
			3, -9, 4, 8,
			div.CompareResultLess,
		}, // -0.333333	 0.500000
		{
			-6, -3, 5, 2,
			div.CompareResultLess,
		}, // 2.000000	 2.500000
		{
			0, 6, 8, -3,
			div.CompareResultGreater,
		}, // 0.000000	 -2.666667
		{
			9, 2, -3, 3,
			div.CompareResultGreater,
		}, // 4.500000	 -1.000000
		{
			-6, -2, 1, -6,
			div.CompareResultGreater,
		}, // 3.000000	 -0.166667
		{
			-6, -1, -7, 8,
			div.CompareResultGreater,
		}, // 6.000000	 -0.875000
		{
			-1, -2, 0, 7,
			div.CompareResultGreater,
		}, // 0.500000	 0.000000
		{
			3, -4, -6, 3,
			div.CompareResultGreater,
		}, // -0.750000	 -2.000000
		{
			-5, 4, -2, 1,
			div.CompareResultGreater,
		}, // -1.250000	 -2.000000
		{
			0, -9, -6, 6,
			div.CompareResultGreater,
		}, // -0.000000	 -1.000000
		{
			-1, 2, -7, 1,
			div.CompareResultGreater,
		}, // -0.500000	 -7.000000
		{
			0, -8, -4, 3,
			div.CompareResultGreater,
		}, // -0.000000	 -1.333333

	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(`%-21s(%3d/%3d)<=>(%3d/%3d)`, tc.want, tc.a, tc.b, tc.c, tc.d), func(t *testing.T) {
			got := div.Compare(tc.a, tc.b, tc.c, tc.d)
			if got != tc.want {
				t.Errorf("Compare(%d, %d, %d, %d) failed. Want: %v, Got: %v", tc.a, tc.b, tc.c, tc.d, tc.want, got)
			}
		})
	}
}

func TestCompare_Near(t *testing.T) {
	testCases := []testcaseCompare{
		{
			-9999999999, -10000000000, -9999999999, -10000000000,
			div.CompareResultEqual,
		},
		{
			-9999999999, -10000000000, -10000000000, -9999999999,
			div.CompareResultLess,
		},
		{
			-9999999999, -10000000000, 9999999999, 10000000000,
			div.CompareResultEqual,
		},
		{
			-9999999999, -10000000000, 10000000000, 9999999999,
			div.CompareResultLess,
		},
		{
			-9999999999, 9999999999, -9999999999, 9999999999,
			div.CompareResultEqual,
		},
		{
			-9999999999, 9999999999, -9999999999, 10000000000,
			div.CompareResultLess,
		},
		{
			-9999999999, 9999999999, -10000000000, 9999999999,
			div.CompareResultGreater,
		},
		{
			-9999999999, 9999999999, -10000000000, 10000000000,
			div.CompareResultEqual,
		},
		{
			-9999999999, 9999999999, 9999999999, -9999999999,
			div.CompareResultEqual,
		},
		{
			-9999999999, 9999999999, 9999999999, -10000000000,
			div.CompareResultLess,
		},
		{
			-9999999999, 9999999999, 10000000000, -9999999999,
			div.CompareResultGreater,
		},
		{
			-9999999999, 9999999999, 10000000000, -10000000000,
			div.CompareResultEqual,
		},
		{
			-9999999999, 10000000000, -9999999999, 9999999999,
			div.CompareResultGreater,
		},
		{
			-9999999999, 10000000000, -9999999999, 10000000000,
			div.CompareResultEqual,
		},
		{
			-9999999999, 10000000000, -10000000000, 9999999999,
			div.CompareResultGreater,
		},
		{
			-9999999999, 10000000000, -10000000000, 10000000000,
			div.CompareResultGreater,
		},
		{
			-9999999999, 10000000000, 9999999999, -9999999999,
			div.CompareResultGreater,
		},
		{
			-9999999999, 10000000000, 9999999999, -10000000000,
			div.CompareResultEqual,
		},
		{
			-9999999999, 10000000000, 10000000000, -9999999999,
			div.CompareResultGreater,
		},
		{
			-9999999999, 10000000000, 10000000000, -10000000000,
			div.CompareResultGreater,
		},
		{
			-10000000000, -9999999999, -9999999999, -10000000000,
			div.CompareResultGreater,
		},
		{
			-10000000000, -9999999999, -10000000000, -9999999999,
			div.CompareResultEqual,
		},
		{
			-10000000000, -9999999999, 9999999999, 10000000000,
			div.CompareResultGreater,
		},
		{
			-10000000000, -9999999999, 10000000000, 9999999999,
			div.CompareResultEqual,
		},
		{
			-10000000000, 9999999999, -9999999999, 9999999999,
			div.CompareResultLess,
		},
		{
			-10000000000, 9999999999, -9999999999, 10000000000,
			div.CompareResultLess,
		},
		{
			-10000000000, 9999999999, -10000000000, 9999999999,
			div.CompareResultEqual,
		},
		{
			-10000000000, 9999999999, -10000000000, 10000000000,
			div.CompareResultLess,
		},
		{
			-10000000000, 9999999999, 9999999999, -9999999999,
			div.CompareResultLess,
		},
		{
			-10000000000, 9999999999, 9999999999, -10000000000,
			div.CompareResultLess,
		},
		{
			-10000000000, 9999999999, 10000000000, -9999999999,
			div.CompareResultEqual,
		},
		{
			-10000000000, 9999999999, 10000000000, -10000000000,
			div.CompareResultLess,
		},
		{
			-10000000000, 10000000000, -9999999999, 9999999999,
			div.CompareResultEqual,
		},
		{
			-10000000000, 10000000000, -9999999999, 10000000000,
			div.CompareResultLess,
		},
		{
			-10000000000, 10000000000, -10000000000, 9999999999,
			div.CompareResultGreater,
		},
		{
			-10000000000, 10000000000, -10000000000, 10000000000,
			div.CompareResultEqual,
		},
		{
			-10000000000, 10000000000, 9999999999, -9999999999,
			div.CompareResultEqual,
		},
		{
			-10000000000, 10000000000, 9999999999, -10000000000,
			div.CompareResultLess,
		},
		{
			-10000000000, 10000000000, 10000000000, -9999999999,
			div.CompareResultGreater,
		},
		{
			-10000000000, 10000000000, 10000000000, -10000000000,
			div.CompareResultEqual,
		},
		{
			9999999999, -9999999999, -9999999999, 9999999999,
			div.CompareResultEqual,
		},
		{
			9999999999, -9999999999, -9999999999, 10000000000,
			div.CompareResultLess,
		},
		{
			9999999999, -9999999999, -10000000000, 9999999999,
			div.CompareResultGreater,
		},
		{
			9999999999, -9999999999, -10000000000, 10000000000,
			div.CompareResultEqual,
		},
		{
			9999999999, -9999999999, 9999999999, -9999999999,
			div.CompareResultEqual,
		},
		{
			9999999999, -9999999999, 9999999999, -10000000000,
			div.CompareResultLess,
		},
		{
			9999999999, -9999999999, 10000000000, -9999999999,
			div.CompareResultGreater,
		},
		{
			9999999999, -9999999999, 10000000000, -10000000000,
			div.CompareResultEqual,
		},
		{
			9999999999, -10000000000, -9999999999, 9999999999,
			div.CompareResultGreater,
		},
		{
			9999999999, -10000000000, -9999999999, 10000000000,
			div.CompareResultEqual,
		},
		{
			9999999999, -10000000000, -10000000000, 9999999999,
			div.CompareResultGreater,
		},
		{
			9999999999, -10000000000, -10000000000, 10000000000,
			div.CompareResultGreater,
		},
		{
			9999999999, -10000000000, 9999999999, -9999999999,
			div.CompareResultGreater,
		},
		{
			9999999999, -10000000000, 9999999999, -10000000000,
			div.CompareResultEqual,
		},
		{
			9999999999, -10000000000, 10000000000, -9999999999,
			div.CompareResultGreater,
		},
		{
			9999999999, -10000000000, 10000000000, -10000000000,
			div.CompareResultGreater,
		},
		{
			9999999999, 10000000000, -9999999999, -10000000000,
			div.CompareResultEqual,
		},
		{
			9999999999, 10000000000, -10000000000, -9999999999,
			div.CompareResultLess,
		},
		{
			9999999999, 10000000000, 9999999999, 10000000000,
			div.CompareResultEqual,
		},
		{
			9999999999, 10000000000, 10000000000, 9999999999,
			div.CompareResultLess,
		},
		{
			10000000000, -9999999999, -9999999999, 9999999999,
			div.CompareResultLess,
		},
		{
			10000000000, -9999999999, -9999999999, 10000000000,
			div.CompareResultLess,
		},
		{
			10000000000, -9999999999, -10000000000, 9999999999,
			div.CompareResultEqual,
		},
		{
			10000000000, -9999999999, -10000000000, 10000000000,
			div.CompareResultLess,
		},
		{
			10000000000, -9999999999, 9999999999, -9999999999,
			div.CompareResultLess,
		},
		{
			10000000000, -9999999999, 9999999999, -10000000000,
			div.CompareResultLess,
		},
		{
			10000000000, -9999999999, 10000000000, -9999999999,
			div.CompareResultEqual,
		},
		{
			10000000000, -9999999999, 10000000000, -10000000000,
			div.CompareResultLess,
		},
		{
			10000000000, -10000000000, -9999999999, 9999999999,
			div.CompareResultEqual,
		},
		{
			10000000000, -10000000000, -9999999999, 10000000000,
			div.CompareResultLess,
		},
		{
			10000000000, -10000000000, -10000000000, 9999999999,
			div.CompareResultGreater,
		},
		{
			10000000000, -10000000000, -10000000000, 10000000000,
			div.CompareResultEqual,
		},
		{
			10000000000, -10000000000, 9999999999, -9999999999,
			div.CompareResultEqual,
		},
		{
			10000000000, -10000000000, 9999999999, -10000000000,
			div.CompareResultLess,
		},
		{
			10000000000, -10000000000, 10000000000, -9999999999,
			div.CompareResultGreater,
		},
		{
			10000000000, -10000000000, 10000000000, -10000000000,
			div.CompareResultEqual,
		},
		{
			10000000000, 9999999999, -9999999999, -10000000000,
			div.CompareResultGreater,
		},
		{
			10000000000, 9999999999, -10000000000, -9999999999,
			div.CompareResultEqual,
		},
		{
			10000000000, 9999999999, 9999999999, 10000000000,
			div.CompareResultGreater,
		},
		{
			10000000000, 9999999999, 10000000000, 9999999999,
			div.CompareResultEqual,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(`%-21s(%12d/%12d)<=>(%12d/%12d)`, tc.want, tc.a, tc.b, tc.c, tc.d), func(t *testing.T) {
			got := div.Compare(tc.a, tc.b, tc.c, tc.d)
			if got != tc.want {
				t.Errorf("Compare(%d, %d, %d, %d) failed. Want: %v, Got: %v", tc.a, tc.b, tc.c, tc.d, tc.want, got)
			}
		})
	}
}

func TestCompare_Small(t *testing.T) {
	testCases := []testcaseCompare{
		{
			1, -10000000000, 1, -10000000000,
			div.CompareResultEqual,
		},
		{
			1, -10000000000, 0, -10000000000,
			div.CompareResultLess,
		},
		{
			1, -10000000000, -1, 10000000000,
			div.CompareResultEqual,
		},
		{
			1, 10000000000, 1, 10000000000,
			div.CompareResultEqual,
		},
		{
			1, 10000000000, 0, 10000000000,
			div.CompareResultGreater,
		},
		{
			1, 10000000000, -1, -10000000000,
			div.CompareResultEqual,
		},
		{
			0, -10000000000, 1, -10000000000,
			div.CompareResultGreater,
		},
		{
			0, -10000000000, 0, -10000000000,
			div.CompareResultEqual,
		},
		{
			0, -10000000000, -1, 10000000000,
			div.CompareResultGreater,
		},
		{
			0, 10000000000, 1, 10000000000,
			div.CompareResultLess,
		},
		{
			0, 10000000000, 0, 10000000000,
			div.CompareResultEqual,
		},
		{
			0, 10000000000, -1, -10000000000,
			div.CompareResultLess,
		},
		{
			-1, -10000000000, 1, 10000000000,
			div.CompareResultEqual,
		},
		{
			-1, -10000000000, 0, 10000000000,
			div.CompareResultGreater,
		},
		{
			-1, -10000000000, -1, -10000000000,
			div.CompareResultEqual,
		},
		{
			-1, 10000000000, 1, -10000000000,
			div.CompareResultEqual,
		},
		{
			-1, 10000000000, 0, -10000000000,
			div.CompareResultLess,
		},
		{
			-1, 10000000000, -1, 10000000000,
			div.CompareResultEqual,
		},
		{
			0, 10000000000, 0, -10000000000,
			div.CompareResultEqual,
		},
		{
			0, -10000000000, 0, 10000000000,
			div.CompareResultEqual,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(`%-21s(%3d/%12d)<=>(%3d/%12d)`, tc.want, tc.a, tc.b, tc.c, tc.d), func(t *testing.T) {
			got := div.Compare(tc.a, tc.b, tc.c, tc.d)
			if got != tc.want {
				t.Errorf("Compare(%d, %d, %d, %d) failed. Want: %v, Got: %v", tc.a, tc.b, tc.c, tc.d, tc.want, got)
			}
		})
	}
}
