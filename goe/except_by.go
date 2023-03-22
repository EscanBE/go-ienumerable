package goe

import "github.com/EscanBE/go-ienumerable/goe/comparers"

type exceptElementHolder struct {
	elementIndex int
	key          any
}

func (src *enumerable[T]) ExceptBy(second IEnumerable[any], requiredKeySelector KeySelector[T], optionalEqualsFunc OptionalEqualsFunc[any]) IEnumerable[T] {
	src.assertSrcNonNil()
	assertSecondIEnumerableNonNil(second)
	assertKeySelectorNonNil(requiredKeySelector)

	var isEquals EqualsFunc[any]
	if optionalEqualsFunc != nil {
		isEquals = EqualsFunc[any](optionalEqualsFunc)
	}

	srcHolders := make([]exceptElementHolder, len(src.data))
	secondData := second.ToArray()

	for i, d1 := range src.data {
		srcHolders[i] = exceptElementHolder{
			elementIndex: i,
			key:          requiredKeySelector(d1),
		}

		if isEquals == nil {
			comparer, found := comparers.TryGetDefaultComparerFromValue(srcHolders[i].key)
			if found {
				isEquals = func(v1, v2 any) bool {
					return comparer.CompareAny(v1, v2) == 0
				}
			}
		}
	}

	if len(secondData) < 1 {
		return src.copyExceptData().withData(distinctByKeySelector(copySlice(src.data), requiredKeySelector, OptionalEqualsFunc[any](isEquals)))
	}

	if len(src.data) < 1 {
		return src.copyExceptData().withEmptyData()
	}

	resultHolders := make([]exceptElementHolder, 0)
	for _, hSource := range srcHolders {
		var foundInAnother bool
		for _, d2 := range secondData {
			if isEquals(hSource.key, d2) {
				foundInAnother = true
				break
			}
		}

		if !foundInAnother {
			var addedPreviously bool

			for _, t := range resultHolders {
				if isEquals(hSource.key, t.key) {
					addedPreviously = true
					break
				}
			}

			if !addedPreviously {
				resultHolders = append(resultHolders, hSource)
			}
		}
	}

	result := make([]T, len(resultHolders))
	for i, resultHolder := range resultHolders {
		result[i] = src.data[resultHolder.elementIndex]
	}

	return src.copyExceptData().withData(result)
}
