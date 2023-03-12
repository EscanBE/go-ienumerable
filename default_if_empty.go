package go_ienumerable

func (src *enumerable[T]) DefaultIfEmpty() IEnumerable[T] {
	src.assertSrcNonNil()

	result := src.copyExceptData()

	if len(src.data) > 0 {
		result = result.withData(copySlice(src.data))
	} else {
		result = result.withData([]T{*new(T)})
	}

	return result
}

func (src *enumerable[T]) DefaultIfEmptyUsing(defaultValue T) IEnumerable[T] {
	src.assertSrcNonNil()

	result := src.copyExceptData()

	if len(src.data) > 0 {
		result = result.withData(copySlice(src.data))
	} else {
		result = result.withData([]T{defaultValue})
	}

	return result
}
