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

	for i, v := range src.data {
		d := selector(v)
		newData[i] = d
		trackerTypes[fmt.Sprintf("%T", d)] = true
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
