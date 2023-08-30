package main

import (
	"fmt"
	"math"
)

func main() {
	var max int64 = math.MaxInt64
	//var hMax int64 = 1<<62 - 1
	var min int64 = math.MinInt64
	//var hMin int64 = -1 << 62
	for _, a := range []int64{-1, 0, 1, max, min} {
		for _, b := range []int64{-1, 0, 1, max, min} {
			if int64(math.Abs(float64(a*b))) <= 1 {
				continue
			}
			if b == 0 {
				fmt.Printf("{\n%d,\t%d,\n%d,div.ErrDivisionByZero,\n}, // error\n", a, b, 0)
			} else if a == math.MinInt64 && b == -1 {
				fmt.Printf("{\n%d,\t%d,\n%d,div.ErrOverflow,\n}, // error\n", a, b, 0)
			} else {
				f := float64(a) / float64(b)
				s := 1.0
				if math.Signbit(f) {
					s = -1.0
					f = -f
				}
				fmt.Printf("{\n%d,\t%d,\n%d,nil,\n}, // %f\n", a, b, int(s*math.Round(f)), f)
			}
		}
	}
}
