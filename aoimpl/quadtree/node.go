package quadtree

import "container/list"

type Node struct {
	Id, Tag                                  int
	Alias                                    string
	BoundWatch, BoundMark                    *Rect
	OwnerLinkedItemWatch, OwnerLikedItemMark list.List
	QuadrantWatch, QuadrantMark              EnumQuadrantType
}
