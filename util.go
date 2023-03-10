package go_ienumerable

import "fmt"

func (src *enumerable[T]) exposeData() []T {
	return src.data
}

func (src *enumerable[T]) assertSrcNonNil() {
	if src == nil {
		panic("src is nil")
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

func (src *enumerable[T]) copyExceptData() *enumerable[T] {
	return &enumerable[T]{
		data:             nil,
		equalsComparator: src.equalsComparator,
		lessComparator:   src.lessComparator,
	}
}

func (src *enumerable[T]) withData(data []T) *enumerable[T] {
	src.data = data
	return src
}

func copySlice[T any](src []T) []T {
	dst := make([]T, len(src))
	if src != nil {
		copy(dst, src)
	}
	return dst
}
