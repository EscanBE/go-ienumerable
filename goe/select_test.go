package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Select(t *testing.T) {
	t.Run("int8", func(t *testing.T) {
		eSrc := NewIEnumerable[int8](2, 3, 4, 5).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(i int8) any {
			return i * 2
		})
		gotData := eGot.exposeData()
		assert.Len(t, gotData, 4)
		assert.Equal(t, int8(4), gotData[0])
		assert.Equal(t, int8(6), gotData[1])
		assert.Equal(t, int8(8), gotData[2])
		assert.Equal(t, int8(10), gotData[3])
		gotE := eGot.(*enumerable[any])
		assert.Nil(t, gotE.equalityComparer)
		assert.Nil(t, gotE.lessComparer)
		assert.Equal(t, "int8", fmt.Sprintf("%T", gotE.data[0]))
		assert.Equal(t, "", gotE.dataType)

		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("string", func(t *testing.T) {
		eSrc := NewIEnumerable[int8](2, 3, 4, 5).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(i int8) any {
			return fmt.Sprintf("%d", i+1)
		})
		gotData := eGot.exposeData()
		assert.Len(t, gotData, 4)
		assert.Equal(t, "3", gotData[0])
		assert.Equal(t, "4", gotData[1])
		assert.Equal(t, "5", gotData[2])
		assert.Equal(t, "6", gotData[3])
		gotE := eGot.(*enumerable[any])
		assert.Nil(t, gotE.equalityComparer)
		assert.Nil(t, gotE.lessComparer)
		assert.Equal(t, "string", fmt.Sprintf("%T", gotE.data[0]))
		assert.Equal(t, "", gotE.dataType)

		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := NewIEnumerable[int8]().WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(i int8) any {
			return int64(i)
		})

		gotData := eGot.exposeData()
		assert.Len(t, gotData, 0)

		gotE := eGot.(*enumerable[any])
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
}
