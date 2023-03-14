package goe

func (src *enumerable[T]) Union(second IEnumerable[T]) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertSecondIEnumerableNonNil(second)

	if src.equalityComparer == nil {
		panicRequire(requireEqualityComparer)
	}

	return src.UnionBy(second, src.equalityComparer)
}

func (src *enumerable[T]) UnionBy(second IEnumerable[T], equalityComparer func(v1, v2 T) bool) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertSecondIEnumerableNonNil(second)
	src.assertComparerNonNil(equalityComparer)

	if len(src.data) < 1 && second.Count() < 1 {
		return src.copyExceptData().withEmptyData()
	}

	result := src.copyExceptData()

	if second.Count() < 1 {
		return result.withData(copySlice(src.ToArray())).internalDistinctBy(equalityComparer)
	}

	if len(src.data) < 1 {
		return result.withData(copySlice(second.ToArray())).internalDistinctBy(equalityComparer)
	}

	return result.withData(append(copySlice(src.data), copySlice(second.ToArray())...)).internalDistinctBy(equalityComparer)
}
