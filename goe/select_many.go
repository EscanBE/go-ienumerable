package goe

import "fmt"

func (src *enumerable[T]) SelectMany(selector func(v T) []any) IEnumerable[any] {
	src.assertSrcNonNil()
	src.assertArraySelectorNonNil(selector)

	if len(src.data) < 1 {
		return NewIEnumerable[any]()
	}

	newData := make([]any, 0)

	trackerTypes := make(map[string]bool)

	for _, d := range src.data {
		a := selector(d)

		if a == nil {
			panic("result array can not be nil")
		}

		if len(a) < 1 {
			continue
		}

		newData = append(newData, a...)

		for _, v := range a {
			trackerTypes[fmt.Sprintf("%T", v)] = true
		}
	}

	uniqueTypes := getMapKeys(trackerTypes)

	result := NewIEnumerable[any](newData...)

	if len(uniqueTypes) == 1 {
		dataType := uniqueTypes[0]
		if len(dataType) > 0 {
			eResult := e[any](result)
			eResult.dataType = dataType
			eResult.injectDefaultComparer()
		}
	}

	return result
}

func (src *enumerable[T]) SelectManyWithSampleValueOfResult(selector func(v T) []any, notNilSampleResultValue any) IEnumerable[any] {
	src.assertSrcNonNil()
	src.assertArraySelectorNonNil(selector)
	src.assertSampleResultValueNonNil(notNilSampleResultValue)

	sampleResultType := fmt.Sprintf("%T", notNilSampleResultValue)

	newData := make([]any, 0)

	if len(src.data) > 0 {
		for i1, d := range src.data {
			a := selector(d)

			if a == nil {
				panic("result array can not be nil")
			}

			if len(a) < 1 {
				continue
			}

			newData = append(newData, a...)

			for i2, v := range a {
				if v == nil {
					continue
				}

				resultType := fmt.Sprintf("%T", v)
				if sampleResultType != resultType {
					panic(fmt.Sprintf("sample result at index %d yields by element at index %d is type [%s] but got result %v of type [%s]", i2, i1, sampleResultType, v, resultType))
				}
			}
		}
	}

	result := NewIEnumerable[any](newData...)

	eResult := e[any](result)
	eResult.dataType = sampleResultType
	eResult.injectDefaultComparer()

	return result
}
