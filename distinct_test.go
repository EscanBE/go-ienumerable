package go_ienumerable

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_Distinct(t *testing.T) {
	fEquals := func(t1, t2 int) bool {
		return t1 == t2
	}
	var tests = []struct {
		name      string
		source    IEnumerable[int]
		fEquals   func(t1, t2 int) bool
		want      IEnumerable[int]
		wantPanic bool
	}{
		{
			name:    "empty source",
			source:  createEmptyIntEnumerable(),
			fEquals: fEquals,
			want:    createEmptyIntEnumerable(),
		},
		{
			name:    "distinct",
			source:  NewIEnumerable[int](2),
			fEquals: fEquals,
			want:    NewIEnumerable[int](2),
		},
		{
			name:      "panic",
			source:    NewIEnumerable[int](2),
			fEquals:   nil,
			wantPanic: true,
		},
		{
			name:    "keep the same order",
			source:  NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			fEquals: fEquals,
			want:    NewIEnumerable[int](1, 2, 3, 6, 5, 4),
		},
		{
			name:    "distinct",
			source:  NewIEnumerable[int](2, 2),
			fEquals: fEquals,
			want:    NewIEnumerable[int](2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			defer deferWantPanicDepends(t, tt.wantPanic)

			got := tt.source.DistinctBy(tt.fEquals)
			gotData := got.exposeData()

			wantData := tt.want.exposeData()

			assert.True(t, reflect.DeepEqual(wantData, gotData))

			bSrc.assertUnchanged(t, tt.source)

		})
	}

	t.Run("equality comparer not set", func(t *testing.T) {
		defer deferWantPanicDepends(t, true)

		NewIEnumerable[int](1, 5, 2, 345, 65, 4574).Distinct()
	})

	t.Run("equality comparer set", func(t *testing.T) {
		assert.Equal(t, 6, NewIEnumerable[int](1, 5, 2, 345, 65, 4574, 1).WithDefaultComparers().Distinct().len())
	})
}
