package goe

func (src *enumerable[T]) FirstOrDefault() T {
	src.assertSrcNonNil()

	if len(src.data) < 1 {
		return *new(T)
	}

	return src.data[0]
}

func (src *enumerable[T]) FirstOrDefaultBy(predicate func(T) bool) T {
	src.assertSrcNonNil()
	src.assertPredicateNonNil(predicate)

	if len(src.data) > 0 {
		for _, d := range src.data {
			if predicate(d) {
				return d
			}
		}
	}

	return *new(T)
}

func (src *enumerable[T]) FirstOrDefaultUsing(defaultValue T) T {
	src.assertSrcNonNil()

	if len(src.data) < 1 {
		return defaultValue
	}

	return src.data[0]
}

func (src *enumerable[T]) FirstOrDefaultByUsing(predicate func(T) bool, defaultValue T) T {
	src.assertSrcNonNil()
	src.assertPredicateNonNil(predicate)

	if len(src.data) > 0 {
		for _, d := range src.data {
			if predicate(d) {
				return d
			}
		}
	}

	return defaultValue
}
