package goe

import "github.com/EscanBE/go-ienumerable/goe/comparers"

func (src *enumerable[T]) Intersect(second IEnumerable[T]) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertSecondIEnumerableNonNil(second)

	comparer := src.defaultComparer
	if comparer == nil {
		comparer = src.findDefaultComparer()
	}

	return src.internalIntersectBy(second, func(v1, v2 T) bool {
		return comparer.Compare(v1, v2) == 0
	})
}

func (src *enumerable[T]) IntersectBy(second IEnumerable[T], equalityComparer func(v1, v2 T) bool) IEnumerable[T] {
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

	return src.internalIntersectBy(second, equalityComparer)
}

func (src *enumerable[T]) IntersectByComparer(second IEnumerable[T], comparer comparers.IComparer[T]) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertSecondIEnumerableNonNil(second)

	if comparer != nil {
		return src.internalIntersectBy(second, func(v1, v2 T) bool {
			return comparer.Compare(v1, v2) == 0
		})
	}

	defaultComparer := src.defaultComparer
	if defaultComparer == nil {
		defaultComparer = src.findDefaultComparer()
	}

	return src.internalIntersectBy(second, func(v1, v2 T) bool {
		return defaultComparer.Compare(v1, v2) == 0
	})
}

func (src *enumerable[T]) internalIntersectBy(second IEnumerable[T], equalityComparer func(v1, v2 T) bool) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertSecondIEnumerableNonNil(second)
	src.assertComparerNonNil(equalityComparer)

	if len(src.data) < 1 || second.Count() < 1 {
		return src.copyExceptData().withEmptyData()
	}

	result := make([]T, 0)
	secondData := second.ToArray()
	for _, fe := range src.data {
		var foundInAnother bool
		for _, se := range secondData {
			if equalityComparer(fe, se) {
				foundInAnother = true
				break
			}
		}
		if foundInAnother {
			var addedPreviously bool

			for _, t := range result {
				if equalityComparer(fe, t) {
					addedPreviously = true
					break
				}
			}

			if !addedPreviously {
				result = append(result, fe)
			}
		}
	}

	return src.copyExceptData().withData(result)
}
