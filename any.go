package go_ienumerable

func (src *enumerable[T]) Any() bool {
	src.assertSrcNonNil()
	return len(src.data) > 0
}

func (src *enumerable[T]) AnyBy(predicate func(T) bool) bool {
	src.assertSrcNonNil()
	src.assertPredicateNonNil(predicate)

	if len(src.data) < 1 {
		return false
	}

	for _, t := range src.data {
		if predicate(t) {
			return true
		}
	}

	return false
}
