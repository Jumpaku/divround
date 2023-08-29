package main

import (
	"fmt"
	"math"
)

func main() {
	for _, a := range []int64{-20000000000000001, -20000000000000000, -19999999999999999, 20000000000000001, 20000000000000000, 19999999999999999} {
		for _, b := range []int64{-20000000000000000, 20000000000000000} {
			if b == 0 {
				continue
			}
			f := float64(a) / float64(b)
			s := 1.0
			if f < 0 {
				s = -1.0
			}
			fmt.Printf("{\n%d,\t%d,\n%d,\n}, // %f\n", a, b, int(s*math.Round(math.Abs(f))), f)
		}
	}
	for _, a := range []int64{-30000000000000001, -30000000000000000, -29999999999999999, 30000000000000001, 30000000000000000, 29999999999999999} {
		for _, b := range []int64{-30000000000000000, 30000000000000000} {
			if b == 0 {
				continue
			}
			f := float64(a) / float64(b)
			s := 1.0
			if f < 0 {
				s = -1.0
			}
			fmt.Printf("{\n%d,\t%d,\n%d,\n}, // %f\n", a, b, int(s*math.Ceil(math.Abs(f))), f)
		}
	}
}
