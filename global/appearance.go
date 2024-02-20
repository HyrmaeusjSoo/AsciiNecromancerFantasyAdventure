package global

import "github.com/gdamore/tcell/v2"

const (
	AsciiHorizon  = ' '  // 地平线
	AsciiHero     = '@'  // 角色
	AsciiPet      = 'd'  // 宠物
	AsciiHWall    = '-'  // 横墙
	AsciiVWall    = '|'  // 竖墙
	AsciiDoor     = '+'  // 门
	AsciiFloorLow = '.'  // 低洼地面
	AsciiFloor    = '·'  // 地板
	AsciiZombie   = 'Z'  // 僵尸
	AsciiCorpse   = '%'  // 尸体
	AsciiCoin     = '$'  // 金币
	AsciiPotion   = '6'  // 治疗药剂
	AsciiPack     = '&'  // 包裹
	AsciiWater    = '~'  // 水
	AsciiFence    = '\'' // 栅栏
	AsciiBush     = '#'  // 灌木从，矮树从
)

// 地图样式
var MapStyle = map[int]tcell.Style{
	MapHouse: tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault),
	MapFloor: tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.Color248),
	MapRoad:  tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault),
}

// 角色样式
var CptStyle = map[int]tcell.Style{
	CptHero: tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.Color199),
	CptPet:  tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.Color140),
}

// 怪物样式
var MstStyle = map[int]tcell.Style{
	MstZombie: tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.Color47),
}

// 尸体样式
var CorpseStyle = tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.Color124)

// 稀有度样式
var RareStyle = map[int]tcell.Style{
	RareCommon:    tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorWhite),
	RareUnCommon:  tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorLime),
	RareRare:      tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorBlue),
	RareVeryRare:  tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorFuchsia),
	RareLegendary: tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.Color202),
}
