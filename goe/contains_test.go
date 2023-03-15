package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Contains_ContainsBy(t *testing.T) {
	fEquals := func(l, r int) bool {
		return l == r
	}
	fCompare := func(l, r int) int {
		return comparers.IntComparer.Compare(l, r)
	}
	var tests = []struct {
		name     string
		source   IEnumerable[int]
		check    int
		fEquals  func(t1, t2 int) bool
		fCompare func(v1, v2 int) int
		comparer comparers.IComparer[int]
		want     bool
	}{
		{
			name:     "empty source",
			source:   createEmptyIntEnumerable(),
			fEquals:  fEquals,
			fCompare: fCompare,
			comparer: comparers.IntComparer,
			want:     false,
		},
		{
			name:     "single",
			source:   NewIEnumerable[int](2),
			check:    1,
			fEquals:  fEquals,
			fCompare: fCompare,
			comparer: comparers.IntComparer,
			want:     false,
		},
		{
			name:     "single",
			source:   NewIEnumerable[int](2),
			check:    2,
			fEquals:  fEquals,
			fCompare: fCompare,
			comparer: comparers.IntComparer,
			want:     true,
		},
		{
			name:     "many",
			source:   NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			check:    99,
			fEquals:  fEquals,
			fCompare: fCompare,
			comparer: comparers.IntComparer,
			want:     false,
		},
		{
			name:     "negative",
			source:   NewIEnumerable[int](-1, -2, -3, -4, 55),
			check:    -4,
			fEquals:  fEquals,
			fCompare: fCompare,
			comparer: comparers.IntComparer,
			want:     true,
		},
		{
			name:     "no equality comparer still ok since int has default comparer",
			source:   NewIEnumerable[int](2),
			check:    2,
			fEquals:  nil,
			fCompare: nil,
			comparer: nil,
			want:     true,
		},
		{
			name:     "no equality comparer still ok since int has default comparer",
			source:   NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			check:    3,
			fEquals:  nil,
			fCompare: nil,
			comparer: nil,
			want:     true,
		},
	}
	for _, tt := range tests {
		t.Run("Contains_"+tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			got := tt.source.Contains(tt.check)

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)
		})

		t.Run("ContainsBy_"+tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			// EqualsFunc
			got := tt.source.ContainsBy(tt.check, tt.fEquals)

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)

			got = tt.source.ContainsBy(tt.check, EqualsFunc[int](tt.fEquals))

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)

			// CompareFunc
			got = tt.source.ContainsBy(tt.check, tt.fCompare)

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)

			got = tt.source.ContainsBy(tt.check, CompareFunc[int](tt.fCompare))

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)

			// IComparer
			got = tt.source.ContainsBy(tt.check, tt.comparer)

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)
		})
	}

	t.Run("auto-resolve comparer if default comparer not set", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4)
		eSrc := e[int](ieSrc)
		eSrc.defaultComparer = nil

		assert.True(t, ieSrc.Contains(3))

		assert.True(t, ieSrc.ContainsBy(3, nil))

		var eff func(v1, v2 int) bool
		assert.True(t, ieSrc.ContainsBy(3, eff))
		var eft EqualsFunc[int]
		assert.True(t, ieSrc.ContainsBy(3, eft))

		var cff func(v1, v2 int) int
		assert.True(t, ieSrc.ContainsBy(3, cff))
		var cft CompareFunc[int]
		assert.True(t, ieSrc.ContainsBy(3, cft))

		var comparer comparers.IComparer[int]
		assert.True(t, ieSrc.ContainsBy(3, comparer))
	})

	t.Run("panic if no default resolver (Contains)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.Contains(MyInt64{})
	})

	t.Run("panic if no default resolver (ContainsBy)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.ContainsBy(MyInt64{}, nil)
	})

	t.Run("panic if not supported comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int]()

		defer deferExpectPanicContains(t, "comparer must be", true)

		var badFunc func(v int) bool
		ieSrc.ContainsBy(1, badFunc)
	})
}
