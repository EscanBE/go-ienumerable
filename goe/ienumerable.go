package goe

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe/comparers"
)

// ensure implementation
var _ IEnumerable[any] = &enumerable[any]{}

// Enumerable is a port from C# linq, support many convenient utilities.
//
// Contract: in methods implementation, the inner array data would not be changed in source enumerable, only perform a soft copy of the data array
type enumerable[T any] struct {
	data     []T
	dataType string

	defaultComparer comparers.IComparer[any]
}

// NewIEnumerable returns an IEnumerable with the same type as data elements
func NewIEnumerable[T any](data ...T) IEnumerable[T] {
	dataType := fmt.Sprintf("%T", *new(T))
	if dataType == "<nil>" {
		dataType = ""
	}

	return (&enumerable[T]{
		data:     copySlice(data),
		dataType: dataType,
	}).injectDefaultComparer()
}

// NewIEnumerableFromMap returns an IEnumerable of KeyValuePair elements
func NewIEnumerableFromMap[K comparable, V any](source map[K]V) IEnumerable[KeyValuePair[K, V]] {
	data := make([]KeyValuePair[K, V], len(source))
	cnt := 0
	for k, v := range source {
		data[cnt] = KeyValuePair[K, V]{
			Key:   k,
			Value: v,
		}
		cnt++
	}
	return NewIEnumerable(data...)
}

// Empty returns an empty IEnumerable with specific type
func Empty[T any]() IEnumerable[T] {
	return NewIEnumerable[T]()
}
