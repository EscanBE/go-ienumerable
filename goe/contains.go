package goe

func (src *enumerable[T]) Contains(value T, optionalEqualsFunc OptionalEqualsFunc[T]) bool {
	src.assertSrcNonNil()

	var equalsFunc EqualsFunc[T]

	if optionalEqualsFunc == nil {
		comparer := src.defaultComparer
		if comparer == nil {
			comparer = src.findDefaultComparer()
		}

		equalsFunc = func(v1, v2 T) bool {
			return comparer.CompareAny(v1, v2) == 0
		}
	} else {
		equalsFunc = EqualsFunc[T](optionalEqualsFunc)
	}

	for _, d := range src.data {
		if equalsFunc(value, d) {
			return true
		}
	}

	return false
}
