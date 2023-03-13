package goe

func (src *enumerable[T]) Intersect(second IEnumerable[T]) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertSecondIEnumerableNonNil(second)

	if src.equalityComparer == nil {
		panicRequire(requireEqualityComparer)
	}

	return src.IntersectBy(second, src.equalityComparer)
}

func (src *enumerable[T]) IntersectBy(second IEnumerable[T], equalityComparer func(v1, v2 T) bool) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertSecondIEnumerableNonNil(second)
	src.assertComparerNonNil(equalityComparer)

	if len(src.data) < 1 || second.len() < 1 {
		return src.copyExceptData().withEmptyData()
	}

	result := make([]T, 0)
	secondData := second.exposeData()
	for _, fe := range src.data {
		var foundInAnother bool
		for _, se := range secondData {
			if equalityComparer(fe, se) {
				foundInAnother = true
				break
			}
		}
		if foundInAnother {
			result = append(result, fe)
		}
	}

	return src.copyExceptData().withData(result)
}
