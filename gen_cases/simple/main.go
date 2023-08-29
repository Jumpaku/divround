package main

import (
	"fmt"
	"math"
)

func main() {
	// Simple
	for a := (-10); a <= (10); a++ {
		for b := (-10); b <= (10); b++ {
			if b == 0 {
				continue
			}
			f := float64(a) / float64(b)
			fmt.Printf("{\n%d,\t%d,\n%d,\n}, // %f\n", a, b, int(math.RoundToEven(f+1)-1), f)
		}
	}
}
