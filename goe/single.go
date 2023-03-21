package goe

func (src *enumerable[T]) Single(optionalPredicate OptionalPredicate[T]) T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

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

	panic(getErrorNoMatch())
}
