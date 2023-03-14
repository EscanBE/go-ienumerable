package goe

import "github.com/EscanBE/go-ienumerable/goe/comparers"

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

func (src *enumerable[T]) ContainsBy(value T, equalityComparer func(v1, v2 T) bool) bool {
	src.assertSrcNonNil()

	if equalityComparer == nil {
		comparer := src.defaultComparer
		if comparer == nil {
			comparer = src.findDefaultComparer()
		}
		equalityComparer = func(v1, v2 T) bool {
			return comparer.Compare(v1, v2) == 0
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

func (src *enumerable[T]) ContainsByComparer(value T, comparer comparers.IComparer[T]) bool {
	src.assertSrcNonNil()

	if comparer != nil {
		return src.ContainsBy(value, func(v1, v2 T) bool {
			return comparer.Compare(v1, v2) == 0
		})
	}

	defaultComparer := src.defaultComparer
	if defaultComparer == nil {
		defaultComparer = src.findDefaultComparer()
	}

	return src.ContainsBy(value, func(v1, v2 T) bool {
		return defaultComparer.Compare(v1, v2) == 0
	})
}
