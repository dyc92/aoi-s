package ECSComponent

import (
	"aoi-s/Game/ECSComponent/ComponentDefine"
	"aoi-s/geo/r3"
)

type CMovement struct {
	CurPosition r3.Vector
	MoveTasks   chan r3.Vector //移动指令队列
	MaxSpeed    float32        //最大速度
	cType       ComponentDefine.BaseComponent
}

func (m *CMovement) Init(tid int) {
	//TODO:tid用来查找配置，生成该组件时检查是否可移动
	m.MoveTasks = make(chan r3.Vector) //可以声明容量，允许卡顿几个帧长
	m.cType.Type = ComponentDefine.Movement
}

// Move 移动后的client发来的坐标,向量
func (m *CMovement) Move(Position, Direction r3.Vector) {
	m.MoveTasks <- Position
}
