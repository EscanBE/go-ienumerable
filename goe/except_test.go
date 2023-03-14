package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_Except(t *testing.T) {
	t.Run("returns correctly", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4)
		ieSecond := NewIEnumerable[int](1, 3, 6, 5, 4)
		ieWant := NewIEnumerable[int](2)
		bSrc := backupForAssetUnchanged(ieSrc)

		got := ieSrc.Except(ieSecond)

		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		bSrc.assertUnchanged(t, ieSrc)
	})

	t.Run("panic if type not registered for default comparer", func(t *testing.T) {
		type MyInt64 struct{}

		defer deferExpectPanicContains(t, "no default comparer registered for [goe.MyInt64]")

		NewIEnumerable[MyInt64]().Except(NewIEnumerable[MyInt64]())
	})
}

func Test_enumerable_ExceptBy(t *testing.T) {
	t.Run("panic if type not registered for default comparer", func(t *testing.T) {
		type MyInt64 struct{}

		defer deferExpectPanicContains(t, "no default comparer registered for [goe.MyInt64]")

		NewIEnumerable[MyInt64]().ExceptBy(NewIEnumerable[MyInt64](), nil)
	})
}

func Test_enumerable_ExceptByComparer(t *testing.T) {
	t.Run("panic if type not registered for default comparer", func(t *testing.T) {
		type MyInt64 struct{}

		defer deferExpectPanicContains(t, "no default comparer registered for [goe.MyInt64]")

		NewIEnumerable[MyInt64]().ExceptByComparer(NewIEnumerable[MyInt64](), nil)
	})
}

func Test_enumerable_Except_ExceptBy_ExceptByComparer(t *testing.T) {
	equalityComparer := func(i1, i2 int) bool {
		return i1 == i2
	}

	tests := []struct {
		name             string
		source           IEnumerable[int]
		second           IEnumerable[int]
		want             IEnumerable[int]
		equalityComparer func(int, int) bool
		comparer         comparers.IComparer[int]
		panic            bool
	}{
		{
			name:             "except not any",
			source:           NewIEnumerable[int](1, 2, 3),
			second:           NewIEnumerable[int](4, 5, 6, 7),
			want:             NewIEnumerable[int](1, 2, 3),
			equalityComparer: equalityComparer,
			comparer:         comparers.IntComparer,
		},
		{
			name:             "except all",
			source:           NewIEnumerable[int](1, 2, 3, 1, 2, 3),
			second:           NewIEnumerable[int](3, 1, 2),
			want:             NewIEnumerable[int](),
			equalityComparer: nil,
			comparer:         nil,
		},
		{
			name:             "auto-resolve comparer",
			source:           NewIEnumerable[int](1, 2, 3, 4),
			second:           NewIEnumerable[int](4, 5, 6, 7),
			want:             NewIEnumerable[int](1, 2, 3),
			equalityComparer: nil,
			comparer:         nil,
		},
		{
			name:             "except one",
			source:           NewIEnumerable[int](1, 2, 3, 4),
			second:           NewIEnumerable[int](4, 5, 6, 7),
			want:             NewIEnumerable[int](1, 2, 3),
			equalityComparer: equalityComparer,
			comparer:         comparers.IntComparer,
		},
		{
			name:             "except some",
			source:           NewIEnumerable[int](1, 2, 3, 5, 6),
			second:           NewIEnumerable[int](4, 5, 6, 7),
			want:             NewIEnumerable[int](1, 2, 3),
			equalityComparer: equalityComparer,
			comparer:         comparers.IntComparer,
		},
		{
			name:             "except when source empty",
			source:           NewIEnumerable[int](),
			second:           NewIEnumerable[int](4, 5, 6, 7),
			want:             NewIEnumerable[int](),
			equalityComparer: equalityComparer,
			comparer:         comparers.IntComparer,
		},
		{
			name:             "except when second empty",
			source:           NewIEnumerable[int](1, 2, 3),
			second:           NewIEnumerable[int](),
			want:             NewIEnumerable[int](1, 2, 3),
			equalityComparer: equalityComparer,
			comparer:         comparers.IntComparer,
		},
		{
			name:             "panic with nil src",
			source:           nil,
			second:           NewIEnumerable[int](4, 5, 6, 7),
			equalityComparer: equalityComparer,
			comparer:         comparers.IntComparer,
			panic:            true,
		},
		{
			name:             "panic with nil second",
			source:           NewIEnumerable[int](1, 2, 3),
			second:           nil,
			equalityComparer: equalityComparer,
			comparer:         comparers.IntComparer,
			panic:            true,
		},
		{
			name:             "panic with both nil",
			source:           nil,
			second:           nil,
			equalityComparer: equalityComparer,
			comparer:         comparers.IntComparer,
			panic:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+"_Except", func(t *testing.T) {
			bSource := backupForAssetUnchanged(tt.source)
			bSecond := backupForAssetUnchanged(tt.second)

			defer deferWantPanicDepends(t, tt.panic)

			// Except
			resultOfExcept2 := tt.source.Except(tt.second)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), resultOfExcept2.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)
		})

		t.Run(tt.name+"_ExceptBy", func(t *testing.T) {
			bSource := backupForAssetUnchanged(tt.source)
			bSecond := backupForAssetUnchanged(tt.second)

			defer deferWantPanicDepends(t, tt.panic)

			// ExceptBy
			resultOfExcept2 := tt.source.ExceptBy(tt.second, tt.equalityComparer)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), resultOfExcept2.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)
		})

		t.Run(tt.name+"_ExceptByComparer", func(t *testing.T) {
			bSource := backupForAssetUnchanged(tt.source)
			bSecond := backupForAssetUnchanged(tt.second)

			defer deferWantPanicDepends(t, tt.panic)

			// ExceptByComparer
			resultOfExcept2 := tt.source.ExceptByComparer(tt.second, tt.comparer)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), resultOfExcept2.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)
		})
	}

	t.Run("except returns distinct", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](5, 2, 2, 6).WithDefaultComparers()
		ieSecond := NewIEnumerable[int](5, 6, 7, 8)
		bSource := backupForAssetUnchanged(ieSrc)
		bSecond := backupForAssetUnchanged(ieSecond)

		ieGot := ieSrc.Except(ieSecond)
		assert.Equal(t, 1, ieGot.Count())
		assert.Equal(t, 2, ieGot.ToArray()[0])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)

		ieGot = ieSrc.ExceptBy(ieSecond, nil)
		assert.Equal(t, 1, ieGot.Count())
		assert.Equal(t, 2, ieGot.ToArray()[0])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)

		ieGot = ieSrc.ExceptBy(ieSecond, func(v1, v2 int) bool {
			return v1 == v2
		})
		assert.Equal(t, 1, ieGot.Count())
		assert.Equal(t, 2, ieGot.ToArray()[0])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)

		ieGot = ieSrc.ExceptByComparer(ieSecond, nil)
		assert.Equal(t, 1, ieGot.Count())
		assert.Equal(t, 2, ieGot.ToArray()[0])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)

		ieGot = ieSrc.ExceptByComparer(ieSecond, comparers.IntComparer)
		assert.Equal(t, 1, ieGot.Count())
		assert.Equal(t, 2, ieGot.ToArray()[0])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)
	})
}
