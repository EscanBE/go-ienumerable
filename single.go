package go_ienumerable

func (src *enumerable[T]) Single() T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

	if len(src.data) > 1 {
		panic(getErrorMoreThanOne())
	}

	return src.data[0]
}

func (src *enumerable[T]) SingleBy(predicate func(T) bool) T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()
	src.assertPredicateNonNil(predicate)

	var result T
	var anyMatch bool
	for _, d := range src.data {
		if predicate(d) {
			if anyMatch {
				panic(getErrorMoreThanOneMatch())
			}
			result = d
			anyMatch = true
		}
	}

	if anyMatch {
		return result
	}

	panic(getErrorNoMatch())
}
