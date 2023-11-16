package game

import (
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
)

func (g *Game) Graph() {
	g.Screen.Clear()
	g.StateBox()

	// g.House.Draw(g.Screen)
	for _, v := range g.Corps {
		v.Draw(g.Screen)
	}
	g.Pet.Draw(g.Screen)
	g.At.Draw(g.Screen)
	for _, v := range g.Msts {
		v.Draw(g.Screen)
	}
	for _, v := range g.Treasures {
		v.Draw(g.Screen)
	}
}

func (g *Game) StateBox() {
	sx, sy := g.Screen.Size()
	for i := 0; i < sx; i++ {
		g.Screen.SetContent(i, sy-3, tcell.RuneHLine, nil, g.Style)
	}
	var at strings.Builder
	at.WriteString("@")
	at.WriteString(strconv.Itoa(g.At.Health))
	if g.At.LastAttacked != 0 {
		at.WriteString(strconv.Itoa(g.At.LastAttacked))
	}
	at.WriteString("/")
	at.WriteString(strconv.Itoa(g.At.Max))
	at.WriteString(" A")
	at.WriteString(strconv.Itoa(g.At.Skill[SKID_CuiDuBiShou]))
	at.WriteString(" S")
	at.WriteString(strconv.Itoa(g.At.Skill[SKID_ShiBao]))
	at.WriteString(" D")
	at.WriteString(strconv.Itoa(g.At.Skill[SKID_Temp]))
	at.WriteString(" F")
	at.WriteString(strconv.Itoa(g.At.Skill[SKID_FaZhen]))
	at.WriteString("  $")
	at.WriteString(strconv.Itoa(g.At.Coins))
	at.WriteString(" #")
	at.WriteString(strconv.Itoa(g.At.Exp))
	at.WriteString("/")
	at.WriteString(strconv.Itoa(g.At.MaxExp))
	at.WriteString("/")
	at.WriteString(strconv.Itoa(int(g.At.Level)))
	g.DrawText(0, sy-2, at.String(), g.Style)
	var ms strings.Builder
	ms.WriteString("M:")
	ms.WriteString(strconv.Itoa(len(g.Msts)))
	ms.WriteString(" %:")
	ms.WriteString(strconv.Itoa(len(g.Corps)))
	ms.WriteString(" S:")
	ms.WriteString(strconv.Itoa(g.Scores))
	g.DrawText(0, sy-1, ms.String(), g.Style)
	for k, v := range g.Msts {
		var mstb strings.Builder
		mstb.WriteString(strconv.Itoa(k))
		mstb.WriteString("=")
		mstb.WriteString(strconv.Itoa(v.Health))
		if v.LastAttacked != 0 {
			mstb.WriteString(strconv.Itoa(v.LastAttacked))
		}
		mstb.WriteString("/")
		mstb.WriteString(strconv.Itoa(v.Max))
		mstb.WriteString(" :")
		mstb.WriteString(strconv.Itoa(v.X))
		mstb.WriteString(",")
		mstb.WriteString(strconv.Itoa(v.Y))
		g.DrawText(sx/2, sy-(1+k), mstb.String(), g.Style)
	}
}

func (g *Game) DrawBox(x1, y1, x2, y2 int, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	tstyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorGray)
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			g.Screen.SetContent(col, row, '.', nil, tstyle)
		}
	}

	for col := x1; col <= x2; col++ {
		g.Screen.SetContent(col, y1, tcell.RuneHLine, nil, g.Style)
		g.Screen.SetContent(col, y2, tcell.RuneHLine, nil, g.Style)
	}
	for row := y1 + 1; row < y2; row++ {
		g.Screen.SetContent(x1, row, tcell.RuneVLine, nil, g.Style)
		g.Screen.SetContent(x2, row, tcell.RuneVLine, nil, g.Style)
	}
	if y1 != y2 && x1 != x2 {
		g.Screen.SetContent(x1, y1, tcell.RuneULCorner, nil, g.Style)
		g.Screen.SetContent(x2, y1, tcell.RuneURCorner, nil, g.Style)
		g.Screen.SetContent(x1, y2, tcell.RuneLLCorner, nil, g.Style)
		g.Screen.SetContent(x2, y2, tcell.RuneLRCorner, nil, g.Style)
	}

	g.DrawBoxText(x1+1, y1+1, x2-1, y2-1, text)
}

func (g *Game) DrawBoxText(x1, y1, x2, y2 int, text string) {
	row, col := y1, x1
	for _, r := range text {
		g.Screen.SetContent(col, row, r, nil, g.Style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func (g *Game) DrawText(x, y int, text string, style tcell.Style) {
	if style == (tcell.Style{}) {
		style = g.Style
	}
	for i := 0; i < len(text); i++ {
		g.Screen.SetContent(x+i, y, rune(text[i]), nil, style)
	}
}
