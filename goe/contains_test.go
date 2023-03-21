package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Contains(t *testing.T) {
	fEquals := func(l, r int) bool {
		return comparers.NumericComparer.CompareAny(l, r) == 0
	}
	var tests = []struct {
		name    string
		source  IEnumerable[int]
		check   int
		fEquals OptionalEqualsFunc[int]
		want    bool
	}{
		{
			name:    "empty source",
			source:  createEmptyIntEnumerable(),
			fEquals: fEquals,
			want:    false,
		},
		{
			name:    "single",
			source:  NewIEnumerable[int](2),
			check:   1,
			fEquals: fEquals,
			want:    false,
		},
		{
			name:    "single",
			source:  NewIEnumerable[int](2),
			check:   2,
			fEquals: fEquals,
			want:    true,
		},
		{
			name:    "many",
			source:  NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			check:   99,
			fEquals: fEquals,
			want:    false,
		},
		{
			name:    "negative",
			source:  NewIEnumerable[int](-1, -2, -3, -4, 55),
			check:   -4,
			fEquals: fEquals,
			want:    true,
		},
		{
			name:    "no equality comparer still ok since int has default comparer",
			source:  NewIEnumerable[int](2),
			check:   2,
			fEquals: nil,
			want:    true,
		},
		{
			name:    "no equality comparer still ok since int has default comparer",
			source:  NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			check:   3,
			fEquals: nil,
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run("Contains_"+tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			// CompareFunc
			got := tt.source.Contains(tt.check, tt.fEquals)

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)
		})
	}

	t.Run("auto-resolve comparer if default comparer not set", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4).
			WithDefaultComparer(nil)

		assert.True(t, ieSrc.Contains(3, nil))

		var cft OptionalEqualsFunc[int]
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
