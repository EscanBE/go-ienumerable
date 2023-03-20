package goe

func (src *enumerable[T]) Contains(value T, optionalCompareFunc CompareFunc[T]) bool {
	src.assertSrcNonNil()

	if optionalCompareFunc == nil {
		comparer := src.defaultComparer
		if comparer == nil {
			comparer = src.findDefaultComparer()
		}

		optionalCompareFunc = func(v1, v2 T) int {
			return comparer.CompareAny(v1, v2)
		}
	}

	for _, d := range src.data {
		if optionalCompareFunc(value, d) == 0 {
			return true
		}
	}

	return false
}
