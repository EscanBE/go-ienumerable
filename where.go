package go_ienumerable

func (src *enumerable[T]) Where(predicate func(T) bool) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertPredicateNonNil(predicate)

	filtered := make([]T, 0)
	if len(src.data) > 0 {
		for _, d := range src.data {
			if predicate(d) {
				filtered = append(filtered, d)
			}
		}
	}

	return src.copyExceptData().withData(filtered)
}
