package global

import (
	"github.com/gdamore/tcell/v2"
)

const (
	FocusPlay     uint8 = iota // 焦点窗口
	FocusSkillBox              // 技能框
)

const (
	AsciiHorizon  = ' ' // 地平线
	AsciiHero     = '@' // 角色
	AsciiPet      = 'd' // 宠物
	AsciiHWall    = '-' // 横墙
	AsciiVWall    = '|' // 竖墙
	AsciiDoor     = '+' // 门
	AsciiFloorLow = '.' // 低洼地面
	AsciiFloor    = '·' // 地面
	AsciiZombie   = 'Z' // 僵尸
	AsciiCorpse   = '%' // 尸体
)

const (
	MapHouse = iota // 地图类型 - 房间
	MapFloor        // 地图类型 - 地面
	MapRoad         // 地图类型 - 路径
)

// 地图样式
var MapStyle = map[int]tcell.Style{
	MapHouse: tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault),
	MapFloor: tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.Color248),
	MapRoad:  tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault),
}

const (
	CptHero = iota // 角色类型 - 主角
	CptPet         // 角色类型 - 宠物
)

// 角色样式
var CptStyle = map[int]tcell.Style{
	CptHero: tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.Color199),
	CptPet:  tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.Color213),
}

const (
	MstZombie = iota // 怪物类型 - 僵尸
)

// 怪物样式
var MstStyle = map[int]tcell.Style{
	MstZombie: tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.Color47),
}

// 尸体样式
var CorpseStyle = tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.Color124)

func IfElse[T any](exp bool, trueValue, falseValue T) T {
	if exp {
		return trueValue
	} else {
		return falseValue
	}
}
