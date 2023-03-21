package goe

func (src *enumerable[T]) Min() T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

	comparer := src.defaultComparer
	if comparer == nil {
		comparer = src.findDefaultComparer()
	}

	greaterComparer := func(v1, v2 T) bool {
		return comparer.CompareAny(v1, v2) < 0
	}

	minIdx := 0

	for i := 1; i < len(src.data); i++ {
		if greaterComparer(src.data[i], src.data[minIdx]) {
			minIdx = i
		}
	}

	return src.data[minIdx]
}
