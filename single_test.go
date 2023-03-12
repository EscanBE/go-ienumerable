package go_ienumerable

import (
	"fmt"
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

		assert.Equal(t, 9, eSrc.Single())
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := NewIEnumerable[int]()
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
			assert.Contains(t, fmt.Sprintf("%v", err), getErrorSrcContainsNoElement().Error())
		}()

		_ = eSrc.Single()
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

		_ = eSrc.Single()
	})
}

func Test_enumerable_SingleBy(t *testing.T) {
	t.Run("match", func(t *testing.T) {
		eSrc := NewIEnumerable[int](6, 9)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		assert.Equal(t, 9, eSrc.SingleBy(func(v int) bool {
			return v >= 8
		}))
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := NewIEnumerable[int]()
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
			assert.Contains(t, fmt.Sprintf("%v", err), getErrorSrcContainsNoElement().Error())
		}()

		_ = eSrc.SingleBy(func(v int) bool {
			return v >= 8
		})
	})

	t.Run("more than one match", func(t *testing.T) {
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

		_ = eSrc.SingleBy(func(v int) bool {
			return v >= 5
		})
	})

	t.Run("no match", func(t *testing.T) {
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
			assert.Contains(t, fmt.Sprintf("%v", err), getErrorNoMatch().Error())
		}()

		_ = eSrc.SingleBy(func(v int) bool {
			return v < 5
		})
	})
}
