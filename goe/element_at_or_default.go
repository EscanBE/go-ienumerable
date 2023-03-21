package goe

func (src *enumerable[T]) ElementAtOrDefault(index int, reverse bool) T {
	src.assertSrcNonNil()

	if index >= len(src.data) {
		return *new(T)
	}

	if reverse {
		index = len(src.data) - 1 - index
	}

	return src.data[index]
}
