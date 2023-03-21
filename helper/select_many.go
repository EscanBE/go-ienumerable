package helper

import (
	"github.com/EscanBE/go-ienumerable/goe"
)

// SelectMany projects each element of a sequence to a goe.IEnumerable[T] and flattens the resulting sequences into one sequence.
func SelectMany[TSource, TResult any](source goe.IEnumerable[TSource], selector func(v TSource) []TResult) goe.IEnumerable[TResult] {
	assertCollectionNotNil(source, "source")
	assertManyResultSelectorFunctionNotNil(selector)

	size := source.Count(nil)
	if size < 1 {
		return goe.NewIEnumerable[TResult]()
	}

	newData := make([]TResult, 0)

	for _, d := range source.ToArray() {
		a := selector(d)

		if a == nil {
			panic("result array can not be nil")
		}

		if len(a) < 1 {
			continue
		}

		newData = append(newData, a...)
	}

	return goe.NewIEnumerable[TResult](newData...)
}

// SelectManyTransform projects each element of a sequence to a goe.IEnumerable[T], flattens the resulting sequences into one sequence, and invokes a result selector function on each element therein. The index of each source element is used in the intermediate projected form of that element.
func SelectManyTransform[TSource, TCollection, TResult any](source goe.IEnumerable[TSource], collectionSelector func(v TSource) []TCollection, resultSelector func(src TSource, col TCollection) TResult) goe.IEnumerable[TResult] {
	assertCollectionNotNil(source, "source")
	assertCollectionSelectorFunctionNotNil(collectionSelector)
	assertResultSelectorFunctionNotNil2(resultSelector)

	size := source.Count(nil)
	if size < 1 {
		return goe.NewIEnumerable[TResult]()
	}

	newData := make([]TResult, 0)

	for _, d := range source.ToArray() {
		col := collectionSelector(d)

		if col == nil {
			panic("result array can not be nil")
		}

		if len(col) < 1 {
			continue
		}

		for _, e := range col {
			newData = append(newData, resultSelector(d, e))
		}
	}

	return goe.NewIEnumerable[TResult](newData...)
}
