package go_ienumerable

func (src *enumerable[T]) All(predicate func(T) bool) bool {
	src.assertSrcNonNil()
	src.assertPredicateNonNil(predicate)

	if len(src.data) < 1 {
		return true
	}

	for _, t := range src.data {
		if !predicate(t) {
			return false
		}
	}

	return true
}
