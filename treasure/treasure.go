package treasure

import (
	"necromancer/global"

	"github.com/gdamore/tcell/v2"
)

type Treasure struct {
	Id    int
	Style tcell.Style
	Name  rune
	Type  int
	Rare  uint8
	X, Y  int
	Val   int
}

func New(x, y int) Treasure {
	typ := (global.Roll(1, 6) + 1) / 2
	t := Treasure{}
	switch typ {
	case global.TreasureTypeCoin:
		t = NewCoin(x, y)
	case global.TreasureTypePotion:
		t = NewPotion(x, y)
	case global.TreasureTypePack:
		t = NewPack(x, y)
	}
	return t
}

func (t *Treasure) Draw(s tcell.Screen) {
	s.SetContent(t.X, t.Y, t.Name, nil, t.Style)
}

func (t *Treasure) Pick() {
	return
}

func (t *Treasure) Use() bool {
	return true
}
