package goe

func (src *enumerable[T]) FirstOrDefault(optionalPredicate OptionalPredicate[T], optionalDefaultValue *T) T {
	src.assertSrcNonNil()

	if len(src.data) > 0 {
		if optionalPredicate == nil {
			return src.data[0]
		} else {
			for _, d := range src.data {
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
