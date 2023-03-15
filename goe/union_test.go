package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_Union_UnionBy_UnionByComparer(t *testing.T) {
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
			name:             "union not any duplicated",
			source:           NewIEnumerable[int](1, 2, 3),
			second:           NewIEnumerable[int](4, 5, 6, 7),
			want:             NewIEnumerable[int](1, 2, 3, 4, 5, 6, 7),
			equalityComparer: equalityComparer,
			comparer:         comparers.IntComparer,
		},
		{
			name:             "union with duplicated",
			source:           NewIEnumerable[int](1, 2, 2, 3),
			second:           NewIEnumerable[int](1, 3),
			want:             NewIEnumerable[int](1, 2, 3),
			equalityComparer: equalityComparer,
			comparer:         comparers.IntComparer,
		},
		{
			name:             "no comparer",
			source:           NewIEnumerable[int](1, 2, 3),
			second:           NewIEnumerable[int](4, 5, 6, 7),
			want:             NewIEnumerable[int](1, 2, 3, 4, 5, 6, 7),
			equalityComparer: nil,
			comparer:         nil,
		},
		{
			name:             "union one",
			source:           NewIEnumerable[int](1, 1, 1, 1, 1),
			second:           NewIEnumerable[int](1, 1, 1, 1),
			want:             NewIEnumerable[int](1),
			equalityComparer: equalityComparer,
			comparer:         comparers.IntComparer,
		},
		{
			name:             "union some",
			source:           NewIEnumerable[int](2, 2, 2, 2, 3, 3, 3, 3),
			second:           NewIEnumerable[int](1, 1, 1, 1, 1, 2, 3),
			want:             NewIEnumerable[int](2, 3, 1),
			equalityComparer: equalityComparer,
			comparer:         comparers.IntComparer,
		},
		{
			name:             "union all",
			source:           NewIEnumerable[int](1, 2, 3, 5, 6, 6),
			second:           NewIEnumerable[int](1, 1, 1, 3, 3, 3, 6, 6, 6, 5, 5, 5, 2, 2, 2),
			want:             NewIEnumerable[int](1, 2, 3, 5, 6),
			equalityComparer: equalityComparer,
			comparer:         comparers.IntComparer,
		},
		{
			name:             "union when source empty",
			source:           NewIEnumerable[int](),
			second:           NewIEnumerable[int](4, 5, 6, 7),
			want:             NewIEnumerable[int](4, 5, 6, 7),
			equalityComparer: equalityComparer,
			comparer:         comparers.IntComparer,
		},
		{
			name:             "union when second empty",
			source:           NewIEnumerable[int](1, 2, 3),
			second:           NewIEnumerable[int](),
			want:             NewIEnumerable[int](1, 2, 3),
			equalityComparer: equalityComparer,
			comparer:         comparers.IntComparer,
		},
		{
			name:             "union when both empty",
			source:           NewIEnumerable[int](),
			second:           NewIEnumerable[int](),
			want:             NewIEnumerable[int](),
			equalityComparer: equalityComparer,
			comparer:         comparers.IntComparer,
		},
		{
			name:             "panic with nil src",
			source:           nil,
			second:           NewIEnumerable[int](4, 5, 6, 7),
			want:             NewIEnumerable[int](),
			equalityComparer: equalityComparer,
			comparer:         comparers.IntComparer,
			panic:            true,
		},
		{
			name:             "panic with nil second",
			source:           NewIEnumerable[int](1, 2, 3),
			second:           nil,
			want:             NewIEnumerable[int](),
			equalityComparer: equalityComparer,
			comparer:         comparers.IntComparer,
			panic:            true,
		},
		{
			name:             "panic with both nil",
			source:           nil,
			second:           nil,
			want:             NewIEnumerable[int](),
			equalityComparer: equalityComparer,
			comparer:         comparers.IntComparer,
			panic:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+"_Union", func(t *testing.T) {
			bSource := backupForAssetUnchanged(tt.source)
			bSecond := backupForAssetUnchanged(tt.second)

			if tt.panic && tt.source == nil {
				return
			}
			defer deferWantPanicDepends(t, tt.panic)

			// Union
			resultOfUnion2 := tt.source.Union(tt.second)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), resultOfUnion2.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)
		})

		t.Run(tt.name+"_UnionBy", func(t *testing.T) {
			bSource := backupForAssetUnchanged(tt.source)
			bSecond := backupForAssetUnchanged(tt.second)

			if tt.panic && tt.source == nil {
				return
			}
			defer deferWantPanicDepends(t, tt.panic)

			// Union
			resultOfUnion2 := tt.source.UnionBy(tt.second, tt.equalityComparer)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), resultOfUnion2.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)
		})

		t.Run(tt.name+"_UnionByComparer", func(t *testing.T) {
			bSource := backupForAssetUnchanged(tt.source)
			bSecond := backupForAssetUnchanged(tt.second)

			if tt.panic && tt.source == nil {
				return
			}
			defer deferWantPanicDepends(t, tt.panic)

			// Union
			resultOfUnion2 := tt.source.UnionByComparer(tt.second, tt.comparer)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), resultOfUnion2.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)
		})
	}

	t.Run("auto-resolve comparer if default comparer is nil", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](5, 2, 2, 6)
		ieSecond := NewIEnumerable[int](1, 2, 2, 3)

		ieSrc.WithDefaultComparer(nil)
		ieSecond.WithDefaultComparer(nil)

		bSource := backupForAssetUnchanged(ieSrc)
		bSecond := backupForAssetUnchanged(ieSecond)

		// Union
		ieGot := ieSrc.Union(ieSecond)
		assert.Equal(t, 5, ieGot.Count())
		assert.Equal(t, 5, ieGot.ToArray()[0])
		assert.Equal(t, 2, ieGot.ToArray()[1])
		assert.Equal(t, 6, ieGot.ToArray()[2])
		assert.Equal(t, 1, ieGot.ToArray()[3])
		assert.Equal(t, 3, ieGot.ToArray()[4])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)

		// UnionBy
		ieGot = ieSrc.UnionBy(ieSecond, nil)
		assert.Equal(t, 5, ieGot.Count())
		assert.Equal(t, 5, ieGot.ToArray()[0])
		assert.Equal(t, 2, ieGot.ToArray()[1])
		assert.Equal(t, 6, ieGot.ToArray()[2])
		assert.Equal(t, 1, ieGot.ToArray()[3])
		assert.Equal(t, 3, ieGot.ToArray()[4])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)

		// UnionByComparer
		ieGot = ieSrc.UnionByComparer(ieSecond, nil)
		assert.Equal(t, 5, ieGot.Count())
		assert.Equal(t, 5, ieGot.ToArray()[0])
		assert.Equal(t, 2, ieGot.ToArray()[1])
		assert.Equal(t, 6, ieGot.ToArray()[2])
		assert.Equal(t, 1, ieGot.ToArray()[3])
		assert.Equal(t, 3, ieGot.ToArray()[4])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)
	})

	t.Run("union returns distinct", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](5, 2, 2, 6)
		ieSecond := NewIEnumerable[int](1, 2, 2, 3)
		bSource := backupForAssetUnchanged(ieSrc)
		bSecond := backupForAssetUnchanged(ieSecond)

		// Union
		ieGot := ieSrc.Union(ieSecond)
		assert.Equal(t, 5, ieGot.Count())
		assert.Equal(t, 5, ieGot.ToArray()[0])
		assert.Equal(t, 2, ieGot.ToArray()[1])
		assert.Equal(t, 6, ieGot.ToArray()[2])
		assert.Equal(t, 1, ieGot.ToArray()[3])
		assert.Equal(t, 3, ieGot.ToArray()[4])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)

		// UnionBy
		ieGot = ieSrc.UnionBy(ieSecond, nil)
		assert.Equal(t, 5, ieGot.Count())
		assert.Equal(t, 5, ieGot.ToArray()[0])
		assert.Equal(t, 2, ieGot.ToArray()[1])
		assert.Equal(t, 6, ieGot.ToArray()[2])
		assert.Equal(t, 1, ieGot.ToArray()[3])
		assert.Equal(t, 3, ieGot.ToArray()[4])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)

		ieGot = ieSrc.UnionBy(ieSecond, func(v1, v2 int) bool {
			return v1 == v2
		})
		assert.Equal(t, 5, ieGot.Count())
		assert.Equal(t, 5, ieGot.ToArray()[0])
		assert.Equal(t, 2, ieGot.ToArray()[1])
		assert.Equal(t, 6, ieGot.ToArray()[2])
		assert.Equal(t, 1, ieGot.ToArray()[3])
		assert.Equal(t, 3, ieGot.ToArray()[4])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)

		// UnionByComparer
		ieGot = ieSrc.UnionByComparer(ieSecond, nil)
		assert.Equal(t, 5, ieGot.Count())
		assert.Equal(t, 5, ieGot.ToArray()[0])
		assert.Equal(t, 2, ieGot.ToArray()[1])
		assert.Equal(t, 6, ieGot.ToArray()[2])
		assert.Equal(t, 1, ieGot.ToArray()[3])
		assert.Equal(t, 3, ieGot.ToArray()[4])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)

		ieGot = ieSrc.UnionByComparer(ieSecond, comparers.IntComparer)
		assert.Equal(t, 5, ieGot.Count())
		assert.Equal(t, 5, ieGot.ToArray()[0])
		assert.Equal(t, 2, ieGot.ToArray()[1])
		assert.Equal(t, 6, ieGot.ToArray()[2])
		assert.Equal(t, 1, ieGot.ToArray()[3])
		assert.Equal(t, 3, ieGot.ToArray()[4])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)
	})
}
