package creature

import (
	"github.com/gdamore/tcell/v2"
)

type Monster struct {
	Creature
}

func NewMonster(style tcell.Style, x, y int, name rune, typ int) *Monster {
	return &Monster{
		Creature{
			10001,
			style, name, typ,
			x, y,
			0, 0,
			50, 50,
			5,
		},
	}
}

func (m *Monster) Attack(x, y int) {}
