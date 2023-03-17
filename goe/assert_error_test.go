package goe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_assertSrcNonNil(t *testing.T) {
	e := new(enumerable[int])
	e = nil

	defer deferExpectPanicContains(t, getErrorSourceIsNil().Error(), true)

	e.assertSrcNonNil()
}

func Test_enumerable_assertSrcNonEmpty(t *testing.T) {
	e := new(enumerable[int])

	defer deferExpectPanicContains(t, getErrorSrcContainsNoElement().Error(), true)

	e.assertSrcNonEmpty()
}

func Test_enumerable_assertPredicateNonNil(t *testing.T) {
	e := new(enumerable[int])

	defer deferExpectPanicContains(t, getErrorNilPredicate().Error(), true)

	e.assertPredicateNonNil(nil)
}

func Test_enumerable_assertSelectorNonNil(t *testing.T) {
	t.Run("selector", func(t *testing.T) {
		e := new(enumerable[int])

		defer deferExpectPanicContains(t, getErrorNilSelector().Error(), true)

		e.assertSelectorNonNil(nil)
	})

	t.Run("array selector", func(t *testing.T) {
		e := new(enumerable[int])

		defer deferExpectPanicContains(t, getErrorNilSelector().Error(), true)

		e.assertArraySelectorNonNil(nil)
	})
}

func Test_enumerable_assertAggregateFuncNonNil(t *testing.T) {
	t.Run("aggregate func", func(t *testing.T) {
		e := new(enumerable[int])

		defer deferExpectPanicContains(t, getErrorNilAggregateFunc().Error(), true)

		e.assertAggregateFuncNonNil(nil)
	})

	t.Run("aggregate any seed func", func(t *testing.T) {
		e := new(enumerable[int])

		defer deferExpectPanicContains(t, getErrorNilAggregateFunc().Error(), true)

		e.assertAggregateAnySeedFuncNonNil(nil)
	})
}

func Test_enumerable_findDefaultComparer(t *testing.T) {
	t.Run("resolve registered", func(t *testing.T) {
		assert.NotNil(t, e[int32](NewIEnumerable[int32]()).findDefaultComparer())
	})

	t.Run("panic if type not registered", func(t *testing.T) {
		type MyInt64 int64

		e := e[MyInt64](NewIEnumerable[MyInt64]())

		defer deferExpectPanicContains(t, "no default comparer registered for [goe.MyInt64]", true)

		e.findDefaultComparer()
	})
}

func Test_enumerable_assertComparerNonNil(t *testing.T) {
	t.Run("comparer func", func(t *testing.T) {
		e := new(enumerable[int])

		defer deferExpectPanicContains(t, getErrorNilComparer().Error(), true)

		e.assertComparerNonNil(nil)
	})
}

func Test_enumerable_assertPredicateNonNil1(t *testing.T) {
	var fPredicate1 func(v int) bool
	var fPredicate2 Predicate[int]
	var fPredicate3 func(v int, i int) bool
	var fPredicate4 PredicateWithIndex[int]
	tests := []struct {
		name      string
		predicate interface{}
	}{
		{
			name:      "nil",
			predicate: nil,
		},
		{
			name:      "nil",
			predicate: fPredicate1,
		},
		{
			name:      "nil",
			predicate: fPredicate2,
		},
		{
			name:      "nil",
			predicate: fPredicate3,
		},
		{
			name:      "nil",
			predicate: fPredicate4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer deferExpectPanicContains(t, getErrorNilPredicate().Error(), true)
			e[int](createEmptyIntEnumerable()).assertPredicateNonNil(tt.predicate)
		})
	}

	t.Run("not a valid predicate", func(t *testing.T) {
		defer deferExpectPanicContains(t, getErrorPredicateMustBePredicate().Error(), true)
		e[int](createEmptyIntEnumerable()).assertPredicateNonNil(func(v1, v2 int) int {
			return v1 * v2
		})
	})
}
