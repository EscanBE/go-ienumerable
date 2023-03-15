package goe

import "github.com/EscanBE/go-ienumerable/goe/comparers"

func (src *enumerable[T]) Max() T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

	comparer := src.defaultComparer
	if comparer == nil {
		comparer = src.findDefaultComparer()
	}

	return src.internalMaxBy(func(v1, v2 T) bool {
		return comparer.Compare(v1, v2) > 0
	})
}

func (src *enumerable[T]) MaxBy(greaterComparer func(left, right T) bool) T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

	if greaterComparer == nil {
		comparer := src.defaultComparer
		if comparer == nil {
			comparer = src.findDefaultComparer()
		}
		greaterComparer = func(v1, v2 T) bool {
			return comparer.Compare(v1, v2) > 0
		}
	}

	return src.internalMaxBy(greaterComparer)
}

func (src *enumerable[T]) MaxByComparer(comparer comparers.IComparer[T]) T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

	if comparer != nil {
		return src.internalMaxBy(func(v1, v2 T) bool {
			return comparer.Compare(v1, v2) > 0
		})
	}

	defaultComparer := src.defaultComparer
	if defaultComparer == nil {
		defaultComparer = src.findDefaultComparer()
	}

	return src.internalMaxBy(func(v1, v2 T) bool {
		return defaultComparer.Compare(v1, v2) > 0
	})
}

func (src *enumerable[T]) internalMaxBy(greaterComparer func(left, right T) bool) T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()
	src.assertComparerNonNil(greaterComparer)

	Max := src.data[0]

	for i := 1; i < len(src.data); i++ {
		ele := src.data[i]
		if greaterComparer(ele, Max) {
			Max = ele
		}
	}

	return Max
}
