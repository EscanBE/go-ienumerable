package goe

func (src *enumerable[T]) First(optionalPredicate OptionalPredicate[T]) T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

	if optionalPredicate == nil {
		return src.data[0]
	}

	for _, d := range src.data {
		if optionalPredicate(d) {
			return d
		}
	}

	panic(getErrorNoMatch())
}
