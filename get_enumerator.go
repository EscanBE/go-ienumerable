package go_ienumerable

func (src *enumerable[T]) GetEnumerator() IEnumerator[T] {
	src.assertSrcNonNil()
	return NewIEnumerator[T](src.exposeData()...)
}
