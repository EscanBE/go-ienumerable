package goe

type IOrderedEnumerable[T any] interface {
	// ThenBy performs a subsequent ordering of the elements in a sequence in ascending order
	// according to the key selector and compareFunc for the selected keys.
	//
	// If omitted the compareFunc, the default comparer for corresponding type will be used,
	// or panic if no default compare found.
	ThenBy(keySelector KeySelector[T], compareFunc CompareFunc[any]) IOrderedEnumerable[T]

	// ThenByDescending performs a subsequent ordering of the elements in a sequence in descending order
	// according to provided comparer and compareFunc for the selected keys.
	//
	// If omitted the compareFunc, the default comparer for corresponding type will be used,
	// or panic if no default compare found.
	ThenByDescending(keySelector KeySelector[T], compareFunc CompareFunc[any]) IOrderedEnumerable[T]

	// GetOrderedEnumerable performs ordering based on provided-chained comparers and return result as IEnumerable[T]
	GetOrderedEnumerable() IEnumerable[T]
}
