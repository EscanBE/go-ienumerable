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

func (src *enumerable[T]) DistinctBy(equalityOrComparer interface{}) IEnumerable[T] {
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

	return src.internalDistinctBy(isEquals)
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
