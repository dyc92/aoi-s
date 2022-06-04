package main

import (
	"aoi-s/aoimpl"
	cmap "aoi-s/aoimpl/map"
	"aoi-s/geo/r3"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	gameMap := cmap.Map{}
	gameMap.Init("StormWall")
	random := rand.New(rand.NewSource(time.Now().UTC().Unix()))

	for i := 1; i < 1001; i++ {
		entity := &aoimpl.Entity{
			Id:       100000 + i*100,
			Name:     "Role" + strconv.Itoa(i),
			Type:     aoimpl.Player,
			Position: r3.Vector{X: float64(random.Intn(1000)), Z: float64(random.Intn(1000))},
		}
		gameMap.EnterMap(entity)
	}
	gameMap.Print()
}

type aoiAlloc func(ud interface{}, ptr interface{}, sz int) interface{}
type aoiCallback func(ud interface{})
type aoiSpace struct{}
