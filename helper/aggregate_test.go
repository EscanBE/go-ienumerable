package helper

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestAggregate(t *testing.T) {
	t.Run("sum", func(t *testing.T) {
		src := goe.NewIEnumerable[int](1, 2, 3, 4)

		result := Aggregate(src, func(pr, v int) int {
			return pr + v
		})

		assert.Equal(t, 10, result)
	})

	t.Run("csv", func(t *testing.T) {
		src := goe.NewIEnumerable[string]("a", "b", "c", "d")

		result := Aggregate(src, func(pr, v string) string {
			return pr + "," + v
		})

		assert.Equal(t, "a,b,c,d", result)
	})

	t.Run("empty", func(t *testing.T) {
		src := goe.NewIEnumerable[int]()

		defer deferExpectPanicContains(t, "collection is empty", true)

		Aggregate(src, func(pr, v int) int {
			return pr + v
		})
	})

	t.Run("function is required", func(t *testing.T) {
		src := goe.NewIEnumerable[string]("a", "b", "c", "d")

		defer deferExpectPanicContains(t, "accumulator function is nil", true)

		_ = Aggregate(src, nil)
	})
}

func TestAggregateSeed(t *testing.T) {
	t.Run("sum", func(t *testing.T) {
		src := goe.NewIEnumerable[int](1, 2, 3, 4)

		result := AggregateSeed(src, 10, func(pr, v int) int {
			return pr + v
		})

		assert.Equal(t, 20, result)
	})

	t.Run("csv", func(t *testing.T) {
		src := goe.NewIEnumerable[string]("a", "b", "c", "d")

		result := AggregateSeed(src, "seed", func(pr, v string) string {
			return pr + "," + v
		})

		assert.Equal(t, "seed,a,b,c,d", result)
	})

	t.Run("empty", func(t *testing.T) {
		src := goe.NewIEnumerable[int]()

		defer deferExpectPanicContains(t, "collection is empty", true)

		AggregateSeed(src, 9, func(pr, v int) int {
			return pr + v
		})
	})

	t.Run("function is required", func(t *testing.T) {
		src := goe.NewIEnumerable[string]("a", "b", "c", "d")

		defer deferExpectPanicContains(t, "accumulator function is nil", true)

		_ = AggregateSeed(src, 9, nil)
	})
}

func TestAggregateSeedTransform(t *testing.T) {
	t.Run("sum", func(t *testing.T) {
		src := goe.NewIEnumerable[int](1, 2, 3, 4)

		result := AggregateSeedTransform(src, 10, func(pr, v int) int {
			return pr + v
		}, func(r int) string {
			return fmt.Sprintf("%d", r)
		})

		assert.Equal(t, "20", result)
	})

	t.Run("csv", func(t *testing.T) {
		src := goe.NewIEnumerable[string]("a", "b", "c", "d")

		result := AggregateSeedTransform(src, "seed", func(pr, v string) string {
			return pr + "," + v
		}, func(r string) string {
			return strings.ToUpper(r)
		})

		assert.Equal(t, "SEED,A,B,C,D", result)
	})

	t.Run("empty", func(t *testing.T) {
		src := goe.NewIEnumerable[int]()

		defer deferExpectPanicContains(t, "collection is empty", true)

		AggregateSeedTransform(src, 9, func(pr, v int) int {
			return pr + v
		}, func(i int) string {
			return fmt.Sprintf("%d", i)
		})
	})

	t.Run("function is required", func(t *testing.T) {
		src := goe.NewIEnumerable[string]("a", "b", "c", "d")

		defer deferExpectPanicContains(t, "accumulator function is nil", true)

		_ = AggregateSeedTransform(src, 9, nil, func(i int) string {
			return fmt.Sprintf("%d", i)
		})
	})

	t.Run("function result selector is required", func(t *testing.T) {
		src := goe.NewIEnumerable[string]("a", "b", "c", "d")

		defer deferExpectPanicContains(t, "result selector function is nil", true)

		_ = AggregateSeedTransform[string, string, string](src, "9", func(pr, v string) string {
			return pr + v
		}, nil)
	})
}
