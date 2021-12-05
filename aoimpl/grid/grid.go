package aoigrid

import (
	. "aoi-s/aoimpl"
)

type Grid struct {
	Id, X, Y int //格子坐标3

	Entities map[EnumEntityType]map[*Entity]struct{}
}

func (g *Grid) Init() {
	g.Entities = make(map[EnumEntityType]map[*Entity]struct{}, MaxType)
}

func (g *Grid) Add(e *Entity) {

	if _, ok := g.Entities[e.Type]; !ok {
		g.Entities[e.Type] = make(map[*Entity]struct{}, 500)
	}
	g.Entities[e.Type][e] = struct{}{}
}

func (g *Grid) Remove(e *Entity) {
	if _, ok := g.Entities[e.Type]; !ok {
		delete(g.Entities[e.Type], e)
	}
}

func (g Grid) GetEntities(t EnumEntityType) map[*Entity]struct{} {
	return g.Entities[t]
}
