package goe

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
	src.assertComparerNonNil(equalityComparer)

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
