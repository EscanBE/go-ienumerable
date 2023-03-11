package go_ienumerable

import "fmt"

var _ IEnumerator[any] = &enumerator[any]{}

type enumerator[T any] struct {
	data     []T
	position int
}

// NewIEnumerator returns a new IEnumerator instance from source slice
func NewIEnumerator[T any](source ...T) IEnumerator[T] {
	return &enumerator[T]{
		data:     copySlice(source),
		position: -1,
	}
}

// Current is implementation for IEnumerator
func (col *enumerator[T]) Current() T {
	col.assertCollectionNonNil()

	result, err := col.CurrentSafe()
	if err != nil {
		panic(err)
	}
	return result
}

// MoveNext is implementation for IEnumerator
func (col *enumerator[T]) MoveNext() bool {
	col.assertCollectionNonNil()

	if col.position >= len(col.data)-1 {
		return false
	}
	col.position++
	return true
}

// Reset is implementation for IEnumerator
func (col *enumerator[T]) Reset() {
	col.assertCollectionNonNil()

	col.position = -1
}

// CurrentSafe is implementation for an extra method of IEnumerator
func (col *enumerator[T]) CurrentSafe() (result T, err error) {
	col.assertCollectionNonNil()
	
	if col.position < 0 || col.position > len(col.data)-1 {
		err = fmt.Errorf("invalid operation")
		return
	}

	result = col.data[col.position]
	return
}
