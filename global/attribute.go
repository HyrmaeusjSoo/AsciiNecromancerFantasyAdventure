package global

// 角色属性
type Attribute struct {
	MaxHealth int
	MaxExp    int
	Damage    int
}

// 预设
var CptAttr = map[int]Attribute{
	CptHero: Attribute{100, 100, 10},
	CptPet:  Attribute{100, 100, 5},
}

var MstAttr = map[int]Attribute{
	MstZombie: Attribute{50, 40, 5},
}

var PotionAttr = map[int]int{
	RareCommon:    10,
	RareUnCommon:  30,
	RareRare:      50,
	RareVeryRare:  75,
	RareLegendary: 100,
}

var CoinAttr = map[int]int{
	RareCommon:    1,
	RareUnCommon:  3,
	RareRare:      5,
	RareVeryRare:  7,
	RareLegendary: 10,
}
