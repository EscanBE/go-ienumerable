package go_ienumerable

import "fmt"

// ensure implementation
var _ IEnumerable[any] = &enumerable[any]{}

// Enumerable is a port from C# linq, support many convenient utilities.
//
// Contract: in methods implementation, the inner array data would not be changed in source enumerable, only perform a soft copy of the data array
type enumerable[T any] struct {
	data     []T
	dataType string

	equalityComparer func(d1, d2 T) bool
	lessComparer     func(d1, d2 T) bool
}

// NewIEnumerable returns an IEnumerable with the same time as data elements
func NewIEnumerable[T any](data ...T) IEnumerable[T] {
	dataType := fmt.Sprintf("%T", *new(T))
	if dataType == "<nil>" {
		dataType = ""
	}
	return &enumerable[T]{
		data:     copySlice(data),
		dataType: dataType,
	}
}

// Empty returns an empty IEnumerable with specific type
func Empty[T any]() IEnumerable[T] {
	return NewIEnumerable[T]()
}
