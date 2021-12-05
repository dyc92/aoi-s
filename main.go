package main

import (
	"aoi-s/aoimpl"
	cmap "aoi-s/aoimpl/map"
	"aoi-s/geo/r3"
	"fmt"

	"strconv"
)

func main() {

	cmap := cmap.CMap{}
	cmap.Init("default")
	for i := 0; i < 100; i++ {
		e := &aoimpl.Entity{
			Id:       i + 1,
			Name:     "Entity" + strconv.Itoa(i),
			Type:     aoimpl.Player,
			Position: r3.Vector{X: float64(i * 2), Z: float64(i * 2)},
			MaxSpeed: 5,
		}
		e.GridID = cmap.GridMgr.Pos2GridIndex(e.Position)
		cmap.GridMgr.EnterMap(e)
	}
	e := cmap.GridMgr.GetEntity(99)
	if e != nil {
		fmt.Printf("%+v", e)
	}
}

type aoiAlloc func(ud interface{}, ptr interface{}, sz int) interface{}
type aoiCallback func(ud interface{})
type aoiSpace struct{}
