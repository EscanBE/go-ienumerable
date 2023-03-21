package goe

func (src *enumerable[T]) Last(optionalPredicate OptionalPredicate[T]) T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

	if optionalPredicate == nil {
		return src.data[len(src.data)-1]
	}

	for i := len(src.data) - 1; i >= 0; i-- {
		d := src.data[i]
		if optionalPredicate(d) {
			return d
		}
	}

	panic(getErrorNoMatch())
}
