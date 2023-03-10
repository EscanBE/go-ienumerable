package go_ienumerable

func (src *enumerable[T]) Count() int {
	src.assertSrcNonNil()
	return len(src.data)
}

func (src *enumerable[T]) CountBy(predicate func(T) bool) int {
	src.assertSrcNonNil()
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
