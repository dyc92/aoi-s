package aoimpl

type EnumEntityType int8

const (
	Player EnumEntityType = iota
	Monster
	Npc
	SubObject
	Loot
	Mine
	Pet
	PlayerMirror
	MaxType
)

type EnumSightType int

const (
	Known   EnumSightType = 0x01 //可被观察（可被感知）,未修改前用world替代的known
	World   EnumSightType = 0x02 //碰撞
	Sight   EnumSightType = 0x04 //视野
	Warning EnumSightType = 0x08 //警戒

	All EnumSightType = 0xff
	Max EnumSightType = 4 //索引
)
