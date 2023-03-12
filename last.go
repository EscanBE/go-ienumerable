package go_ienumerable

func (src *enumerable[T]) Last() T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

	return src.data[len(src.data)-1]
}

func (src *enumerable[T]) LastBy(predicate func(T) bool) T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()
	src.assertPredicateNonNil(predicate)

	for i := len(src.data) - 1; i >= 0; i-- {
		if predicate(src.data[i]) {
			return src.data[i]
		}
	}

	panic(getErrorNoMatch())
}
