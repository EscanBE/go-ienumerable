package go_ienumerable

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Min(t *testing.T) {
	lessComparer := func(l, r int) bool {
		return l < r
	}

	t.Run("Min", func(t *testing.T) {
		src := injectIntComparers(NewIEnumerable[int](4, 7, 5, 6, 3, 2))
		bSrc := backupForAssetUnchanged(src)

		got := src.Min()
		assert.Equal(t, 2, got)

		bSrc.assertUnchanged(t, src)
	})

	t.Run("Min panic empty", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "source contains no element", fmt.Sprintf("%v", err))
		}()
		_ = NewIEnumerable[int]().Min()
	})

	t.Run("Min panic no comparer", func(t *testing.T) {

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), "the following comparer must be set")
		}()
		_ = NewIEnumerable[int](1, 2, 3).Min()
	})

	t.Run("MinBy", func(t *testing.T) {
		src := NewIEnumerable[int](4, 7, 5, 6, 3, 2)
		bSrc := backupForAssetUnchanged(src)

		got := src.MinBy(lessComparer)
		assert.Equal(t, 2, got)

		bSrc.assertUnchanged(t, src)
	})

	t.Run("MinBy panic empty", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "source contains no element", fmt.Sprintf("%v", err))
		}()
		_ = NewIEnumerable[int]().MinBy(lessComparer)
	})

	t.Run("MinBy panic no comparer", func(t *testing.T) {

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "comparer is nil", fmt.Sprintf("%v", err))
		}()
		_ = NewIEnumerable[int](1, 2, 3).MinBy(nil)
	})
}
