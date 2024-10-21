package game

import (
	"math"
	"necromancer/global"
	"strconv"
	"strings"
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
	Name     string // 法术
	Cast     int    // 骰子数
	Dice     int    // 骰子面数
	HighCast int    // 升环施法加成 骰子数
	HighDice int    // 升环施法加成 骰子面数
	Used     bool   // 是否使用
	TargetX  int    // 目标坐标x
	TargetY  int    // 目标坐标y
}

type Skill struct {
	Current uint8
	Skills  map[uint8]*Spell
}

func NewSkill() Skill {
	return Skill{
		0,
		map[uint8]*Spell{
			SKID_CuiDuBiShou: {SKID_CuiDuBiShou, "[a]. CuiDuBiShou (2d10 +1d10)",
				2, 10, 1, 10, false, 0, 0},
			SKID_ShiBao: {SKID_ShiBao, "[s]. ShiBao (8d6 +2d6)",
				8, 6, 2, 6, false, 0, 0},
			SKID_Temp: {SKID_Temp, "[d]. Temp",
				0, 0, 0, 0, false, 0, 0},
			SKID_FaZhen: {SKID_FaZhen, "[f]. FaZhen (1d4 +2d4)",
				1, 4, 2, 4, false, 0, 0},
		},
	}
}

func (s *Skill) Select(g *Game, key tcell.Key) {
	showAnime := true
	switch key {
	case tcell.KeyUp, 'k':
		s.Current = global.IfElse(s.Current <= 1, uint8(len(s.Skills)), s.Current-1)
	case tcell.KeyDown, 'j':
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
		style := g.Style
		if g.At.Skill[k] == SKMAXLevel {
			style = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.Color245)
		}
		if s.Current == k {
			style = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.Color202)
		}
		name := strings.Builder{}
		name.WriteString(" ")
		name.WriteString(strconv.Itoa(g.At.Skill[k]))
		name.WriteString(" - ")
		name.WriteString(v.Name)
		g.DrawText(x1+1, y1+int(k), name.String(), style)
	}
	g.Screen.Show()
}

func spellDamage(spell *Spell, lv int) int {
	dmg := global.LaunchAttack(spell.Cast, spell.Dice)
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
		g.Corps = append(g.Corps[:bk], g.Corps[bk+1:]...)

		spell := g.SK.Skills[SKID_ShiBao]
		// spell.Used, spell.TargetX, spell.TargetY = true, x, y

		lv := g.At.Skill[SKID_ShiBao]
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

	SKID_Temp: func(g *Game) {},

	/*
		~.~'~.~'~.~'~.~'~.~'~.~
		'       .       .       '
		~    .             .    ~
		.           .           .
		~                       ~
		'                       '
		~     .     #     .     ~
		.                       .
		~                       ~
		'           .           '
		~    .             .    ~
		.       .       .       .
		~.~'~.~'~.~'~.~'~.~'~.~
	*/
	SKID_FaZhen: func(g *Game) {
		if len(g.Msts) == 0 {
			return
		}
		var (
			tx, ty int         // 目标点
			b, bk  = 100.0, -1 // 距离， 怪堆下标
		)
		for k, v := range g.Msts {
			a := math.Abs(float64(v.X)-float64(g.At.X)) + math.Abs(float64(v.Y)-float64(g.At.Y))
			if a < b {
				b, bk = a, k
				tx, ty = v.X, v.Y
			}
		}
		if b >= 100 || bk == -1 {
			return
		}

		spell := g.SK.Skills[SKID_FaZhen]
		spell.Used, spell.TargetX, spell.TargetY = true, tx, ty

		FaZhenFlash(g)
		DrawFaZhen(g)
	},
}

func FaZhenFlash(g *Game) {
	lv := g.At.Skill[SKID_FaZhen]
	if lv < 1 {
		return
	}
	spell := g.SK.Skills[SKID_FaZhen]
	if !spell.Used {
		return
	}
	tx, ty := spell.TargetX, spell.TargetY
	x1, x2, y1, y2 := tx-12, tx+12, ty-6, ty+6
	for k, v := range g.Msts {
		if !(v.X >= x1 && v.X <= x2 && v.Y >= y1 && v.Y <= y2) {
			continue
		}

		dmg := spellDamage(spell, lv)
		if hp := v.Heal(-dmg); hp <= 0 {
			g.MstDeath(k)
		}
	}
}

func DrawFaZhen(g *Game) {
	spell := g.SK.Skills[SKID_FaZhen]
	if !spell.Used {
		return
	}
	offsetX := spell.TargetX - 12
	colorStyle := tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.Color124)
	g.DrawText(offsetX, spell.TargetY-6, " ~.~'~.~'~.~'~.~'~.~'~.~ ", colorStyle)
	g.DrawText(offsetX, spell.TargetY-5, "'       .       .       '", colorStyle)
	g.DrawText(offsetX, spell.TargetY-4, "~    .             .    ~", colorStyle)
	g.DrawText(offsetX, spell.TargetY-3, ".           .           .", colorStyle)
	g.DrawText(offsetX, spell.TargetY-2, "~                       ~", colorStyle)
	g.DrawText(offsetX, spell.TargetY-1, "'                       '", colorStyle)
	g.DrawText(offsetX, spell.TargetY, "~     .     #     .     ~", colorStyle)
	g.DrawText(offsetX, spell.TargetY+1, ".                       .", colorStyle)
	g.DrawText(offsetX, spell.TargetY+2, "~                       ~", colorStyle)
	g.DrawText(offsetX, spell.TargetY+3, "'           .           '", colorStyle)
	g.DrawText(offsetX, spell.TargetY+4, "~    .             .    ~", colorStyle)
	g.DrawText(offsetX, spell.TargetY+5, ".       .       .       .", colorStyle)
	g.DrawText(offsetX, spell.TargetY+6, " ~.~'~.~'~.~'~.~'~.~'~.~ ", colorStyle)
	g.Screen.Show()
}
