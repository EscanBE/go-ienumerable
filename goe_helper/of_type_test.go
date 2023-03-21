package goe_helper

import (
	"github.com/EscanBE/go-ienumerable/goe"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOfType(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ie := goe.NewIEnumerable[any](1, "Hello", "World", 2, true)

		got1 := OfType[string](ie).ToArray()
		assert.Len(t, got1, 2)
		assert.Equal(t, "Hello", got1[0])
		assert.Equal(t, "World", got1[1])

		got2 := OfType[int](ie).ToArray()
		assert.Len(t, got2, 2)
		assert.Equal(t, 1, got2[0])
		assert.Equal(t, 2, got2[1])

		got3 := OfType[bool](ie).ToArray()
		assert.Len(t, got3, 1)
		assert.True(t, got3[0])

		got4 := OfType[int8](ie).ToArray()
		assert.Len(t, got4, 0)
	})

	t.Run("success empty", func(t *testing.T) {
		ie := goe.NewIEnumerable[any]()

		got1 := OfType[string](ie).ToArray()
		assert.Len(t, got1, 0)

		got2 := OfType[int](ie).ToArray()
		assert.Len(t, got2, 0)

		got3 := OfType[bool](ie).ToArray()
		assert.Len(t, got3, 0)

		got4 := OfType[int8](ie).ToArray()
		assert.Len(t, got4, 0)
	})

	t.Run("panic when source is nil", func(t *testing.T) {
		var ie goe.IEnumerable[any]
		defer deferExpectPanicContains(t, "source collection is nil", true)
		OfType[string](ie)
	})
}
