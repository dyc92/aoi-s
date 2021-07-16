package aoimpl

type EntityType int8

const (
	Player EntityType = iota
	Begin  EntityType = Player
	Monster
	Npc
	SubObject
	Loot
	Mine
	Pet
	PlayerMirror
	Gfx
	MaxType
)

const (
	Known   = 0x01 //可被观察（可被感知）,未修改前用world替代的known
	World   = 0x02 //碰撞
	Sight   = 0x04 //视野
	Warning = 0x08 //警戒

	All = 0xff
	Max = 4 //索引
)
