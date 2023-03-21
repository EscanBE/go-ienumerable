package goe

import "github.com/EscanBE/go-ienumerable/goe/comparers"

func (src *enumerable[T]) Intersect(second IEnumerable[T], optionalEqualsFunc OptionalEqualsFunc[T]) IEnumerable[T] {
	src.assertSrcNonNil()
	assertSecondIEnumerableNonNil(second)

	var isEquals EqualsFunc[T]
	if optionalEqualsFunc == nil {
		defaultComparer := src.defaultComparer
		if defaultComparer == nil {
			defaultComparer = src.findDefaultComparer()
		}
		isEquals = func(v1, v2 T) bool {
			return defaultComparer.CompareAny(v1, v2) == 0
		}
	} else {
		isEquals = EqualsFunc[T](optionalEqualsFunc)
	}

	result := src.copyExceptData()

	if len(src.data) < 1 || second.Count(nil) < 1 {
		result = result.withEmptyData()
	} else {
		intersect := make([]T, 0)
		secondData := second.ToArray()

		for _, fe := range src.data {
			for _, se := range secondData {
				if isEquals(fe, se) {
					intersect = append(intersect, fe)
					break
				}
			}
		}

		result = result.withData(distinct(intersect, OptionalEqualsFunc[T](isEquals)))
	}

	return result
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
		type holder struct {
			key   any
			index int
		}

		intersectIndex := make([]int, 0)
		sourceKeys := make([]holder, len(src.data))
		secondKeys := make([]any, second.Count(nil))

		for i, t := range src.data {
			sourceKeys[i] = holder{
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
