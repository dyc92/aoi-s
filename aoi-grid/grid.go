package aoigrid

type Manager struct {
}

type Entity struct {
	Type                 EntityType
	HeartbeatInterval    int
	NeightborPositionLen float32
	DefaultKnownRadius   float32
	DefaultSightRadius   float32
	MaxSpeed             float32
}

func (e *Entity) init() {
	e.NeightborPositionLen = 3.0
	e.DefaultKnownRadius = 0.5
	e.DefaultSightRadius = 20.0
	e.MaxSpeed = 15.0
}

type Grid struct {
	Entities map[EntityType]map[Entity]struct{}
}
