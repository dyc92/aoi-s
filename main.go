package main

import (
	"aoi-s/common"
	"fmt"
	"github.com/golang/geo/s1"
)

func main() {

	v:=common.Vector{X: 1.0, Y: 1.0, Z: 1.0}
	angle:=geo.Angle(v,common.Vector{X: 2.0, Y: 2.0, Z: 2.0})

	fmt.Println(angle)
}

type aoiAlloc func(ud interface{}, ptr interface{}, sz int) interface{}
type aoiCallback func(ud interface{})
type aoiSpace struct{}
