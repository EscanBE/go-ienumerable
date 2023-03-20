package goe

func (src *enumerable[T]) Distinct(optionalCompareFunc CompareFunc[T]) IEnumerable[T] {
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

	return src.internalDistinctBy(func(v1, v2 T) bool {
		return optionalCompareFunc(v1, v2) == 0
	})
}

func (src *enumerable[T]) internalDistinctBy(equalityComparer func(v1, v2 T) bool) IEnumerable[T] {
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
