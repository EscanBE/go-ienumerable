package goe_helper

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe"
	"math"
)

// Range generates a sequence of int32 numbers within a specified range.
func Range(start, count int) goe.IEnumerable[int] {
	if count < 0 {
		panic("count is less than 0")
	}
	if start+count-1 > math.MaxInt32 {
		panic(fmt.Sprintf("can not larger than %d", math.MaxInt32))
	}
	data := make([]int, count)
	for i := 0; i < count; i++ {
		data[i] = start + i
	}
	return goe.NewIEnumerable[int](data...)
}
