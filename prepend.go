package go_ienumerable

func (src *enumerable[T]) Prepend(element T) IEnumerable[T] {
	src.assertSrcNonNil()

	result := src.copyExceptData()

	if len(src.data) < 1 {
		result.withData([]T{element})
	} else {
		dst := make([]T, len(src.data)+1)
		copy(dst[1:], src.data)
		dst[0] = element
		result.withData(dst)
	}

	return result
}
