package go_ienumerable

func (src *enumerable[T]) Distinct() IEnumerable[T] {
	src.assertSrcNonNil()

	if src.equalityComparer == nil {
		panicRequire(requireEqualityComparer)
	}

	return src.DistinctBy(src.equalityComparer)
}

func (src *enumerable[T]) DistinctBy(equalityComparer func(v1, v2 T) bool) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertComparerNonNil(equalityComparer)

	if len(src.data) < 1 {
		return src.copyExceptData().withEmptyData()
	}

	if len(src.data) < 2 {
		return src.copyExceptData().withData(copySlice(src.data))
	}

	uniqueSet := []T{src.data[0]}

	for i1 := 1; i1 < len(src.data); i1++ {
		ele := src.data[i1]

		var exists bool
		for _, uniq := range uniqueSet {
			if equalityComparer(ele, uniq) {
				exists = true
				break
			}
		}

		if !exists {
			uniqueSet = append(uniqueSet, ele)
		}
	}

	return src.copyExceptData().withData(uniqueSet)
}
