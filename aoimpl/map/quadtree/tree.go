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

type Quadtree struct {
	_width, _height, slotNodeCount              int
	_maxSlotDepth, _tagMax, _defaultSightRadius int
	_dicNode                                    map[int]*Node
	_dicSlot                                    *Slot
}

func (q *Quadtree) Init(xMin, yMin, xMax, yMax, maxNode, maxDepth, tagMax int) {
	q._width = xMax - xMin
	q._height = yMax - yMin
	q.slotNodeCount = maxNode
	q._maxSlotDepth = maxDepth
	q._tagMax = tagMax
	q._dicNode = make(map[int]*Node)
}

func (q *Quadtree) Add(id, tag int, watchRect, markRect *Rect) bool {
	if _, ok := q._dicNode[id]; ok {
		return false
	}

	node := NewNode(id, tag, watchRect, markRect)
	q._dicNode[id] = node
	q._dicSlot.Add(node)
	return true
}

func (q *Quadtree) Remove(id int) bool {
	if _, ok := q._dicNode[id]; !ok {
		return false
	}

	node := q._dicNode[id]
	q._dicSlot.Remove(node)
	delete(q._dicNode, id)
	return true
}

func (q *Quadtree) Clear() {
	q._width = 0
	q._height = 0
	q.slotNodeCount = 0
	q._maxSlotDepth = 0
	q._tagMax = 0
	q._dicNode = nil
	q._dicSlot = nil
}
