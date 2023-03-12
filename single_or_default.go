package go_ienumerable

func (src *enumerable[T]) SingleOrDefault() T {
	src.assertSrcNonNil()

	if len(src.data) > 1 {
		panic(getErrorMoreThanOne())
	}

	if len(src.data) < 1 {
		return *new(T)
	}

	return src.data[0]
}

func (src *enumerable[T]) SingleOrDefaultBy(predicate func(T) bool) T {
	src.assertSrcNonNil()
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

	return *new(T)
}

func (src *enumerable[T]) SingleOrDefaultUsing(defaultValue T) T {
	src.assertSrcNonNil()

	if len(src.data) > 1 {
		panic(getErrorMoreThanOne())
	}

	if len(src.data) < 1 {
		return defaultValue
	}

	return src.data[0]
}

func (src *enumerable[T]) SingleOrDefaultByUsing(predicate func(T) bool, defaultValue T) T {
	src.assertSrcNonNil()
	src.assertPredicateNonNil(predicate)

	if len(src.data) < 1 {
		return defaultValue
	}

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

	return defaultValue
}
