package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/EscanBE/go-ienumerable/goe/reflection"
)

func (src *enumerable[T]) Select(selector func(v T) any) IEnumerable[any] {
	src.assertSrcNonNil()
	src.assertSelectorNonNil(selector)

	if len(src.data) < 1 {
		return NewIEnumerable[any]()
	}

	newData := make([]any, len(src.data))

	for i, d := range src.data {
		v := selector(d)
		newData[i] = v
	}

	result := NewIEnumerable[any](newData...)

	if len(newData) > 0 {
		var nextDefaultComparer comparers.IComparer[any]
		var nextDateType string

		for _, d := range newData {
			if d == nil {
				continue
			}

			if nextDefaultComparer != nil && len(nextDateType) > 0 {
				break
			}

			if nextDefaultComparer == nil {
				comparer, found := comparers.TryGetDefaultComparerFromValue(d)
				if found {
					nextDefaultComparer = comparer
				}
			}

			if len(nextDateType) < 1 {
				vo, isNil := reflection.RootValueExtractor(d)
				if !isNil {
					nextDateType = vo.Type().String()
				}
			}
		}

		eResult := e[any](result)
		if nextDefaultComparer != nil {
			eResult.defaultComparer = nextDefaultComparer
		}
		if len(nextDateType) > 0 {
			eResult.dataType = nextDateType
		}
	}

	return result
}

func (src *enumerable[T]) SelectNewValue(selector func(v T) T) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertSelectorSameNonNil(selector)

	result := src.copyExceptData()

	if len(src.data) < 1 {
		result = result.withEmptyData()
	} else {
		newData := make([]T, len(src.data))

		for i, d := range src.data {
			newData[i] = selector(d)
		}

		result = result.withData(newData)
	}

	return result
}
