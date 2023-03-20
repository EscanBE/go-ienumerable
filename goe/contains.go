package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
)

func (src *enumerable[T]) Contains(value T) bool {
	src.assertSrcNonNil()

	comparer := src.defaultComparer
	if comparer == nil {
		comparer = src.findDefaultComparer()
	}

	if len(src.data) < 1 {
		return false
	}

	for _, d := range src.data {
		if comparer.CompareAny(value, d) == 0 {
			return true
		}
	}

	return false
}

func (src *enumerable[T]) ContainsBy(value T, equalityOrComparer interface{}) bool {
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

	if len(src.data) < 1 {
		return false
	}

	for _, d := range src.data {
		if isEquals(value, d) {
			return true
		}
	}

	return false
}
