package global

import (
	"github.com/gdamore/tcell/v2"
)

const (
	AsciiHorizon  = ' '
	AsciiHero     = '@'
	AsciiPet      = 'd'
	AsciiHWall    = '-'
	AsciiVWall    = '|'
	AsciiDoor     = '+'
	AsciiFloorLow = '.'
	AsciiFloor    = '·'
	AsciiZombie   = 'Z'
	AsciiCorpse   = '%'
)

const (
	MapHouse = iota
	MapFloor
	MapRoad
)

var MapStyle = map[int]tcell.Style{
	MapHouse: tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault),
	MapFloor: tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.Color248),
	MapRoad:  tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault),
}

const (
	CptHero = iota
	CptPet
)

var CptStyle = map[int]tcell.Style{
	CptHero: tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.Color199),
	CptPet:  tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.Color213),
}

const (
	MstZombie = iota
)

var CorpseStyle = tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.Color124)

var MstStyle = map[int]tcell.Style{
	MstZombie: tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.Color47),
}

func IfElse[T any](exp bool, trueValue, falseValue T) T {
	if exp {
		return trueValue
	} else {
		return falseValue
	}
}
