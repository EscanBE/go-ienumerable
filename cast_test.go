package go_ienumerable

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_unbox_alias(t *testing.T) {
	t.Run("can not unbox custom alias not equals", func(t *testing.T) {
		type MyInt64 int64
		assert.Equal(t, "go_ienumerable.MyInt64", NewIEnumerable[MyInt64]().exposeDataType())
		var mv MyInt64 = 1000
		v, unboxOk := any(mv).(int64)
		assert.False(t, unboxOk)
		assert.Equal(t, int64(0), v)
	})

	t.Run("can unbox custom alias of equals", func(t *testing.T) {
		type MyInt64 = int64
		assert.Equal(t, "int64", NewIEnumerable[MyInt64]().exposeDataType())
		var mv MyInt64 = 1000
		v, unboxOk := any(mv).(int64)
		assert.True(t, unboxOk)
		assert.Equal(t, int64(1000), v)
	})

	t.Run("can unbox alias byte of uint8", func(t *testing.T) {
		assert.Equal(t, "uint8", NewIEnumerable[byte]().exposeDataType())
		var mv byte = 255
		v, unboxOk := any(mv).(uint8)
		assert.True(t, unboxOk)
		assert.Equal(t, uint8(255), v)
	})

	t.Run("can unbox alias rune of int32", func(t *testing.T) {
		assert.Equal(t, "int32", NewIEnumerable[rune]().exposeDataType())
		var mv rune = math.MaxInt32
		v, unboxOk := any(mv).(int32)
		assert.True(t, unboxOk)
		assert.Equal(t, int32(math.MaxInt32), v)
	})
}

func Test_enumerable_Cast(t *testing.T) {
	t.Run("byte", func(t *testing.T) {
		eSrc := NewIEnumerable[byte](5, 3).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v byte) any {
			return v
		}).CastByte().(*enumerable[byte])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, byte(5), eGot.data[0])
		assert.Equal(t, byte(3), eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[int](999, 1).Select(func(v int) any {
			return v
		}).CastByte()
	})

	t.Run("int32", func(t *testing.T) {
		eSrc := NewIEnumerable[int32](5, 3).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v int32) any {
			return v
		}).CastInt32().(*enumerable[int32])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, int32(5), eGot.data[0])
		assert.Equal(t, int32(3), eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[int64](math.MaxInt32+1, 3).Select(func(v int64) any {
			return v
		}).CastInt32()
	})

	t.Run("int64", func(t *testing.T) {
		eSrc := NewIEnumerable[int64](5, 3).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v int64) any {
			return v
		}).CastInt64().(*enumerable[int64])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, int64(5), eGot.data[0])
		assert.Equal(t, int64(3), eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[uint64](math.MaxUint64, 3).Select(func(v uint64) any {
			return v
		}).CastInt64()
	})

	t.Run("int", func(t *testing.T) {
		eSrc := NewIEnumerable[int](5, 3).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v int) any {
			return v
		}).CastInt().(*enumerable[int])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, 5, eGot.data[0])
		assert.Equal(t, 3, eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[uint](math.MaxUint, 3).Select(func(v uint) any {
			return v
		}).CastInt()
	})

	t.Run("float64", func(t *testing.T) {
		eSrc := NewIEnumerable[float64](5.0, 3.0).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v float64) any {
			return v
		}).CastFloat64().(*enumerable[float64])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, float64(5.0), eGot.data[0])
		assert.Equal(t, float64(3.0), eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("float64", func(t *testing.T) {
		eSrc := NewIEnumerable[any](uint32(math.MaxUint32), int8(3))

		eGot := eSrc.Select(func(v any) any {
			return v
		}).CastFloat64().(*enumerable[float64])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, float64(math.MaxUint32), eGot.data[0])
		assert.Equal(t, float64(3.0), eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)
	})

	t.Run("string", func(t *testing.T) {
		eSrc := NewIEnumerable[string]("5", "3").WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v string) any {
			return v
		}).CastString().(*enumerable[string])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, "5", eGot.data[0])
		assert.Equal(t, "3", eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[int64](5, 3).Select(func(v int64) any {
			return v
		}).CastString()
	})

	t.Run("bool", func(t *testing.T) {
		eSrc := NewIEnumerable[bool](true, false).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v bool) any {
			return v
		}).CastBool().(*enumerable[bool])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, true, eGot.data[0])
		assert.Equal(t, false, eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[int64](5, 3).Select(func(v int64) any {
			return v
		}).CastBool()
	})

	t.Run("panic message when nil value", func(t *testing.T) {
		eSrc := NewIEnumerable[int32](2, 3).WithDefaultComparers()

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic but not panicked")
				return
			}

			errMsg := fmt.Sprintf("%v", err)
			assert.Contains(t, errMsg, "value <nil> of type <nil> cannot be casted to int64")
		}()

		eSrc.Select(func(v int32) any {
			return nil
		}).CastInt64()
	})
}
