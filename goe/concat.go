package goe

func (src *enumerable[T]) Concat(second IEnumerable[T]) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertSecondIEnumerableNonNil(second)

	result := src.copyExceptData()

	if len(src.data) == 0 && second.Count() > 0 {
		result = result.withData(copySlice(second.exposeData()))
	} else if len(src.data) > 0 && second.Count() == 0 {
		result = result.withData(copySlice(src.data))
	} else if len(src.data) > 0 && second.Count() > 0 {
		result = result.withData(append(copySlice(src.data), copySlice(second.exposeData())...))
	} else {
		result = result.withEmptyData()
	}

	return result
}
