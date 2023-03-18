package goe

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe/comparers2"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_Distinct_DistinctBy(t *testing.T) {
	fEquals := func(v1, v2 int) bool {
		return v1 == v2
	}
	fCompare := func(v1, v2 int) int {
		return comparers.NumericComparer.CompareAny(v1, v2)
	}
	var tests = []struct {
		name     string
		source   IEnumerable[int]
		fEquals  func(t1, t2 int) bool
		fCompare func(t1, t2 int) int
		comparer comparers.IComparer[any]
		want     IEnumerable[int]
	}{
		{
			name:     "empty source",
			source:   createEmptyIntEnumerable(),
			fEquals:  fEquals,
			fCompare: fCompare,
			comparer: comparers.NumericComparer,
			want:     createEmptyIntEnumerable(),
		},
		{
			name:     "distinct",
			source:   NewIEnumerable[int](2),
			fEquals:  fEquals,
			fCompare: fCompare,
			comparer: comparers.NumericComparer,
			want:     NewIEnumerable[int](2),
		},
		{
			name:     "distinct",
			source:   NewIEnumerable[int](2, 2),
			fEquals:  fEquals,
			fCompare: fCompare,
			comparer: comparers.NumericComparer,
			want:     NewIEnumerable[int](2),
		},
		{
			name:     "no equality comparer still ok since int has default comparer",
			source:   NewIEnumerable[int](2),
			fEquals:  nil,
			fCompare: nil,
			comparer: nil,
			want:     NewIEnumerable[int](2),
		},
		{
			name:     "keep the same order",
			source:   NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			fEquals:  fEquals,
			fCompare: fCompare,
			comparer: comparers.NumericComparer,
			want:     NewIEnumerable[int](1, 2, 3, 6, 5, 4),
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Distinct-%s", tt.name), func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			got := tt.source.Distinct()

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))

			bSrc.assertUnchanged(t, tt.source)
		})
		t.Run(fmt.Sprintf("DistinctBy-%s", tt.name), func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			// EqualsFunc
			got := tt.source.DistinctBy(tt.fEquals)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))

			bSrc.assertUnchanged(t, tt.source)

			got = tt.source.DistinctBy(EqualsFunc[int](tt.fEquals))

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))

			bSrc.assertUnchanged(t, tt.source)

			// CompareFunc
			got = tt.source.DistinctBy(tt.fCompare)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))

			bSrc.assertUnchanged(t, tt.source)

			got = tt.source.DistinctBy(CompareFunc[int](tt.fCompare))

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))

			bSrc.assertUnchanged(t, tt.source)

			// IComparer
			got = tt.source.DistinctBy(tt.comparer)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))

			bSrc.assertUnchanged(t, tt.source)
		})
	}

	t.Run("auto-resolve comparer if default comparer not set", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4).
			WithDefaultComparer(nil)
		ieWant := NewIEnumerable[int](1, 2, 3, 6, 5, 4)

		bSrc := backupForAssetUnchanged(ieSrc)

		got := ieSrc.Distinct()
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		got = ieSrc.DistinctBy(nil)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		var eff func(v1, v2 int) bool
		got = ieSrc.DistinctBy(eff)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))
		var eft EqualsFunc[int]
		got = ieSrc.DistinctBy(eft)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		var cff func(v1, v2 int) int
		got = ieSrc.DistinctBy(cff)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))
		var cft CompareFunc[int]
		got = ieSrc.DistinctBy(cft)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		var comparer comparers.IComparer[int]
		got = ieSrc.DistinctBy(comparer)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		assert.Nil(t, e[int](ieSrc).defaultComparer)

		bSrc.assertUnchanged(t, ieSrc)
	})

	t.Run("panic if no default resolver (Distinct)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.Distinct()
	})

	t.Run("panic if no default resolver (DistinctBy)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.DistinctBy(nil)
	})

	t.Run("panic if not supported comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int]()

		defer deferExpectPanicContains(t, "comparer must be", true)

		var badFunc func(v int) bool
		ieSrc.DistinctBy(badFunc)
	})

	t.Run("panic if not supported comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1)

		defer deferExpectPanicContains(t, "comparer must be", true)

		var badFunc LessFunc[int]
		ieSrc.DistinctBy(badFunc)
	})

	t.Run("panic if not supported comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1)

		defer deferExpectPanicContains(t, "comparer must be", true)

		var badFunc GreaterFunc[int]
		ieSrc.DistinctBy(badFunc)
	})
}
