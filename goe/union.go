package goe

func (src *enumerable[T]) Union(second IEnumerable[T], optionalEqualsFunc OptionalEqualsFunc[T]) IEnumerable[T] {
	src.assertSrcNonNil()
	assertSecondIEnumerableNonNil(second)

	var equalsFunc EqualsFunc[T]

	if optionalEqualsFunc == nil {
		comparer := src.defaultComparer
		if comparer == nil {
			comparer = src.findDefaultComparer()
		}
		equalsFunc = func(v1, v2 T) bool {
			return comparer.CompareAny(v1, v2) == 0
		}
	} else {
		equalsFunc = EqualsFunc[T](optionalEqualsFunc)
	}

	result := src.copyExceptData()

	if len(src.data) < 1 && second.Count(nil) < 1 {
		return result.withEmptyData()
	}

	uniqueData := distinct(
		append(copySlice(src.data), copySlice(second.ToArray())...),
		OptionalEqualsFunc[T](equalsFunc),
	)

	return result.withData(uniqueData)
}
