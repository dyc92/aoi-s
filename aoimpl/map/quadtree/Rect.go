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

func NewRect(minX, maxX, minY, maxY float64) *Rect {
	r := &Rect{}
	r.Min = r2.Point{X: minX, Y: minY}
	r.Max = r2.Point{X: maxX, Y: maxY}
	r.Width = math.Abs(r.Max.X - r.Min.X)
	r.Height = math.Abs(r.Max.Y - r.Min.Y)
	r.CenterX = r.Min.X + r.Width*0.5
	r.CenterY = r.Min.Y + r.Height*0.5
	return r
}

func (r Rect) IsCollision(rc *Rect) bool {
	return r.Min.X <= rc.Max.X &&
		r.Max.X >= rc.Min.X &&
		r.Min.Y <= rc.Max.Y &&
		r.Max.Y >= rc.Min.Y
}

func CheckQuadType(r1, r2 *Rect) (t EnumQuadrantType) {
	top := r2.Min.Y >= r1.CenterY
	bottom := r2.Max.Y <= r1.CenterY
	left := r2.Min.X <= r1.CenterX
	right := r2.Max.X >= r1.CenterX
	if top && left {
		if !left && right {
			t = quadrant1
		} else if left && !right {
			t = quadrant2
		}
	} else if !top && bottom {
		if !left && right {
			t = quadrant4
		} else if left && !right {
			t = quadrant3
		}
	}
	return t
}

func SplitQuadRect(rc *Rect, quadrantType EnumQuadrantType) *Rect {
	minX, minY, maxX, maxY := rc.Min.X, rc.Min.Y, rc.Max.X, rc.Max.Y
	hw := rc.Width * 0.5
	hh := rc.Height * 0.5

	switch quadrantType {
	case quadrant1:
		minX += hw
		minY += hh
	case quadrant2:
		maxX -= hw
		minY += hh
	case quadrant3:
		maxX -= hw
		maxY -= hh
	case quadrant4:
		minX += hw
		maxY -= hh
	default:
		return nil
	}
	return NewRect(minX, maxX, minY, maxY)
}
