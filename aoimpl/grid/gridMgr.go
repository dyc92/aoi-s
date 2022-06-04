package aoigrid

import (
	. "aoi-s/aoimpl"
	"aoi-s/geo/r3"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

type GridManager struct {
	grids     []*Grid
	entityMap sync.Map

	width, height             float64
	gridSize                  float64
	minX, minZ, maxX, maxZ    float64
	maxRow, maxCol, maxGridID int
}

func InitGridManager(width, height float64, gridSize float64) (gm *GridManager) {
	gm = &GridManager{
		width:     width,
		height:    height,
		gridSize:  gridSize,
		maxX:      width / 2,
		maxZ:      height / 2,
		minX:      -width / 2,
		minZ:      -height / 2,
		maxRow:    int(math.Ceil(height / gridSize)),
		maxCol:    int(math.Ceil(width / gridSize)),
		entityMap: sync.Map{},
	}

	gm.maxGridID = gm.maxRow * gm.maxCol
	gm.grids = make([]*Grid, 0, gm.maxGridID)
	for i := 0; i < gm.maxGridID; i++ {
		g := &Grid{Id: i}
		g.Init()
		gm.grids = append(gm.grids, g)
	}

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

func (gm *GridManager) Release() {
	str := strings.Builder{}
	str.Grow(1000)
	for i := 0; i < len(gm.grids); i++ {
		es := gm.grids[i].GetEntities(Player)
		str.WriteString("index:+" + strconv.Itoa(i) + ",count:" + strconv.Itoa(len(es)) + "\r\n")
	}
	fmt.Println(str.String())
	gm.grids = make([]*Grid, gm.maxGridID)
	gm.entityMap = sync.Map{}
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
	e.SkipStep--
	if e.SkipStep > 0 {
		return
	}
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
