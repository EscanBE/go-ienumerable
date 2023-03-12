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

	if len(src.data) < 1 && second.len() < 1 {
		return src.copyExceptData().withEmptyData()
	}

	result := src.copyExceptData()

	if second.len() < 1 {
		return result.withData(copySlice(src.exposeData())).internalDistinctBy(equalityComparer)
	}

	if len(src.data) < 1 {
		return result.withData(copySlice(second.exposeData())).internalDistinctBy(equalityComparer)
	}

	return result.withData(append(copySlice(src.data), copySlice(second.exposeData())...)).internalDistinctBy(equalityComparer)
}
