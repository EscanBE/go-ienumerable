package go_ienumerable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Unbox(t *testing.T) {
	t.Run("int8", func(t *testing.T) {
		eSrc := NewIEnumerable[int8](5, 3).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v int8) any {
			return v
		}).UnboxInt8().(*enumerable[int8])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, int8(5), eGot.data[0])
		assert.Equal(t, int8(3), eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[int64](5, 3).Select(func(v int64) any {
			return v
		}).UnboxInt8()
	})

	t.Run("uint8", func(t *testing.T) {
		eSrc := NewIEnumerable[uint8](5, 3).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v uint8) any {
			return v
		}).UnboxUInt8().(*enumerable[uint8])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, uint8(5), eGot.data[0])
		assert.Equal(t, uint8(3), eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[uint64](5, 3).Select(func(v uint64) any {
			return v
		}).UnboxUInt8()
	})

	t.Run("int16", func(t *testing.T) {
		eSrc := NewIEnumerable[int16](5, 3).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v int16) any {
			return v
		}).UnboxInt16().(*enumerable[int16])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, int16(5), eGot.data[0])
		assert.Equal(t, int16(3), eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[int8](5, 3).Select(func(v int8) any {
			return v
		}).UnboxInt16()
	})

	t.Run("uint16", func(t *testing.T) {
		eSrc := NewIEnumerable[uint16](5, 3).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v uint16) any {
			return v
		}).UnboxUInt16().(*enumerable[uint16])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, uint16(5), eGot.data[0])
		assert.Equal(t, uint16(3), eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[uint8](5, 3).Select(func(v uint8) any {
			return v
		}).UnboxUInt16()
	})

	t.Run("int32", func(t *testing.T) {
		eSrc := NewIEnumerable[int32](5, 3).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v int32) any {
			return v
		}).UnboxInt32().(*enumerable[int32])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, int32(5), eGot.data[0])
		assert.Equal(t, int32(3), eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[int8](5, 3).Select(func(v int8) any {
			return v
		}).UnboxInt32()
	})

	t.Run("uint32", func(t *testing.T) {
		eSrc := NewIEnumerable[uint32](5, 3).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v uint32) any {
			return v
		}).UnboxUInt32().(*enumerable[uint32])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, uint32(5), eGot.data[0])
		assert.Equal(t, uint32(3), eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[uint8](5, 3).Select(func(v uint8) any {
			return v
		}).UnboxUInt32()
	})

	t.Run("int64", func(t *testing.T) {
		eSrc := NewIEnumerable[int64](5, 3).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v int64) any {
			return v
		}).UnboxInt64().(*enumerable[int64])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, int64(5), eGot.data[0])
		assert.Equal(t, int64(3), eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[int8](5, 3).Select(func(v int8) any {
			return v
		}).UnboxInt64()
	})

	t.Run("uint64", func(t *testing.T) {
		eSrc := NewIEnumerable[uint64](5, 3).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v uint64) any {
			return v
		}).UnboxUInt64().(*enumerable[uint64])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, uint64(5), eGot.data[0])
		assert.Equal(t, uint64(3), eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[uint8](5, 3).Select(func(v uint8) any {
			return v
		}).UnboxUInt64()
	})

	t.Run("int", func(t *testing.T) {
		eSrc := NewIEnumerable[int](5, 3).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v int) any {
			return v
		}).UnboxInt().(*enumerable[int])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, 5, eGot.data[0])
		assert.Equal(t, 3, eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[int8](5, 3).Select(func(v int8) any {
			return v
		}).UnboxInt()
	})

	t.Run("uint", func(t *testing.T) {
		eSrc := NewIEnumerable[uint](5, 3).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v uint) any {
			return v
		}).UnboxUInt().(*enumerable[uint])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, uint(5), eGot.data[0])
		assert.Equal(t, uint(3), eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[uint8](5, 3).Select(func(v uint8) any {
			return v
		}).UnboxUInt()
	})

	t.Run("uintptr", func(t *testing.T) {
		eSrc := NewIEnumerable[uintptr](5, 3).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v uintptr) any {
			return v
		}).UnboxUIntptr().(*enumerable[uintptr])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, uintptr(5), eGot.data[0])
		assert.Equal(t, uintptr(3), eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[uint8](5, 3).Select(func(v uint8) any {
			return v
		}).UnboxUIntptr()
	})

	t.Run("float32", func(t *testing.T) {
		eSrc := NewIEnumerable[float32](5.0, 3.0).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v float32) any {
			return v
		}).UnboxFloat32().(*enumerable[float32])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, float32(5.0), eGot.data[0])
		assert.Equal(t, float32(3.0), eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[float64](5.0, 3.0).Select(func(v float64) any {
			return v
		}).UnboxFloat32()
	})

	t.Run("float64", func(t *testing.T) {
		eSrc := NewIEnumerable[float64](5.0, 3.0).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v float64) any {
			return v
		}).UnboxFloat64().(*enumerable[float64])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, float64(5.0), eGot.data[0])
		assert.Equal(t, float64(3.0), eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[float32](5.0, 3.0).Select(func(v float32) any {
			return v
		}).UnboxFloat64()
	})

	t.Run("complex64", func(t *testing.T) {
		eSrc := NewIEnumerable[complex64](5.0, 3.0).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v complex64) any {
			return v
		}).UnboxComplex64().(*enumerable[complex64])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, complex64(5.0), eGot.data[0])
		assert.Equal(t, complex64(3.0), eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[complex128](5.0, 3.0).Select(func(v complex128) any {
			return v
		}).UnboxComplex64()
	})

	t.Run("complex128", func(t *testing.T) {
		eSrc := NewIEnumerable[complex128](5.0, 3.0).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v complex128) any {
			return v
		}).UnboxComplex128().(*enumerable[complex128])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, complex128(5.0), eGot.data[0])
		assert.Equal(t, complex128(3.0), eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[complex64](5.0, 3.0).Select(func(v complex64) any {
			return v
		}).UnboxComplex128()
	})

	t.Run("string", func(t *testing.T) {
		eSrc := NewIEnumerable[string]("5", "3").WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v string) any {
			return v
		}).UnboxString().(*enumerable[string])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, "5", eGot.data[0])
		assert.Equal(t, "3", eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[int64](5, 3).Select(func(v int64) any {
			return v
		}).UnboxString()
	})

	t.Run("bool", func(t *testing.T) {
		eSrc := NewIEnumerable[bool](true, false).WithDefaultComparers()
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Select(func(v bool) any {
			return v
		}).UnboxBool().(*enumerable[bool])
		assert.Len(t, eGot.data, 2)
		assert.Equal(t, true, eGot.data[0])
		assert.Equal(t, false, eGot.data[1])

		assert.Nil(t, eGot.equalityComparer)
		assert.Nil(t, eGot.lessComparer)

		bSrc.assertUnchanged(t, eSrc)

		defer deferWantPanicDepends(t, true)
		NewIEnumerable[int64](5, 3).Select(func(v int64) any {
			return v
		}).UnboxBool()
	})
}
