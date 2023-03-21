package goe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_SequenceEqual(t *testing.T) {
	fEquals := func(i1, i2 int) bool {
		return i1 == i2
	}

	tests := []struct {
		name    string
		source  IEnumerable[int]
		second  IEnumerable[int]
		fEquals OptionalEqualsFunc[int]
		want    bool
	}{
		{
			name:    "same",
			source:  createIntEnumerable(5, 10),
			second:  createIntEnumerable(5, 10),
			fEquals: nil,
			want:    true,
		},
		{
			name:    "same with comparer",
			source:  createIntEnumerable(5, 10),
			second:  createIntEnumerable(5, 10),
			fEquals: fEquals,
			want:    true,
		},
		{
			name:    "same",
			source:  createIntEnumerable(5, 10),
			second:  createIntEnumerable(5, 10),
			fEquals: nil,
			want:    true,
		},
		{
			name:   "different size",
			source: createIntEnumerable(5, 11),
			second: createIntEnumerable(5, 10),
			want:   false,
		},
		{
			name:   "different element, same size",
			source: createIntEnumerable(5, 10),
			second: createIntEnumerable(6, 11),
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.source.SequenceEqual(tt.second, tt.fEquals))
		})
	}

	t.Run("panic unable detect comparer", func(t *testing.T) {
		type MyStruct struct{}
		ie1 := NewIEnumerable[MyStruct](MyStruct{})
		ie2 := NewIEnumerable[MyStruct](MyStruct{})

		defer deferExpectPanicContains(t, getErrorFailedCompare2ElementsInArray().Error(), true)

		_ = ie1.SequenceEqual(ie2, nil)
	})

	t.Run("resolve comparer from first collection", func(t *testing.T) {
		ie1 := NewIEnumerable[any](nil, nil, "")
		ie2 := NewIEnumerable[any](nil, nil, "")

		assert.True(t, ie1.SequenceEqual(ie2, nil))
	})

	t.Run("resolve comparer from second collection", func(t *testing.T) {
		ie1 := NewIEnumerable[any](nil, nil, nil)
		ie2 := NewIEnumerable[any](nil, nil, "")

		assert.False(t, ie1.SequenceEqual(ie2, nil))
	})
}
