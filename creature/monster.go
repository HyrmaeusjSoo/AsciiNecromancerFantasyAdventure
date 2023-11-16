package creature

import (
	"necromancer/global"
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
			style, name, typ, 1,
			x, y,
			0, 0,
			global.MstAttr[typ].MaxHealth, global.MstAttr[typ].MaxHealth,
			global.MstAttr[typ].Damage,
			0,
		},
	}
}

func (m *Monster) Attack(x, y int) {}
