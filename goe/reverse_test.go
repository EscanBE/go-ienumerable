package goe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_Reverse(t *testing.T) {
	tests := []struct {
		name   string
		source []int
		want   []int
	}{
		{
			name:   "empty ok",
			source: []int{},
			want:   []int{},
		},
		{
			name:   "single ok",
			source: []int{2},
			want:   []int{2},
		},
		{
			name:   "reverse",
			source: []int{2, 3, 4, 5, 6},
			want:   []int{6, 5, 4, 3, 2},
		},
		{
			name:   "reverse",
			source: []int{2, 3, 4, 5, 6, 9, 8, 7, 6, 5},
			want:   []int{5, 6, 7, 8, 9, 6, 5, 4, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eSource := NewIEnumerable[int](tt.source...)
			eWant := NewIEnumerable[int](tt.want...)
			backSrc := backupForAssetUnchanged(eSource)

			eGot := eSource.Reverse()

			assert.True(t, reflect.DeepEqual(eWant.exposeData(), eGot.exposeData()))

			backSrc.assertUnchanged(t, eSource)
		})
	}
}
