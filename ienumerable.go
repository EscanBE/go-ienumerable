package go_ienumerable

var _ IEnumerable[any] = &enumerable[any]{}

// Enumerable is a port from C# linq, support many convenient utilities
//
// Contract: in method implementation, the inner array data would not be changed in source enumerable, only perform a soft copy of the data array
type enumerable[T any] struct {
	data []T
}

func NewIEnumerable[T any](data ...T) IEnumerable[T] {
	return &enumerable[T]{
		data: copySlice(data),
	}
}
