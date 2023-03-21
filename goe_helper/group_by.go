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

// GroupByTransform groups the elements of a sequence according to a specified key selector function and creates a result value from each group and its key. Key values are compared by using a specified comparer, and the elements of each group are projected by using a specified function.
func GroupByTransform[TSource, TKey, TElement, TResult any](source goe.IEnumerable[TSource], keySelector func(TSource) TKey, elementSelector func(TSource) TElement, resultSelector func(TKey, goe.IEnumerable[TElement]) TResult, optionalKeyEqualityFunc goe.OptionalEqualsFunc[TKey]) goe.IEnumerable[TResult] {
	if resultSelector == nil {
		panic("result selector is nil")
	}

	groups := GroupBy(source, keySelector, elementSelector, optionalKeyEqualityFunc)
	results := make([]TResult, groups.Count(nil))

	for i, group := range groups.ToArray() {
		results[i] = resultSelector(group.Key, group.Elements)
	}

	return goe.NewIEnumerable[TResult](results...)
}
