package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Max_MaxBy(t *testing.T) {
	fGreater := func(l, r int) bool {
		return l > r
	}
	fCompare := func(l, r int) int {
		return comparers.NumericComparer.CompareAny(l, r)
	}

	var tests = []struct {
		name     string
		source   IEnumerable[int]
		want     int
		fGreater func(t1, t2 int) bool
		fCompare func(v1, v2 int) int
		comparer comparers.IComparer[any]
	}{
		{
			name:     "single",
			source:   NewIEnumerable[int](2),
			want:     2,
			fGreater: fGreater,
			fCompare: fCompare,
			comparer: comparers.NumericComparer,
		},
		{
			name:     "single but duplicated",
			source:   NewIEnumerable[int](2, 2),
			want:     2,
			fGreater: fGreater,
			fCompare: fCompare,
			comparer: comparers.NumericComparer,
		},
		{
			name:     "many",
			source:   NewIEnumerable[int](-99, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			want:     6,
			fGreater: fGreater,
			fCompare: fCompare,
			comparer: comparers.NumericComparer,
		},
		{
			name:     "negative",
			source:   NewIEnumerable[int](-1, -2, -3, -4, -55),
			want:     -1,
			fGreater: fGreater,
			fCompare: fCompare,
			comparer: comparers.NumericComparer,
		},
		{
			name:     "no equality comparer still ok since int has default comparer",
			source:   NewIEnumerable[int](2, 2),
			want:     2,
			fGreater: nil,
			fCompare: nil,
			comparer: nil,
		},
		{
			name:     "no equality comparer still ok since int has default comparer",
			source:   NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, -5, 4, 4),
			want:     6,
			fGreater: nil,
			fCompare: nil,
			comparer: nil,
		},
	}
	for _, tt := range tests {
		t.Run("Max_"+tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			got := tt.source.Max()

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)
		})

		t.Run("MaxBy_"+tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			// GreaterFunc
			got := tt.source.MaxBy(tt.fGreater)

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)

			got = tt.source.MaxBy(GreaterFunc[int](tt.fGreater))

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)

			// CompareFunc
			got = tt.source.MaxBy(tt.fCompare)

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)

			got = tt.source.MaxBy(CompareFunc[int](tt.fCompare))

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)

			// IComparer
			got = tt.source.MaxBy(tt.comparer)

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)
		})
	}

	t.Run("auto-resolve comparer if default comparer not set", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1, 2, 2, 3, -3, 6, 6, 6, 5, 4, 4).
			WithDefaultComparer(nil)

		assert.Equal(t, 6, ieSrc.Max())

		assert.Equal(t, 6, ieSrc.MaxBy(nil))

		var gff func(v1, v2 int) bool
		assert.Equal(t, 6, ieSrc.MaxBy(gff))
		var gft GreaterFunc[int]
		assert.Equal(t, 6, ieSrc.MaxBy(gft))

		var cff func(v1, v2 int) int
		assert.Equal(t, 6, ieSrc.MaxBy(cff))
		var cft CompareFunc[int]
		assert.Equal(t, 6, ieSrc.MaxBy(cft))

		var comparer comparers.IComparer[int]
		assert.Equal(t, 6, ieSrc.MaxBy(comparer))

		assert.Nil(t, e[int](ieSrc).defaultComparer)
	})

	t.Run("panic if no default resolver (Max)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64](MyInt64{})

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.Max()
	})

	t.Run("panic if no default resolver (MaxBy)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64](MyInt64{})

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.MaxBy(nil)
	})

	t.Run("panic if not supported comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1)

		defer deferExpectPanicContains(t, "comparer must be", true)

		var badFunc func(v int) bool
		ieSrc.MaxBy(badFunc)
	})

	t.Run("panic if not supported comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1)

		defer deferExpectPanicContains(t, "comparer must be", true)

		var badFunc LessFunc[int]
		ieSrc.MaxBy(badFunc)
	})

	t.Run("panic if not supported comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1)

		defer deferExpectPanicContains(t, "comparer must be", true)

		var badFunc EqualsFunc[int]
		ieSrc.MaxBy(badFunc)
	})

	t.Run("panic if empty source (Max)", func(t *testing.T) {
		ieSrc := NewIEnumerable[int]()

		defer deferExpectPanicContains(t, getErrorSrcContainsNoElement().Error(), true)

		ieSrc.Max()
	})

	t.Run("panic if empty source (MaxBy)", func(t *testing.T) {
		ieSrc := NewIEnumerable[int]()

		defer deferExpectPanicContains(t, getErrorSrcContainsNoElement().Error(), true)

		ieSrc.MaxBy(nil)
	})
}
