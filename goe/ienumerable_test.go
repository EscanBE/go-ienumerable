package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewIEnumerable_dataType(t *testing.T) {
	type ManyString []string
	type SampleStruct struct {
	}

	assert.Equal(t, "", e[any](NewIEnumerable[any]()).dataType)
	assert.Equal(t, "int8", e[int8](NewIEnumerable[int8]()).dataType)
	assert.Equal(t, "uint8", e[uint8](NewIEnumerable[uint8]()).dataType)
	assert.Equal(t, "uint8", e[byte](NewIEnumerable[byte]()).dataType)
	assert.Equal(t, "int16", e[int16](NewIEnumerable[int16]()).dataType)
	assert.Equal(t, "uint16", e[uint16](NewIEnumerable[uint16]()).dataType)
	assert.Equal(t, "int32", e[int32](NewIEnumerable[int32]()).dataType)
	assert.Equal(t, "uint32", e[uint32](NewIEnumerable[uint32]()).dataType)
	assert.Equal(t, "int32", e[rune](NewIEnumerable[rune]()).dataType)
	assert.Equal(t, "int64", e[int64](NewIEnumerable[int64]()).dataType)
	assert.Equal(t, "uint64", e[uint64](NewIEnumerable[uint64]()).dataType)
	assert.Equal(t, "int", e[int](NewIEnumerable[int]()).dataType)
	assert.Equal(t, "uint", e[uint](NewIEnumerable[uint]()).dataType)
	assert.Equal(t, "float32", e[float32](NewIEnumerable[float32]()).dataType)
	assert.Equal(t, "float64", e[float64](NewIEnumerable[float64]()).dataType)
	assert.Equal(t, "complex64", e[complex64](NewIEnumerable[complex64]()).dataType)
	assert.Equal(t, "complex128", e[complex128](NewIEnumerable[complex128]()).dataType)
	assert.Equal(t, "string", e[string](NewIEnumerable[string]()).dataType)
	assert.Equal(t, "goe.ManyString", e[ManyString](NewIEnumerable[ManyString]()).dataType)
	assert.Equal(t, "goe.SampleStruct", e[SampleStruct](NewIEnumerable[SampleStruct]()).dataType)
	assert.Equal(t, "", e[interface{}](NewIEnumerable[interface{}]()).dataType)
	assert.Equal(t, "[]interface {}", e[[]any](NewIEnumerable[[]any]()).dataType)
	assert.Equal(t, "[]int8", e[[]int8](NewIEnumerable[[]int8]()).dataType)
	assert.Equal(t, "[]uint8", e[[]uint8](NewIEnumerable[[]uint8]()).dataType)
	assert.Equal(t, "[]uint8", e[[]uint8](NewIEnumerable[[]byte]()).dataType)
	assert.Equal(t, "[]int16", e[[]int16](NewIEnumerable[[]int16]()).dataType)
	assert.Equal(t, "[]uint16", e[[]uint16](NewIEnumerable[[]uint16]()).dataType)
	assert.Equal(t, "[]int32", e[[]int32](NewIEnumerable[[]int32]()).dataType)
	assert.Equal(t, "[]uint32", e[[]uint32](NewIEnumerable[[]uint32]()).dataType)
	assert.Equal(t, "[]int32", e[[]int32](NewIEnumerable[[]rune]()).dataType)
	assert.Equal(t, "[]int64", e[[]int64](NewIEnumerable[[]int64]()).dataType)
	assert.Equal(t, "[]uint64", e[[]uint64](NewIEnumerable[[]uint64]()).dataType)
	assert.Equal(t, "[]int", e[[]int](NewIEnumerable[[]int]()).dataType)
	assert.Equal(t, "[]uint", e[[]uint](NewIEnumerable[[]uint]()).dataType)
	assert.Equal(t, "[]float32", e[[]float32](NewIEnumerable[[]float32]()).dataType)
	assert.Equal(t, "[]float64", e[[]float64](NewIEnumerable[[]float64]()).dataType)
	assert.Equal(t, "[]complex64", e[[]complex64](NewIEnumerable[[]complex64]()).dataType)
	assert.Equal(t, "[]complex128", e[[]complex128](NewIEnumerable[[]complex128]()).dataType)
	assert.Equal(t, "[]string", e[[]string](NewIEnumerable[[]string]()).dataType)

	testInitWithCorrectDataTypeAndComparer[byte](t, 2, 3)

	testInitWithCorrectDataTypeAndComparer[int32](t, 2, 3)

	testInitWithCorrectDataTypeAndComparer[int64](t, 2, 3)

	testInitWithCorrectDataTypeAndComparer[int](t, 2, 3)

	testInitWithCorrectDataTypeAndComparer[float64](t, 2, 3)

	testInitWithCorrectDataTypeAndComparer[string](t, "2", "3")

	testInitWithCorrectDataTypeAndComparer[bool](t, false, true)
}

func testInitWithCorrectDataTypeAndComparer[T any](t *testing.T, lessValue, greaterValue T) {
	dataType := fmt.Sprintf("%T", *new(T))
	t.Run(fmt.Sprintf("cast correct type & comparer [%s]", dataType), func(t *testing.T) {
		ieSrc := NewIEnumerable[T]()

		eSrc := e[T](ieSrc)
		assert.Equal(t, dataType, eSrc.dataType)
		assert.NotNil(t, eSrc.defaultComparer)

		assert.Equal(t, -1, eSrc.defaultComparer.CompareAny(lessValue, greaterValue))
		assert.Equal(t, 0, eSrc.defaultComparer.CompareAny(lessValue, lessValue))
		assert.Equal(t, 0, eSrc.defaultComparer.CompareAny(greaterValue, greaterValue))
		assert.Equal(t, 1, eSrc.defaultComparer.CompareAny(greaterValue, lessValue))

		assert.Equal(t, -1, eSrc.defaultComparer.CompareAny(any(lessValue), any(greaterValue)))
		assert.Equal(t, 0, eSrc.defaultComparer.CompareAny(any(lessValue), any(lessValue)))
		assert.Equal(t, 0, eSrc.defaultComparer.CompareAny(any(greaterValue), any(greaterValue)))
		assert.Equal(t, 1, eSrc.defaultComparer.CompareAny(any(greaterValue), any(lessValue)))
	})
}

func TestNewIEnumerableFromMap(t *testing.T) {
	tests := []struct {
		name   string
		source map[string]int
		want   IEnumerable[KeyValuePair[string, int]]
	}{
		{
			name:   "nil yields empty",
			source: nil,
			want:   NewIEnumerable[KeyValuePair[string, int]](),
		},
		{
			name:   "empty yields empty",
			source: map[string]int{},
			want:   NewIEnumerable[KeyValuePair[string, int]](),
		},
		{
			name: "one yields one",
			source: map[string]int{
				"1": 1,
			},
			want: NewIEnumerable[KeyValuePair[string, int]](KeyValuePair[string, int]{
				Key:   "1",
				Value: 1,
			}),
		},
		{
			name: "two yields two",
			source: map[string]int{
				"1": 1,
				"2": 2,
			},
			want: NewIEnumerable[KeyValuePair[string, int]](
				KeyValuePair[string, int]{
					Key:   "1",
					Value: 1,
				},
				KeyValuePair[string, int]{
					Key:   "2",
					Value: 2,
				},
			),
		},
		{
			name: "five yields five",
			source: map[string]int{
				"1": 1,
				"2": 2,
				"3": 3,
				"4": 4,
				"5": 5,
			},
			want: NewIEnumerable[KeyValuePair[string, int]](
				KeyValuePair[string, int]{
					Key:   "1",
					Value: 1,
				},
				KeyValuePair[string, int]{
					Key:   "2",
					Value: 2,
				},
				KeyValuePair[string, int]{
					Key:   "3",
					Value: 3,
				},
				KeyValuePair[string, int]{
					Key:   "4",
					Value: 4,
				},
				KeyValuePair[string, int]{
					Key:   "5",
					Value: 5,
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewIEnumerableFromMap(tt.source).OrderBy(func(p KeyValuePair[string, int]) any {
				return p.Key
			}, nil).GetOrderedEnumerable().ToArray()
			want := tt.want.OrderBy(func(p KeyValuePair[string, int]) any {
				return p.Key
			}, nil).GetOrderedEnumerable().ToArray()
			assert.True(t, reflect.DeepEqual(got, want))
		})
	}
}
