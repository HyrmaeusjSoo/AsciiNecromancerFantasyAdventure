package game

import (
	"necromancer/global"

	"github.com/gdamore/tcell/v2"
)

type House struct {
	style          tcell.Style
	x1, y1, x2, y2 int
	door           []DoorPosition
}
type DoorPosition struct {
	x, y int
}

func NewHouse(x1, y1, x2, y2 int, dp []DoorPosition) House {
	return House{
		global.MapStyle[global.MapHouse],
		x1, y1, x2, y2, dp,
	}
}

func (h *House) Draw(s tcell.Screen) {
	if h.y2 < h.y1 {
		h.y1, h.y2 = h.y2, h.y1
	}
	if h.x2 < h.x1 {
		h.x1, h.x2 = h.x2, h.x1
	}
	//填充地板
	for row := h.y1; row <= h.y2; row++ {
		for col := h.x1; col < h.x2; col++ {
			s.SetContent(col, row, '.', nil, global.MapStyle[global.MapFloor])
		}
	}
	//画墙
	for col := h.x1; col <= h.x2; col++ { //tcell.RuneHLine
		s.SetContent(col, h.y1, '-', nil, h.style)
		s.SetContent(col, h.y2, '-', nil, h.style)
	}
	for row := h.y1 + 1; row < h.y2; row++ { //tcell.RuneVLine
		s.SetContent(h.x1, row, '|', nil, h.style)
		s.SetContent(h.x2, row, '|', nil, h.style)
	}
	// 画墙角
	/* if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, '-', nil, h.style) //tcell.RuneULCorner
		s.SetContent(x2, y1, '-', nil, h.style) //tcell.RuneURCorner
		s.SetContent(x1, y2, '-', nil, h.style) //tcell.RuneLLCorner
		s.SetContent(x2, y2, '-', nil, h.style) //tcell.RuneLRCorner
	} */

	//画门
	for _, p := range h.door {
		s.SetContent(p.x, p.y, '+', nil, h.style)
	}
}
