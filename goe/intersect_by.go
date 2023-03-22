package goe

import "github.com/EscanBE/go-ienumerable/goe/comparers"

type intersectElementHolder struct {
	key   any
	index int
}

func (src *enumerable[T]) IntersectBy(second IEnumerable[T], keySelector KeySelector[T], optionalEqualsFunc OptionalEqualsFunc[any]) IEnumerable[T] {
	src.assertSrcNonNil()
	assertSecondIEnumerableNonNil(second)
	assertKeySelectorNonNil(keySelector)

	var isEquals EqualsFunc[any]
	if optionalEqualsFunc != nil {
		isEquals = EqualsFunc[any](optionalEqualsFunc)
	}

	result := src.copyExceptData()

	if len(src.data) < 1 || second.Count(nil) < 1 {
		result = result.withEmptyData()
	} else {
		intersectIndex := make([]int, 0)
		sourceKeys := make([]intersectElementHolder, len(src.data))
		secondKeys := make([]any, second.Count(nil))

		for i, t := range src.data {
			sourceKeys[i] = intersectElementHolder{
				key:   keySelector(t),
				index: i,
			}

			if isEquals == nil {
				defaultComparer, found := comparers.TryGetDefaultComparerFromValue(sourceKeys[i].key)
				if found {
					isEquals = func(v1, v2 any) bool {
						return defaultComparer.CompareAny(v1, v2) == 0
					}
				}
			}
		}

		for i, t := range second.ToArray() {
			secondKeys[i] = keySelector(t)
			if isEquals == nil {
				defaultComparer, found := comparers.TryGetDefaultComparerFromValue(secondKeys[i])
				if found {
					isEquals = func(v1, v2 any) bool {
						return defaultComparer.CompareAny(v1, v2) == 0
					}
				}
			}
		}

		if isEquals == nil {
			panic(getErrorFailedCompare2ElementsInArray())
		}

		for _, fe := range sourceKeys {
			for _, se := range secondKeys {
				if isEquals(fe.key, se) {
					intersectIndex = append(intersectIndex, fe.index)
					break
				}
			}
		}

		intersect := make([]T, len(intersectIndex))
		for i, index := range intersectIndex {
			intersect[i] = src.data[index]
		}

		result = result.withData(distinctByKeySelector(intersect, keySelector, OptionalEqualsFunc[any](isEquals)))
	}

	return result
}
