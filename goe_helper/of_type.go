package goe_helper

import "github.com/EscanBE/go-ienumerable/goe"

// OfType filters the elements of an IEnumerable[any] based on a specified type.
func OfType[TResult any](source goe.IEnumerable[any]) goe.IEnumerable[TResult] {
	assertCollectionNotNil(source, "source")

	result := make([]TResult, 0)

	for _, src := range source.ToArray() {
		if casted, ok := src.(TResult); ok {
			result = append(result, casted)
		}
	}

	return goe.NewIEnumerable[TResult](result...)
}
