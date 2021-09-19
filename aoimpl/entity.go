package aoimpl

import (
	. "aoi-s/common"
	"aoi-s/geo/r3"
)

type Entity struct {
	Id        int
	GridID    int
	SightGrid int
	KnowGrid  int
	Name      string
	Type      EntityType

	Position    r3.Vector //位置
	Rotation    r3.Vector //旋转
	Orientation r3.Vector //方向
	MaxSpeed    float32   //最大速度

	worldBox, sightBox, knownBox A3DAABB

	collisionRadius float32 //碰撞半径
	KnownRadius     float32
	SightRadius     float32

	HideSelf  bool
	HideOther bool
}

func (e *Entity) init() {
	e.KnownRadius = 0.5
	e.SightRadius = 20.0
	e.MaxSpeed = 15.0
	e.worldBox = NewBox()
	e.sightBox = NewBox()
	e.knownBox = NewBox()
	e.collisionRadius = 0.1
}

type GridManager struct {
}
