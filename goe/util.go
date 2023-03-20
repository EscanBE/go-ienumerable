package goe

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe/comparers"
)

func (src *enumerable[T]) copyExceptData() *enumerable[T] {
	if src == nil {
		return nil
	}
	return &enumerable[T]{
		data:            nil,
		dataType:        src.dataType,
		defaultComparer: src.defaultComparer,
	}
}

func (src *enumerable[T]) withData(data []T) *enumerable[T] {
	if src == nil {
		return nil
	}
	src.data = data
	return src
}

func (src *enumerable[T]) withEmptyData() *enumerable[T] {
	if src == nil {
		return nil
	}
	src.data = make([]T, 0)
	return src
}

func copySlice[T any](src []T) []T {
	dst := make([]T, len(src))
	if len(src) > 0 {
		copy(dst, src)
	}
	return dst
}

func (src *enumerable[T]) findDefaultComparer() comparers.IComparer[any] {
	comparer, found := src.tryFindDefaultComparer()
	if found {
		return comparer
	}

	panic(fmt.Errorf("no default comparer registered for [%s]", src.dataType))
}

func (src *enumerable[T]) tryFindDefaultComparer() (comparers.IComparer[any], bool) {
	if comparer, found := comparers.TryGetDefaultComparer[T](); found {
		return comparer, true
	}

	return nil, false
}

func (src *enumerable[T]) injectDefaultComparer() IEnumerable[T] {
	if comparer, found := src.tryFindDefaultComparer(); found {
		src.defaultComparer = comparer
	}

	return src
}

// cast IEnumerable back to *enumerable for accessing private fields.
func e[T any](ie IEnumerable[T]) *enumerable[T] {
	return ie.(*enumerable[T])
}

// cast IEnumerable[T] to IEnumerable[any]
func asIEnumerableAny[T any](ie IEnumerable[T]) IEnumerable[any] {
	if ie == nil {
		return nil
	}
	return ie.Select(func(v T) any {
		return v
	})
}

// Ptr convert a value into pointer form,
// usually used to provide default value for methods like FirstOrDefault, LastOrDefault, SingleOrDefault,...
func Ptr[T any](value T) *T {
	return &value
}
