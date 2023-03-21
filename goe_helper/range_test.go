package goe_helper

import (
	"github.com/EscanBE/go-ienumerable/goe"
	"github.com/stretchr/testify/assert"
	"math"
	"reflect"
	"testing"
)

func TestRange(t *testing.T) {
	tests := []struct {
		name  string
		start int
		count int
		want  goe.IEnumerable[int]
	}{
		{
			name:  "normal",
			start: 0,
			count: 3,
			want:  goe.NewIEnumerable[int](0, 1, 2),
		},
		{
			name:  "normal",
			start: -2,
			count: 3,
			want:  goe.NewIEnumerable[int](-2, -1, 0),
		},
		{
			name:  "one",
			start: -2,
			count: 1,
			want:  goe.NewIEnumerable[int](-2),
		},
		{
			name:  "none",
			start: -2,
			count: 0,
			want:  goe.NewIEnumerable[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Range(tt.start, tt.count)
			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))
		})
	}

	t.Run("negative count", func(t *testing.T) {
		Range(3, 3)
		Range(3, 2)
		Range(3, 1)
		Range(3, 0)

		defer deferExpectPanicContains(t, "count is less than 0", true)

		Range(3, -1)
	})

	t.Run("over int32", func(t *testing.T) {
		Range(math.MaxInt32, 1)
		Range(math.MaxInt32-1, 2)
		Range(math.MaxInt32-2, 3)

		defer deferExpectPanicContains(t, "can not larger than", true)

		Range(math.MaxInt32-2, 4)
	})
}
