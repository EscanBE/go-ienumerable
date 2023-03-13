package goe

func (src *enumerable[T]) Take(count int) IEnumerable[T] {
	src.assertSrcNonNil()

	if count < 1 {
		return src.copyExceptData().withEmptyData()
	}

	if len(src.data) <= count { // take all
		return src.copyExceptData().withData(copySlice(src.data))
	}

	copied := make([]T, count)
	copy(copied, src.data)
	return src.copyExceptData().withData(copied)
}

func (src *enumerable[T]) TakeLast(count int) IEnumerable[T] {
	src.assertSrcNonNil()

	if count < 1 {
		return src.copyExceptData().withEmptyData()
	}

	if len(src.data) <= count { // take all
		return src.copyExceptData().withData(copySlice(src.data))
	}

	copied := make([]T, count)
	copy(copied, src.data[len(src.data)-count:])
	return src.copyExceptData().withData(copied)
}
