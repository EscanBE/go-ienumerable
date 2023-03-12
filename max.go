package go_ienumerable

func (src *enumerable[T]) Max() T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

	if src.lessComparer == nil {
		panicRequire(requireLessComparer)
	}

	return src.MaxBy(src.lessComparer)
}

func (src *enumerable[T]) MaxBy(lessComparer func(left, right T) bool) T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()
	src.assertComparerNonNil(lessComparer)

	max := src.data[0]

	for i := 1; i < len(src.data); i++ {
		ele := src.data[i]
		if lessComparer(max, ele) {
			max = ele
		}
	}

	return max
}
