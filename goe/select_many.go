package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/EscanBE/go-ienumerable/goe/reflection"
)

func (src *enumerable[T]) SelectMany(selector func(v T) []any) IEnumerable[any] {
	src.assertSrcNonNil()
	src.assertArraySelectorNonNil(selector)

	if len(src.data) < 1 {
		return NewIEnumerable[any]()
	}

	newData := make([]any, 0)

	for _, d := range src.data {
		a := selector(d)

		if a == nil {
			panic("result array can not be nil")
		}

		if len(a) < 1 {
			continue
		}

		newData = append(newData, a...)
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
