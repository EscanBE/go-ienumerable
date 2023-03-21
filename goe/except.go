package goe

func (src *enumerable[T]) Except(second IEnumerable[T], optionalEqualsFunc OptionalEqualsFunc[T]) IEnumerable[T] {
	src.assertSrcNonNil()
	assertSecondIEnumerableNonNil(second)

	var isEquals EqualsFunc[T]
	if optionalEqualsFunc == nil {
		defaultComparer := src.defaultComparer
		if defaultComparer == nil {
			defaultComparer = src.findDefaultComparer()
		}
		isEquals = func(v1, v2 T) bool {
			return defaultComparer.CompareAny(v1, v2) == 0
		}
	} else {
		isEquals = EqualsFunc[T](optionalEqualsFunc)
	}

	if second.Count(nil) < 1 {
		return src.copyExceptData().withData(distinct(copySlice(src.data), OptionalEqualsFunc[T](isEquals)))
	}

	if len(src.data) < 1 {
		return src.copyExceptData().withEmptyData()
	}

	result := make([]T, 0)
	secondData := second.ToArray()
	for _, d := range src.data {
		var foundInAnother bool
		for _, t := range secondData {
			if isEquals(d, t) {
				foundInAnother = true
				break
			}
		}
		if !foundInAnother {
			var addedPreviously bool

			for _, t := range result {
				if isEquals(d, t) {
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
