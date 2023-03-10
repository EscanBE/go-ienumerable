package go_ienumerable

func (e *enumerable[T]) Count() int {
	return len(e.data)
}

func (e *enumerable[T]) CountBy(predicate func(T) bool) int {
	if len(e.data) < 1 {
		return 0
	}

	count := 0
	for _, t := range e.data {
		if predicate(t) {
			count++
		}
	}
	return count
}
