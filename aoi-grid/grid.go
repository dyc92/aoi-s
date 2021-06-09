package aoigrid

type Manager struct {
}

type Grid struct {
	Entities map[EntityType]map[Entity]struct{}
}

func (g *Grid) AddEntity(e Entity) {

	if _, ok := g.Entities[e.Type]; !ok {
		g.Entities[e.Type] = map[Entity]struct{}{e: {}}
	}
	g.Entities[e.Type][e] = struct{}{}
}

func (g *Grid) Remove(e Entity) {
	if _, ok := g.Entities[e.Type]; !ok {
		delete(g.Entities[e.Type], e)
	}
}

func (g Grid) GetEntity(t EntityType) map[Entity]struct{} {
	return g.Entities[t]
}
