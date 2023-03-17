package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Min_MinBy(t *testing.T) {
	fLess := func(l, r int) bool {
		return l < r
	}
	fCompare := func(l, r int) int {
		return comparers.IntComparer.Compare(l, r)
	}

	var tests = []struct {
		name     string
		source   IEnumerable[int]
		want     int
		fLess    func(t1, t2 int) bool
		fCompare func(v1, v2 int) int
		comparer comparers.IComparer[int]
	}{
		{
			name:     "single",
			source:   NewIEnumerable[int](2),
			want:     2,
			fLess:    fLess,
			fCompare: fCompare,
			comparer: comparers.IntComparer,
		},
		{
			name:     "single but duplicated",
			source:   NewIEnumerable[int](2, 2),
			want:     2,
			fLess:    fLess,
			fCompare: fCompare,
			comparer: comparers.IntComparer,
		},
		{
			name:     "many",
			source:   NewIEnumerable[int](-99, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			want:     -99,
			fLess:    fLess,
			fCompare: fCompare,
			comparer: comparers.IntComparer,
		},
		{
			name:     "negative",
			source:   NewIEnumerable[int](-1, -2, -3, -4, 55),
			want:     -4,
			fLess:    fLess,
			fCompare: fCompare,
			comparer: comparers.IntComparer,
		},
		{
			name:     "no equality comparer still ok since int has default comparer",
			source:   NewIEnumerable[int](2, 2),
			want:     2,
			fLess:    nil,
			fCompare: nil,
			comparer: nil,
		},
		{
			name:     "no equality comparer still ok since int has default comparer",
			source:   NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, -5, 4, 4),
			want:     -5,
			fLess:    nil,
			fCompare: nil,
			comparer: nil,
		},
	}
	for _, tt := range tests {
		t.Run("Min_"+tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			got := tt.source.Min()

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)
		})

		t.Run("MinBy_"+tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			// LessFunc
			got := tt.source.MinBy(tt.fLess)

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)

			got = tt.source.MinBy(LessFunc[int](tt.fLess))

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)

			// CompareFunc
			got = tt.source.MinBy(tt.fCompare)

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)

			got = tt.source.MinBy(CompareFunc[int](tt.fCompare))

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)

			// IComparer
			got = tt.source.MinBy(tt.comparer)

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)
		})
	}

	t.Run("auto-resolve comparer if default comparer not set", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1, 2, 2, 3, -3, 6, 6, 6, 5, 4, 4).
			WithDefaultComparer(nil)

		assert.Equal(t, -3, ieSrc.Min())

		assert.Equal(t, -3, ieSrc.MinBy(nil))

		var lff func(v1, v2 int) bool
		assert.Equal(t, -3, ieSrc.MinBy(lff))
		var lft LessFunc[int]
		assert.Equal(t, -3, ieSrc.MinBy(lft))

		var cff func(v1, v2 int) int
		assert.Equal(t, -3, ieSrc.MinBy(cff))
		var cft CompareFunc[int]
		assert.Equal(t, -3, ieSrc.MinBy(cft))

		var comparer comparers.IComparer[int]
		assert.Equal(t, -3, ieSrc.MinBy(comparer))
	})

	t.Run("panic if no default resolver (Min)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64](MyInt64{})

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.Min()
	})

	t.Run("panic if no default resolver (MinBy)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64](MyInt64{})

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.MinBy(nil)
	})

	t.Run("panic if not supported comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1)

		defer deferExpectPanicContains(t, "comparer must be", true)

		var badFunc func(v int) bool
		ieSrc.MinBy(badFunc)
	})

	t.Run("panic if not supported comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1)

		defer deferExpectPanicContains(t, "comparer must be", true)

		var badFunc GreaterFunc[int]
		ieSrc.MinBy(badFunc)
	})

	t.Run("panic if not supported comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1)

		defer deferExpectPanicContains(t, "comparer must be", true)

		var badFunc EqualsFunc[int]
		ieSrc.MinBy(badFunc)
	})

	t.Run("panic if empty source (Min)", func(t *testing.T) {
		ieSrc := NewIEnumerable[int]()

		defer deferExpectPanicContains(t, getErrorSrcContainsNoElement().Error(), true)

		ieSrc.Min()
	})

	t.Run("panic if empty source (MinBy)", func(t *testing.T) {
		ieSrc := NewIEnumerable[int]()

		defer deferExpectPanicContains(t, getErrorSrcContainsNoElement().Error(), true)

		ieSrc.MinBy(nil)
	})
}
