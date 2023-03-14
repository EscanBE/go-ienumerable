package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Min_MinBy_MinByComparer(t *testing.T) {
	lessComparer := func(l, r int) bool {
		return l < r
	}

	t.Run("Min* as normal", func(t *testing.T) {
		src := NewIEnumerable[int](4, 7, 5, 6, 3, 2)
		bSrc := backupForAssetUnchanged(src)

		got := src.Min()
		assert.Equal(t, 2, got)

		bSrc.assertUnchanged(t, src)

		got = src.MinBy(lessComparer)
		assert.Equal(t, 2, got)

		bSrc.assertUnchanged(t, src)

		got = src.MinByComparer(comparers.IntComparer)
		assert.Equal(t, 2, got)

		bSrc.assertUnchanged(t, src)
	})

	t.Run("Min* without comparer, using default", func(t *testing.T) {
		src := NewIEnumerable[int](4, 7, 5, 6, 3, 2)
		bSrc := backupForAssetUnchanged(src)

		got := src.MinBy(nil)
		assert.Equal(t, 2, got)

		bSrc.assertUnchanged(t, src)

		got = src.MinByComparer(nil)
		assert.Equal(t, 2, got)

		bSrc.assertUnchanged(t, src)
	})

	t.Run("Min* as without default comparer", func(t *testing.T) {
		src := NewIEnumerable[int](4, 7, 5, 6, 3, 2)
		src.WithDefaultComparer(nil)
		bSrc := backupForAssetUnchanged(src)

		got := src.Min()
		assert.Equal(t, 2, got)

		bSrc.assertUnchanged(t, src)

		got = src.MinBy(nil)
		assert.Equal(t, 2, got)

		bSrc.assertUnchanged(t, src)

		got = src.MinBy(lessComparer)
		assert.Equal(t, 2, got)

		bSrc.assertUnchanged(t, src)

		got = src.MinByComparer(nil)
		assert.Equal(t, 2, got)

		bSrc.assertUnchanged(t, src)

		got = src.MinByComparer(comparers.IntComparer)
		assert.Equal(t, 2, got)

		bSrc.assertUnchanged(t, src)
	})

	t.Run("Min panic empty", func(t *testing.T) {
		defer deferExpectPanicContains(t, "source contains no element", true)
		_ = NewIEnumerable[int]().Min()
	})

	t.Run("MinBy panic empty", func(t *testing.T) {
		defer deferExpectPanicContains(t, "source contains no element", true)
		_ = NewIEnumerable[int]().MinBy(nil)
	})

	t.Run("MinByComparer panic empty", func(t *testing.T) {
		defer deferExpectPanicContains(t, "source contains no element", true)
		_ = NewIEnumerable[int]().MinByComparer(nil)
	})

	t.Run("Min panic no comparer", func(t *testing.T) {
		type MyInt64 struct{}
		defer deferExpectPanicContains(t, "no default comparer registered", true)
		_ = NewIEnumerable[MyInt64](MyInt64{}).Min()
	})

	t.Run("MinBy panic no comparer", func(t *testing.T) {
		type MyInt64 struct{}
		defer deferExpectPanicContains(t, "no default comparer registered", true)
		_ = NewIEnumerable[MyInt64](MyInt64{}).MinBy(nil)
	})

	t.Run("MinByComparer panic no comparer", func(t *testing.T) {
		type MyInt64 struct{}
		defer deferExpectPanicContains(t, "no default comparer registered", true)
		_ = NewIEnumerable[MyInt64](MyInt64{}).MinByComparer(nil)
	})
}
