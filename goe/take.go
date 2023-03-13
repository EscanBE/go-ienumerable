package goe

func (src *enumerable[T]) Take(count int) IEnumerable[T] {
	src.assertSrcNonNil()

	if count < 1 {
		return src.copyExceptData().withEmptyData()
	}

	if len(src.data) <= count { // take all
		return src.copyExceptData().withData(copySlice(src.data))
	}

	copied := make([]T, count)
	copy(copied, src.data)
	return src.copyExceptData().withData(copied)
}

func (src *enumerable[T]) TakeLast(count int) IEnumerable[T] {
	src.assertSrcNonNil()

	if count < 1 {
		return src.copyExceptData().withEmptyData()
	}

	if len(src.data) <= count { // take all
		return src.copyExceptData().withData(copySlice(src.data))
	}

	copied := make([]T, count)
	copy(copied, src.data[len(src.data)-count:])
	return src.copyExceptData().withData(copied)
}

func (src *enumerable[T]) TakeWhile(predicate func(value T) bool) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertPredicateNonNil(predicate)

	filtered := make([]T, 0)
	if len(src.data) > 0 {
		for _, d := range src.data {
			if predicate(d) {
				filtered = append(filtered, d)
			} else {
				break
			}
		}
	}

	return src.copyExceptData().withData(filtered)
}

func (src *enumerable[T]) TakeWhileWidx(predicate func(value T, index int) bool) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertPredicate2NonNil(predicate)

	filtered := make([]T, 0)
	if len(src.data) > 0 {
		for i, d := range src.data {
			if predicate(d, i) {
				filtered = append(filtered, d)
			} else {
				break
			}
		}
	}

	return src.copyExceptData().withData(filtered)
}
