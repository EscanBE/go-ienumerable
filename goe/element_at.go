package goe

func (src *enumerable[T]) ElementAt(index int, reverse bool) T {
	src.assertSrcNonNil()
	src.assertIndex(index)

	if reverse {
		index = len(src.data) - 1 - index
	}

	return src.data[index]
}
