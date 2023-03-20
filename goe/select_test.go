package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_enumerable_Select(t *testing.T) {
	t.Run("int8", func(t *testing.T) {
		eSrc := NewIEnumerable[int8](2, 3, 4, 5)
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
		assert.NotNil(t, gotE.defaultComparer)
		assert.Equal(t, "int8", fmt.Sprintf("%T", gotE.data[0]))
		assert.Equal(t, "int8", gotE.dataType)

		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("string", func(t *testing.T) {
		eSrc := NewIEnumerable[int8](2, 3, 4, 5)
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
		assert.NotNil(t, gotE.defaultComparer)
		assert.Equal(t, "string", fmt.Sprintf("%T", gotE.data[0]))
		assert.Equal(t, "string", gotE.dataType)

		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := NewIEnumerable[int8]()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(i int8) any {
			return int64(i)
		})

		gotData := eGot.ToArray()
		assert.Len(t, gotData, 0)

		gotE := e[any](eGot)
		assert.Nil(t, gotE.defaultComparer) // not able to detect default comparer because no value to detect type
		assert.Equal(t, "", gotE.dataType)

		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("nil selector", func(t *testing.T) {
		eSrc := NewIEnumerable[int8]()

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
		assert.Equal(t, 1, eGot.defaultComparer.CompareAny(gotArray[0], gotArray[1]))
		assert.Equal(t, -1, eGot.defaultComparer.CompareAny(gotArray[1], gotArray[0]))
		assert.Zero(t, eGot.defaultComparer.CompareAny(gotArray[0], 3*time.Minute))
		assert.Zero(t, eGot.defaultComparer.CompareAny(gotArray[1], 1*time.Minute))
	})

	t.Run("not panic if not able to detect comparer", func(t *testing.T) {
		type MyInt struct {
			Value int
		}

		ieSrc := NewIEnumerable[int](3, 1)

		ieGot := ieSrc.Select(func(i int) any {
			return MyInt{
				Value: i,
			}
		})

		gotArray := ieGot.ToArray()

		assert.Equal(t, 3, gotArray[0].(MyInt).Value)
		assert.Equal(t, 1, gotArray[1].(MyInt).Value)

		eGot := e[any](ieGot)
		assert.Equal(t, "goe.MyInt", eGot.dataType)
		assert.Nil(t, eGot.defaultComparer)
	})
}

func Test_enumerable_SelectNewValue(t *testing.T) {
	t.Run("int8", func(t *testing.T) {
		eSrc := NewIEnumerable[int8](2, 3, 4, 5)
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.SelectNewValue(func(i int8) int8 {
			return i * 2
		})
		gotData := eGot.ToArray()
		assert.Len(t, gotData, 4)
		assert.Equal(t, int8(4), gotData[0])
		assert.Equal(t, int8(6), gotData[1])
		assert.Equal(t, int8(8), gotData[2])
		assert.Equal(t, int8(10), gotData[3])
		gotE := e[int8](eGot)
		assert.NotNil(t, gotE.defaultComparer)
		assert.Equal(t, "int8", fmt.Sprintf("%T", gotE.data[0]))
		assert.Equal(t, "int8", gotE.dataType)

		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := NewIEnumerable[int8]()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.SelectNewValue(func(i int8) int8 {
			return i
		})

		gotData := eGot.ToArray()
		assert.Len(t, gotData, 0)

		gotE := e[int8](eGot)
		assert.NotNil(t, gotE.defaultComparer)
		assert.Equal(t, "int8", gotE.dataType)

		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("nil selector", func(t *testing.T) {
		eSrc := NewIEnumerable[int8]()

		defer deferWantPanicDepends(t, true)

		_ = eSrc.SelectNewValue(nil)
	})

	t.Run("automatically inject type and comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](5, 3)

		ieGot := ieSrc.SelectNewValue(func(i int) int {
			return i + 3
		})

		gotArray := ieGot.ToArray()

		assert.Equal(t, 8, gotArray[0])
		assert.Equal(t, 6, gotArray[1])

		eGot := e[int](ieGot)
		assert.Equal(t, "int", eGot.dataType)
		assert.NotNil(t, eGot.defaultComparer)
		assert.Equal(t, 1, eGot.defaultComparer.CompareAny(gotArray[0], gotArray[1]))
		assert.Equal(t, -1, eGot.defaultComparer.CompareAny(gotArray[1], gotArray[0]))
		assert.Zero(t, eGot.defaultComparer.CompareAny(gotArray[0], 8))
		assert.Zero(t, eGot.defaultComparer.CompareAny(gotArray[1], 6))
	})
}
