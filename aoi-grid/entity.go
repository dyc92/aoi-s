package aoigrid

import (
	. "aoi/common"
)

type Entity struct {
	ID        int
	GridID    int
	SightGrid int
	KnowGrid  int
	Name      string
	Type      EntityType

	Position    Vector  //位置
	Rotation    Vector  //旋转
	Orientation Vector  //方向
	MaxSpeed    float32 //最大速度

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
