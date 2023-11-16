package global

import (
	"math/rand"
)

func Roll(n, d int) int {
	val := 0
	for i := 1; i <= n; i++ {
		x := rand.Intn(d) + 1
		val += x
	}
	return val
}

func AttackRoll() int {
	return rand.Intn(20) + 1
}
