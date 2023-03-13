package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_ElementAt(t *testing.T) {
	src := NewIEnumerable[int](1, 2, 3, 4, 5, 6, 7, 8, 9)
	tests := []struct {
		index       int
		want        int
		wantReverse int
		wantPanic   bool
		wantDefault int
	}{
		{
			index:       0,
			want:        1,
			wantReverse: 9,
		},
		{
			index:       1,
			want:        2,
			wantReverse: 8,
		},
		{
			index:       2,
			want:        3,
			wantReverse: 7,
		},
		{
			index:       3,
			want:        4,
			wantReverse: 6,
		},
		{
			index:       4,
			want:        5,
			wantReverse: 5,
		},
		{
			index:       5,
			want:        6,
			wantReverse: 4,
		},
		{
			index:       10,
			wantPanic:   true,
			wantDefault: 0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.index), func(t *testing.T) {
			bSrc := backupForAssetUnchanged(src)
			defer deferWantPanicDepends(t, tt.wantPanic)
			assert.Equalf(t, tt.want, src.ElementAt(tt.index), "ElementAt(%v)", tt.index)
			assert.Equalf(t, tt.wantReverse, src.ElementAtReverse(tt.index), "ElementAtReverse(%v)", tt.index)
			bSrc.assertUnchanged(t, src)
		})
		t.Run(fmt.Sprintf("%d or default", tt.index), func(t *testing.T) {
			bSrc := backupForAssetUnchanged(src)
			want := tt.want
			wantReverse := tt.wantReverse
			if tt.wantPanic {
				want = tt.wantDefault
				wantReverse = tt.wantDefault
			}
			assert.Equalf(t, want, src.ElementAtOrDefault(tt.index), "ElementAtOrDefault(%v)", tt.index)
			assert.Equalf(t, wantReverse, src.ElementAtReverseOrDefault(tt.index), "ElementAtReverseOrDefault(%v)", tt.index)
			bSrc.assertUnchanged(t, src)
		})
	}

	for _, tt := range []int{-2, -1, 0} {
		t.Run("out of bound", func(t *testing.T) {
			defer func() {
				err := recover()
				if err == nil {
					t.Errorf("expect error")
					return
				}

				assert.Equal(t, "index out of bound", fmt.Sprintf("%v", err))
			}()

			NewIEnumerable[int]().ElementAt(tt)
		})
	}
}
