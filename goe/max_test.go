package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Max_MaxBy_MaxByComparer(t *testing.T) {
	greaterComparer := func(l, r int) bool {
		return l > r
	}

	t.Run("Max* as normal", func(t *testing.T) {
		src := NewIEnumerable[int](4, 7, 5, 6, 3, 2)
		bSrc := backupForAssetUnchanged(src)

		got := src.Max()
		assert.Equal(t, 7, got)

		bSrc.assertUnchanged(t, src)

		got = src.MaxBy(greaterComparer)
		assert.Equal(t, 7, got)

		bSrc.assertUnchanged(t, src)

		got = src.MaxByComparer(comparers.IntComparer)
		assert.Equal(t, 7, got)

		bSrc.assertUnchanged(t, src)
	})

	t.Run("Max* without comparer, using default", func(t *testing.T) {
		src := NewIEnumerable[int](4, 7, 5, 6, 3, 2)
		bSrc := backupForAssetUnchanged(src)

		got := src.MaxBy(nil)
		assert.Equal(t, 7, got)

		bSrc.assertUnchanged(t, src)

		got = src.MaxByComparer(nil)
		assert.Equal(t, 7, got)

		bSrc.assertUnchanged(t, src)
	})

	t.Run("Max* as without default comparer", func(t *testing.T) {
		src := NewIEnumerable[int](4, 7, 5, 6, 3, 2)
		src.WithDefaultComparer(nil)
		bSrc := backupForAssetUnchanged(src)

		got := src.Max()
		assert.Equal(t, 7, got)

		bSrc.assertUnchanged(t, src)

		got = src.MaxBy(nil)
		assert.Equal(t, 7, got)

		bSrc.assertUnchanged(t, src)

		got = src.MaxBy(greaterComparer)
		assert.Equal(t, 7, got)

		bSrc.assertUnchanged(t, src)

		got = src.MaxByComparer(nil)
		assert.Equal(t, 7, got)

		bSrc.assertUnchanged(t, src)

		got = src.MaxByComparer(comparers.IntComparer)
		assert.Equal(t, 7, got)

		bSrc.assertUnchanged(t, src)
	})

	t.Run("Max panic empty", func(t *testing.T) {
		defer deferExpectPanicContains(t, "source contains no element", true)
		_ = NewIEnumerable[int]().Max()
	})

	t.Run("MaxBy panic empty", func(t *testing.T) {
		defer deferExpectPanicContains(t, "source contains no element", true)
		_ = NewIEnumerable[int]().MaxBy(nil)
	})

	t.Run("MaxByComparer panic empty", func(t *testing.T) {
		defer deferExpectPanicContains(t, "source contains no element", true)
		_ = NewIEnumerable[int]().MaxByComparer(nil)
	})

	t.Run("Max panic no comparer", func(t *testing.T) {
		type MyInt64 struct{}
		defer deferExpectPanicContains(t, "no default comparer registered", true)
		_ = NewIEnumerable[MyInt64](MyInt64{}).Max()
	})

	t.Run("MaxBy panic no comparer", func(t *testing.T) {
		type MyInt64 struct{}
		defer deferExpectPanicContains(t, "no default comparer registered", true)
		_ = NewIEnumerable[MyInt64](MyInt64{}).MaxBy(nil)
	})

	t.Run("MaxByComparer panic no comparer", func(t *testing.T) {
		type MyInt64 struct{}
		defer deferExpectPanicContains(t, "no default comparer registered", true)
		_ = NewIEnumerable[MyInt64](MyInt64{}).MaxByComparer(nil)
	})
}
