package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Contains(t *testing.T) {
	fCompare := func(l, r int) int {
		return comparers.NumericComparer.CompareAny(l, r)
	}
	var tests = []struct {
		name     string
		source   IEnumerable[int]
		check    int
		fCompare func(v1, v2 int) int
		want     bool
	}{
		{
			name:     "empty source",
			source:   createEmptyIntEnumerable(),
			fCompare: fCompare,
			want:     false,
		},
		{
			name:     "single",
			source:   NewIEnumerable[int](2),
			check:    1,
			fCompare: fCompare,
			want:     false,
		},
		{
			name:     "single",
			source:   NewIEnumerable[int](2),
			check:    2,
			fCompare: fCompare,
			want:     true,
		},
		{
			name:     "many",
			source:   NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			check:    99,
			fCompare: fCompare,
			want:     false,
		},
		{
			name:     "negative",
			source:   NewIEnumerable[int](-1, -2, -3, -4, 55),
			check:    -4,
			fCompare: fCompare,
			want:     true,
		},
		{
			name:     "no equality comparer still ok since int has default comparer",
			source:   NewIEnumerable[int](2),
			check:    2,
			fCompare: nil,
			want:     true,
		},
		{
			name:     "no equality comparer still ok since int has default comparer",
			source:   NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			check:    3,
			fCompare: nil,
			want:     true,
		},
	}
	for _, tt := range tests {
		t.Run("Contains_"+tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			// CompareFunc
			got := tt.source.Contains(tt.check, tt.fCompare)

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)
		})
	}

	t.Run("auto-resolve comparer if default comparer not set", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4).
			WithDefaultComparer(nil)

		assert.True(t, ieSrc.Contains(3, nil))

		var cff func(v1, v2 int) int
		assert.True(t, ieSrc.Contains(3, cff))
		var cft CompareFunc[int]
		assert.True(t, ieSrc.Contains(3, cft))

		assert.Nil(t, e[int](ieSrc).defaultComparer)
	})

	t.Run("panic if no default resolver (Contains)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.Contains(MyInt64{}, nil)
	})

}
