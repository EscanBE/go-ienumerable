package goe_helper

import (
	"github.com/EscanBE/go-ienumerable/goe"
	"github.com/EscanBE/go-ienumerable/goe/comparers"
)

// GroupJoin correlates the elements of two sequences based on key equality and groups the results. An option equality function is used to compare keys.
func GroupJoin[TOuter, TInner, TKey, TResult any](outer goe.IEnumerable[TOuter], inner goe.IEnumerable[TInner], outerKeySelector func(TOuter) TKey, innerKeySelector func(TInner) TKey, resultSelector func(TOuter, goe.IEnumerable[TInner]) TResult, optionalKeyEqualityFunc goe.OptionalEqualsFunc[TKey]) goe.IEnumerable[TResult] {
	if outer == nil {
		panic("outer collection is nil")
	}
	if inner == nil {
		panic("inner collection is nil")
	}
	if outerKeySelector == nil {
		panic("outer key selector is nil")
	}
	if innerKeySelector == nil {
		panic("inner key selector is nil")
	}
	if resultSelector == nil {
		panic("result selector is nil")
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

	type outerHolder struct {
		outer TOuter
		key   TKey
	}

	type innerHolder struct {
		inner TInner
		key   TKey
	}

	outerHolders := Select(outer, func(v TOuter) outerHolder {
		return outerHolder{
			outer: v,
			key:   outerKeySelector(v),
		}
	}).ToArray()

	innerHolders := Select(inner, func(v TInner) innerHolder {
		return innerHolder{
			inner: v,
			key:   innerKeySelector(v),
		}
	}).ToArray()

	results := make([]TResult, 0)

	for _, outer := range outerHolders {
		groupsOfInner := make([]TInner, 0)
		for _, inner := range innerHolders {
			if keyEqualityComparer(outer.key, inner.key) {
				groupsOfInner = append(groupsOfInner, inner.inner)
			}
		}
		results = append(results, resultSelector(outer.outer, goe.NewIEnumerable[TInner](groupsOfInner...)))
	}

	return goe.NewIEnumerable[TResult](results...)
}
