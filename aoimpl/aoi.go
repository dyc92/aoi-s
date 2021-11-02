package aoimpl

import (
	"aoi-s/geo/r2"
)

type AOI interface {
	AddEntity(e Entity)
	Remove(e Entity)
	GetEntity(t EnumEntityType)
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
