package quadtree

import "container/list"

type Slot struct {
	tree     *Quadtree
	parent   *Slot
	bound    *Rect //边界
	quadType EnumQuadrantType
	depth    int64
	child    []*Slot
	leaf     list.List
}

func (s *Slot) Init() {
	s.child = make([]*Slot, 4)
}

func (s *Slot) Insert(n *Node) {
	parseQuad(s.bound, n.BoundWatch)
}

func parseQuad(bound, node *Rect) EnumQuadrantType {

	top := node.Min.Y >= bound.CenterY
	bottom := node.Max.Y <= bound.CenterY
	left := node.Max.X <= bound.CenterX
	right := node.Min.X >= bound.CenterX

	if top && !bottom {
		if !left && right {
			return quadrant1
		} else if left && !right {
			return quadrant2
		}
	} else if !top && bottom {
		if !left && right {
			return quadrant4
		} else if left && !right {
			return quadrant3
		}
	}
	return quadrantNone
}
