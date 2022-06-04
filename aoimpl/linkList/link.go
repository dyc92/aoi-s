package linkList

import (
	. "aoi-s/aoimpl"
)

type LinkMgr struct {
	LinkX []*Link
	LinkY []*Link
	Size  float64
}

func (m *LinkMgr) Init() {

}

type Link struct {
	Entities []*Entity
}

func (l *Link) Sort() {

}
