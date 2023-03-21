package goe_helper

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe"
)

// ToDictionary creates a map[TKey]TElement from an IEnumerable[TSource] according to a specified key selector function, and an element selector function.
// Due to limitation of Go, only comparable types can be used as map key so no comparer is needed.
func ToDictionary[TSource any, TKey comparable, TElement any](source goe.IEnumerable[TSource], keySelector func(TSource) TKey, elementSelector func(TSource) TElement) map[TKey]TElement {
	assertCollectionNotNil(source, "source")
	if keySelector == nil {
		panic("key selector is nil")
	}
	if elementSelector == nil {
		panic("element selector is nil")
	}

	result := make(map[TKey]TElement)
	for _, src := range source.ToArray() {
		key := keySelector(src)
		if _, exists := result[key]; exists {
			panic(fmt.Sprintf("duplicated key %v", key))
		}
		result[key] = elementSelector(src)
	}

	return result
}
