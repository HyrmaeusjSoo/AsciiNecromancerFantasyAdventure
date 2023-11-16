package treasure

import "necromancer/global"

func NewPotion(x, y int) Treasure {
	rare := (global.Roll(1, 10) + 1) / 2
	return Treasure{
		Id:    1,
		Style: global.RareStyle[rare],
		Name:  global.AsciiPotion,
		Type:  global.TreasureTypePotion,
		Rare:  uint8(rare),
		X:     x,
		Y:     y,
		Val:   global.PotionAttr[rare],
	}
}
