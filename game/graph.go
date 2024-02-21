package game

import (
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
)

func (g *Game) Graph() {
	g.Screen.Clear()

	DrawFaZhen(g)

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

func (g *Game) OpeningAnimation() {
	x, y := g.Screen.Size()
	x, y = x/2, y/2
	g.DrawText(x-17, y, "Ascii", tcell.StyleDefault.Background(tcell.Color16).Foreground(tcell.ColorReset))
	g.DrawText(x-11, y, "Necromancer", tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.Color43))
	g.DrawText(x+1, y, "F", tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.Color160))
	g.DrawText(x+2, y, "a", tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.Color208))
	g.DrawText(x+3, y, "n", tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.Color226))
	g.DrawText(x+4, y, "t", tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.Color40))
	g.DrawText(x+5, y, "a", tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.Color45))
	g.DrawText(x+6, y, "s", tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.Color21))
	g.DrawText(x+7, y, "y", tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.Color165))
	g.DrawText(x+9, y, "Adventure", tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.Color199))
	g.Screen.Show()
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
	at.WriteString(strconv.Itoa(int(g.At.Level)))
	at.WriteString("/")
	at.WriteString(strconv.Itoa(g.At.Exp))
	at.WriteString("/")
	at.WriteString(strconv.Itoa(g.At.MaxExp))
	g.DrawText(0, sy-2, at.String(), g.Style)
	var ms strings.Builder
	ms.WriteString("M:")
	ms.WriteString(strconv.Itoa(len(g.Msts)))
	ms.WriteString(" %:")
	ms.WriteString(strconv.Itoa(len(g.Corps)))
	ms.WriteString(" S:")
	ms.WriteString(strconv.Itoa(g.Scores))
	ms.WriteString("   ")
	spell := g.SK.Skills[SKID_FaZhen]
	ms.WriteString(strconv.Itoa(spell.TargetX))
	ms.WriteString("/")
	ms.WriteString(strconv.Itoa(spell.TargetY))
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
	for k, v := range text {
		g.Screen.SetContent(x+k, y, v, nil, style)
	}
}
