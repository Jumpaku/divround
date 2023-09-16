package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/Jumpaku/divround"
)

func main() {
	type Example struct {
		a int64
		b int64
	}
	examples := []Example{
		{55, 10},
		{41, 10},
		{40, 10},
		{39, 10},
		{25, 10},
		{-25, 10},
		{-39, 10},
		{-40, 10},
		{-41, 10},
		{-55, 10},
	}
	funcs := []func(int64, int64) (int64, error){
		divround.DivCeil,
		divround.DivFloor,
		divround.DivTrunc,
		divround.DivRoundAwayFromZero,
		divround.DivRound,
		divround.DivRoundHalfZero,
		divround.DivRoundHalfDown,
		divround.DivRoundHalfUp,
		divround.DivRoundHalfEven,
		divround.DivRoundHalfOdd,
	}

	funcNames := []string{}
	for _, f := range funcs {
		funcNames = append(funcNames, fmt.Sprintf("%17s", funcName(f))[28:])
	}
	fmt.Printf("| a  /  b  | %s |\n", strings.Join(funcNames, " | "))
	for _, example := range examples {
		results := []string{}
		for _, f := range funcs {
			r, _ := f(example.a, example.b)
			results = append(results, fmt.Sprintf("%3d", r))
		}
		fmt.Printf("| %3d / %-3d | %s |\n", example.a, example.b, strings.Join(results, " | "))
	}
}

func funcName(f func(int64, int64) (int64, error)) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
