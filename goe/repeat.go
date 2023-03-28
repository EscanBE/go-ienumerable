package goe

func (src *enumerable[T]) Repeat(element T, count int) IEnumerable[T] {
	src.assertSrcNonNil()

	if count < 0 {
		panic("count is less than 0")
	}

	data := make([]T, count)
	for i := 0; i < count; i++ {
		ele := element // copy
		data[i] = ele
	}

	return src.copyExceptData().withData(data)
}
