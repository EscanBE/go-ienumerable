package helper

import (
	"github.com/EscanBE/go-ienumerable/goe"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestChunk(t *testing.T) {
	tests := []struct {
		name string
		src  goe.IEnumerable[int]
		size int
		want goe.IEnumerable[[]int]
	}{
		{
			name: "empty",
			src:  goe.NewIEnumerable[int](),
			size: rand.Intn(30) + 2,
			want: goe.NewIEnumerable[[]int](),
		},
		{
			name: "1",
			src:  goe.NewIEnumerable[int](1, 2, 3, 4),
			size: 1,
			want: goe.NewIEnumerable[[]int]([]int{1}, []int{2}, []int{3}, []int{4}),
		},
		{
			name: "2",
			src:  goe.NewIEnumerable[int](1, 2, 3, 4),
			size: 2,
			want: goe.NewIEnumerable[[]int]([]int{1, 2}, []int{3, 4}),
		},
		{
			name: "4",
			src:  goe.NewIEnumerable[int](1, 2, 3, 4),
			size: 4,
			want: goe.NewIEnumerable[[]int]([]int{1, 2, 3, 4}),
		},
		{
			name: "odd 1",
			src:  goe.NewIEnumerable[int](1, 2, 3, 4, 5),
			size: 4,
			want: goe.NewIEnumerable[[]int]([]int{1, 2, 3, 4}, []int{5}),
		},
		{
			name: "odd 2",
			src:  goe.NewIEnumerable[int](1, 2, 3, 4, 5, 6),
			size: 4,
			want: goe.NewIEnumerable[[]int]([]int{1, 2, 3, 4}, []int{5, 6}),
		},
		{
			name: "odd 3",
			src:  goe.NewIEnumerable[int](1, 2, 3, 4, 5, 6, 7),
			size: 4,
			want: goe.NewIEnumerable[[]int]([]int{1, 2, 3, 4}, []int{5, 6, 7}),
		},
		{
			name: "4",
			src:  goe.NewIEnumerable[int](1, 2, 3, 4, 5, 6, 7, 8),
			size: 4,
			want: goe.NewIEnumerable[[]int]([]int{1, 2, 3, 4}, []int{5, 6, 7, 8}),
		},
		{
			name: "odd 1",
			src:  goe.NewIEnumerable[int](1, 2, 3, 4, 5, 6, 7, 8, 9),
			size: 4,
			want: goe.NewIEnumerable[[]int]([]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Chunk[int](tt.src, tt.size)

			if got.Count() != tt.want.Count() {
				t.Errorf("length expected %d got %d", tt.want.Count(), got.Count())
				return
			}

			if tt.want.Count() > 0 {
				gotData := got.ToArray()
				for i0, aw := range tt.want.ToArray() {
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
		})
	}
}
