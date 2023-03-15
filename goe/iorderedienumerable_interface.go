package goe

type IOrderedEnumerable[T any] interface {
	// ThenBy performs a subsequent ordering of the elements in a sequence in ascending order according to provided comparer.
	ThenBy(compareFuncOrComparer interface{}) IOrderedEnumerable[T]

	// ThenByDescending performs a subsequent ordering of the elements in a sequence in descending order according to provided comparer.
	ThenByDescending(compareFuncOrComparer interface{}) IOrderedEnumerable[T]

	// GetOrderedEnumerable performs ordering based on provided-chained comparers and return result as IEnumerable[T]
	GetOrderedEnumerable() IEnumerable[T]
}
