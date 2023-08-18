package creature

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type Monster struct {
	Creature
}

func NewMonster(style tcell.Style, x, y int, name rune, typ int) *Monster {
	return &Monster{
		Creature{
			int(time.Now().UnixMilli()),
			style, name, typ,
			x, y,
			0, 0,
			50, 50,
			5,
		},
	}
}

func (m *Monster) Attack(x, y int) {}
