package go_ienumerable

func (src *enumerable[T]) Contains(value T) bool {
	src.assertSrcNonNil()

	if src.equalityComparer == nil {
		panicRequire(requireEqualityComparer)
	}

	return src.ContainsBy(value, src.equalityComparer)
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
