package common

//轴对齐碰撞箱
type A3DAABB struct {
	Center, //坐标中心
	Extents, //线段 箱形体边长
	Mins,
	Maxs Vector
}

func NewBox() A3DAABB {
	box := A3DAABB{}
	box.Clear()
	return box
}

func (a *A3DAABB) Clear() {
	a.Maxs.Set(999999.0, 999999.0, 999999.0)
	a.Maxs.Set(-999999.0, -999999.0, -999999.0)
}

func (a *A3DAABB) CompleteCenterExts() {
	a.Center = a.Mins.Add(a.Maxs).Mul(0.5)
	a.Extents = a.Maxs.Sub(a.Center)
}

func (a *A3DAABB) IsValid() bool {
	return a.Mins.X <= a.Maxs.X && a.Mins.Y <= a.Maxs.Y && a.Mins.Z <= a.Maxs.Z
}
