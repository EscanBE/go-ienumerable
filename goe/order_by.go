package goe

func (src *enumerable[T]) OrderBy(requiredKeySelector KeySelector[T], optionalCompareFunc OptionalCompareFunc[any]) IOrderedEnumerable[T] {
	src.assertSrcNonNil()
	assertKeySelectorNonNil(requiredKeySelector)

	return newIOrderedEnumerable[T](src, requiredKeySelector, optionalCompareFunc, CLC_ASC)
}
