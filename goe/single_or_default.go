package goe

func (src *enumerable[T]) SingleOrDefault(optionalPredicate OptionalPredicate[T], defaultValue *T) T {
	src.assertSrcNonNil()

	if len(src.data) > 0 {
		if optionalPredicate == nil {
			if len(src.data) > 1 {
				panic(getErrorMoreThanOne())
			} else {
				return src.data[0]
			}
		} else {
			var result T
			var anyMatch bool
			for _, d := range src.data {
				if optionalPredicate(d) {
					if anyMatch {
						panic(getErrorMoreThanOneMatch())
					}
					result = d
					anyMatch = true
				}
			}

			if anyMatch {
				return result
			}
		}
	}

	if defaultValue == nil {
		return *new(T)
	} else {
		return *defaultValue
	}
}
