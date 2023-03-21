package goe

//goland:noinspection SpellCheckingInspection
func (src *enumerable[T]) SkipWhile(predicate interface{}) IEnumerable[T] {
	src.assertSrcNonNil()

	var selector PredicateWithIndex[T]

	if predicate != nil {
		if pff, okPff := predicate.(func(value T) bool); okPff {
			if pff != nil {
				selector = func(value T, _ int) bool {
					return pff(value)
				}
			}
		} else if pft, okPft := predicate.(Predicate[T]); okPft {
			if pft != nil {
				selector = func(value T, _ int) bool {
					return pft(value)
				}
			}
		} else if piff, okPiff := predicate.(func(value T, index int) bool); okPiff {
			if piff != nil {
				selector = piff
			}
		} else if pift, okPift := predicate.(PredicateWithIndex[T]); okPift {
			if pift != nil {
				selector = pift
			}
		} else {
			panic(getErrorPredicateMustBePredicate())
		}
	}

	src.assertPredicateNonNil(selector)

	if len(src.data) > 0 {
		for i, d := range src.data {
			if selector(d, i) {
				continue
			} else {
				copied := copySlice(src.data[i:])
				return src.copyExceptData().withData(copied)
			}
		}
	}

	return src.copyExceptData().withEmptyData()
}
