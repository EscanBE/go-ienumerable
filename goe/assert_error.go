package goe

import "fmt"

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

func (src *enumerable[T]) assertSampleResultValueNonNil(sampleResultValue any) {
	if sampleResultValue == nil {
		panic(getErrorSampleValueIsNil())
	}
}

func getErrorSampleValueIsNil() error {
	return fmt.Errorf("sample result value is nil")
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