package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_SingleOrDefault(t *testing.T) {
	t.Run("match", func(t *testing.T) {
		eSrc := NewIEnumerable[int](9)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		assert.Equal(t, 9, eSrc.SingleOrDefault(nil, nil))
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := NewIEnumerable[int]()
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		assert.Zero(t, eSrc.SingleOrDefault(nil, nil))
	})

	t.Run("more than one", func(t *testing.T) {
		eSrc := NewIEnumerable[int](6, 9)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), getErrorMoreThanOne().Error())
		}()

		_ = eSrc.SingleOrDefault(nil, nil)
	})

	t.Run("match", func(t *testing.T) {
		eSrc := NewIEnumerable[int](6, 9)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		var predicate OptionalPredicate[int] = func(v int) bool {
			return v >= 8
		}

		assert.Equal(t, 9, eSrc.SingleOrDefault(predicate, nil))
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := NewIEnumerable[int]()
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		var predicate OptionalPredicate[int] = func(v int) bool {
			return v >= 8
		}

		got := eSrc.SingleOrDefault(predicate, nil)

		assert.Zero(t, got)
	})

	t.Run("more than one match", func(t *testing.T) {
		eSrc := NewIEnumerable[int](6, 9)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		defer deferExpectPanicContains(t, getErrorMoreThanOne().Error(), true)

		var predicate OptionalPredicate[int] = func(v int) bool {
			return v >= 5
		}
		_ = eSrc.SingleOrDefault(predicate, nil)
	})

	t.Run("no match", func(t *testing.T) {
		eSrc := NewIEnumerable[int](6, 9)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		var predicate OptionalPredicate[int] = func(v int) bool {
			return v < 5
		}

		got := eSrc.SingleOrDefault(predicate, nil)

		assert.Zero(t, got)
	})

	t.Run("match", func(t *testing.T) {
		eSrc := NewIEnumerable[int](9)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		assert.Equal(t, 9, eSrc.SingleOrDefault(nil, Ptr(66)))
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := NewIEnumerable[int]()
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		assert.Equal(t, 9, eSrc.SingleOrDefault(nil, Ptr(9)))
	})

	t.Run("more than one", func(t *testing.T) {
		eSrc := NewIEnumerable[int](6, 9)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		defer deferExpectPanicContains(t, getErrorMoreThanOne().Error(), true)

		_ = eSrc.SingleOrDefault(nil, Ptr(88))
	})

	t.Run("match", func(t *testing.T) {
		eSrc := NewIEnumerable[int](6, 9)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		var predicate OptionalPredicate[int] = func(v int) bool {
			return v >= 8
		}

		assert.Equal(t, 9, eSrc.SingleOrDefault(predicate, Ptr(999)))
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := NewIEnumerable[int]()
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		var predicate OptionalPredicate[int] = func(v int) bool {
			return v >= 8
		}

		got := eSrc.SingleOrDefault(predicate, Ptr(999))

		assert.Equal(t, 999, got)
	})

	t.Run("more than one match", func(t *testing.T) {
		eSrc := NewIEnumerable[int](6, 9)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		defer deferExpectPanicContains(t, getErrorMoreThanOne().Error(), true)

		var predicate OptionalPredicate[int] = func(v int) bool {
			return v >= 5
		}

		_ = eSrc.SingleOrDefault(predicate, Ptr(99))
	})

	t.Run("no match", func(t *testing.T) {
		eSrc := NewIEnumerable[int](6, 9)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		var predicate OptionalPredicate[int] = func(v int) bool {
			return v < 5
		}

		got := eSrc.SingleOrDefault(predicate, Ptr(999))

		assert.Equal(t, 999, got)
	})
}
