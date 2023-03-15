package goe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_DefaultIfEmpty(t *testing.T) {
	t.Run("non-empty", func(t *testing.T) {
		eSrc := NewIEnumerable[int8](3, 4, 5)
		bSrc := backupForAssetUnchanged(eSrc)

		srcD := eSrc.ToArray()

		got := eSrc.DefaultIfEmpty()
		gotData := got.ToArray()
		assert.Len(t, gotData, eSrc.Count())
		assert.Equal(t, srcD[0], gotData[0])
		assert.Equal(t, srcD[1], gotData[1])
		assert.Equal(t, srcD[2], gotData[2])

		bSrc.assertUnchanged(t, eSrc)
		bSrc.assertUnchangedIgnoreData(t, got)
	})

	t.Run("uint32", func(t *testing.T) {
		eSrc := NewIEnumerable[uint32]()
		bSrc := backupForAssetUnchanged(eSrc)

		got := eSrc.DefaultIfEmpty()
		gotData := got.ToArray()
		assert.Len(t, gotData, 1)
		assert.Zero(t, gotData[0])

		bSrc.assertUnchanged(t, eSrc)
		bSrc.assertUnchangedIgnoreData(t, got)
	})

	t.Run("float32", func(t *testing.T) {
		eSrc := NewIEnumerable[float32]()
		bSrc := backupForAssetUnchanged(eSrc)

		got := eSrc.DefaultIfEmpty()
		gotData := got.ToArray()
		assert.Len(t, gotData, 1)
		assert.Zero(t, gotData[0])

		bSrc.assertUnchanged(t, eSrc)
		bSrc.assertUnchangedIgnoreData(t, got)
	})

	t.Run("string", func(t *testing.T) {
		eSrc := NewIEnumerable[string]()
		bSrc := backupForAssetUnchanged(eSrc)

		got := eSrc.DefaultIfEmpty()
		gotData := got.ToArray()
		assert.Len(t, gotData, 1)
		assert.Empty(t, gotData[0])

		bSrc.assertUnchanged(t, eSrc)
		bSrc.assertUnchangedIgnoreData(t, got)
	})

	t.Run("bool", func(t *testing.T) {
		eSrc := NewIEnumerable[bool]()
		bSrc := backupForAssetUnchanged(eSrc)

		got := eSrc.DefaultIfEmpty()
		gotData := got.ToArray()
		assert.Len(t, gotData, 1)
		assert.False(t, gotData[0])

		bSrc.assertUnchanged(t, eSrc)
		bSrc.assertUnchangedIgnoreData(t, got)
	})

	t.Run("complex64", func(t *testing.T) {
		eSrc := NewIEnumerable[complex64]()
		bSrc := backupForAssetUnchanged(eSrc)

		got := eSrc.DefaultIfEmpty()
		gotData := got.ToArray()
		assert.Len(t, gotData, 1)
		assert.Zero(t, gotData[0])

		bSrc.assertUnchanged(t, eSrc)
		bSrc.assertUnchangedIgnoreData(t, got)
	})

	type MyStruct struct {
		V1 int
		V2 bool
		V3 string
	}
	t.Run("struct", func(t *testing.T) {
		eSrc := NewIEnumerable[MyStruct]()
		bSrc := backupForAssetUnchanged(eSrc)

		got := eSrc.DefaultIfEmpty()
		gotData := got.ToArray()
		assert.Len(t, gotData, 1)
		assert.Zero(t, gotData[0].V1)
		assert.False(t, gotData[0].V2)
		assert.Empty(t, gotData[0].V3)

		bSrc.assertUnchanged(t, eSrc)
		bSrc.assertUnchangedIgnoreData(t, got)
	})

	t.Run("pointer", func(t *testing.T) {
		eSrc := NewIEnumerable[*MyStruct]()
		bSrc := backupForAssetUnchanged(eSrc)

		got := eSrc.DefaultIfEmpty()
		gotData := got.ToArray()
		assert.Len(t, gotData, 1)
		assert.Nil(t, gotData[0])

		bSrc.assertUnchanged(t, eSrc)
		bSrc.assertUnchangedIgnoreData(t, got)
	})
}

func Test_enumerable_DefaultIfEmptyUsing(t *testing.T) {
	t.Run("non-empty", func(t *testing.T) {
		defaultValue := int8(6)

		eSrc := NewIEnumerable[int8](3, 4, 5)
		bSrc := backupForAssetUnchanged(eSrc)

		srcD := eSrc.ToArray()

		got := eSrc.DefaultIfEmptyUsing(defaultValue)
		gotData := got.ToArray()
		assert.Len(t, gotData, eSrc.Count())
		assert.Equal(t, srcD[0], gotData[0])
		assert.Equal(t, srcD[1], gotData[1])
		assert.Equal(t, srcD[2], gotData[2])

		bSrc.assertUnchanged(t, eSrc)
		bSrc.assertUnchangedIgnoreData(t, got)
	})

	t.Run("uint32", func(t *testing.T) {
		defaultValue := uint32(99)

		eSrc := NewIEnumerable[uint32]()
		bSrc := backupForAssetUnchanged(eSrc)

		got := eSrc.DefaultIfEmptyUsing(defaultValue)
		gotData := got.ToArray()
		assert.Len(t, gotData, 1)
		assert.Equal(t, defaultValue, gotData[0])

		bSrc.assertUnchanged(t, eSrc)
		bSrc.assertUnchangedIgnoreData(t, got)
	})

	t.Run("float32", func(t *testing.T) {
		defaultValue := float32(99.99)

		eSrc := NewIEnumerable[float32]()
		bSrc := backupForAssetUnchanged(eSrc)

		got := eSrc.DefaultIfEmptyUsing(defaultValue)
		gotData := got.ToArray()
		assert.Len(t, gotData, 1)
		assert.Equal(t, defaultValue, gotData[0])

		bSrc.assertUnchanged(t, eSrc)
		bSrc.assertUnchangedIgnoreData(t, got)
	})

	t.Run("string", func(t *testing.T) {
		defaultValue := "99"

		eSrc := NewIEnumerable[string]()
		bSrc := backupForAssetUnchanged(eSrc)

		got := eSrc.DefaultIfEmptyUsing(defaultValue)
		gotData := got.ToArray()
		assert.Len(t, gotData, 1)
		assert.Equal(t, defaultValue, gotData[0])

		bSrc.assertUnchanged(t, eSrc)
		bSrc.assertUnchangedIgnoreData(t, got)
	})

	t.Run("bool", func(t *testing.T) {
		defaultValue := true

		eSrc := NewIEnumerable[bool]()
		bSrc := backupForAssetUnchanged(eSrc)

		got := eSrc.DefaultIfEmptyUsing(defaultValue)
		gotData := got.ToArray()
		assert.Len(t, gotData, 1)
		assert.Equal(t, defaultValue, gotData[0])

		bSrc.assertUnchanged(t, eSrc)
		bSrc.assertUnchangedIgnoreData(t, got)
	})

	t.Run("complex64", func(t *testing.T) {
		defaultValue := complex64(99)

		eSrc := NewIEnumerable[complex64]()
		bSrc := backupForAssetUnchanged(eSrc)

		got := eSrc.DefaultIfEmptyUsing(defaultValue)
		gotData := got.ToArray()
		assert.Len(t, gotData, 1)
		assert.Equal(t, defaultValue, gotData[0])

		bSrc.assertUnchanged(t, eSrc)
		bSrc.assertUnchangedIgnoreData(t, got)
	})

	type MyStruct struct {
		V1 int
		V2 bool
		V3 string
	}

	t.Run("struct", func(t *testing.T) {
		defaultValue := MyStruct{
			V1: 99,
			V2: true,
			V3: "99",
		}

		eSrc := NewIEnumerable[MyStruct]()
		bSrc := backupForAssetUnchanged(eSrc)

		got := eSrc.DefaultIfEmptyUsing(defaultValue)
		gotData := got.ToArray()
		assert.Len(t, gotData, 1)
		assert.Equal(t, 99, gotData[0].V1)
		assert.True(t, gotData[0].V2)
		assert.Equal(t, "99", gotData[0].V3)

		bSrc.assertUnchanged(t, eSrc)
		bSrc.assertUnchangedIgnoreData(t, got)
	})

	t.Run("pointer", func(t *testing.T) {
		defaultValue := &MyStruct{
			V1: 99,
			V2: true,
			V3: "99",
		}

		eSrc := NewIEnumerable[*MyStruct]()
		bSrc := backupForAssetUnchanged(eSrc)

		got := eSrc.DefaultIfEmptyUsing(defaultValue)
		gotData := got.ToArray()
		assert.Len(t, gotData, 1)
		assert.Equal(t, 99, gotData[0].V1)
		assert.True(t, gotData[0].V2)
		assert.Equal(t, "99", gotData[0].V3)

		bSrc.assertUnchanged(t, eSrc)
		bSrc.assertUnchangedIgnoreData(t, got)
	})
}
