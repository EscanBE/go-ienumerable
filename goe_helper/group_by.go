package goe_helper

import (
	"github.com/EscanBE/go-ienumerable/goe"
	"github.com/EscanBE/go-ienumerable/goe/comparers"
)

// GroupBy groups the elements of a sequence according to a key selector function. The keys are compared by using a comparer and each group's elements are projected by using a specified function.
func GroupBy[TSource, TKey, TElement any](source goe.IEnumerable[TSource], keySelector func(TSource) TKey, elementSelector func(TSource) TElement, optionalKeyEqualityFunc goe.OptionalEqualsFunc[TKey]) goe.IEnumerable[goe.Group[TKey, goe.IEnumerable[TElement]]] {
	assertCollectionNotNil(source, "source")
	if keySelector == nil {
		panic("key selector is nil")
	}
	if elementSelector == nil {
		panic("element selector is nil")
	}

	var keyEqualityComparer goe.EqualsFunc[TKey]
	if optionalKeyEqualityFunc != nil {
		keyEqualityComparer = goe.EqualsFunc[TKey](optionalKeyEqualityFunc)
	} else {
		comparer, found := comparers.TryGetDefaultComparer[TKey]()
		if !found {
			panic("no default comparer registered for key type")
		}
		keyEqualityComparer = func(k1, k2 TKey) bool {
			return comparer.CompareAny(k1, k2) == 0
		}
	}

	type holder struct {
		key      TKey
		elements []TElement
	}

	holders := make([]holder, 0)

	for _, src := range source.ToArray() {
		key := keySelector(src)

		var exist = false
		for i, h := range holders {
			if keyEqualityComparer(h.key, key) {
				exist = true
				h.elements = append(h.elements, elementSelector(src))
				holders[i] = h
				break
			}
		}

		if !exist {
			holders = append(holders, holder{
				key:      key,
				elements: []TElement{elementSelector(src)},
			})
		}
	}

	data := make([]goe.Group[TKey, goe.IEnumerable[TElement]], len(holders))

	for i, h := range holders {
		data[i] = goe.Group[TKey, goe.IEnumerable[TElement]]{
			Key:      h.key,
			Elements: goe.NewIEnumerable[TElement](h.elements...),
		}
	}

	return goe.NewIEnumerable[goe.Group[TKey, goe.IEnumerable[TElement]]](data...)
}
