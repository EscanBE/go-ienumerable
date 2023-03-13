package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Max(t *testing.T) {
	lessComparer := func(l, r int) bool {
		return l < r
	}

	t.Run("Max", func(t *testing.T) {
		src := injectIntComparers(NewIEnumerable[int](4, 7, 5, 6, 3, 2))
		bSrc := backupForAssetUnchanged(src)

		got := src.Max()
		assert.Equal(t, 7, got)

		bSrc.assertUnchanged(t, src)
	})

	t.Run("Max panic empty", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "source contains no element", fmt.Sprintf("%v", err))
		}()
		_ = NewIEnumerable[int]().Max()
	})

	t.Run("Max panic no comparer", func(t *testing.T) {

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), "the following comparer must be set")
		}()
		_ = NewIEnumerable[int](1, 2, 3).Max()
	})

	t.Run("MaxBy", func(t *testing.T) {
		src := NewIEnumerable[int](4, 7, 5, 6, 3, 2)
		bSrc := backupForAssetUnchanged(src)

		got := src.MaxBy(lessComparer)
		assert.Equal(t, 7, got)

		bSrc.assertUnchanged(t, src)
	})

	t.Run("MaxBy panic empty", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "source contains no element", fmt.Sprintf("%v", err))
		}()
		_ = NewIEnumerable[int]().MaxBy(lessComparer)
	})

	t.Run("MaxBy panic no comparer", func(t *testing.T) {

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "comparer is nil", fmt.Sprintf("%v", err))
		}()
		_ = NewIEnumerable[int](1, 2, 3).MaxBy(nil)
	})
}
