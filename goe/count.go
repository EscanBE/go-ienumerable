package goe

func (src *enumerable[T]) Count() int {
	src.assertSrcNonNil()
	return len(src.data)
}

func (src *enumerable[T]) LongCount() int64 {
	src.assertSrcNonNil()
	return int64(len(src.data))
}

func (src *enumerable[T]) CountBy(predicate func(T) bool) int {
	src.assertSrcNonNil()
	src.assertPredicateNonNil(predicate)

	if len(src.data) < 1 {
		return 0
	}

	count := 0
	for _, t := range src.data {
		if predicate(t) {
			count++
		}
	}
	return count
}

func (src *enumerable[T]) LongCountBy(predicate func(T) bool) int64 {
	src.assertSrcNonNil()
	src.assertPredicateNonNil(predicate)

	if len(src.data) < 1 {
		return 0
	}

	var count int64 = 0
	for _, t := range src.data {
		if predicate(t) {
			count++
		}
	}
	return count
}
