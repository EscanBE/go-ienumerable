package helper

import (
	"github.com/EscanBE/go-ienumerable/goe"
	"github.com/EscanBE/go-ienumerable/goe/comparers"
)

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

	result := goe.NewIEnumerable[TResult](newData...)

	if len(newData) > 0 {
		var nextDefaultComparer comparers.IComparer[any]

		for _, d := range newData {
			if any(d) == nil {
				continue
			}

			hasNextDefaultComparer := nextDefaultComparer != nil

			if hasNextDefaultComparer {
				break
			}

			if !hasNextDefaultComparer {
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
