package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_Except_ExceptBy(t *testing.T) {
	fEquals := func(i1, i2 int) bool {
		return i1 == i2
	}
	fCompare := func(i1, i2 int) int {
		return comparers.IntComparer.Compare(i1, i2)
	}

	tests := []struct {
		name     string
		source   IEnumerable[int]
		second   IEnumerable[int]
		want     IEnumerable[int]
		fEquals  func(int, int) bool
		fCompare func(int, int) int
		comparer comparers.IComparer[int]
		panic    bool
	}{
		{
			name:     "except not any",
			source:   NewIEnumerable[int](1, 2, 3),
			second:   NewIEnumerable[int](4, 5, 6, 7),
			want:     NewIEnumerable[int](1, 2, 3),
			fEquals:  fEquals,
			fCompare: fCompare,
			comparer: comparers.IntComparer,
		},
		{
			name:     "returns distinct",
			source:   NewIEnumerable[int](1, 2, 2, 2, 3),
			second:   NewIEnumerable[int](1, 3),
			want:     NewIEnumerable[int](2),
			fEquals:  fEquals,
			fCompare: fCompare,
			comparer: comparers.IntComparer,
		},
		{
			name:     "except all",
			source:   NewIEnumerable[int](1, 2, 3, 1, 2, 3),
			second:   NewIEnumerable[int](3, 1, 2),
			want:     NewIEnumerable[int](),
			fEquals:  nil,
			fCompare: nil,
			comparer: nil,
		},
		{
			name:     "auto-resolve comparer",
			source:   NewIEnumerable[int](1, 2, 3, 4),
			second:   NewIEnumerable[int](4, 5, 6, 7),
			want:     NewIEnumerable[int](1, 2, 3),
			fEquals:  nil,
			fCompare: nil,
			comparer: nil,
		},
		{
			name:     "except one",
			source:   NewIEnumerable[int](1, 2, 3, 4),
			second:   NewIEnumerable[int](4, 5, 6, 7),
			want:     NewIEnumerable[int](1, 2, 3),
			fEquals:  fEquals,
			fCompare: fCompare,
			comparer: comparers.IntComparer,
		},
		{
			name:     "except some",
			source:   NewIEnumerable[int](1, 2, 3, 5, 6),
			second:   NewIEnumerable[int](4, 5, 6, 7),
			want:     NewIEnumerable[int](1, 2, 3),
			fEquals:  fEquals,
			fCompare: fCompare,
			comparer: comparers.IntComparer,
		},
		{
			name:     "except when source empty",
			source:   NewIEnumerable[int](),
			second:   NewIEnumerable[int](4, 5, 6, 7),
			want:     NewIEnumerable[int](),
			fEquals:  fEquals,
			fCompare: fCompare,
			comparer: comparers.IntComparer,
		},
		{
			name:     "except when second empty",
			source:   NewIEnumerable[int](1, 2, 3),
			second:   NewIEnumerable[int](),
			want:     NewIEnumerable[int](1, 2, 3),
			fEquals:  fEquals,
			fCompare: fCompare,
			comparer: comparers.IntComparer,
		},
		{
			name:     "panic with nil src",
			source:   nil,
			second:   NewIEnumerable[int](4, 5, 6, 7),
			want:     NewIEnumerable[int](),
			fEquals:  fEquals,
			fCompare: fCompare,
			comparer: comparers.IntComparer,
			panic:    true,
		},
		{
			name:     "panic with nil second",
			source:   NewIEnumerable[int](1, 2, 3),
			second:   nil,
			want:     NewIEnumerable[int](),
			fEquals:  fEquals,
			fCompare: fCompare,
			comparer: comparers.IntComparer,
			panic:    true,
		},
		{
			name:     "panic with both nil",
			source:   nil,
			second:   nil,
			want:     NewIEnumerable[int](),
			fEquals:  fEquals,
			fCompare: fCompare,
			comparer: comparers.IntComparer,
			panic:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+"_Except", func(t *testing.T) {
			bSource := backupForAssetUnchanged(tt.source)
			bSecond := backupForAssetUnchanged(tt.second)

			if tt.panic && tt.source == nil {
				return
			}
			defer deferWantPanicDepends(t, tt.panic)

			// Except
			result := tt.source.Except(tt.second)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), result.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)
		})

		t.Run(tt.name+"_ExceptBy", func(t *testing.T) {
			bSource := backupForAssetUnchanged(tt.source)
			bSecond := backupForAssetUnchanged(tt.second)

			if tt.panic && tt.source == nil {
				return
			}
			defer deferWantPanicDepends(t, tt.panic)

			// nil
			result := tt.source.ExceptBy(tt.second, nil)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), result.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)

			// EqualsFunc
			result = tt.source.ExceptBy(tt.second, tt.fEquals)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), result.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)

			result = tt.source.ExceptBy(tt.second, EqualsFunc[int](tt.fEquals))

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), result.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)

			// CompareFunc
			result = tt.source.ExceptBy(tt.second, tt.fCompare)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), result.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)

			result = tt.source.ExceptBy(tt.second, CompareFunc[int](tt.fCompare))

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), result.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)

			// IComparer
			result = tt.source.ExceptBy(tt.second, tt.comparer)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), result.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)
		})
	}

	t.Run("auto-resolve comparer if default comparer not set", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](5, 2, 2, 6)
		ieSecond := NewIEnumerable[int](5, 6, 7, 8)
		ieWant := NewIEnumerable[int](2)

		eSrc := e[int](ieSrc)
		eSrc.defaultComparer = nil

		bSrc := backupForAssetUnchanged(ieSrc)

		got := ieSrc.ExceptBy(ieSecond, nil)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		var eff func(v1, v2 int) bool
		got = ieSrc.ExceptBy(ieSecond, eff)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))
		var eft EqualsFunc[int]
		got = ieSrc.ExceptBy(ieSecond, eft)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		var cff func(v1, v2 int) int
		got = ieSrc.ExceptBy(ieSecond, cff)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))
		var cft CompareFunc[int]
		got = ieSrc.ExceptBy(ieSecond, cft)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		var comparer comparers.IComparer[int]
		got = ieSrc.ExceptBy(ieSecond, comparer)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		bSrc.assertUnchanged(t, ieSrc)
	})

	t.Run("panic if no default resolver (Except)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.Except(ieSrc)
	})

	t.Run("panic if no default resolver (ExceptBy)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.ExceptBy(ieSrc, nil)
	})

	t.Run("panic if not supported comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int]()

		defer deferExpectPanicContains(t, "comparer must be", true)

		var badFunc func(v int) bool
		ieSrc.ExceptBy(ieSrc, badFunc)
	})
}
