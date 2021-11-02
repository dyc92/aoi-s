package quadtree

import (
	"testing"
	"unsafe"
)

type TName struct {
	A int64
	B int8
	C bool
}
type TBame struct {
	B int8
	A int64
	C bool
}

func TestMemRise(t *testing.T) {
	t.Log(unsafe.Alignof(TName{}))
	t.Log(unsafe.Alignof(TBame{}))

	tn := TName{}
	t.Log(unsafe.Sizeof(tn))

	tb := TBame{}
	t.Log(unsafe.Sizeof(tb))

}
