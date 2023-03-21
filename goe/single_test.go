package goe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Single(t *testing.T) {
	t.Run("match", func(t *testing.T) {
		eSrc := NewIEnumerable[int](9)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		assert.Equal(t, 9, eSrc.Single(nil))
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := NewIEnumerable[int]()
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		defer deferExpectPanicContains(t, getErrorSrcContainsNoElement().Error(), true)

		_ = eSrc.Single(nil)
	})

	t.Run("more than one", func(t *testing.T) {
		eSrc := NewIEnumerable[int](6, 9)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		defer deferExpectPanicContains(t, getErrorMoreThanOne().Error(), true)

		_ = eSrc.Single(nil)
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

		assert.Equal(t, 9, eSrc.Single(predicate))
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := NewIEnumerable[int]()
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		defer deferExpectPanicContains(t, getErrorSrcContainsNoElement().Error(), true)

		var predicate OptionalPredicate[int] = func(v int) bool {
			return v >= 8
		}
		_ = eSrc.Single(predicate)
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
		_ = eSrc.Single(predicate)
	})

	t.Run("no match", func(t *testing.T) {
		eSrc := NewIEnumerable[int](6, 9)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		defer deferExpectPanicContains(t, getErrorNoMatch().Error(), true)

		var predicate OptionalPredicate[int] = func(v int) bool {
			return v < 5
		}
		_ = eSrc.Single(predicate)
	})
}
