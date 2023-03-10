package go_ienumerable

func (src *enumerable[T]) ToArray() []T {
	src.assertSrcNonNil()

	if len(src.data) < 1 {
		return make([]T, 0)
	}

	return copySlice(src.data)
}
