package goe

func (src *enumerable[T]) Order() IOrderedEnumerable[T] {
	src.assertSrcNonNil()

	return newIOrderedEnumerable[T](src, func(ele T) any {
		return ele
	}, nil, CLC_ASC)
}

func (src *enumerable[T]) OrderBy(keySelector KeySelector[T], compareFunc CompareFunc[any]) IOrderedEnumerable[T] {
	src.assertSrcNonNil()
	assertKeySelectorNonNil(keySelector)

	return newIOrderedEnumerable[T](src, keySelector, compareFunc, CLC_ASC)
}

func (src *enumerable[T]) OrderByDescending() IOrderedEnumerable[T] {
	src.assertSrcNonNil()

	return newIOrderedEnumerable[T](src, func(ele T) any {
		return ele
	}, nil, CLC_DESC)
}

func (src *enumerable[T]) OrderByDescendingBy(keySelector KeySelector[T], compareFunc CompareFunc[any]) IOrderedEnumerable[T] {
	src.assertSrcNonNil()
	assertKeySelectorNonNil(keySelector)

	return newIOrderedEnumerable[T](src, keySelector, compareFunc, CLC_DESC)
}
