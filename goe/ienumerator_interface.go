package goe

// IEnumerator supports a simple iteration over a collection.
type IEnumerator[T any] interface {
	// Current gets the element in the collection at the current position of the enumerator.
	Current() T

	// MoveNext advances the enumerator to the next element of the collection.
	MoveNext() bool

	// Reset sets the enumerator to its initial position, which is before the first element in the collection.
	Reset()

	// CurrentSafe gets the element in the collection at the current position of the enumerator,
	// or error if current position is out of bound, or MoveNext has not been called
	CurrentSafe() (T, error)
}
