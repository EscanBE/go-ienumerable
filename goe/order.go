package goe

func (src *enumerable[T]) Order() IOrderedEnumerable[T] {
	src.assertSrcNonNil()

	comparer := src.defaultComparer
	if comparer == nil {
		comparer = src.findDefaultComparer()
	}

	return newIOrderedEnumerable[T](src, func(ele T) any {
		return ele
	}, func(v1, v2 any) int {
		return comparer.CompareAny(v1, v2)
	}, CLC_ASC)
}

func (src *enumerable[T]) OrderBy(requiredKeySelector KeySelector[T], optionalCompareFunc CompareFunc[any]) IOrderedEnumerable[T] {
	src.assertSrcNonNil()
	assertKeySelectorNonNil(requiredKeySelector)

	return newIOrderedEnumerable[T](src, requiredKeySelector, optionalCompareFunc, CLC_ASC)
}

func (src *enumerable[T]) OrderByDescending() IOrderedEnumerable[T] {
	src.assertSrcNonNil()

	comparer := src.defaultComparer
	if comparer == nil {
		comparer = src.findDefaultComparer()
	}

	return newIOrderedEnumerable[T](src, func(ele T) any {
		return ele
	}, func(v1, v2 any) int {
		return comparer.CompareAny(v1, v2)
	}, CLC_DESC)
}

func (src *enumerable[T]) OrderByDescendingBy(requiredKeySelector KeySelector[T], optionalCompareFunc CompareFunc[any]) IOrderedEnumerable[T] {
	src.assertSrcNonNil()
	assertKeySelectorNonNil(requiredKeySelector)

	return newIOrderedEnumerable[T](src, requiredKeySelector, optionalCompareFunc, CLC_DESC)
}
