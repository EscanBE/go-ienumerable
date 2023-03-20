package goe

func (src *enumerable[T]) LastOrDefault(optionalPredicate OptionalPredicate[T], optionalDefaultValue *T) T {
	src.assertSrcNonNil()

	if len(src.data) > 0 {
		if optionalPredicate == nil {
			return src.data[len(src.data)-1]
		} else {
			for i := len(src.data) - 1; i >= 0; i-- {
				d := src.data[i]
				if optionalPredicate(d) {
					return d
				}
			}
		}
	}

	if optionalDefaultValue == nil {
		return *new(T)
	} else {
		return *optionalDefaultValue
	}
}
