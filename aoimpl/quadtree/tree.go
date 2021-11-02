package quadtree

//### 原四叉树每tick需要清空再播入。可以使用slice和index，只用插入不清空。单个删除赋nil#
/* 象限划分
 * 横轴: x, 增长方向:向右
 * 纵轴: y, 增长方向:向下
 *      -
 *    3 | 4
 *  - ----- +
 *    2 | 1
 *      +
 */

type EnumQuadrantType int8

const (
	quadrantNone EnumQuadrantType = iota //未被任何象限包含
	quadrant1
	quadrant2
	quadrant3
	quadrant4
)

const (
	watcher = iota
	marker
	all
)

type Quadtree struct {
	_width, _height, _maxSoltNode, _maxSoltDepth, _tagMax, _defaultSightRadius int
	_dicNode                                                                   map[int]*Node
}

func (q Quadtree) Init(xMin, yMin, xMax, yMax, maxNode, maxDepth, tagMax, sightRadius int) {
	q._width = xMax - xMin
	q._height = yMax - yMin
	q._maxSoltNode = maxNode
	q._maxSoltDepth = maxDepth
	q._tagMax = tagMax
	q._dicNode = make(map[int]*Node)

}
