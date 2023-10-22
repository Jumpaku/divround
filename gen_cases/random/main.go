package main

import (
	"fmt"
	"math"

	"github.com/Jumpaku/divround"
)

func main() {
	for x := int64(0); x <= int64(1000); x++ {
		for y := int64(-1000); y <= int64(1000); y++ {
			z := int64(100)
			{
				got := int64(math.Round(float64(x) * (float64(y) / float64(z))))
				want, _ := divround.DivRound(x*y, z)
				if got != want {
					fmt.Printf("Round (%+06d * (%+06d / %d)): want = %+07d, got = %07d\n", x, y, z, want, got)
				}
			}
			{
				got := int64(math.Ceil(float64(x) * (float64(y) / float64(z))))
				want, _ := divround.DivCeil(x*y, z)
				if got != want {
					fmt.Printf("Ceil  (%+06d * (%+06d / %d)): want = %+07d, got = %07d\n", x, y, z, want, got)
				}
			}
			{
				got := int64(math.Floor(float64(x) * (float64(y) / float64(z))))
				want, _ := divround.DivFloor(x*y, z)
				if got != want {
					fmt.Printf("Floor (%+06d * (%+06d / %d)): want = %+07d, got = %07d\n", x, y, z, want, got)
				}
			}
			{
				got := int64(math.Trunc(float64(x) * (float64(y) / float64(z))))
				want, _ := divround.DivTrunc(x*y, z)
				if got != want {
					fmt.Printf("Trunc (%+06d * (%+06d / %d)): want = %+07d, got = %07d\n", x, y, z, want, got)
				}
			}
		}
	}
}
