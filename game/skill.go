package game

import (
	"math"
	"necromancer/global"
	"time"

	"github.com/gdamore/tcell/v2"
)

const (
	SKID_CuiDuBiShou uint8 = iota + 1
	SKID_ShiBao
	SKID_Temp
	SKID_FaZhen

	SKMAXLevel = 4
)

type Spell struct {
	Id       uint8
	Name     string
	Cast     int
	Dice     int
	HighCast int
	HighDice int
}

type Skill struct {
	Current uint8
	Skills  map[uint8]Spell
}

func NewSkill() Skill {
	return Skill{
		0,
		map[uint8]Spell{
			SKID_CuiDuBiShou: Spell{SKID_CuiDuBiShou, "a. Cui Du Bi Shou (2d10 up 1d10)",
				2, 10, 1, 10},
			SKID_ShiBao: Spell{SKID_ShiBao, "s. Shi Bao (8d6 up 2d6)",
				8, 6, 2, 6},
			SKID_Temp: Spell{SKID_Temp, "d. Temp",
				0, 0, 0, 0},
			SKID_FaZhen: Spell{SKID_FaZhen, "f. Fa Zhen",
				0, 0, 0, 0},
		},
	}
}

func (s *Skill) Select(g *Game, key tcell.Key) {
	showAnime := true
	switch key {
	case tcell.KeyUp:
		s.Current = global.IfElse(s.Current <= 1, uint8(len(s.Skills)), s.Current-1)
	case tcell.KeyDown:
		s.Current = global.IfElse(s.Current >= uint8(len(s.Skills)), 1, s.Current+1)
	case tcell.KeyClear:
		s.Current = 1
	case tcell.KeyEnter:
		if s.Current == 0 {
			return
		}
		if g.At.Skill[s.Current] >= SKMAXLevel {
			return
		}
		g.At.Skill[s.Current]++
		g.Focus = global.FocusPlay
		g.Graph()
		showAnime = false
	}
	if showAnime {
		s.SelectAnime(g)
	}
}

func (s *Skill) SelectAnime(g *Game) {
	sx, sy := g.Screen.Size()
	sx, sy = sx-1, sy-3
	x1, y1 := sx/3, sy/3
	x2, y2 := x1*2, y1+len(s.Skills)+1 //y1*2
	g.DrawBox(x1, y1, x2, y2, "")
	for k, v := range s.Skills {
		style := global.IfElse(s.Current == k, tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.Color202), g.Style)
		g.DrawText(x1+1, y1+int(k), v.Name, style)
	}
	g.Screen.Show()
}

func spellDamage(spell Spell, lv int) int {
	ar := global.AttackRoll()
	if ar == 1 {
		return 0
	}
	if ar == 20 {
		spell.Cast *= 2
	}
	dmg := global.Roll(spell.Cast, spell.Dice)
	dmg += global.Roll((lv-1)*spell.HighCast, spell.HighDice)
	return dmg
}

var SkillFunc = map[uint8]func(g *Game){
	/*
		1 2 3 4 5 6 7 8 9 10
		2
		3   @
		4     ·
		5       ·
		6         ·
		7           ·
		8             *
		9               Z
		10
	*/
	SKID_CuiDuBiShou: func(g *Game) {
		if len(g.Msts) == 0 {
			return
		}
		var (
			ax, ay = float64(g.At.X), float64(g.At.Y)
			tx, ty float64
			b, bk  = 50.0, -1
		)
		for k, v := range g.Msts {
			a := math.Abs(float64(v.X)-ax) + math.Abs(float64(v.Y)-ay)
			if a < b {
				b, bk = a, k
				tx, ty = float64(v.X), float64(v.Y)
			}
		}
		if b >= 50 || bk == -1 {
			return
		}
		dmg := spellDamage(g.SK.Skills[SKID_CuiDuBiShou], g.At.Skill[SKID_CuiDuBiShou])
		if hp := g.Msts[bk].Heal(-dmg); hp <= 0 {
			g.MstDeath(bk)
		}

		step := math.Max(math.Abs(ax-tx), math.Abs(ay-ty))
		increX, increY := -(ax-tx)/step, -(ay-ty)/step
		x, y := ax, ay

		for i := 1; i < int(step); i++ {
			x += increX
			y += increY
			w := global.IfElse(x == tx-increX && y == ty-increY, '*', '·')
			g.Screen.SetContent(int(x+0.5), int(y+0.5), w, nil, g.Style)
		}
		g.Screen.Show()
	},

	/*
		· · * · ·
		· * ~ * ·
		* ~ % ~ *
		· * ~ * ·
		· · * · ·
	*/
	SKID_ShiBao: func(g *Game) {
		if len(g.Corps) <= 0 {
			return
		}
		scr, x, y := g.Screen, 0, 0
		b, bk := 100, -1
		ax1, ay1, ax2, ay2 := g.At.X-20, g.At.Y-10, g.At.X+20, g.At.Y+10
		for k, v := range g.Corps {
			if !(v.X >= ax1 && v.X <= ax2 && v.Y >= ay1 && v.Y <= ay2) {
				continue
			}
			a := int(math.Abs(float64(v.X-g.At.X)) + math.Abs(float64(v.Y-g.At.Y)))
			if a < b {
				b, bk = a, k
				x, y = v.X, v.Y
			}
		}
		if b >= 100 || bk == -1 {
			return
		}

		spell := g.SK.Skills[SKID_ShiBao]
		lv := g.At.Skill[SKID_ShiBao]
		g.Corps = append(g.Corps[:bk], g.Corps[bk+1:]...)
		x1, x2, y1, y2 := x-2, x+2, y-2, y+2
		for k, v := range g.Msts {
			if !(v.X >= x1 && v.X <= x2 && v.Y >= y1 && v.Y <= y2) {
				continue
			}
			dmg := spellDamage(spell, lv)
			if hp := v.Heal(-dmg); hp <= 0 {
				g.MstDeath(k)
			}
		}

		for row := y - 1; row <= y+1; row++ {
			for col := x - 1; col <= x+1; col++ {
				scr.SetContent(col, row, '·', nil, g.Style)
			}
		}
		scr.SetContent(x, y, '%', nil, global.CorpseStyle)
		scr.Show()
		time.Sleep(150 * time.Millisecond)
		for row := y1; row <= y2; row++ {
			for col := x1; col <= x2; col++ {
				w := global.IfElse(row >= y-1 && row <= y+1 && col >= x-1 && col <= x+1, '*', '·')
				scr.SetContent(col, row, w, nil, g.Style)
			}
		}
		scr.SetContent(x, y, '%', nil, global.CorpseStyle)
		scr.Show()
		time.Sleep(150 * time.Millisecond)
		for row := y1; row <= y2; row++ {
			for col := x1; col <= x2; col++ {
				w := global.IfElse(row == y1 || row == y2 || col == x1 || col == x2, '*', '~')
				scr.SetContent(col, row, w, nil, g.Style)
			}
		}
		scr.Show()
		time.Sleep(150 * time.Millisecond)
		for row := y1; row <= y2; row++ {
			for col := x1; col <= x2; col++ {
				scr.SetContent(col, row, '~', nil, g.Style)
			}
		}
		scr.Show()
	},

	/*
		             ··~·~··
		        ·~·           ·~·
		     ·~·                ·~·
		   ·~·                    ·~·
		  ·~·                      ·~·
		 ·~·                        ·~·
		·~·             +            ·~·
		 ·~·                        ·~·
		  ·~·                      ·~·
		   ·~·                    ·~·
		     ·~·                ·~·
		        ·~·           ·~·
		             ··~·~··
	*/
	SKID_FaZhen: func(g *Game) {

	},
}
