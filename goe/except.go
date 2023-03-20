package goe

import "github.com/EscanBE/go-ienumerable/goe/comparers"

func (src *enumerable[T]) Except(second IEnumerable[T]) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertSecondIEnumerableNonNil(second)

	comparer := src.defaultComparer
	if comparer == nil {
		comparer = src.findDefaultComparer()
	}

	return src.internalExceptBy(second, func(v1, v2 T) bool {
		return comparer.CompareAny(v1, v2) == 0
	})
}

func (src *enumerable[T]) ExceptBy(second IEnumerable[T], equalityOrComparer interface{}) IEnumerable[T] {
	src.assertSrcNonNil()

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
					return cpr.CompareAny(v1, v2) == 0
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
			return defaultComparer.CompareAny(v1, v2) == 0
		}
	}

	return src.internalExceptBy(second, isEquals)
}

func (src *enumerable[T]) internalExceptBy(second IEnumerable[T], equalityComparer func(v1, v2 T) bool) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertSecondIEnumerableNonNil(second)
	src.assertComparerNonNil(equalityComparer)

	if second.Count() < 1 {
		return src.copyExceptData().withData(copySlice(src.ToArray()))
	}

	if len(src.data) < 1 {
		return src.copyExceptData().withEmptyData()
	}

	result := make([]T, 0)
	secondData := second.ToArray()
	for _, d := range src.data {
		var foundInAnother bool
		for _, t := range secondData {
			if equalityComparer(d, t) {
				foundInAnother = true
				break
			}
		}
		if !foundInAnother {
			var addedPreviously bool

			for _, t := range result {
				if equalityComparer(d, t) {
					addedPreviously = true
					break
				}
			}

			if !addedPreviously {
				result = append(result, d)
			}
		}
	}

	return src.copyExceptData().withData(result)
}
