package go_ienumerable

func (src *enumerable[T]) Reverse() IEnumerable[T] {
	src.assertSrcNonNil()

	size := len(src.data)
	reverseData := make([]T, size)

	if size > 0 {
		for i, t := range src.data {
			reverseData[size-1-i] = t
		}
	}

	return src.copyExceptData().withData(reverseData)
}
