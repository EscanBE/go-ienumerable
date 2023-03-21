package goe

func (src *enumerable[T]) UnionBy(second IEnumerable[T], keySelector KeySelector[T], optionalEqualsFunc OptionalEqualsFunc[any]) IEnumerable[T] {
	src.assertSrcNonNil()
	assertSecondIEnumerableNonNil(second)
	assertKeySelectorNonNil(keySelector)

	result := src.copyExceptData()

	if len(src.data) < 1 && second.Count(nil) < 1 {
		result = result.withEmptyData()
	} else {
		uniqueData := distinctByKeySelector(
			append(copySlice(src.data), copySlice(second.ToArray())...),
			keySelector,
			optionalEqualsFunc,
		)

		result = result.withData(uniqueData)
	}

	return result
}
