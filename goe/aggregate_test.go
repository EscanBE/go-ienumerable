package goe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Aggregate(t *testing.T) {
	t.Run("sum", func(t *testing.T) {
		e := NewIEnumerable[int](1, 2, 3, 4)

		result := e.Aggregate(func(pr, v int) int {
			return pr + v
		})

		assert.Equal(t, 10, result)
	})

	t.Run("csv", func(t *testing.T) {
		e := NewIEnumerable[string]("a", "b", "c", "d")

		result := e.Aggregate(func(pr, v string) string {
			return pr + "," + v
		})

		assert.Equal(t, "a,b,c,d", result)
	})

	t.Run("empty", func(t *testing.T) {
		e := NewIEnumerable[int]()

		defer deferExpectPanicContains(t, getErrorSrcContainsNoElement().Error(), true)

		e.Aggregate(func(pr, v int) int {
			return pr + v
		})
	})
}

func Test_enumerable_AggregateWithSeed(t *testing.T) {
	t.Run("sum", func(t *testing.T) {
		e := NewIEnumerable[int](1, 2, 3, 4)

		result := e.AggregateWithSeed(10, func(pr, v int) int {
			return pr + v
		})

		assert.Equal(t, 20, result)
	})

	t.Run("csv", func(t *testing.T) {
		e := NewIEnumerable[string]("a", "b", "c", "d")

		result := e.AggregateWithSeed("seed", func(pr, v string) string {
			return pr + "," + v
		})

		assert.Equal(t, "seed,a,b,c,d", result)
	})

	t.Run("empty", func(t *testing.T) {
		e := NewIEnumerable[int]()

		defer deferExpectPanicContains(t, getErrorSrcContainsNoElement().Error(), true)

		e.AggregateWithSeed(9, func(pr, v int) int {
			return pr + v
		})
	})
}

func Test_enumerable_AggregateWithAnySeed(t *testing.T) {
	t.Run("sum", func(t *testing.T) {
		e := NewIEnumerable[int](1, 2, 3, 4)

		result := e.AggregateWithAnySeed(10, func(pr any, v int) any {
			return pr.(int) + v
		})

		assert.Equal(t, 20, result)
	})

	t.Run("csv", func(t *testing.T) {
		e := NewIEnumerable[string]("a", "b", "c", "d")

		result := e.AggregateWithAnySeed("seed", func(pr any, v string) any {
			return pr.(string) + "," + v
		})

		assert.Equal(t, "seed,a,b,c,d", result)
	})

	t.Run("empty", func(t *testing.T) {
		e := NewIEnumerable[int]()

		defer deferExpectPanicContains(t, getErrorSrcContainsNoElement().Error(), true)

		e.AggregateWithAnySeed(9, func(pr any, v int) any {
			return pr.(int) + v
		})
	})
}
