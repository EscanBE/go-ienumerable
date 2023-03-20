package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Min_MinBy(t *testing.T) {
	fCompare := func(l, r any) int {
		return comparers.NumericComparer.CompareAny(l, r)
	}

	var tests = []struct {
		name     string
		source   IEnumerable[int]
		want     int
		fCompare func(v1, v2 any) int
		comparer comparers.IComparer[any]
	}{
		{
			name:     "single",
			source:   NewIEnumerable[int](2),
			want:     2,
			fCompare: fCompare,
			comparer: comparers.NumericComparer,
		},
		{
			name:     "single but duplicated",
			source:   NewIEnumerable[int](2, 2),
			want:     2,
			fCompare: fCompare,
			comparer: comparers.NumericComparer,
		},
		{
			name:     "many",
			source:   NewIEnumerable[int](-99, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			want:     -99,
			fCompare: fCompare,
			comparer: comparers.NumericComparer,
		},
		{
			name:     "negative",
			source:   NewIEnumerable[int](-1, -2, -3, -4, -55),
			want:     -55,
			fCompare: fCompare,
			comparer: comparers.NumericComparer,
		},
		{
			name:     "no input compare func still ok since int has default comparer",
			source:   NewIEnumerable[int](2, 2),
			want:     2,
			fCompare: nil,
			comparer: nil,
		},
		{
			name:     "no input compare func still ok since int has default comparer",
			source:   NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, -5, 4, 4),
			want:     -5,
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

			// CompareFunc
			got := tt.source.MinBy(test_getSelfSelector[int](), tt.fCompare)

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)
		})
	}

	t.Run("auto-resolve comparer if default comparer not set", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1, 2, 2, 3, -3, 6, 6, 6, 5, 4, 4).
			WithDefaultComparer(nil)

		assert.Equal(t, -3, ieSrc.Min())

		assert.Equal(t, -3, ieSrc.MinBy(test_getSelfSelector[int](), nil))

		var cff func(v1, v2 any) int
		assert.Equal(t, -3, ieSrc.MinBy(test_getSelfSelector[int](), cff))

		assert.Nil(t, e[int](ieSrc).defaultComparer)
	})

	t.Run("order nil types which support comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[*string](nil, nil, nil).WithDefaultComparer(nil)

		var nilString *string
		assert.Equal(t, nilString, ieSrc.Min())

		assert.Equal(t, nilString, ieSrc.MinBy(test_getSelfSelector[*string](), nil))

		assert.Nil(t, e[*string](ieSrc).defaultComparer)
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

		defer deferExpectPanicContains(t, getErrorFailedCompare2ElementsInArray().Error(), true)

		ieSrc.MinBy(test_getSelfSelector[MyInt64](), nil)
	})

	t.Run("panic if empty source (Min)", func(t *testing.T) {
		ieSrc := NewIEnumerable[int]()

		defer deferExpectPanicContains(t, getErrorSrcContainsNoElement().Error(), true)

		ieSrc.Min()
	})

	t.Run("panic if empty source (MinBy)", func(t *testing.T) {
		ieSrc := NewIEnumerable[int]()

		defer deferExpectPanicContains(t, getErrorSrcContainsNoElement().Error(), true)

		ieSrc.MinBy(test_getSelfSelector[int](), nil)
	})
}
