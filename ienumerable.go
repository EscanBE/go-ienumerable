package go_ienumerable

var _ IEnumerable[any] = &enumerable[any]{}

// Enumerable is a port from C# linq, support many convenient utilities
//
// Contract: in method implementation, the inner array data would not be changed in source enumerable, only perform a soft copy of the data array
type enumerable[T any] struct {
	data []T

	equalityComparer func(d1, d2 T) bool
	lessComparer     func(d1, d2 T) bool
}

// NewIEnumerable returns an IEnumerable with the same time as data elements
func NewIEnumerable[T any](data ...T) IEnumerable[T] {
	return &enumerable[T]{
		data: copySlice(data),
	}
}

// Empty returns an empty IEnumerable with specific type
func Empty[T any]() IEnumerable[T] {
	return &enumerable[T]{
		data: make([]T, 0),
	}
}
