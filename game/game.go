package game

import (
	"math/rand"
	"necromancer/creature"
	"necromancer/global"
	"necromancer/treasure"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	//maximum int
	Screen    tcell.Screen
	Style     tcell.Style
	Focus     uint8
	At        creature.Character
	Pet       creature.Character
	House     House
	Msts      []*creature.Monster
	Corps     []*creature.Corpse
	Scores    int
	SK        Skill
	Treasures []treasure.Treasure
}

func InitGameObject(x, y int) (g Game) {
	msts := []*creature.Monster{
		creature.NewMonster(global.MstStyle[global.MstZombie], x-x, y-y, global.AsciiZombie, global.MstZombie),
		creature.NewMonster(global.MstStyle[global.MstZombie], x-1, y-4, global.AsciiZombie, global.MstZombie),
	}
	return Game{
		Focus:  global.FocusPlay,
		At:     creature.NewCharacter(global.CptStyle[global.CptHero], (x-1)/2, (y-3)/2, global.AsciiHero, global.CptHero),
		Pet:    creature.NewCharacter(global.CptStyle[global.CptPet], (x-1)/2+1, (y-3)/2+1, global.AsciiPet, global.CptPet),
		House:  NewHouse(10, 10, 20, 20, []DoorPosition{{10, 15}, {20, 15}}),
		Msts:   msts,
		Scores: 0,
		SK:     NewSkill(),
	}
}

func NewGame() *Game {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	game := InitGameObject(s.Size())
	game.Screen = s
	game.Style = defStyle
	return &game
}

func (g *Game) Start() {
	g.OpeningAnimation()
	// g.Graph()
	Input(g)
}

func (g *Game) Turn(act tcell.Key) {
	if g.Focus == global.FocusPlay {
		tx, ty := g.At.X, g.At.Y
		switch act {
		case tcell.KeyUp:
			ty--
		case tcell.KeyDown:
			ty++
		case tcell.KeyLeft:
			tx--
		case tcell.KeyRight:
			tx++
		default:
			if act == 'a' {
				go g.UseSkill(SKID_CuiDuBiShou)
			} else if act == 's' {
				go g.UseSkill(SKID_ShiBao)
			} else if act == 'd' {
				go g.UseSkill(SKID_Temp)
			} else if act == 'f' {
				go g.UseSkill(SKID_FaZhen)
			}
		}
		g.turnSched(tx, ty)
		g.Graph()
	} else if g.Focus == global.FocusSkillBox {
		g.SK.Select(g, act)
	}

	g.At.LastAttacked = 0
	for _, v := range g.Msts {
		v.LastAttacked = 0
	}
}

func (g *Game) UseSkill(skid uint8) {
	if g.At.Skill[skid] > 0 {
		SkillFunc[skid](g)
	}
}

func (g *Game) turnSched(tx, ty int) {
	FaZhenFlash(g)

	if g.At.TurnRound(tx, ty); CanMove(g.Screen, tx, ty) {
		g.At.Move(tx, ty)
	}
	if ptx, pty, canMove := g.Pet.Creature.TurnRound(tx, ty); canMove {
		g.Pet.Follow(ptx, pty)
	}
	for k, v := range g.Msts {
		if mtx, mty, canMove := v.TurnRound(g.At.Tx, g.At.Ty); canMove && CanMove(g.Screen, mtx, mty) {
			v.Move(mtx, mty)
		}
		if v.X == g.At.Tx && v.Y == g.At.Ty {
			if hp := v.Heal(-g.At.HitDamage()); hp <= 0 {
				g.MstDeath(k)
				continue
			}
		}
		if v.Tx == g.At.X && v.Ty == g.At.Y {
			if hp := g.At.Heal(-v.HitDamage()); hp <= 0 {
				g.Screen.Clear()
				g.DrawText(tx, ty, "The necromancer should know what death is.", g.Style)
				g.Screen.Show()
				time.Sleep(5 * time.Second)
			}
		}
	}
	if len(g.Msts) < 2 {
		g.GenerateMst()
	}

	for k, v := range g.Treasures {
		if tx != v.X || ty != v.Y {
			continue
		}
		used := false
		switch v.Type {
		case global.TreasureTypeCoin:
			g.At.Coins += v.Val
			used = true
		case global.TreasureTypePotion:
			g.At.Heal(v.Val)
			used = true
		case global.TreasureTypePack:
			used = true
		}
		if used {
			g.Treasures = append(g.Treasures[:k], g.Treasures[k+1:]...)
			return
		}
	}
}

/*
s(0x, ny)
s(mx, ny)
s(0y, nx)
s(my, nx)
*/
func (g *Game) GenerateMst() {
	mx, my := g.Screen.Size()
	mx, my = mx-1, my-4
	x, y := 0, 0
	if rand.Intn(4) < 2 {
		x = []int{0, mx}[rand.Intn(2)]
		y = rand.Intn(my)
	} else {
		x = rand.Intn(mx)
		y = []int{0, my}[rand.Intn(2)]
	}
	g.Msts = append(g.Msts, creature.NewMonster(global.MstStyle[global.MstZombie], x, y, global.AsciiZombie, global.MstZombie))
}

func (g *Game) MstDeath(i int) {
	if len(g.Msts) < (i+1) || g.Msts == nil {
		return
	}
	mst := g.Msts[i]
	x, y := mst.X, mst.Y
	g.Corps = append(g.Corps, &creature.Corpse{Id: 10001, Style: global.CorpseStyle, Name: global.AsciiCorpse, Type: global.MstZombie, X: x, Y: y})
	g.Msts = append(g.Msts[:i], g.Msts[i+1:]...)
	g.Reward(x, y)
	g.Scores += 1
	lvUp := g.At.AddExp(global.MstAttr[mst.Type].MaxExp)
	if lvUp {
		maxLv := 0
		for _, v := range g.At.Skill {
			maxLv += v
		}
		if maxLv < SKMAXLevel*len(g.SK.Skills) {
			g.Focus = global.FocusSkillBox
			g.SK.Select(g, tcell.KeyClear)
		}
	}
}

func (g *Game) Reward(x, y int) {
	if global.Roll(1, 6) != 6 {
		return
	}

	if g.At.X <= x {
		x += 1
	} else {
		x -= 1
	}
	if g.At.Y <= y {
		y += 1
	} else {
		y -= 1
	}

	g.Treasures = append(g.Treasures, treasure.New(x, y))
}
