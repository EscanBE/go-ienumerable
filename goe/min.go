package goe

import "github.com/EscanBE/go-ienumerable/goe/comparers2"

func (src *enumerable[T]) Min() T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

	comparer := src.defaultComparer
	if comparer == nil {
		comparer = src.findDefaultComparer()
	}

	return src.internalMinBy(func(v1, v2 T) bool {
		return comparer.CompareAny(v1, v2) < 0
	})
}

func (src *enumerable[T]) MinBy(lessThanOrComparer interface{}) T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

	var isLessThan LessFunc[T]

	if lessThanOrComparer != nil {
		if lff, okLff := lessThanOrComparer.(func(v1, v2 T) bool); okLff {
			if lff != nil {
				isLessThan = lff
			}
		} else if lft, okLft := lessThanOrComparer.(LessFunc[T]); okLft {
			if lft != nil {
				isLessThan = lft
			}
		} else if cff, okCff := lessThanOrComparer.(func(v1, v2 T) int); okCff {
			if cff != nil {
				isLessThan = func(v1, v2 T) bool {
					return cff(v1, v2) < 0
				}
			}
		} else if cft, okCft := lessThanOrComparer.(CompareFunc[T]); okCft {
			if cft != nil {
				isLessThan = func(v1, v2 T) bool {
					return cft(v1, v2) < 0
				}
			}
		} else if cpr, okCpr := lessThanOrComparer.(comparers.IComparer[T]); okCpr {
			if cpr != nil {
				isLessThan = func(v1, v2 T) bool {
					return cpr.CompareAny(v1, v2) < 0
				}
			}
		} else {
			panic(getErrorComparerMustBeLessThanFuncOrIComparer())
		}
	}

	if isLessThan == nil {
		defaultComparer := src.defaultComparer
		if defaultComparer == nil {
			defaultComparer = src.findDefaultComparer()
		}
		isLessThan = func(v1, v2 T) bool {
			return defaultComparer.CompareAny(v1, v2) < 0
		}
	}

	return src.internalMinBy(isLessThan)
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
