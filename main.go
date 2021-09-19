package main

import (
	"aoi-s/geo/r3"
	"fmt"
)

func main() {

	v := r3.Vector{X: 1.0, Y: 1.0, Z: 1.0}
	angle := v.Angle(r3.Vector{X: 2.0, Y: 2.0, Z: 2.0})

	fmt.Println(angle)
}

type aoiAlloc func(ud interface{}, ptr interface{}, sz int) interface{}
type aoiCallback func(ud interface{})
type aoiSpace struct{}
