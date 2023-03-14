package goe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_Order(t *testing.T) {
	lessComparer := func(l, r int) bool {
		return l < r
	}

	t.Run("OrderBy", func(t *testing.T) {
		src := NewIEnumerable[int](4, 7, 5, 6, 3, 2)
		bSrc := backupForAssetUnchanged(src)

		got := src.OrderBy(lessComparer)
		assert.True(t, reflect.DeepEqual([]int{2, 3, 4, 5, 6, 7}, got.ToArray()))

		bSrc.assertUnchanged(t, src)
		bSrc.assertUnchangedIgnoreData(t, got)

		src = createEmptyIntEnumerable()
		got = src.OrderBy(lessComparer)
		assert.Empty(t, got.ToArray())

		defer deferWantPanicDepends(t, true)
		_ = src.OrderBy(nil)
	})

	t.Run("OrderByDescendingBy", func(t *testing.T) {
		src := NewIEnumerable[int](4, 7, 5, 6, 3, 2)
		bSrc := backupForAssetUnchanged(src)

		got := src.OrderByDescendingBy(lessComparer)
		assert.True(t, reflect.DeepEqual([]int{7, 6, 5, 4, 3, 2}, got.ToArray()))

		bSrc.assertUnchanged(t, src)
		bSrc.assertUnchangedIgnoreData(t, got)

		src = createEmptyIntEnumerable()
		got = src.OrderByDescendingBy(lessComparer)
		assert.Empty(t, got.ToArray())

		defer deferWantPanicDepends(t, true)
		_ = src.OrderByDescendingBy(nil)
	})

	t.Run("Order", func(t *testing.T) {
		src := injectIntComparers(NewIEnumerable[int](4, 7, 5, 6, 3, 2))
		bSrc := backupForAssetUnchanged(src)

		got := src.Order()
		assert.True(t, reflect.DeepEqual([]int{2, 3, 4, 5, 6, 7}, got.ToArray()))

		bSrc.assertUnchanged(t, src)
		bSrc.assertUnchangedIgnoreData(t, got)

		src = injectIntComparers(createEmptyIntEnumerable())
		got = src.Order()
		assert.Empty(t, got.ToArray())

		defer deferWantPanicDepends(t, true)
		_ = createEmptyIntEnumerable().Order()
	})

	t.Run("OrderByDescending", func(t *testing.T) {
		src := injectIntComparers(NewIEnumerable[int](4, 7, 5, 6, 3, 2))
		bSrc := backupForAssetUnchanged(src)

		got := src.OrderByDescending()
		assert.True(t, reflect.DeepEqual([]int{7, 6, 5, 4, 3, 2}, got.ToArray()))

		bSrc.assertUnchanged(t, src)
		bSrc.assertUnchangedIgnoreData(t, got)

		src = injectIntComparers(createEmptyIntEnumerable())
		got = src.OrderByDescending()
		assert.Empty(t, got.ToArray())

		defer deferWantPanicDepends(t, true)
		_ = createEmptyIntEnumerable().OrderByDescending()
	})

	t.Run("Same order", func(t *testing.T) {
		wrongLessComparer := func(l, r int) bool {
			return l == r
		}
		src := createRandomIntEnumerable(1_000)
		bSrc := backupForAssetUnchanged(src)

		srcData := copySlice(src.ToArray())

		got := src.OrderBy(wrongLessComparer)
		assert.True(t, reflect.DeepEqual(srcData, got.ToArray()))

		got = src.OrderByDescendingBy(wrongLessComparer)
		assert.True(t, reflect.DeepEqual(srcData, got.ToArray()))

		src.WithLessComparer(wrongLessComparer)

		got = src.Order()
		assert.True(t, reflect.DeepEqual(srcData, got.ToArray()))

		got = src.OrderByDescending()
		assert.True(t, reflect.DeepEqual(srcData, got.ToArray()))

		bSrc.assertUnchanged(t, src)
	})

	t.Run("Order string", func(t *testing.T) {
		src := NewIEnumerable[string]("2", "22", "11").WithDefaultComparers()
		bSrc := backupForAssetUnchanged(src)

		got := src.Order().ToArray()
		assert.Equal(t, "11", got[0])
		assert.Equal(t, "2", got[1])
		assert.Equal(t, "22", got[2])

		bSrc.assertUnchanged(t, src)
	})

	t.Run("Order boolean", func(t *testing.T) {
		src := NewIEnumerable[bool](true, false, true, false).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(src)

		got := src.Order().ToArray()
		assert.False(t, got[0])
		assert.False(t, got[1])
		assert.True(t, got[2])
		assert.True(t, got[3])

		bSrc.assertUnchanged(t, src)
	})
}
