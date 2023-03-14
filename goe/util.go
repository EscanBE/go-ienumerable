package goe

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe/comparers"
)

func (src *enumerable[T]) assertSrcNonNil() {
	if src == nil {
		panic(getErrorSourceIsNil())
	}
}

func getErrorSourceIsNil() error {
	return fmt.Errorf("source is nil")
}

func (src *enumerable[T]) assertSrcNonEmpty() {
	if len(src.data) < 1 {
		panic(getErrorSrcContainsNoElement())
	}
}

func getErrorSrcContainsNoElement() error {
	return fmt.Errorf("source contains no element")
}

func (_ *enumerable[T]) assertSizeGt0(size int) {
	if size < 1 {
		panic("size is below 1")
	}
}

func (src *enumerable[T]) assertIndex(index int) {
	if 0 > index || index >= len(src.data) {
		panic("index out of bound")
	}
}

func (col *enumerator[T]) assertCollectionNonNil() {
	if col == nil {
		panic("collection is nil")
	}
}

func (_ *enumerable[T]) assertSecondIEnumerableNonNil(second IEnumerable[T]) {
	if second == nil {
		panic("second IEnumerable is nil")
	}
}

func (src *enumerable[T]) assertPredicateNonNil(predicate func(T) bool) {
	if predicate == nil {
		panic(getErrorNilPredicate())
	}
}

func (src *enumerable[T]) assertPredicate2NonNil(predicate func(T, int) bool) {
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

func getErrorNilComparer() error {
	return fmt.Errorf("comparer is nil")
}

func (src *enumerable[T]) assertSelectorNonNil(selector func(T) any) {
	if selector == nil {
		panic(getErrorNilSelector())
	}
}

func (src *enumerable[T]) assertArraySelectorNonNil(selector func(T) []any) {
	if selector == nil {
		panic(getErrorNilSelector())
	}
}

func getErrorNilSelector() error {
	return fmt.Errorf("selector is nil")
}

func (src *enumerable[T]) assertAggregateFuncNonNil(f func(T, T) T) {
	if f == nil {
		panic(getErrorNilAggregateFunc())
	}
}

func (src *enumerable[T]) assertAggregateAnySeedFuncNonNil(f func(any, T) any) {
	if f == nil {
		panic(getErrorNilAggregateFunc())
	}
}

// TODO move all errors into a separate file
func getErrorNilAggregateFunc() error {
	return fmt.Errorf("aggregate function is nil")
}

func getErrorMoreThanOne() error {
	return fmt.Errorf("more than one element")
}

func getErrorMoreThanOneMatch() error {
	return fmt.Errorf("more than one element satisfies the condition in predicate")
}

func getErrorNoMatch() error {
	return fmt.Errorf("no element satisfies the condition in predicate")
}

func (src *enumerable[T]) copyExceptData() *enumerable[T] {
	if src == nil {
		return nil
	}
	return &enumerable[T]{
		data:             nil,
		dataType:         src.dataType,
		equalityComparer: src.equalityComparer,
		lessComparer:     src.lessComparer,
		defaultComparer:  src.defaultComparer,
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

func getMapKeys[T comparable](m map[T]bool) []T {
	keys := make([]T, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

func (src *enumerable[T]) findDefaultComparer() comparers.IComparer[any] {
	comparer, found := src.tryFindDefaultComparer()
	if found {
		return comparer
	}

	panic(fmt.Errorf("no default comparer registered for [%s]", src.dataType))
}

func (src *enumerable[T]) tryFindDefaultComparer() (comparers.IComparer[any], bool) {
	if len(src.dataType) > 0 {
		if comparer, found := comparers.TryGetDefaultComparerByTypeName(src.dataType); found {
			return comparer, true
		}
	}

	if comparer, found := comparers.TryGetDefaultComparer[T](); found {
		return comparer, true
	}

	return nil, false
}

func (src *enumerable[T]) injectDefaultComparer() IEnumerable[T] {
	if comparer, found := src.tryFindDefaultComparer(); found {
		src.defaultComparer = comparer
	}

	return src
}

// cast IEnumerable back to *enumerable for accessing private fields.
func e[T any](ie IEnumerable[T]) *enumerable[T] {
	return ie.(*enumerable[T])
}
