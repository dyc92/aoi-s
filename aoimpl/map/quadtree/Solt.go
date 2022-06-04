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

func (s *Slot) Add(n *Node) {
	if n.BoundWatch == nil {
		return
	}
	quadType := CheckQuadType(s.bound, n.BoundWatch)
	if s.child != nil && quadType != quadrantNone {
		s.child[quadType].Add(n)
	}
}
func (s *Slot) Remove(n *Node) {
	quadType := CheckQuadType(s.bound, n.BoundWatch)
	s.child[quadType].Remove(n)
}

type Node struct {
	Id, Tag                                  int
	Alias                                    string
	BoundWatch, BoundMark                    *Rect
	OwnerLinkedItemWatch, OwnerLikedItemMark *list.List
	QuadrantWatch, QuadrantMark              EnumQuadrantType
}

func NewNode(id, tag int, wRc, mRc *Rect) *Node {
	return &Node{
		BoundWatch:           wRc,
		BoundMark:            mRc,
		OwnerLinkedItemWatch: nil,
		OwnerLikedItemMark:   nil,
		QuadrantWatch:        quadrantNone,
		QuadrantMark:         quadrantNone,
	}
}
