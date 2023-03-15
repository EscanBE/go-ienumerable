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
		if comparer.Compare(value, d) == 0 {
			return true
		}
	}

	return false
}

func (src *enumerable[T]) ContainsBy(value T, comparer interface{}) bool {
	src.assertSrcNonNil()

	var equalityComparer EqualsFunc[T]

	if comparer != nil {
		if eff, okEff := comparer.(func(v1, v2 T) bool); okEff {
			if eff != nil {
				equalityComparer = eff
			}
		} else if eft, okEft := comparer.(EqualsFunc[T]); okEft {
			if eft != nil {
				equalityComparer = eft
			}
		} else if cff, okCff := comparer.(func(v1, v2 T) int); okCff {
			if cff != nil {
				equalityComparer = func(v1, v2 T) bool {
					return cff(v1, v2) == 0
				}
			}
		} else if cft, okCft := comparer.(CompareFunc[T]); okCft {
			if cft != nil {
				equalityComparer = func(v1, v2 T) bool {
					return cft(v1, v2) == 0
				}
			}
		} else if cpr, okCpr := comparer.(comparers.IComparer[T]); okCpr {
			if cpr != nil {
				equalityComparer = func(v1, v2 T) bool {
					return cpr.Compare(v1, v2) == 0
				}
			}
		} else {
			panic(getErrorComparerMustBeEqualsFuncOrIComparer())
		}
	}

	if equalityComparer == nil {
		defaultComparer := src.defaultComparer
		if defaultComparer == nil {
			defaultComparer = src.findDefaultComparer()
		}
		equalityComparer = func(v1, v2 T) bool {
			return defaultComparer.Compare(v1, v2) == 0
		}
	}

	if len(src.data) < 1 {
		return false
	}

	for _, d := range src.data {
		if equalityComparer(value, d) {
			return true
		}
	}

	return false
}
