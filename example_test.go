package divround_test

import (
	"errors"
	"fmt"
	"log"

	. "github.com/Jumpaku/divround"
)

func ExampleDivRound() {
	var (
		attack        int64 = 74
		effectPercent int64 = 25
		defense       int64 = 80
	)

	// damage = attack * (1 + effectRate / 100) - defense
	damage, err := DivRound(attack*(100+effectPercent)-100*defense, 100)
	if err != nil {
		if errors.Is(err, ErrOverflow) {
			log.Panicf(`overflow occurred: %v`, err)
		}
		if errors.Is(err, ErrZeroDivision) {
			log.Panicf(`zero division occurred: %v`, err)
		}
	}

	fmt.Printf("Damage: %d\n", damage)
	// Output: Damage: 13
}
