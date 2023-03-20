package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_Except(t *testing.T) {
	fCompare := func(i1, i2 int) int {
		return comparers.NumericComparer.CompareTyped(i1, i2)
	}

	tests := []struct {
		name     string
		source   IEnumerable[int]
		second   IEnumerable[int]
		want     IEnumerable[int]
		fCompare func(int, int) int
		panic    bool
	}{
		{
			name:     "except not any",
			source:   NewIEnumerable[int](1, 2, 3),
			second:   NewIEnumerable[int](4, 5, 6, 7),
			want:     NewIEnumerable[int](1, 2, 3),
			fCompare: fCompare,
		},
		{
			name:     "returns distinct",
			source:   NewIEnumerable[int](1, 2, 2, 2, 3),
			second:   NewIEnumerable[int](1, 3),
			want:     NewIEnumerable[int](2),
			fCompare: fCompare,
		},
		{
			name:     "except all",
			source:   NewIEnumerable[int](1, 2, 3, 1, 2, 3),
			second:   NewIEnumerable[int](3, 1, 2),
			want:     NewIEnumerable[int](),
			fCompare: nil,
		},
		{
			name:     "auto-resolve comparer",
			source:   NewIEnumerable[int](1, 2, 3, 4),
			second:   NewIEnumerable[int](4, 5, 6, 7),
			want:     NewIEnumerable[int](1, 2, 3),
			fCompare: nil,
		},
		{
			name:     "except one",
			source:   NewIEnumerable[int](1, 2, 3, 4),
			second:   NewIEnumerable[int](4, 5, 6, 7),
			want:     NewIEnumerable[int](1, 2, 3),
			fCompare: fCompare,
		},
		{
			name:     "except some",
			source:   NewIEnumerable[int](1, 2, 3, 5, 6),
			second:   NewIEnumerable[int](4, 5, 6, 7),
			want:     NewIEnumerable[int](1, 2, 3),
			fCompare: fCompare,
		},
		{
			name:     "except when source empty",
			source:   NewIEnumerable[int](),
			second:   NewIEnumerable[int](4, 5, 6, 7),
			want:     NewIEnumerable[int](),
			fCompare: fCompare,
		},
		{
			name:     "except when second empty",
			source:   NewIEnumerable[int](1, 2, 3),
			second:   NewIEnumerable[int](),
			want:     NewIEnumerable[int](1, 2, 3),
			fCompare: fCompare,
		},
		{
			name:     "panic with nil src",
			source:   nil,
			second:   NewIEnumerable[int](4, 5, 6, 7),
			want:     NewIEnumerable[int](),
			fCompare: fCompare,
			panic:    true,
		},
		{
			name:     "panic with nil second",
			source:   NewIEnumerable[int](1, 2, 3),
			second:   nil,
			want:     NewIEnumerable[int](),
			fCompare: fCompare,
			panic:    true,
		},
		{
			name:     "panic with both nil",
			source:   nil,
			second:   nil,
			want:     NewIEnumerable[int](),
			fCompare: fCompare,
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

			// nil
			result := tt.source.Except(tt.second, nil)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), result.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)

			// CompareFunc
			result = tt.source.Except(tt.second, tt.fCompare)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), result.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)
		})
	}

	t.Run("auto-resolve comparer if default comparer not set", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](5, 2, 2, 6).
			WithDefaultComparer(nil)
		ieSecond := NewIEnumerable[int](5, 6, 7, 8)
		ieWant := NewIEnumerable[int](2)

		bSrc := backupForAssetUnchanged(ieSrc)

		got := ieSrc.Except(ieSecond, nil)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		var cff func(v1, v2 int) int
		got = ieSrc.Except(ieSecond, cff)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))
		var cft CompareFunc[int]
		got = ieSrc.Except(ieSecond, cft)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		bSrc.assertUnchanged(t, ieSrc)

		assert.Nil(t, e[int](ieSrc).defaultComparer)
	})

	t.Run("panic if no default resolver (Except)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.Except(ieSrc, nil)
	})
}
