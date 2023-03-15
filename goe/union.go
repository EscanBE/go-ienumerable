package goe

import "github.com/EscanBE/go-ienumerable/goe/comparers"

func (src *enumerable[T]) Union(second IEnumerable[T]) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertSecondIEnumerableNonNil(second)

	comparer := src.defaultComparer
	if comparer == nil {
		comparer = src.findDefaultComparer()
	}

	return src.internalUnionBy(second, func(v1, v2 T) bool {
		return comparer.Compare(v1, v2) == 0
	})
}

func (src *enumerable[T]) UnionBy(second IEnumerable[T], equalityComparer func(v1, v2 T) bool) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertSecondIEnumerableNonNil(second)

	if equalityComparer == nil {
		comparer := src.defaultComparer
		if comparer == nil {
			comparer = src.findDefaultComparer()
		}
		equalityComparer = func(v1, v2 T) bool {
			return comparer.Compare(v1, v2) == 0
		}
	}

	return src.internalUnionBy(second, equalityComparer)
}

func (src *enumerable[T]) UnionByComparer(second IEnumerable[T], comparer comparers.IComparer[T]) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertSecondIEnumerableNonNil(second)

	if comparer != nil {
		return src.internalUnionBy(second, func(v1, v2 T) bool {
			return comparer.Compare(v1, v2) == 0
		})
	}

	defaultComparer := src.defaultComparer
	if defaultComparer == nil {
		defaultComparer = src.findDefaultComparer()
	}

	return src.internalUnionBy(second, func(v1, v2 T) bool {
		return defaultComparer.Compare(v1, v2) == 0
	})
}

func (src *enumerable[T]) internalUnionBy(second IEnumerable[T], equalityComparer func(v1, v2 T) bool) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertSecondIEnumerableNonNil(second)
	src.assertComparerNonNil(equalityComparer)

	result := src.copyExceptData()

	if len(src.data) < 1 && second.Count() < 1 {
		return result.withEmptyData()
	}

	return result.
		withData(append(copySlice(src.data), copySlice(second.ToArray())...)).
		internalDistinctBy(equalityComparer)
}
