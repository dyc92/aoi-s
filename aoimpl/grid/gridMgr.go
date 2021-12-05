package aoigrid

import (
	. "aoi-s/aoimpl"
	"aoi-s/geo/r3"
	"math"
	"sync"
)

type GridManager struct {
	width, height             float64
	gridSize                  int
	minX, minZ, maxX, maxZ    float64
	maxRow, maxCol, maxGridID int
	grids                     []*Grid
	entityMap                 sync.Map
}

func InitGridManager(width, height float64, gridSize int) (gm *GridManager) {
	gm = &GridManager{}
	gm.width, gm.height, gm.gridSize = width, height, gridSize
	gm.maxX, gm.maxZ = width/2, height/2
	gm.minX, gm.minZ = -gm.maxX, -gm.maxZ

	gm.maxRow = int(math.Floor(height/float64(gridSize))) + 1
	gm.maxCol = int(math.Floor(width/float64(gridSize))) + 1
	gm.maxGridID = gm.maxRow * gm.maxCol
	gm.grids = make([]*Grid, 0, gm.maxGridID)
	for i := 0; i < gm.maxGridID; i++ {
		g := &Grid{Id: i}
		g.Init()
		gm.grids = append(gm.grids, g)
	}
	gm.entityMap = sync.Map{}
	return
}

func (gm GridManager) GetEntity(id int) *Entity {
	obj, ok := gm.entityMap.Load(id)
	if !ok {
		return nil
	}
	return obj.(*Entity)
}

func (gm GridManager) LengthInGrid(l float64) int {
	return int(math.Ceil(l / GridSize))
}

func (gm GridManager) Pos2GridIndex(pos r3.Vector) int {
	row := int(math.Floor(pos.Z-gm.minZ) / float64(gm.gridSize))
	col := int(math.Floor(pos.X-gm.minX) / float64(gm.gridSize))
	return row*gm.maxCol + col
}

func (gm *GridManager) CalcGridDistance(grid1, grid2 int) float64 {
	row1 := grid1 / gm.maxCol
	col1 := grid1 % gm.maxCol

	row2 := grid2 / gm.maxCol
	col2 := grid2 % gm.maxCol
	return math.Max(math.Abs(float64(row1-row2)), math.Abs(float64(col1-col2)))
}

func (gm GridManager) IsPosValid(pos r3.Vector) bool {
	return math.Abs(pos.X) <= gm.maxX && math.Abs(pos.Z) <= gm.maxZ
}

func (gm *GridManager) Release() {
	gm.grids = make([]*Grid, gm.maxGridID)
}

func (gm *GridManager) EnterMap(e *Entity) {
	if !gm.IsPosValid(e.Position) {
		return
	}
	e.GridID = gm.Pos2GridIndex(e.Position)
	gm.grids[e.GridID].Add(e)
	gm.entityMap.Store(e.Id, e)
}

func (gm *GridManager) LeaveMap(e *Entity) {
	if e.GridID < 0 || e.GridID >= gm.maxGridID {
		return
	}
	gm.grids[e.GridID].Remove(e)
	e.GridID = -1
}

func (gm *GridManager) UpdatePos(e *Entity) {
	if !gm.IsPosValid(e.Position) {
		return
	}
	if e.GridID < 0 || e.GridID >= gm.maxGridID {
		return
	}
	index := gm.Pos2GridIndex(e.Position)
	if index == e.GridID {
		return
	}
	gm.grids[e.GridID].Remove(e)
	gm.grids[index].Add(e)
}
