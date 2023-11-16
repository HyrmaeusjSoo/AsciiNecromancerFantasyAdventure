package treasure

import "necromancer/global"

func NewCoin(x, y int) Treasure {
	rare := (global.Roll(1, 10) + 1) / 2
	return Treasure{
		Id:    1,
		Style: global.RareStyle[rare],
		Name:  global.AsciiCoin,
		Type:  global.TreasureTypeCoin,
		Rare:  uint8(rare),
		X:     x,
		Y:     y,
		Val:   global.CoinAttr[rare],
	}
}
