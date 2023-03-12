package goe

func (src *enumerable[T]) Empty() IEnumerable[T] {
	return src.copyExceptData().withEmptyData()
}
