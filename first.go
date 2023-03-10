package go_ienumerable

import "fmt"

func (src *enumerable[T]) First() T {
	result, err := src.FirstSafe()
	if err != nil {
		panic(err)
	}
	return result
}

func (src *enumerable[T]) FirstBy(predicate func(T) bool) T {
	result, err := src.FirstSafeBy(predicate)
	if err != nil {
		panic(err)
	}
	return result
}

func (src *enumerable[T]) FirstSafe() (result T, err error) {
	src.assertSrcNonNil()

	if len(src.data) < 1 {
		err = fmt.Errorf("sequence contains no element")
		return
	}

	result = src.data[0]
	return
}

func (src *enumerable[T]) FirstSafeBy(predicate func(T) bool) (result T, err error) {
	src.assertSrcNonNil()
	if predicate == nil {
		err = getErrorNilPredicate()
		return
	}

	if len(src.data) < 1 {
		err = fmt.Errorf("sequence contains no element")
		return
	}

	if len(src.data) > 0 {
		for _, d := range src.data {
			if predicate(d) {
				result = d
				return
			}
		}
	}

	err = fmt.Errorf("no element satisfies the condition in predicate")
	return
}
