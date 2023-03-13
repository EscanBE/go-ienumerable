package goe

type IOrderedEnumerable[T any] interface {
	// ThenBy performs a subsequent ordering of the elements in a sequence in ascending order according to provided comparer.
	ThenBy(comparer func(v1, v2 T) int) IOrderedEnumerable[T]

	// ThenByDescending performs a subsequent ordering of the elements in a sequence in descending order according to provided comparer.
	ThenByDescending(comparer func(v1, v2 T) int) IOrderedEnumerable[T]

	// GetEnumerable performs ordering based on provided-chained comparers and return result as IEnumerable[T]
	GetEnumerable() IEnumerable[T]
}
