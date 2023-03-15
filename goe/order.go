package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"sort"
)

func (src *enumerable[T]) Order() IEnumerable[T] {
	src.assertSrcNonNil()

	comparer := src.defaultComparer
	if comparer == nil {
		comparer = src.findDefaultComparer()
	}

	return src.internalOrderBy(func(v1, v2 T) bool {
		return comparer.Compare(v1, v2) < 0
	})
}

func (src *enumerable[T]) OrderBy(lessComparer func(left, right T) bool) IEnumerable[T] {
	src.assertSrcNonNil()

	if lessComparer == nil {
		comparer := src.defaultComparer
		if comparer == nil {
			comparer = src.findDefaultComparer()
		}
		lessComparer = func(v1, v2 T) bool {
			return comparer.Compare(v1, v2) < 0
		}
	}

	return src.internalOrderBy(lessComparer)
}

func (src *enumerable[T]) OrderByComparer(comparer comparers.IComparer[T]) IEnumerable[T] {
	src.assertSrcNonNil()

	if comparer != nil {
		return src.internalOrderBy(func(v1, v2 T) bool {
			return comparer.Compare(v1, v2) < 0
		})
	}

	defaultComparer := src.defaultComparer
	if defaultComparer == nil {
		defaultComparer = src.findDefaultComparer()
	}

	return src.internalOrderBy(func(v1, v2 T) bool {
		return defaultComparer.Compare(v1, v2) < 0
	})
}

func (src *enumerable[T]) internalOrderBy(lessComparer func(left, right T) bool) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertComparerNonNil(lessComparer)

	if len(src.data) < 1 {
		return src.copyExceptData().withEmptyData()
	}

	copied := copySlice(src.data)

	sort.SliceStable(copied, func(i, j int) bool {
		return lessComparer(copied[i], copied[j])
	})

	return src.copyExceptData().withData(copied)
}

func (src *enumerable[T]) OrderByDescending() IEnumerable[T] {
	src.assertSrcNonNil()

	if src.lessComparer == nil {
		panicRequire(requireLessComparer)
	}

	return src.OrderByDescendingBy(src.lessComparer)
}

func (src *enumerable[T]) OrderByDescendingBy(lessComparer func(left, right T) bool) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertComparerNonNil(lessComparer)

	if len(src.data) < 1 {
		return src.copyExceptData().withEmptyData()
	}

	copied := copySlice(src.data)

	sort.SliceStable(copied, func(i, j int) bool {
		return lessComparer(copied[j], copied[i])
	})

	return src.copyExceptData().withData(copied)
}
