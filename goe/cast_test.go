package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_unbox_alias(t *testing.T) {
	t.Run("can not unbox custom alias not equals", func(t *testing.T) {
		type MyInt64 int64
		assert.Equal(t, "goe.MyInt64", e[MyInt64](NewIEnumerable[MyInt64]()).dataType)
		var mv MyInt64 = 1000
		v, unboxOk := any(mv).(int64)
		assert.False(t, unboxOk)
		assert.Equal(t, int64(0), v)
	})

	t.Run("can unbox custom alias of equals", func(t *testing.T) {
		type MyInt64 = int64
		assert.Equal(t, "int64", e[MyInt64](NewIEnumerable[MyInt64]()).dataType)
		var mv MyInt64 = 1000
		v, unboxOk := any(mv).(int64)
		assert.True(t, unboxOk)
		assert.Equal(t, int64(1000), v)
	})

	t.Run("can unbox alias byte of uint8", func(t *testing.T) {
		assert.Equal(t, "uint8", e[byte](NewIEnumerable[byte]()).dataType)
		var mv byte = 255
		v, unboxOk := any(mv).(uint8)
		assert.True(t, unboxOk)
		assert.Equal(t, uint8(255), v)
	})

	t.Run("can unbox alias rune of int32", func(t *testing.T) {
		assert.Equal(t, "int32", e[rune](NewIEnumerable[rune]()).dataType)
		var mv rune = math.MaxInt32
		v, unboxOk := any(mv).(int32)
		assert.True(t, unboxOk)
		assert.Equal(t, int32(math.MaxInt32), v)
	})
}

//goland:noinspection GoRedundantConversion
func Test_enumerable_Cast(t *testing.T) {
	t.Run("byte", func(t *testing.T) {
		eSrc := NewIEnumerable[byte](5, 3)
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v byte) any {
			return v
		}).CastByte().(*enumerable[byte])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, byte(5), eGot.data[0])
		assert.Equal(t, byte(3), eGot.data[1])

		assert.NotNil(t, eGot.defaultComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[int](999, 1).Select(func(v int) any {
			return v
		}).CastByte()
	})

	t.Run("int32", func(t *testing.T) {
		eSrc := NewIEnumerable[int32](5, 3)
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v int32) any {
			return v
		}).CastInt32().(*enumerable[int32])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, int32(5), eGot.data[0])
		assert.Equal(t, int32(3), eGot.data[1])

		assert.NotNil(t, eGot.defaultComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[int64](math.MaxInt32+1, 3).Select(func(v int64) any {
			return v
		}).CastInt32()
	})

	t.Run("int64", func(t *testing.T) {
		eSrc := NewIEnumerable[int64](5, 3)
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v int64) any {
			return v
		}).CastInt64().(*enumerable[int64])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, int64(5), eGot.data[0])
		assert.Equal(t, int64(3), eGot.data[1])

		assert.NotNil(t, eGot.defaultComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[uint64](math.MaxUint64, 3).Select(func(v uint64) any {
			return v
		}).CastInt64()
	})

	t.Run("int", func(t *testing.T) {
		eSrc := NewIEnumerable[int](5, 3)
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v int) any {
			return v
		}).CastInt().(*enumerable[int])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, 5, eGot.data[0])
		assert.Equal(t, 3, eGot.data[1])

		assert.NotNil(t, eGot.defaultComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[uint](math.MaxUint, 3).Select(func(v uint) any {
			return v
		}).CastInt()
	})

	t.Run("float64", func(t *testing.T) {
		eSrc := NewIEnumerable[float64](5.0, 3.0)
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v float64) any {
			return v
		}).CastFloat64().(*enumerable[float64])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, float64(5.0), eGot.data[0])
		assert.Equal(t, float64(3.0), eGot.data[1])

		assert.NotNil(t, eGot.defaultComparer)

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

		assert.NotNil(t, eGot.defaultComparer)
	})

	t.Run("string", func(t *testing.T) {
		eSrc := NewIEnumerable[string]("5", "3")
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v string) any {
			return v
		}).CastString().(*enumerable[string])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, "5", eGot.data[0])
		assert.Equal(t, "3", eGot.data[1])

		assert.NotNil(t, eGot.defaultComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[int64](5, 3).Select(func(v int64) any {
			return v
		}).CastString()
	})

	t.Run("bool", func(t *testing.T) {
		eSrc := NewIEnumerable[bool](true, false)
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v bool) any {
			return v
		}).CastBool().(*enumerable[bool])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, true, eGot.data[0])
		assert.Equal(t, false, eGot.data[1])

		assert.NotNil(t, eGot.defaultComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[int64](5, 3).Select(func(v int64) any {
			return v
		}).CastBool()
	})

	t.Run("no panic message when nil value", func(t *testing.T) {
		eSrc := createRandomIntEnumerable(30)

		got := eSrc.Select(func(v int) any {
			var p *string
			return p
		}).CastInt64().SumInt()

		assert.Zero(t, got)
	})

	testCastCorrectDataTypeAndComparer[byte](t, func(eAny IEnumerable[any]) IEnumerable[byte] {
		return eAny.CastByte()
	})

	testCastCorrectDataTypeAndComparer[int32](t, func(eAny IEnumerable[any]) IEnumerable[int32] {
		return eAny.CastInt32()
	})

	testCastCorrectDataTypeAndComparer[int64](t, func(eAny IEnumerable[any]) IEnumerable[int64] {
		return eAny.CastInt64()
	})

	testCastCorrectDataTypeAndComparer[int](t, func(eAny IEnumerable[any]) IEnumerable[int] {
		return eAny.CastInt()
	})

	testCastCorrectDataTypeAndComparer[float64](t, func(eAny IEnumerable[any]) IEnumerable[float64] {
		return eAny.CastFloat64()
	})

	testCastCorrectDataTypeAndComparer[string](t, func(eAny IEnumerable[any]) IEnumerable[string] {
		return eAny.CastString()
	})

	testCastCorrectDataTypeAndComparer[bool](t, func(eAny IEnumerable[any]) IEnumerable[bool] {
		return eAny.CastBool()
	})
}

func testCastCorrectDataTypeAndComparer[T any](t *testing.T, cast func(IEnumerable[any]) IEnumerable[T]) {
	dataType := fmt.Sprintf("%T", *new(T))
	t.Run(fmt.Sprintf("cast correct type & comparer [%s]", dataType), func(t *testing.T) {
		ieSrc := NewIEnumerable[T]()

		eSrc := e[T](ieSrc)
		assert.Equal(t, dataType, eSrc.dataType)
		assert.NotNil(t, eSrc.defaultComparer)

		casted := cast(eSrc.Select(func(input T) any {
			return input
		}))

		eCasted := e[T](casted)
		assert.Equal(t, dataType, eCasted.dataType)
		assert.NotNil(t, eCasted.defaultComparer)
	})
}
