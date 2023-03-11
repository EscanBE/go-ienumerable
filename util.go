package go_ienumerable

import "fmt"

func (src *enumerable[T]) exposeData() []T {
	return src.data
}

func (src *enumerable[T]) len() int {
	return len(src.data)
}

func (src *enumerable[T]) assertSrcNonNil() {
	if src == nil {
		panic("src is nil")
	}
}

func (col *enumerator[T]) assertCollectionNonNil() {
	if col == nil {
		panic("collection is nil")
	}
}

func (src *enumerable[T]) assertSecondIEnumerableNonNil(second IEnumerable[T]) {
	if second == nil {
		panic("second IEnumerable is nil")
	}
}

func (src *enumerable[T]) assertPredicateNonNil(predicate func(T) bool) {
	if predicate == nil {
		panic(getErrorNilPredicate())
	}
}

func getErrorNilPredicate() error {
	return fmt.Errorf("predicate is nil")
}

func (src *enumerable[T]) assertComparerNonNil(comparer func(T, T) bool) {
	if comparer == nil {
		panic(getErrorNilComparer())
	}
}

func (src *enumerable[T]) assertSelectorNonNil(selector func(T) any) {
	if selector == nil {
		panic(getErrorNilSelector())
	}
}

func getErrorNilComparer() error {
	return fmt.Errorf("comparer is nil")
}

func getErrorNilSelector() error {
	return fmt.Errorf("selector is nil")
}

func (src *enumerable[T]) copy() IEnumerable[T] {
	return src.copyExceptData().withData(copySlice(src.data))
}

func (src *enumerable[T]) copyExceptData() *enumerable[T] {
	if src == nil {
		return nil
	}
	return &enumerable[T]{
		data:             nil,
		equalityComparer: src.equalityComparer,
		lessComparer:     src.lessComparer,
	}
}

func (src *enumerable[T]) withData(data []T) *enumerable[T] {
	if src == nil {
		return nil
	}
	src.data = data
	return src
}

func (src *enumerable[T]) withEmptyData() *enumerable[T] {
	if src == nil {
		return nil
	}
	src.data = make([]T, 0)
	return src
}

func copySlice[T any](src []T) []T {
	dst := make([]T, len(src))
	if src != nil {
		copy(dst, src)
	}
	return dst
}
