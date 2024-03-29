package goe

func (src *enumerable[T]) Append(element T) IEnumerable[T] {
	src.assertSrcNonNil()

	result := src.copyExceptData()

	if len(src.data) < 1 {
		result.withData([]T{element})
	} else {
		dst := make([]T, len(src.data)+1)
		copy(dst, src.data)
		dst[len(dst)-1] = element
		result.withData(dst)
	}

	return result
}
