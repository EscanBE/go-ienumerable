package helper

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZipF(t *testing.T) {
	first := goe.NewIEnumerable[int](4, 5, 6, 7)
	second := goe.NewIEnumerable[string]("9", "8", "7")
	emptyIntSeq := goe.NewIEnumerable[int]()
	emptyStrSeq := goe.NewIEnumerable[string]()

	resultSelector := func(i1 int, i2 string) string {
		return fmt.Sprintf("%d %s", i1, i2)
	}

	t.Run("one empty sequence", func(t *testing.T) {
		assert.Empty(t, ZipF(first, emptyStrSeq, resultSelector).ToArray())
		assert.Empty(t, ZipF(emptyIntSeq, second, resultSelector).ToArray())
	})

	t.Run("zip success", func(t *testing.T) {
		zipped := ZipF(first, second, resultSelector).ToArray()
		assert.Len(t, zipped, 3)
		assert.Equal(t, "4 9", zipped[0])
		assert.Equal(t, "5 8", zipped[1])
		assert.Equal(t, "6 7", zipped[2])
	})

	t.Run("require result selector", func(t *testing.T) {
		defer deferExpectPanicContains(t, "result selector is required", true)

		ZipF[int, string, string](first, second, nil)
	})

	t.Run("require first sequence", func(t *testing.T) {
		defer deferExpectPanicContains(t, "first collection is nil", true)

		ZipF[int, string, string](nil, second, resultSelector)
	})

	t.Run("require second sequence", func(t *testing.T) {
		defer deferExpectPanicContains(t, "second collection is nil", true)

		ZipF[int, string, string](first, nil, resultSelector)
	})
}

func TestZip2(t *testing.T) {
	first := goe.NewIEnumerable[int](4, 5, 6, 7)
	second := goe.NewIEnumerable[string]("9", "8", "7")
	emptyIntSeq := goe.NewIEnumerable[int]()
	emptyStrSeq := goe.NewIEnumerable[string]()

	t.Run("one empty sequence", func(t *testing.T) {
		assert.Empty(t, Zip2(emptyIntSeq, second).ToArray())
		assert.Empty(t, Zip2(first, emptyStrSeq).ToArray())
	})

	t.Run("zip success", func(t *testing.T) {
		zipped := Zip2(first, second).ToArray()
		assert.Len(t, zipped, 3)
		assert.Equal(t, 4, zipped[0].First)
		assert.Equal(t, "9", zipped[0].Second)
		assert.Equal(t, 5, zipped[1].First)
		assert.Equal(t, "8", zipped[1].Second)
		assert.Equal(t, 6, zipped[2].First)
		assert.Equal(t, "7", zipped[2].Second)
	})

	t.Run("require first sequence", func(t *testing.T) {
		defer deferExpectPanicContains(t, "first collection is nil", true)

		Zip2[int, string](nil, second)
	})

	t.Run("require second sequence", func(t *testing.T) {
		defer deferExpectPanicContains(t, "second collection is nil", true)

		Zip2[int, string](first, nil)
	})
}

func TestZip3(t *testing.T) {
	first := goe.NewIEnumerable[int](4, 5, 6, 7)
	second := goe.NewIEnumerable[string]("9", "8", "7")
	third := goe.NewIEnumerable[float64](9.9, 6.6, 8.8)
	emptyIntSeq := goe.NewIEnumerable[int]()
	emptyStrSeq := goe.NewIEnumerable[string]()
	emptyFloatSeq := goe.NewIEnumerable[float64]()

	t.Run("one empty sequence", func(t *testing.T) {
		assert.Empty(t, Zip3(emptyIntSeq, second, third).ToArray())
		assert.Empty(t, Zip3(first, emptyStrSeq, third).ToArray())
		assert.Empty(t, Zip3(first, second, emptyFloatSeq).ToArray())
	})

	t.Run("zip success", func(t *testing.T) {
		zipped := Zip3(first, second, third).ToArray()
		assert.Len(t, zipped, 3)
		assert.Equal(t, 4, zipped[0].First)
		assert.Equal(t, "9", zipped[0].Second)
		assert.Equal(t, 9.9, zipped[0].Third)
		assert.Equal(t, 5, zipped[1].First)
		assert.Equal(t, "8", zipped[1].Second)
		assert.Equal(t, 6.6, zipped[1].Third)
		assert.Equal(t, 6, zipped[2].First)
		assert.Equal(t, "7", zipped[2].Second)
		assert.Equal(t, 8.8, zipped[2].Third)
	})

	t.Run("require first sequence", func(t *testing.T) {
		defer deferExpectPanicContains(t, "first collection is nil", true)

		Zip3[int, string, float64](nil, second, third)
	})

	t.Run("require second sequence", func(t *testing.T) {
		defer deferExpectPanicContains(t, "second collection is nil", true)

		Zip3[int, string, float64](first, nil, third)
	})

	t.Run("require third sequence", func(t *testing.T) {
		defer deferExpectPanicContains(t, "third collection is nil", true)

		Zip3[int, string, float64](first, second, nil)
	})
}
