package main

import (
	"aoi/slicetool"
	"encoding/json"
	"math/rand"
)

func main() {
	var s = make([]int, 1000)

	for i := 0; i < 100; i++ {
		s[i] = rand.Intn(100)
	}
	var r = slicetool.RepeatingElement(s)
	var m, _ = json.Marshal(r)
	print(string(m))
}
