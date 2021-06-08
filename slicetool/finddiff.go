package slicetool

import (
	"encoding/json"
	"math/rand"
	"testing"
)

// RepeatingElement 重复元素出现次数(查重多个数组也类似，dic计数>1为交集，=1为差集)
func RepeatingElement(arr []int) map[int]int {

	var dic = make(map[int]int)
	for _, v := range arr {
		if _,ok:= dic[v];ok {
			dic[v]++
		}else {
			dic[v]=1
		}
	}
	return dic
}



func TestRE(t *testing.T) {
	var s = make([]int,1000)

	for i := 0; i < 1000; i++ {
		s[i]=rand.Intn(100)
	}
	var r= RepeatingElement(s)
	var m,_ =json.Marshal(r)
	t.Log(m)
}