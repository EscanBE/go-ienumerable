package goe

func (src *enumerable[T]) Intersect(second IEnumerable[T], optionalCompareFunc CompareFunc[T]) IEnumerable[T] {
	src.assertSrcNonNil()
	assertSecondIEnumerableNonNil(second)

	if optionalCompareFunc == nil {
		defaultComparer := src.defaultComparer
		if defaultComparer == nil {
			defaultComparer = src.findDefaultComparer()
		}
		optionalCompareFunc = func(v1, v2 T) int {
			return defaultComparer.CompareAny(v1, v2)
		}
	}

	return src.internalIntersectBy(second, func(v1, v2 T) bool {
		return optionalCompareFunc(v1, v2) == 0
	})
}

func (src *enumerable[T]) internalIntersectBy(second IEnumerable[T], equalityComparer func(v1, v2 T) bool) IEnumerable[T] {
	src.assertSrcNonNil()
	assertSecondIEnumerableNonNil(second)
	src.assertComparerNonNil(equalityComparer)

	if len(src.data) < 1 || second.Count() < 1 {
		return src.copyExceptData().withEmptyData()
	}

	result := make([]T, 0)
	secondData := second.ToArray()
	for _, fe := range src.data {
		var foundInAnother bool
		for _, se := range secondData {
			if equalityComparer(fe, se) {
				foundInAnother = true
				break
			}
		}
		if foundInAnother {
			var addedPreviously bool

			for _, t := range result {
				if equalityComparer(fe, t) {
					addedPreviously = true
					break
				}
			}

			if !addedPreviously {
				result = append(result, fe)
			}
		}
	}

	return src.copyExceptData().withData(result)
}
