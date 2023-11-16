package global

const (
	FocusPlay     uint8 = iota // 焦点窗口
	FocusSkillBox              // 技能框
)

const (
	MapHouse = iota // 地图类型 - 房间
	MapFloor        // 地图类型 - 地面
	MapRoad         // 地图类型 - 路径
)

const (
	CptHero = iota // 角色类型 - 主角
	CptPet         // 角色类型 - 宠物
)

const (
	MstZombie = iota // 怪物类型 - 僵尸
)

const (
	RareCommon = iota + 1
	RareUnCommon
	RareRare
	RareVeryRare
	RareLegendary
)

const (
	TreasureTypeCoin = iota + 1
	TreasureTypePotion
	TreasureTypePack
)

func IfElse[T any](exp bool, trueValue, falseValue T) T {
	if exp {
		return trueValue
	} else {
		return falseValue
	}
}
