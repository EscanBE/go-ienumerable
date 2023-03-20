package goe

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_Distinct(t *testing.T) {
	fCompare := func(v1, v2 int) int {
		return comparers.NumericComparer.CompareAny(v1, v2)
	}
	var tests = []struct {
		name     string
		source   IEnumerable[int]
		fCompare func(t1, t2 int) int
		want     IEnumerable[int]
	}{
		{
			name:     "empty source",
			source:   createEmptyIntEnumerable(),
			fCompare: fCompare,
			want:     createEmptyIntEnumerable(),
		},
		{
			name:     "distinct",
			source:   NewIEnumerable[int](2),
			fCompare: fCompare,
			want:     NewIEnumerable[int](2),
		},
		{
			name:     "distinct",
			source:   NewIEnumerable[int](2, 2),
			fCompare: fCompare,
			want:     NewIEnumerable[int](2),
		},
		{
			name:     "no equality comparer still ok since int has default comparer",
			source:   NewIEnumerable[int](2),
			fCompare: nil,
			want:     NewIEnumerable[int](2),
		},
		{
			name:     "keep the same order",
			source:   NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			fCompare: fCompare,
			want:     NewIEnumerable[int](1, 2, 3, 6, 5, 4),
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Distinct-%s", tt.name), func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			// CompareFunc
			got := tt.source.Distinct(tt.fCompare)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))

			bSrc.assertUnchanged(t, tt.source)
		})
	}

	t.Run("auto-resolve comparer if default comparer not set", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4).
			WithDefaultComparer(nil)
		ieWant := NewIEnumerable[int](1, 2, 3, 6, 5, 4)

		bSrc := backupForAssetUnchanged(ieSrc)

		got := ieSrc.Distinct(nil)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		var cff func(v1, v2 int) int
		got = ieSrc.Distinct(cff)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))
		var cft CompareFunc[int]
		got = ieSrc.Distinct(cft)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		assert.Nil(t, e[int](ieSrc).defaultComparer)

		bSrc.assertUnchanged(t, ieSrc)
	})

	t.Run("panic if no default resolver (Distinct)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.Distinct(nil)
	})
}
