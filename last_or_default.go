package go_ienumerable

func (src *enumerable[T]) LastOrDefault(defaultValue T) T {
	src.assertSrcNonNil()

	if len(src.data) < 1 {
		return defaultValue
	}

	return src.data[len(src.data)-1]
}

func (src *enumerable[T]) LastOrDefaultBy(predicate func(T) bool, defaultValue T) T {
	result, err := src.LastOrDefaultSafeBy(predicate, defaultValue)
	if err != nil {
		panic(err)
	}
	return result
}

func (src *enumerable[T]) LastOrDefaultSafeBy(predicate func(T) bool, defaultValue T) (result T, err error) {
	src.assertSrcNonNil()
	if predicate == nil {
		err = getErrorNilPredicate()
		return
	}

	if len(src.data) > 0 {
		for i := len(src.data) - 1; i >= 0; i-- {
			if predicate(src.data[i]) {
				result = src.data[i]
				return
			}
		}
	}

	result = defaultValue
	return
}
