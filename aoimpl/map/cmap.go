package cmap

import "aoi-s/aoimpl/grid"

type CMap struct {
	Name               string
	MapTemplateId      int
	LineId             int
	LimitMaxRoleCount  int
	LimitWarnRoleCount int
	Tag                string
	GridSize           float32
	MaxKnow            float32
	ModeType           EnumMapModeType
	StateType          EnumMapStateType
	GridMgr            aoigrid.GridManager
	SightRadius        int
}

func (m *CMap) Init() {
	m.ModeType = MapModeUnKnown
	m.StateType = MapStateUnKnown
}
