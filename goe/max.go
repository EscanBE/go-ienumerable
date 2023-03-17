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

func (src *enumerable[T]) MaxBy(greaterThanOrComparer interface{}) T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

	var isGreaterThan GreaterFunc[T]

	if greaterThanOrComparer != nil {
		if gff, okGff := greaterThanOrComparer.(func(v1, v2 T) bool); okGff {
			if gff != nil {
				isGreaterThan = gff
			}
		} else if gft, okGft := greaterThanOrComparer.(GreaterFunc[T]); okGft {
			if gft != nil {
				isGreaterThan = gft
			}
		} else if cff, okCff := greaterThanOrComparer.(func(v1, v2 T) int); okCff {
			if cff != nil {
				isGreaterThan = func(v1, v2 T) bool {
					return cff(v1, v2) > 0
				}
			}
		} else if cft, okCft := greaterThanOrComparer.(CompareFunc[T]); okCft {
			if cft != nil {
				isGreaterThan = func(v1, v2 T) bool {
					return cft(v1, v2) > 0
				}
			}
		} else if cpr, okCpr := greaterThanOrComparer.(comparers.IComparer[T]); okCpr {
			if cpr != nil {
				isGreaterThan = func(v1, v2 T) bool {
					return cpr.Compare(v1, v2) > 0
				}
			}
		} else {
			panic(getErrorComparerMustBeGreaterThanFuncOrIComparer())
		}
	}

	if isGreaterThan == nil {
		defaultComparer := src.defaultComparer
		if defaultComparer == nil {
			defaultComparer = src.findDefaultComparer()
		}
		isGreaterThan = func(v1, v2 T) bool {
			return defaultComparer.Compare(v1, v2) > 0
		}
	}

	return src.internalMaxBy(isGreaterThan)
}

func (src *enumerable[T]) internalMaxBy(greaterComparer func(left, right T) bool) T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()
	src.assertComparerNonNil(greaterComparer)

	max := src.data[0]

	for i := 1; i < len(src.data); i++ {
		ele := src.data[i]
		if greaterComparer(ele, max) {
			max = ele
		}
	}

	return max
}
