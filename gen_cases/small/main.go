package main

import (
	"fmt"
	"math"
)

func main() {
	for _, a := range []int64{-2, -1, 0, 1, 2} {
		for _, b := range []int64{-1234567890123456789, 1234567890123456789} {
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
}
