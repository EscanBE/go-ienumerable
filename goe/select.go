package goe

import "fmt"

func (src *enumerable[T]) Select(selector func(v T) any) IEnumerable[any] {
	src.assertSrcNonNil()
	src.assertSelectorNonNil(selector)

	if len(src.data) < 1 {
		return NewIEnumerable[any]()
	}

	newData := make([]any, len(src.data))

	trackerTypes := make(map[string]bool)

	for i, d := range src.data {
		v := selector(d)
		newData[i] = v
		trackerTypes[fmt.Sprintf("%T", v)] = true
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

func (src *enumerable[T]) SelectWithSampleValueOfResult(selector func(v T) any, notNilSampleResultValue any) IEnumerable[any] {
	src.assertSrcNonNil()
	src.assertSelectorNonNil(selector)
	src.assertSampleResultValueNonNil(notNilSampleResultValue)

	sampleResultType := fmt.Sprintf("%T", notNilSampleResultValue)

	newData := make([]any, len(src.data))

	if len(src.data) > 0 {
		for i, d := range src.data {
			v := selector(d)
			newData[i] = v
			resultType := fmt.Sprintf("%T", v)
			if sampleResultType != resultType {
				panic(fmt.Sprintf("sample result is type [%s] but got result %v of type [%s]", sampleResultType, v, resultType))
			}
		}
	}

	result := NewIEnumerable[any](newData...)

	eResult := e[any](result)
	eResult.dataType = sampleResultType
	eResult.injectDefaultComparer()

	return result
}
