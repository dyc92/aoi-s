package cmap

import "aoi-s/aoimpl/map/quadtree"

type IAoiScan interface {
	/// <summary> self的视野与other的可见碰撞 </summary>
	onCollidedSightWithKnown(n quadtree.Node) bool
	/// <summary> self的可见与other的视野碰撞 </summary>
	onCollidedKnownWithSight(n quadtree.Node) bool
}

type AoiScan struct {
	_map                         *Map
	_sight                       quadtree.Rect
	_known                       quadtree.Rect
	MaxDepth                     int
	HitCount                     int
	SightGrid, KnownGrid         int
	GridID                       int
	MaxGridRadius, MinGridRadius int
}
