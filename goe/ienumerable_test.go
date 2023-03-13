package goe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewIEnumerable_dataType(t *testing.T) {
	assert.Equal(t, "", NewIEnumerable[any]().exposeDataType())
	assert.Equal(t, "", NewIEnumerable[interface{}]().exposeDataType())
	assert.Equal(t, "int8", NewIEnumerable[int8]().exposeDataType())
	assert.Equal(t, "uint8", NewIEnumerable[uint8]().exposeDataType())
	assert.Equal(t, "uint8", NewIEnumerable[byte]().exposeDataType())
	assert.Equal(t, "int16", NewIEnumerable[int16]().exposeDataType())
	assert.Equal(t, "uint16", NewIEnumerable[uint16]().exposeDataType())
	assert.Equal(t, "int32", NewIEnumerable[int32]().exposeDataType())
	assert.Equal(t, "uint32", NewIEnumerable[uint32]().exposeDataType())
	assert.Equal(t, "int32", NewIEnumerable[rune]().exposeDataType())
	assert.Equal(t, "int64", NewIEnumerable[int64]().exposeDataType())
	assert.Equal(t, "uint64", NewIEnumerable[uint64]().exposeDataType())
	assert.Equal(t, "int", NewIEnumerable[int]().exposeDataType())
	assert.Equal(t, "uint", NewIEnumerable[uint]().exposeDataType())
	assert.Equal(t, "float32", NewIEnumerable[float32]().exposeDataType())
	assert.Equal(t, "float64", NewIEnumerable[float64]().exposeDataType())
	assert.Equal(t, "complex64", NewIEnumerable[complex64]().exposeDataType())
	assert.Equal(t, "complex128", NewIEnumerable[complex128]().exposeDataType())
	assert.Equal(t, "string", NewIEnumerable[string]().exposeDataType())
	type ManyString []string
	assert.Equal(t, "goe.ManyString", NewIEnumerable[ManyString]().exposeDataType())
	type SampleStruct struct {
	}
	assert.Equal(t, "goe.SampleStruct", NewIEnumerable[SampleStruct]().exposeDataType())
	assert.Equal(t, "[]interface {}", NewIEnumerable[[]any]().exposeDataType())
	assert.Equal(t, "[]int8", NewIEnumerable[[]int8]().exposeDataType())
	assert.Equal(t, "[]uint8", NewIEnumerable[[]uint8]().exposeDataType())
	assert.Equal(t, "[]uint8", NewIEnumerable[[]byte]().exposeDataType())
	assert.Equal(t, "[]int16", NewIEnumerable[[]int16]().exposeDataType())
	assert.Equal(t, "[]uint16", NewIEnumerable[[]uint16]().exposeDataType())
	assert.Equal(t, "[]int32", NewIEnumerable[[]int32]().exposeDataType())
	assert.Equal(t, "[]uint32", NewIEnumerable[[]uint32]().exposeDataType())
	assert.Equal(t, "[]int32", NewIEnumerable[[]rune]().exposeDataType())
	assert.Equal(t, "[]int64", NewIEnumerable[[]int64]().exposeDataType())
	assert.Equal(t, "[]uint64", NewIEnumerable[[]uint64]().exposeDataType())
	assert.Equal(t, "[]int", NewIEnumerable[[]int]().exposeDataType())
	assert.Equal(t, "[]uint", NewIEnumerable[[]uint]().exposeDataType())
	assert.Equal(t, "[]float32", NewIEnumerable[[]float32]().exposeDataType())
	assert.Equal(t, "[]float64", NewIEnumerable[[]float64]().exposeDataType())
	assert.Equal(t, "[]complex64", NewIEnumerable[[]complex64]().exposeDataType())
	assert.Equal(t, "[]complex128", NewIEnumerable[[]complex128]().exposeDataType())
	assert.Equal(t, "[]string", NewIEnumerable[[]string]().exposeDataType())
}
