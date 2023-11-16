package treasure

import "necromancer/global"

func NewPack(x, y int) Treasure {
	rare := (global.Roll(1, 10) + 1) / 2
	return Treasure{
		Id:    1,
		Style: global.RareStyle[rare],
		Name:  global.AsciiPack,
		Type:  global.TreasureTypePack,
		Rare:  uint8(rare),
		X:     x,
		Y:     y,
		Val:   0,
	}
}
