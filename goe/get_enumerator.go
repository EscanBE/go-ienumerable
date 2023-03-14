package goe

func (src *enumerable[T]) GetEnumerator() IEnumerator[T] {
	src.assertSrcNonNil()
	return NewIEnumerator[T](src.ToArray()...)
}
