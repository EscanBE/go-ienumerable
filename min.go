package go_ienumerable

func (src *enumerable[T]) Min() T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

	if src.lessComparer == nil {
		panicRequire(requireLessComparer)
	}

	return src.MinBy(src.lessComparer)
}

func (src *enumerable[T]) MinBy(lessComparer func(left, right T) bool) T {
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
