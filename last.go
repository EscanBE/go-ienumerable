package go_ienumerable

import "fmt"

func (src *enumerable[T]) Last() T {
	result, err := src.LastSafe()
	if err != nil {
		panic(err)
	}
	return result
}

func (src *enumerable[T]) LastBy(predicate func(T) bool) T {
	result, err := src.LastSafeBy(predicate)
	if err != nil {
		panic(err)
	}
	return result
}

func (src *enumerable[T]) LastSafe() (result T, err error) {
	src.assertSrcNonNil()

	if len(src.data) < 1 {
		err = fmt.Errorf("sequence contains no element")
		return
	}

	result = src.data[len(src.data)-1]
	return
}

func (src *enumerable[T]) LastSafeBy(predicate func(T) bool) (result T, err error) {
	src.assertSrcNonNil()
	if predicate == nil {
		err = getErrorNilPredicate()
		return
	}

	if len(src.data) < 1 {
		err = fmt.Errorf("sequence contains no element")
		return
	}

	for i := len(src.data) - 1; i >= 0; i-- {
		if predicate(src.data[i]) {
			result = src.data[i]
			return
		}
	}

	err = fmt.Errorf("no element satisfies the condition in predicate")
	return
}
