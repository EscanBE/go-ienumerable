package goe

func (src *enumerable[T]) Except(second IEnumerable[T], optionalCompareFunc CompareFunc[T]) IEnumerable[T] {
	src.assertSrcNonNil()

	if optionalCompareFunc == nil {
		defaultComparer := src.defaultComparer
		if defaultComparer == nil {
			defaultComparer = src.findDefaultComparer()
		}
		optionalCompareFunc = func(v1, v2 T) int {
			return defaultComparer.CompareAny(v1, v2)
		}
	}

	return src.internalExceptBy(second, func(v1, v2 T) bool {
		return optionalCompareFunc(v1, v2) == 0
	})
}

func (src *enumerable[T]) internalExceptBy(second IEnumerable[T], equalityComparer func(v1, v2 T) bool) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertSecondIEnumerableNonNil(second)
	src.assertComparerNonNil(equalityComparer)

	if second.Count() < 1 {
		return src.copyExceptData().withData(copySlice(src.ToArray()))
	}

	if len(src.data) < 1 {
		return src.copyExceptData().withEmptyData()
	}

	result := make([]T, 0)
	secondData := second.ToArray()
	for _, d := range src.data {
		var foundInAnother bool
		for _, t := range secondData {
			if equalityComparer(d, t) {
				foundInAnother = true
				break
			}
		}
		if !foundInAnother {
			var addedPreviously bool

			for _, t := range result {
				if equalityComparer(d, t) {
					addedPreviously = true
					break
				}
			}

			if !addedPreviously {
				result = append(result, d)
			}
		}
	}

	return src.copyExceptData().withData(result)
}
