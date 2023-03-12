package go_ienumerable

func (src *enumerable[T]) First() T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

	return src.data[0]
}

func (src *enumerable[T]) FirstBy(predicate func(T) bool) T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()
	src.assertPredicateNonNil(predicate)

	for _, d := range src.data {
		if predicate(d) {
			return d
		}
	}

	panic(getErrorNoMatch())
}
