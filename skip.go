package go_ienumerable

func (src *enumerable[T]) Skip(count int) IEnumerable[T] {
	src.assertSrcNonNil()

	if count < 1 {
		return src.copyExceptData().withData(copySlice(src.data))
	}

	if len(src.data) <= count { // skipped all
		return src.copyExceptData()
	}

	return src.copyExceptData().withData(src.data[count:])
}
