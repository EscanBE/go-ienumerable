package goe

import "fmt"

func (src *enumerable[T]) Take(count int) IEnumerable[T] {
	src.assertSrcNonNil()

	if count < 1 {
		return src.copyExceptData().withEmptyData()
	}

	if len(src.data) <= count { // take all
		return src.copyExceptData().withData(copySlice(src.data))
	}

	copied := make([]T, count)
	copy(copied, src.data)
	return src.copyExceptData().withData(copied)
}

func (src *enumerable[T]) TakeLast(count int) IEnumerable[T] {
	src.assertSrcNonNil()

	if count < 1 {
		return src.copyExceptData().withEmptyData()
	}

	if len(src.data) <= count { // take all
		return src.copyExceptData().withData(copySlice(src.data))
	}

	copied := make([]T, count)
	copy(copied, src.data[len(src.data)-count:])
	return src.copyExceptData().withData(copied)
}

//goland:noinspection SpellCheckingInspection
func (src *enumerable[T]) TakeWhile(predicate interface{}) IEnumerable[T] {
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

	fmt.Printf("selector nil %t\n", selector == nil)

	src.assertPredicateNonNil(selector)

	filtered := make([]T, 0)
	if len(src.data) > 0 {
		for i, d := range src.data {
			if selector(d, i) {
				filtered = append(filtered, d)
			} else {
				break
			}
		}
	}

	return src.copyExceptData().withData(filtered)
}
