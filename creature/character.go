package creature

import (
	"necromancer/global"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Status struct {
	Exp    int
	MaxExp int
	Skill  map[uint8]int
	Coins  int
}

type Character struct {
	Creature
	Status
}

func NewCharacter(style tcell.Style, x, y int, name rune, typ int) Character {
	return Character{
		Creature{
			int(time.Now().UnixMilli()),
			style, name, typ, 1,
			x, y,
			0, 0,
			global.CptAttr[typ].MaxHealth, global.CptAttr[typ].MaxHealth,
			global.CptAttr[typ].Damage,
			0,
		},
		Status{
			0, global.CptAttr[typ].MaxExp,
			map[uint8]int{
				1: 0, 2: 0, 3: 0, 4: 0,
			},
			0,
		},
	}
}

func (c *Character) TurnRound(x, y int) {
	c.Tx, c.Ty = x, y
}

func (c *Character) Move(x, y int) {
	c.X, c.Y = x, y
}

func (c *Character) Follow(x, y int) {
	c.Creature.Move(x, y)
	// if math.Sqrt(math.Pow(float64(c.X-x), 2)+math.Pow(float64(c.Y-y), 2)) < 2 {
	// 	return
	// }
	// xdis, ydis := math.Abs(float64(c.X-x)), math.Abs(float64(c.Y-y))
	// if xdis < 3 && ydis == 0 {
	// 	return
	// } else if ydis < 3 && xdis == 0 {
	// 	return
	// }
	// if xdis > ydis {
	// 	if c.X > x {
	// 		c.X--
	// 	} else {
	// 		c.X++
	// 	}
	// } else {
	// 	if c.Y > y {
	// 		c.Y--
	// 	} else {
	// 		c.Y++
	// 	}
	// }
}

func (c *Character) AddExp(n int) bool {
	c.Exp += n
	if c.Exp >= c.MaxExp {
		c.LevelUp()
		c.Exp -= c.MaxExp
		return true
	}
	return false
}

func (c *Character) LevelUp() {
	if c.Level >= 20 {
		return
	}
	c.Level++
	c.Damage++
	c.Max += 10
	c.MaxExp += 10
}
