package aoimpl

type EnumMapModeType int8

const (
	MapModeUnKnown EnumMapModeType = iota // 未知，用于初始化默认值
	MapModeNormal                         // 常规地图
	MapModeClone                          // 副本地图
)

type EnumMapStateType int8

const (
	MapStateUnKnown EnumMapStateType = iota // 位置状态
	MapStateOpen                            // 开启状态
	MapStateClose                           // 关闭状态
	MapStateNormal                          // 普通状态
)

const (
	GridSize    float64 = 100
	MaxKnown    float64 = 0
	MinBigKnown float64 = 50
)
