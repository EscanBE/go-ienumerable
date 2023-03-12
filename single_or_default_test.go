package go_ienumerable

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

		assert.Equal(t, 9, eSrc.SingleOrDefault())
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := NewIEnumerable[int]()
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		assert.Zero(t, eSrc.SingleOrDefault())
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

		_ = eSrc.SingleOrDefault()
	})
}

func Test_enumerable_SingleOrDefaultBy(t *testing.T) {
	t.Run("match", func(t *testing.T) {
		eSrc := NewIEnumerable[int](6, 9)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		assert.Equal(t, 9, eSrc.SingleOrDefaultBy(func(v int) bool {
			return v >= 8
		}))
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := NewIEnumerable[int]()
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		got := eSrc.SingleOrDefaultBy(func(v int) bool {
			return v >= 8
		})

		assert.Zero(t, got)
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

		_ = eSrc.SingleOrDefaultBy(func(v int) bool {
			return v >= 5
		})
	})

	t.Run("no match", func(t *testing.T) {
		eSrc := NewIEnumerable[int](6, 9)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		got := eSrc.SingleOrDefaultBy(func(v int) bool {
			return v < 5
		})

		assert.Zero(t, got)
	})

	t.Run("nil predicate", func(t *testing.T) {
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
			assert.Contains(t, fmt.Sprintf("%v", err), getErrorNilPredicate().Error())
		}()

		_ = eSrc.SingleOrDefaultBy(nil)
	})
}

func Test_enumerable_SingleOrDefaultUsing(t *testing.T) {
	t.Run("match", func(t *testing.T) {
		eSrc := NewIEnumerable[int](9)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		assert.Equal(t, 9, eSrc.SingleOrDefaultUsing(66))
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := NewIEnumerable[int]()
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		assert.Equal(t, 9, eSrc.SingleOrDefaultUsing(9))
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

		_ = eSrc.SingleOrDefaultUsing(88)
	})
}

func Test_enumerable_SingleOrDefaultByUsing(t *testing.T) {
	t.Run("match", func(t *testing.T) {
		eSrc := NewIEnumerable[int](6, 9)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		assert.Equal(t, 9, eSrc.SingleOrDefaultByUsing(func(v int) bool {
			return v >= 8
		}, 999))
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := NewIEnumerable[int]()
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		got := eSrc.SingleOrDefaultByUsing(func(v int) bool {
			return v >= 8
		}, 999)

		assert.Equal(t, 999, got)
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

		_ = eSrc.SingleOrDefaultByUsing(func(v int) bool {
			return v >= 5
		}, 99)
	})

	t.Run("no match", func(t *testing.T) {
		eSrc := NewIEnumerable[int](6, 9)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		got := eSrc.SingleOrDefaultByUsing(func(v int) bool {
			return v < 5
		}, 999)

		assert.Equal(t, 999, got)
	})

	t.Run("nil predicate", func(t *testing.T) {
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
			assert.Contains(t, fmt.Sprintf("%v", err), getErrorNilPredicate().Error())
		}()

		_ = eSrc.SingleOrDefaultByUsing(nil, 9)
	})
}
