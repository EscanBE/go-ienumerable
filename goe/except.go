package goe

import "github.com/EscanBE/go-ienumerable/goe/comparers"

func (src *enumerable[T]) Except(second IEnumerable[T], optionalEqualsFunc OptionalEqualsFunc[T]) IEnumerable[T] {
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

	if second.Count(nil) < 1 {
		return src.copyExceptData().withData(distinct(copySlice(src.data), OptionalEqualsFunc[T](isEquals)))
	}

	if len(src.data) < 1 {
		return src.copyExceptData().withEmptyData()
	}

	result := make([]T, 0)
	secondData := second.ToArray()
	for _, d := range src.data {
		var foundInAnother bool
		for _, t := range secondData {
			if isEquals(d, t) {
				foundInAnother = true
				break
			}
		}
		if !foundInAnother {
			var addedPreviously bool

			for _, t := range result {
				if isEquals(d, t) {
					addedPreviously = true
					break
				}
			}

			if !addedPreviously {
				result = append(result, d)
			}
		}
	}

	return src.copyExceptData().withData(result)
}

func (src *enumerable[T]) ExceptBy(second IEnumerable[any], requiredKeySelector KeySelector[T], optionalEqualsFunc OptionalEqualsFunc[any]) IEnumerable[T] {
	src.assertSrcNonNil()
	assertSecondIEnumerableNonNil(second)
	assertKeySelectorNonNil(requiredKeySelector)

	var isEquals EqualsFunc[any]
	if optionalEqualsFunc != nil {
		isEquals = EqualsFunc[any](optionalEqualsFunc)
	}

	type holder struct {
		elementIndex int
		key          any
	}

	srcHolders := make([]holder, len(src.data))
	secondData := second.ToArray()

	for i, d1 := range src.data {
		srcHolders[i] = holder{
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

	resultHolders := make([]holder, 0)
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
