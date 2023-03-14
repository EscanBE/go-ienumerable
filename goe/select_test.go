package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_enumerable_Select(t *testing.T) {
	t.Run("int8", func(t *testing.T) {
		eSrc := NewIEnumerable[int8](2, 3, 4, 5).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(i int8) any {
			return i * 2
		})
		gotData := eGot.ToArray()
		assert.Len(t, gotData, 4)
		assert.Equal(t, int8(4), gotData[0])
		assert.Equal(t, int8(6), gotData[1])
		assert.Equal(t, int8(8), gotData[2])
		assert.Equal(t, int8(10), gotData[3])
		gotE := e[any](eGot)
		assert.Nil(t, gotE.equalityComparer)
		assert.Nil(t, gotE.lessComparer)
		assert.Equal(t, "int8", fmt.Sprintf("%T", gotE.data[0]))
		assert.Equal(t, "int8", gotE.dataType)

		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("string", func(t *testing.T) {
		eSrc := NewIEnumerable[int8](2, 3, 4, 5).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(i int8) any {
			return fmt.Sprintf("%d", i+1)
		})
		gotData := eGot.ToArray()
		assert.Len(t, gotData, 4)
		assert.Equal(t, "3", gotData[0])
		assert.Equal(t, "4", gotData[1])
		assert.Equal(t, "5", gotData[2])
		assert.Equal(t, "6", gotData[3])
		gotE := e[any](eGot)
		assert.Nil(t, gotE.equalityComparer)
		assert.Nil(t, gotE.lessComparer)
		assert.Equal(t, "string", fmt.Sprintf("%T", gotE.data[0]))
		assert.Equal(t, "string", gotE.dataType)

		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := NewIEnumerable[int8]().WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(i int8) any {
			return int64(i)
		})

		gotData := eGot.ToArray()
		assert.Len(t, gotData, 0)

		gotE := e[any](eGot)
		assert.Nil(t, gotE.equalityComparer)
		assert.Nil(t, gotE.lessComparer)
		assert.Equal(t, "", gotE.dataType)

		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("nil selector", func(t *testing.T) {
		eSrc := NewIEnumerable[int8]().WithDefaultComparers()

		defer deferWantPanicDepends(t, true)

		_ = eSrc.Select(nil)
	})

	t.Run("automatically inject type and comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](3, 1)

		ieGot := ieSrc.Select(func(i int) any {
			return time.Duration(i) * time.Minute
		})

		gotArray := ieGot.ToArray()

		assert.Equal(t, 3*time.Minute, gotArray[0])
		assert.Equal(t, 1*time.Minute, gotArray[1])

		eGot := e[any](ieGot)
		assert.Equal(t, "time.Duration", eGot.dataType)
		assert.NotNil(t, eGot.defaultComparer)
		assert.Equal(t, 1, eGot.defaultComparer.Compare(gotArray[0], gotArray[1]))
		assert.Equal(t, -1, eGot.defaultComparer.Compare(gotArray[1], gotArray[0]))
		assert.Zero(t, eGot.defaultComparer.Compare(gotArray[0], 3*time.Minute))
		assert.Zero(t, eGot.defaultComparer.Compare(gotArray[1], 1*time.Minute))
	})
}
