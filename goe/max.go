package goe

import comparers "github.com/EscanBE/go-ienumerable/goe/comparers"

func (src *enumerable[T]) Max() T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

	comparer := src.defaultComparer
	if comparer == nil {
		comparer = src.findDefaultComparer()
	}

	greaterComparer := func(v1, v2 T) bool {
		return comparer.CompareAny(v1, v2) > 0
	}

	maxIdx := 0

	for i := 1; i < len(src.data); i++ {
		if greaterComparer(src.data[i], src.data[maxIdx]) {
			maxIdx = i
		}
	}

	return src.data[maxIdx]
}

func (src *enumerable[T]) MaxBy(requiredKeySelector KeySelector[T], optionalCompareFunc OptionalCompareFunc[any]) T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()
	assertKeySelectorNonNil(requiredKeySelector)

	keys := make([]any, len(src.data))
	for i, t := range src.data {
		keys[i] = requiredKeySelector(t)
		if optionalCompareFunc == nil {
			comparer, found := comparers.TryGetDefaultComparerFromValue(keys[i])
			if found {
				optionalCompareFunc = func(v1, v2 any) int {
					return comparer.CompareAny(v1, v2)
				}
			}
		}
	}

	if optionalCompareFunc == nil {
		panic(getErrorFailedCompare2ElementsInArray())
	}

	keyIdx := 0
	for i := 1; i < len(src.data); i++ {
		if optionalCompareFunc(keys[i], keys[keyIdx]) > 0 {
			keyIdx = i
		}
	}

	return src.data[keyIdx]
}
