package go_ienumerable

func (src *enumerable[T]) ElementAt(index int) T {
	src.assertSrcNonNil()
	src.assertIndex(index)

	return src.data[index]
}

func (src *enumerable[T]) ElementAtReverse(reverseIndex int) T {
	src.assertSrcNonNil()
	src.assertIndex(reverseIndex)

	return src.data[len(src.data)-1-reverseIndex]
}

func (src *enumerable[T]) ElementAtOrDefault(index int) T {
	src.assertSrcNonNil()

	if index >= len(src.data) {
		return *new(T)
	}

	return src.data[index]
}

func (src *enumerable[T]) ElementAtReverseOrDefault(reverseIndex int) T {
	src.assertSrcNonNil()

	if reverseIndex >= len(src.data) {
		return *new(T)
	}

	return src.data[len(src.data)-1-reverseIndex]
}
