package goe_helper

import "github.com/EscanBE/go-ienumerable/goe"

// Enumerable returns a new instance of EnumerableHelper[T]
// which provide utility methods
func Enumerable[T any]() EnumerableHelper[T] {
	return EnumerableHelper[T]{}
}

// EnumerableHelper provides utility methods like Empty, Repeat,...
type EnumerableHelper[T any] struct {
}

// Empty returns an empty IEnumerable[T] that has the specified type argument.
func (e EnumerableHelper[T]) Empty() goe.IEnumerable[T] {
	return goe.NewIEnumerable[T]()
}

// Repeat generates a sequence that contains one repeated value.
func (e EnumerableHelper[T]) Repeat(element T, count int) goe.IEnumerable[T] {
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
