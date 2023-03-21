package goe

func (src *enumerable[T]) Count(optionalPredicate OptionalPredicate[T]) int {
	src.assertSrcNonNil()

	if optionalPredicate == nil {
		return len(src.data)
	}

	count := 0
	for _, t := range src.data {
		if optionalPredicate(t) {
			count++
		}
	}
	return count
}
