package goe

func (src *enumerable[T]) Skip(count int) IEnumerable[T] {
	src.assertSrcNonNil()

	if count < 1 {
		return src.copyExceptData().withData(copySlice(src.data))
	}

	if len(src.data) <= count { // skipped all
		return src.copyExceptData().withEmptyData()
	}

	return src.copyExceptData().withData(copySlice(src.data[count:]))
}

func (src *enumerable[T]) SkipLast(count int) IEnumerable[T] {
	src.assertSrcNonNil()

	if count < 1 {
		return src.copyExceptData().withData(copySlice(src.data))
	}

	if len(src.data) <= count { // skipped all
		return src.copyExceptData().withEmptyData()
	}

	return src.copyExceptData().withData(copySlice(src.data[:len(src.data)-count]))
}

func (src *enumerable[T]) SkipWhile(predicate func(value T) bool) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertPredicateNonNil(predicate)

	if len(src.data) > 0 {
		for i, d := range src.data {
			if predicate(d) {
				continue
			} else {
				copied := copySlice(src.data[i:])
				return src.copyExceptData().withData(copied)
			}
		}
	}

	return src.copyExceptData().withEmptyData()
}

func (src *enumerable[T]) SkipWhileWidx(predicate func(value T, index int) bool) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertPredicate2NonNil(predicate)

	if len(src.data) > 0 {
		for i, d := range src.data {
			if predicate(d, i) {
				continue
			} else {
				copied := copySlice(src.data[i:])
				return src.copyExceptData().withData(copied)
			}
		}
	}

	return src.copyExceptData().withEmptyData()
}
