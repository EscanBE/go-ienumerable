package goe

import "github.com/EscanBE/go-ienumerable/goe/comparers"

func (src *enumerable[T]) SequenceEqual(second IEnumerable[T], optionalEqualsFunc OptionalEqualsFunc[T]) bool {
	src.assertSrcNonNil()
	assertSecondIEnumerableNonNil(second)

	var secondData []T

	var isEquals EqualsFunc[T]
	if optionalEqualsFunc != nil {
		isEquals = EqualsFunc[T](optionalEqualsFunc)
	} else {
		defaultComparer := src.defaultComparer
		if defaultComparer != nil {
			isEquals = func(v1, v2 T) bool {
				return defaultComparer.CompareAny(v1, v2) == 0
			}
		} else {
			if len(src.data) > 0 {
				for _, ds := range src.data {
					comparer, found := comparers.TryGetDefaultComparerFromValue(ds)
					if found {
						isEquals = func(v1, v2 T) bool {
							return comparer.CompareAny(v1, v2) == 0
						}
						break
					}
				}
			}
			if isEquals == nil && second.Count(nil) > 0 {
				secondData = second.ToArray()
				for _, ds := range secondData {
					comparer, found := comparers.TryGetDefaultComparerFromValue(ds)
					if found {
						isEquals = func(v1, v2 T) bool {
							return comparer.CompareAny(v1, v2) == 0
						}
						break
					}
				}
			}
		}
	}

	if isEquals == nil {
		panic(getErrorFailedCompare2ElementsInArray())
	}

	if len(src.data) != second.Count(nil) {
		return false
	}

	if secondData == nil {
		secondData = second.ToArray()
	}

	for i := 0; i < len(src.data); i++ {
		d1 := src.data[i]
		d2 := secondData[i]

		if !isEquals(d1, d2) {
			return false
		}
	}

	return true
}
