package goe

func (src *enumerable[T]) Intersect(second IEnumerable[T], optionalEqualsFunc OptionalEqualsFunc[T]) IEnumerable[T] {
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

	result := src.copyExceptData()

	if len(src.data) < 1 || second.Count(nil) < 1 {
		result = result.withEmptyData()
	} else {
		intersect := make([]T, 0)
		secondData := second.ToArray()

		for _, fe := range src.data {
			for _, se := range secondData {
				if isEquals(fe, se) {
					intersect = append(intersect, fe)
					break
				}
			}
		}

		result = result.withData(distinct(intersect, OptionalEqualsFunc[T](isEquals)))
	}

	return result
}

func (src *enumerable[T]) IntersectBy(second IEnumerable[T], keySelector KeySelector[T], optionalEqualsFunc OptionalEqualsFunc[any]) IEnumerable[T] {
	src.assertSrcNonNil()
	assertSecondIEnumerableNonNil(second)
	assertKeySelectorNonNil(keySelector)

	return nil
}
