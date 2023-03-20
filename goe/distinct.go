package goe

import "github.com/EscanBE/go-ienumerable/goe/comparers"

func (src *enumerable[T]) Distinct(optionalEqualsFunc OptionalEqualsFunc[T]) IEnumerable[T] {
	src.assertSrcNonNil()

	var equalityFunc OptionalEqualsFunc[T]

	if optionalEqualsFunc == nil {
		defaultComparer := src.defaultComparer
		if defaultComparer == nil {
			defaultComparer = src.findDefaultComparer()
		}
		equalityFunc = func(v1, v2 T) bool {
			return defaultComparer.CompareAny(v1, v2) == 0
		}
	} else {
		equalityFunc = optionalEqualsFunc
	}

	uniqueData := distinct(src.data, equalityFunc)

	return src.copyExceptData().withData(uniqueData)
}

func (src *enumerable[T]) DistinctBy(requiredKeySelector KeySelector[T], optionalEqualsFunc OptionalEqualsFunc[any]) IEnumerable[T] {
	src.assertSrcNonNil()
	assertKeySelectorNonNil(requiredKeySelector)

	unique := distinctByKeySelector(copySlice(src.data), requiredKeySelector, optionalEqualsFunc)
	return src.copyExceptData().withData(unique)
}

func distinct[T any](data []T, optionalEqualityComparer OptionalEqualsFunc[T]) []T {
	var equalityComparer EqualsFunc[T]

	if optionalEqualityComparer != nil {
		equalityComparer = EqualsFunc[T](optionalEqualityComparer)
	} else {
		comparer, found := comparers.TryGetDefaultComparer[T]()
		if !found {
			panic(getErrorFailedCompare2ElementsInArray())
		}
		equalityComparer = func(v1, v2 T) bool {
			return comparer.CompareAny(v1, v2) == 0
		}
	}

	if len(data) < 2 {
		return data
	}

	uniqueSet := []T{data[0]}

	for i1 := 1; i1 < len(data); i1++ {
		ele := data[i1]

		var exists bool
		for _, uniq := range uniqueSet {
			if equalityComparer(ele, uniq) {
				exists = true
				break
			}
		}

		if !exists {
			uniqueSet = append(uniqueSet, ele)
		}
	}

	return uniqueSet
}

func distinctByKeySelector[T any](data []T, requiredKeySelector KeySelector[T], optionalEqualityComparer OptionalEqualsFunc[any]) []T {
	if requiredKeySelector == nil {
		panic(getErrorKeySelectorNotNil())
	}

	var equalityComparer EqualsFunc[any]
	if optionalEqualityComparer != nil {
		equalityComparer = EqualsFunc[any](optionalEqualityComparer)
	}

	type holder struct {
		elementIndex int
		key          any
	}

	holders := make([]holder, len(data))

	for i, d := range data {
		holders[i] = holder{
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

	uniqueSet := []holder{holders[0]}

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
