package aoigrid

import . "aoi/aoimpl"

type Grid struct {
	Id,X,Y int //格子坐标3

	Entities map[EntityType]map[int]*Entity
}

func (g *Grid) AddEntity(e Entity) {

	if _, ok := g.Entities[e.Type]; !ok {
		g.Entities[e.Type] = make(map[int]*Entity,500)
	}
	g.Entities[e.Type][e.Id] =&e
}

func (g *Grid) Remove(e Entity) {
	if _, ok := g.Entities[e.Type]; !ok {
		delete(g.Entities[e.Type], e.Id)
	}
}

func (g Grid) GetEntity(t EntityType) map[int]*Entity {
	return g.Entities[t]
}