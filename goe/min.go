package goe

import "github.com/EscanBE/go-ienumerable/goe/comparers"

func (src *enumerable[T]) Min() T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

	comparer := src.defaultComparer
	if comparer == nil {
		comparer = src.findDefaultComparer()
	}

	return src.internalMinBy(func(v1, v2 T) bool {
		return comparer.Compare(v1, v2) < 0
	})
}

func (src *enumerable[T]) MinBy(lessComparer func(left, right T) bool) T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

	if lessComparer == nil {
		comparer := src.defaultComparer
		if comparer == nil {
			comparer = src.findDefaultComparer()
		}
		lessComparer = func(v1, v2 T) bool {
			return comparer.Compare(v1, v2) < 0
		}
	}

	return src.internalMinBy(lessComparer)
}

func (src *enumerable[T]) MinByComparer(comparer comparers.IComparer[T]) T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

	if comparer != nil {
		return src.internalMinBy(func(v1, v2 T) bool {
			return comparer.Compare(v1, v2) < 0
		})
	}

	defaultComparer := src.defaultComparer
	if defaultComparer == nil {
		defaultComparer = src.findDefaultComparer()
	}

	return src.internalMinBy(func(v1, v2 T) bool {
		return defaultComparer.Compare(v1, v2) < 0
	})
}

func (src *enumerable[T]) internalMinBy(lessComparer func(left, right T) bool) T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()
	src.assertComparerNonNil(lessComparer)

	min := src.data[0]

	for i := 1; i < len(src.data); i++ {
		ele := src.data[i]
		if lessComparer(ele, min) {
			min = ele
		}
	}

	return min
}
