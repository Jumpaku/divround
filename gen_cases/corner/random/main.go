package main

import (
	"fmt"
	"math"
	"math/rand"

	div "github.com/Jumpaku/divround"
)

func main() {
	r := rand.New(rand.NewSource(2))
	for i := 0; i < 12345678901; i++ {
		m := int64(1000000000000000000)
		a := r.Int63n(2*m) - m
		b := r.Int63n(2*m) - m
		if b == 0 {
			continue
		}
		var found bool
		if q, _ := div.DivFloor(a, b); q != int64(math.Floor(float64(a)/float64(b))) {
			found = true
		}
		if q, _ := div.DivCeil(a, b); q != int64(math.Ceil(float64(a)/float64(b))) {
			found = true
		}
		if q, _ := div.DivRound(a, b); q != int64(math.Round(float64(a)/float64(b))) {
			found = true
		}
		if q, _ := div.DivRoundHalfEven(a, b); q != int64(math.RoundToEven(float64(a)/float64(b))) {
			found = true
		}
		if found {
			f := float64(a) / float64(b)
			fmt.Printf("{\n%d,\t%d\n//%f\n},\n", a, b, f)
		}
	}
}
