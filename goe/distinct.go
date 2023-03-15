package goe

import "github.com/EscanBE/go-ienumerable/goe/comparers"

func (src *enumerable[T]) Distinct() IEnumerable[T] {
	src.assertSrcNonNil()

	comparer := src.defaultComparer
	if comparer == nil {
		comparer = src.findDefaultComparer()
	}

	return src.internalDistinctBy(func(v1, v2 T) bool {
		return comparer.Compare(v1, v2) == 0
	})
}

func (src *enumerable[T]) DistinctBy(equalityComparer func(v1, v2 T) bool) IEnumerable[T] {
	if equalityComparer == nil {
		comparer := src.defaultComparer
		if comparer == nil {
			comparer = src.findDefaultComparer()
		}
		equalityComparer = func(v1, v2 T) bool {
			return comparer.Compare(v1, v2) == 0
		}
	}

	return src.internalDistinctBy(equalityComparer)
}

func (src *enumerable[T]) DistinctByComparer(comparer comparers.IComparer[T]) IEnumerable[T] {
	src.assertSrcNonNil()

	if comparer != nil {
		return src.internalDistinctBy(func(v1, v2 T) bool {
			return comparer.Compare(v1, v2) == 0
		})
	}

	defaultComparer := src.defaultComparer
	if defaultComparer == nil {
		defaultComparer = src.findDefaultComparer()
	}

	return src.internalDistinctBy(func(v1, v2 T) bool {
		return defaultComparer.Compare(v1, v2) == 0
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
