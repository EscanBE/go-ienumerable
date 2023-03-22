package goe

import "github.com/EscanBE/go-ienumerable/goe/comparers"

func (src *enumerable[T]) DistinctBy(keySelector KeySelector[T], optionalEqualsFunc OptionalEqualsFunc[any]) IEnumerable[T] {
	src.assertSrcNonNil()
	assertKeySelectorNonNil(keySelector)

	unique := distinctByKeySelector(copySlice(src.data), keySelector, optionalEqualsFunc)
	return src.copyExceptData().withData(unique)
}

type distinctElementHolder struct {
	elementIndex int
	key          any
}

func distinctByKeySelector[T any](data []T, requiredKeySelector KeySelector[T], optionalEqualityComparer OptionalEqualsFunc[any]) []T {
	if requiredKeySelector == nil {
		panic(getErrorKeySelectorNotNil())
	}

	var equalityComparer EqualsFunc[any]
	if optionalEqualityComparer != nil {
		equalityComparer = EqualsFunc[any](optionalEqualityComparer)
	}

	holders := make([]distinctElementHolder, len(data))

	for i, d := range data {
		holders[i] = distinctElementHolder{
			elementIndex: i,
			key:          requiredKeySelector(d),
		}

		if equalityComparer == nil {
			comparer, found := comparers.TryGetDefaultComparerFromValue(holders[i].key)
			if found {
				equalityComparer = func(v1, v2 any) bool {
					return comparer.CompareAny(v1, v2) == 0
				}
			}
		}
	}

	if equalityComparer == nil {
		panic(getErrorFailedCompare2ElementsInArray())
	}

	if len(data) < 1 {
		return data
	}

	uniqueSet := []distinctElementHolder{holders[0]}

	for i1 := 1; i1 < len(data); i1++ {
		ele := holders[i1]

		var exists bool
		for _, uniq := range uniqueSet {
			if equalityComparer(ele.key, uniq.key) {
				exists = true
				break
			}
		}

		if !exists {
			uniqueSet = append(uniqueSet, ele)
		}
	}

	uniqueResult := make([]T, len(uniqueSet))
	for i, h := range uniqueSet {
		uniqueResult[i] = data[h.elementIndex]
	}

	return uniqueResult
}
