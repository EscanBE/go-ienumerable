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
		ds := selector(d)
		if len(ds) < 1 {
			continue
		}

		newData = append(newData, ds...)

		for _, v := range ds {
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
