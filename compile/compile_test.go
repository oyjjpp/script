package compile

import (
	"testing"
	"unsafe"
)

func TestArraySize(t *testing.T) {
	data := [200]interface{}{}
	rs := unsafe.Sizeof(data)
	t.Log(rs)

	dataByte := [200]byte{}
	rsByte := unsafe.Sizeof(dataByte)
	t.Log(rsByte)
}

// go:noinline
func newArray() *[4]int {
	a := [4]int{1, 2, 3, 4}
	return &a
}

func TestNewArray(t *testing.T) {
	rs := newArray()
	t.Log(rs)
}
