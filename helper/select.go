package helper

import (
	"github.com/EscanBE/go-ienumerable/goe"
)

// Select projects each element of a sequence into a new form.
func Select[TSource, TResult any](source goe.IEnumerable[TSource], selector func(v TSource) TResult) goe.IEnumerable[TResult] {
	assertCollectionNotNil(source, "source")
	assertResultSelectorFunctionNotNil(selector)

	size := source.Count(nil)
	if size < 1 {
		return goe.NewIEnumerable[TResult]()
	}

	newData := make([]TResult, size)

	for i, d := range source.ToArray() {
		v := selector(d)
		newData[i] = v
	}

	return goe.NewIEnumerable[TResult](newData...)
}
