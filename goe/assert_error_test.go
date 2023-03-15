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
