package go_ienumerable

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
