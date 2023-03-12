package go_ienumerable

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func Test_enumerable_ChunkToAny(t *testing.T) {
	tests := []struct {
		name string
		src  IEnumerable[int]
		size int
		want IEnumerable[[]any]
	}{
		{
			name: "empty",
			src:  NewIEnumerable[int](),
			size: rand.Intn(30) + 2,
			want: NewIEnumerable[[]any](),
		},
		{
			name: "1",
			src:  NewIEnumerable[int](1, 2, 3, 4),
			size: 1,
			want: NewIEnumerable[[]any]([]any{1}, []any{2}, []any{3}, []any{4}),
		},
		{
			name: "2",
			src:  NewIEnumerable[int](1, 2, 3, 4),
			size: 2,
			want: NewIEnumerable[[]any]([]any{1, 2}, []any{3, 4}),
		},
		{
			name: "4",
			src:  NewIEnumerable[int](1, 2, 3, 4),
			size: 4,
			want: NewIEnumerable[[]any]([]any{1, 2, 3, 4}),
		},
		{
			name: "odd 1",
			src:  NewIEnumerable[int](1, 2, 3, 4, 5),
			size: 4,
			want: NewIEnumerable[[]any]([]any{1, 2, 3, 4}, []any{5}),
		},
		{
			name: "odd 2",
			src:  NewIEnumerable[int](1, 2, 3, 4, 5, 6),
			size: 4,
			want: NewIEnumerable[[]any]([]any{1, 2, 3, 4}, []any{5, 6}),
		},
		{
			name: "odd 3",
			src:  NewIEnumerable[int](1, 2, 3, 4, 5, 6, 7),
			size: 4,
			want: NewIEnumerable[[]any]([]any{1, 2, 3, 4}, []any{5, 6, 7}),
		},
		{
			name: "4",
			src:  NewIEnumerable[int](1, 2, 3, 4, 5, 6, 7, 8),
			size: 4,
			want: NewIEnumerable[[]any]([]any{1, 2, 3, 4}, []any{5, 6, 7, 8}),
		},
		{
			name: "odd 1",
			src:  NewIEnumerable[int](1, 2, 3, 4, 5, 6, 7, 8, 9),
			size: 4,
			want: NewIEnumerable[[]any]([]any{1, 2, 3, 4}, []any{5, 6, 7, 8}, []any{9}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.src)

			got := tt.src.ChunkToAny(tt.size)

			assert.Equal(t, "[]interface {}", got.exposeDataType())

			if got.len() != tt.want.len() {
				t.Errorf("length expected %d got %d", tt.want.len(), got.len())
				return
			}

			if tt.want.len() > 0 {
				gotData := got.exposeData()
				for i0, aw := range tt.want.exposeData() {
					ag := gotData[i0]
					if len(aw) != len(ag) {
						t.Errorf("data[%d] length expected %d got %d", i0, len(aw), len(ag))
						continue
					}
					for i1, vw := range aw {
						vg := ag[i1]
						assert.Equalf(t, vw, vg, "data[%d][%d]", i0, i1)
					}
				}
			}

			bSrc.assertUnchanged(t, tt.src)
		})
	}

	t.Run("panic chunk size zero", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}

			assert.Equal(t, "size is below 1", fmt.Sprintf("%v", err))
		}()

		NewIEnumerable[int]().ChunkToAny(0)
	})
}
