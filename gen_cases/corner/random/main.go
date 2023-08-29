package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/Jumpaku/go-div"
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
		if div.DivFloor(a, b) != int64(math.Floor(float64(a)/float64(b))) {
			found = true
		}
		if div.DivCeil(a, b) != int64(math.Ceil(float64(a)/float64(b))) {
			found = true
		}
		if div.DivRound(a, b) != int64(math.Round(float64(a)/float64(b))) {
			found = true
		}
		if div.DivRoundHalfToEven(a, b) != int64(math.RoundToEven(float64(a)/float64(b))) {
			found = true
		}
		if found {
			f := float64(a) / float64(b)
			fmt.Printf("{\n%d,\t%d\n//%f\n},\n", a, b, f)
		}
	}
}
