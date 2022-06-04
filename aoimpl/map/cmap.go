package cmap

import (
	"aoi-s/aoimpl"
	"aoi-s/aoimpl/grid"
)

type Map struct {
	Name               string
	MapTemplateId      int
	LineId             int
	LimitMaxRoleCount  int
	LimitWarnRoleCount int
	Tag                string
	SightRadius        int
	gridMgr            aoimpl.AOI
	ModeType           aoimpl.EnumMapModeType
	StateType          aoimpl.EnumMapStateType
}

func (m *Map) Init(name string) {
	m.ModeType = aoimpl.MapModeUnKnown
	m.StateType = aoimpl.MapStateUnKnown
	m.Name = name
	m.gridMgr = aoigrid.InitGridManager(1000, 1000, aoimpl.GridSize)
}

func (m *Map) EnterMap(entity *aoimpl.Entity) {
	m.gridMgr.EnterMap(entity)
}

func (m *Map) GetEntity(id int) *aoimpl.Entity {
	return m.gridMgr.GetEntity(id)
}

func (m *Map) Print() {
	m.gridMgr.Release()
}
