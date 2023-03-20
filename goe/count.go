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

func (src *enumerable[T]) LongCount(optionalPredicate OptionalPredicate[T]) int64 {
	count := src.Count(optionalPredicate)
	return int64(count)
}
