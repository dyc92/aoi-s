package quadtree

import (
	"aoi-s/geo/r2"
	"math"
)

type Rect struct {
	Min              r2.Point
	Max              r2.Point
	Width, Height    float64
	CenterX, CenterY float64
}

func (r *Rect) Init(minX, maxX, minY, maxY float64) {
	r.Min = r2.Point{X: minX, Y: minY}
	r.Max = r2.Point{X: maxX, Y: maxY}
	r.Width = math.Abs(r.Max.X - r.Min.X)
	r.Height = math.Abs(r.Max.Y - r.Min.Y)
	r.CenterX = r.Min.X + r.Width*0.5
	r.CenterY = r.Min.Y + r.Height*0.5
}

func (r Rect) IsCollision(rc *Rect) bool {
	return r.Min.X <= rc.Max.X &&
		r.Max.X >= rc.Min.X &&
		r.Min.Y <= rc.Max.Y &&
		r.Max.Y >= rc.Min.Y
}
