package goe_helper

import "github.com/EscanBE/go-ienumerable/goe"

// Repeat generates a sequence that contains one repeated value.
func Repeat[T any](element T, count int) goe.IEnumerable[T] {
	if count < 0 {
		panic("count is less than 0")
	}
	data := make([]T, count)
	for i := 0; i < count; i++ {
		ele := element // copy
		data[i] = ele
	}
	return goe.NewIEnumerable[T](data...)
}
