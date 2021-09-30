package aoigrid

import (
	. "aoi-s/aoimpl"
	"aoi-s/aoimpl/map"
	"aoi-s/geo/r3"
	"math"
)

type GridManager struct {
	width, height, gridSize   float64
	minX, minZ, maxX, maxZ    float64
	maxRow, maxCol, maxGridID int
	grids                     []Grid
}

func (gm GridManager) LengthInGrid(l float64) int {
	return int(math.Ceil(l / cmap.GridSize))
}

func (gm GridManager) Pos2GridIndex(pos r3.Vector) int {
	row := int(math.Floor(pos.Z-gm.minZ) / gm.gridSize)
	col := int(math.Floor(pos.X-gm.minX) / gm.gridSize)
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

func (gm *GridManager) Init(width, height, gridSize float64) {
	gm.width, gm.height, gm.gridSize = width, height, gridSize
	gm.maxX, gm.maxZ = width/2, height/2
	gm.minX, gm.minZ = gm.maxX, gm.maxZ

	gm.maxRow, gm.maxCol = int(math.Floor(height/gridSize)+1), int(math.Floor(width/gridSize)+1)
	gm.maxGridID = gm.maxRow * gm.maxCol
	gm.grids = make([]Grid, gm.maxGridID)
}

func (gm *GridManager) Release() {
	gm.grids = make([]Grid, gm.maxGridID)
}

func (gm *GridManager) EnterMap(e *Entity) {
	if !gm.IsPosValid(e.Position) {
		return
	}
	e.GridID = gm.Pos2GridIndex(e.Position)
	gm.grids[e.GridID].AddEntity(e)
}

func (gm *GridManager) LeaveMap(e *Entity) {
	if e.GridID < 0 || e.GridID >= gm.maxGridID {
		return
	}
	gm.grids[e.GridID].RemoveEntity(e)
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
	gm.grids[e.GridID].RemoveEntity(e)
	gm.grids[index].AddEntity(e)
}
