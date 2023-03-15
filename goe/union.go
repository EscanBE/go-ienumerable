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

func (src *enumerable[T]) UnionBy(second IEnumerable[T], equalityOrComparer interface{}) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertSecondIEnumerableNonNil(second)

	var isEquals EqualsFunc[T]

	if equalityOrComparer != nil {
		if eff, okEff := equalityOrComparer.(func(v1, v2 T) bool); okEff {
			if eff != nil {
				isEquals = eff
			}
		} else if eft, okEft := equalityOrComparer.(EqualsFunc[T]); okEft {
			if eft != nil {
				isEquals = eft
			}
		} else if cff, okCff := equalityOrComparer.(func(v1, v2 T) int); okCff {
			if cff != nil {
				isEquals = func(v1, v2 T) bool {
					return cff(v1, v2) == 0
				}
			}
		} else if cft, okCft := equalityOrComparer.(CompareFunc[T]); okCft {
			if cft != nil {
				isEquals = func(v1, v2 T) bool {
					return cft(v1, v2) == 0
				}
			}
		} else if cpr, okCpr := equalityOrComparer.(comparers.IComparer[T]); okCpr {
			if cpr != nil {
				isEquals = func(v1, v2 T) bool {
					return cpr.Compare(v1, v2) == 0
				}
			}
		} else {
			panic(getErrorComparerMustBeEqualsFuncOrIComparer())
		}
	}

	if isEquals == nil {
		defaultComparer := src.defaultComparer
		if defaultComparer == nil {
			defaultComparer = src.findDefaultComparer()
		}
		isEquals = func(v1, v2 T) bool {
			return defaultComparer.Compare(v1, v2) == 0
		}
	}

	return src.internalUnionBy(second, isEquals)
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
