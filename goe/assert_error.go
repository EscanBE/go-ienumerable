package goe

import (
	"fmt"
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
		panic(getErrorCollectionIsNil())
	}
}

func getErrorCollectionIsNil() error {
	return fmt.Errorf("collection is nil")
}

func assertSecondIEnumerableNonNil[T any](second IEnumerable[T]) {
	if second == nil {
		panic("second IEnumerable is nil")
	}
}

//goland:noinspection SpellCheckingInspection
func (src *enumerable[T]) assertPredicateNonNil(predicate interface{}) {
	if predicate == nil {
		panic(getErrorNilPredicate())
	} else if pff, okPff := predicate.(func(value T) bool); okPff {
		if pff == nil {
			panic(getErrorNilPredicate())
		}
	} else if pft, okPft := predicate.(Predicate[T]); okPft {
		if pft == nil {
			panic(getErrorNilPredicate())
		}
	} else if piff, okPiff := predicate.(func(value T, index int) bool); okPiff {
		if piff == nil {
			panic(getErrorNilPredicate())
		}
	} else if pift, okPift := predicate.(PredicateWithIndex[T]); okPift {
		if pift == nil {
			panic(getErrorNilPredicate())
		}
	} else {
		panic(getErrorPredicateMustBePredicate())
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

func (src *enumerable[T]) assertSelectorSameNonNil(selector func(T) T) {
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

func getErrorPredicateMustBePredicate() error {
	return fmt.Errorf("predicate must be\n- single predicate function: func(value T) bool\n- or predicate with index: func(value T, index int) bool")
}

func (o *orderedEnumerable[T]) assertSrcNonNil() {
	if o == nil {
		panic(getErrorSourceIsNil())
	}
}

func assertKeySelectorNonNil[T any](keySelector KeySelector[T]) {
	if keySelector == nil {
		panic(getErrorKeySelectorNotNil())
	}
}

func getErrorKeySelectorNotNil() error {
	return fmt.Errorf("key selector must not be nil")
}

func getErrorFailedCompare2ElementsInArray() error {
	return fmt.Errorf("failed to compare two elements in the array, at least one object must have registered default comparer")
}
