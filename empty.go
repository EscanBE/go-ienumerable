package go_ienumerable

func (src *enumerable[T]) Empty() IEnumerable[T] {
	return src.copyExceptData().withEmptyData()
}
