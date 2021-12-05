package aoimpl

import (
	"aoi-s/geo/r2"
	"aoi-s/geo/r3"
)

type AOI interface {
	GetEntity(int) *Entity
	EnterMap(e *Entity)
	LeaveMap(e *Entity)
	UpdatePos(e *Entity)
	Release()
	Pos2GridIndex(pos r3.Vector) int
}

type IAoiScan struct {
	MaxDepth   int
	CheckCount int
	HitCount   int
	Sight      r2.Rect
	Known      r2.Rect
	MarkerTag  []int
	WatcherTag []int
}
