package go_ienumerable

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_Union_UnionBy(t *testing.T) {
	fEquals := func(t1, t2 int) bool {
		return t1 == t2
	}
	var tests = []struct {
		name      string
		source    IEnumerable[int]
		second    IEnumerable[int]
		fEquals   func(t1, t2 int) bool
		want      IEnumerable[int]
		wantPanic bool
	}{
		{
			name:    "union empty with empty",
			source:  createEmptyIntEnumerable(),
			second:  createEmptyIntEnumerable(),
			fEquals: fEquals,
			want:    createEmptyIntEnumerable(),
		},
		{
			name:    "union empty with non-empty",
			source:  createEmptyIntEnumerable(),
			second:  NewIEnumerable[int](2, 2),
			fEquals: fEquals,
			want:    NewIEnumerable[int](2),
		},
		{
			name:    "union non-empty with empty",
			source:  NewIEnumerable[int](2, 2, 3),
			second:  createEmptyIntEnumerable(),
			fEquals: fEquals,
			want:    NewIEnumerable[int](2, 3),
		},
		{
			name:    "union non-empty with non-empty",
			source:  NewIEnumerable[int](2, 2, 4, 5, 6),
			second:  NewIEnumerable[int](9, 8, 7),
			fEquals: fEquals,
			want:    NewIEnumerable[int](2, 4, 5, 6, 9, 8, 7),
		},
		{
			name:    "distinct",
			source:  NewIEnumerable[int](2),
			second:  NewIEnumerable[int](2),
			fEquals: fEquals,
			want:    NewIEnumerable[int](2),
		},
		{
			name:      "panic due to second is nil",
			source:    NewIEnumerable[int](2),
			fEquals:   fEquals,
			wantPanic: true,
		},
		{
			name:      "panic due to no equality comparer",
			source:    NewIEnumerable[int](2),
			second:    NewIEnumerable[int](2),
			fEquals:   nil,
			wantPanic: true,
		},
		{
			name:    "keep the same order",
			source:  NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			second:  NewIEnumerable[int](4, 5, 6, 7),
			fEquals: fEquals,
			want:    NewIEnumerable[int](1, 2, 3, 6, 5, 4, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)
			bSecond := backupForAssetUnchanged(tt.second)

			defer deferWantPanicDepends(t, tt.wantPanic)

			got := tt.source.UnionBy(tt.second, tt.fEquals)
			gotData := got.exposeData()

			wantData := tt.want.exposeData()

			assert.True(t, reflect.DeepEqual(wantData, gotData))

			bSrc.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)
		})
	}

	t.Run("equality comparer not set on source", func(t *testing.T) {
		second := NewIEnumerable[int](1).WithDefaultComparers()

		defer deferWantPanicDepends(t, true)

		NewIEnumerable[int](1, 5, 2, 345, 65, 4574).Union(second)
	})

	t.Run("equality comparer set on first, not require on second", func(t *testing.T) {
		second := NewIEnumerable[int](1, 2, 99)

		assert.Equal(t, 7, NewIEnumerable[int](1, 5, 2, 345, 65, 4574, 1).WithDefaultComparers().Union(second).len())
	})
}
