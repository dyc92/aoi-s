package cmap

import (
	"aoi-s/aoimpl"
	"aoi-s/aoimpl/grid"
)

type CMap struct {
	Name               string
	MapTemplateId      int
	LineId             int
	LimitMaxRoleCount  int
	LimitWarnRoleCount int
	Tag                string
	ModeType           aoimpl.EnumMapModeType
	StateType          aoimpl.EnumMapStateType
	GridMgr            aoimpl.AOI
	SightRadius        int
}

func (m *CMap) Init(name string) {
	m.ModeType = aoimpl.MapModeUnKnown
	m.StateType = aoimpl.MapStateUnKnown
	m.Name = name
	m.GridMgr = aoigrid.InitGridManager(1000, 1000, int(aoimpl.GridSize))
}
