package go_ienumerable

func (src *enumerable[T]) FirstOrDefault(defaultValue T) T {
	src.assertSrcNonNil()

	if len(src.data) < 1 {
		return defaultValue
	}

	return src.data[0]
}

func (src *enumerable[T]) FirstOrDefaultBy(predicate func(T) bool, defaultValue T) T {
	result, err := src.FirstOrDefaultSafeBy(predicate, defaultValue)
	if err != nil {
		panic(err)
	}
	return result
}

func (src *enumerable[T]) FirstOrDefaultSafeBy(predicate func(T) bool, defaultValue T) (result T, err error) {
	src.assertSrcNonNil()
	if predicate == nil {
		err = getErrorNilPredicate()
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

	result = defaultValue
	return
}
