package goe

func (src *enumerable[T]) Except(second IEnumerable[T]) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertSecondIEnumerableNonNil(second)

	if src.equalityComparer == nil {
		panicRequire(requireEqualityComparer)
	}

	return src.ExceptBy(second, src.equalityComparer)
}

func (src *enumerable[T]) ExceptBy(second IEnumerable[T], equalityComparer func(v1, v2 T) bool) IEnumerable[T] {
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
			result = append(result, d)
		}
	}

	return src.copyExceptData().withData(result)
}
