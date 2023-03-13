package goe

func (src *enumerable[T]) LastOrDefault() T {
	src.assertSrcNonNil()

	if len(src.data) < 1 {
		return *new(T)
	}

	return src.data[len(src.data)-1]
}

func (src *enumerable[T]) LastOrDefaultBy(predicate func(T) bool) T {
	src.assertSrcNonNil()
	src.assertPredicateNonNil(predicate)

	if len(src.data) > 0 {
		for i := len(src.data) - 1; i >= 0; i-- {
			if predicate(src.data[i]) {
				return src.data[i]
			}
		}
	}

	return *new(T)
}

func (src *enumerable[T]) LastOrDefaultUsing(defaultValue T) T {
	src.assertSrcNonNil()

	if len(src.data) < 1 {
		return defaultValue
	}

	return src.data[len(src.data)-1]
}

func (src *enumerable[T]) LastOrDefaultByUsing(predicate func(T) bool, defaultValue T) T {
	src.assertSrcNonNil()
	src.assertPredicateNonNil(predicate)

	if len(src.data) > 0 {
		for i := len(src.data) - 1; i >= 0; i-- {
			if predicate(src.data[i]) {
				return src.data[i]
			}
		}
	}

	return defaultValue
}