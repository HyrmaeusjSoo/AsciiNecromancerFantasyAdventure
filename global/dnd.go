package global

import (
	"math/rand/v2"
)

func Roll(n, d int) int {
	dnd := 0
	for i := 1; i <= n; i++ {
		x := rand.IntN(d) + 1
		dnd += x
	}
	return dnd
}

func AttackRoll() int {
	return Roll(1, 20)
}

func LaunchAttack(cast, dice int) (hitDamage int) {
	ar := AttackRoll()
	if ar == 1 {
		return 0
	}
	if ar == 20 {
		cast *= 2
	}
	return Roll(cast, dice)
}
