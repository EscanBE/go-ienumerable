package helper

import (
	"github.com/EscanBE/go-ienumerable/goe"
	"github.com/EscanBE/go-ienumerable/goe/comparers"
)

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

	result := goe.NewIEnumerable[TResult](newData...)

	if len(newData) > 0 {
		var nextDefaultComparer comparers.IComparer[any]

		for _, d := range newData {
			if any(d) == nil {
				continue
			}

			if nextDefaultComparer != nil {
				break
			}

			if nextDefaultComparer == nil {
				comparer, found := comparers.TryGetDefaultComparerFromValue(d)
				if found {
					nextDefaultComparer = comparer
				}
			}
		}

		if nextDefaultComparer != nil {
			result = result.WithDefaultComparerAny(nextDefaultComparer)
		}
	}

	return result
}
