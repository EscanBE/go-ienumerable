package goe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_Intersect_IntersectBy(t *testing.T) {
	equalityComparer := func(i1, i2 int) bool {
		return i1 == i2
	}

	tests := []struct {
		name                 string
		source               IEnumerable[int]
		second               IEnumerable[int]
		want                 IEnumerable[int]
		wantPanicIntersect   bool
		equalityComparer     func(int, int) bool
		wantPanicIntersectBy bool
	}{
		{
			name:             "intersect not any",
			source:           injectIntComparers(NewIEnumerable[int](1, 2, 3)),
			second:           NewIEnumerable[int](4, 5, 6, 7),
			want:             NewIEnumerable[int](),
			equalityComparer: equalityComparer,
		},
		{
			name:                 "no comparer",
			source:               NewIEnumerable[int](1, 2, 3),
			second:               NewIEnumerable[int](4, 5, 6, 7),
			wantPanicIntersect:   true,
			wantPanicIntersectBy: true,
		},
		{
			name:             "intersect one",
			source:           injectIntComparers(NewIEnumerable[int](1, 2, 3, 4)),
			second:           NewIEnumerable[int](4, 5, 6, 7),
			want:             NewIEnumerable[int](4),
			equalityComparer: equalityComparer,
		},
		{
			name:             "intersect some",
			source:           injectIntComparers(NewIEnumerable[int](1, 2, 3, 5, 6)),
			second:           NewIEnumerable[int](4, 5, 6, 7),
			want:             NewIEnumerable[int](5, 6),
			equalityComparer: equalityComparer,
		},
		{
			name:             "intersect when source empty",
			source:           injectIntComparers(NewIEnumerable[int]()),
			second:           NewIEnumerable[int](4, 5, 6, 7),
			want:             NewIEnumerable[int](),
			equalityComparer: equalityComparer,
		},
		{
			name:             "intersect when second empty",
			source:           injectIntComparers(NewIEnumerable[int](1, 2, 3)),
			second:           NewIEnumerable[int](),
			want:             NewIEnumerable[int](),
			equalityComparer: equalityComparer,
		},
		{
			name:             "intersect when both empty",
			source:           injectIntComparers(NewIEnumerable[int]()),
			second:           NewIEnumerable[int](),
			want:             NewIEnumerable[int](),
			equalityComparer: equalityComparer,
		},
		{
			name:                 "panic with nil src",
			source:               nil,
			second:               NewIEnumerable[int](4, 5, 6, 7),
			wantPanicIntersect:   true,
			equalityComparer:     equalityComparer,
			wantPanicIntersectBy: true,
		},
		{
			name:                 "panic with nil second",
			source:               NewIEnumerable[int](1, 2, 3),
			second:               nil,
			wantPanicIntersect:   true,
			equalityComparer:     equalityComparer,
			wantPanicIntersectBy: true,
		},
		{
			name:                 "panic with both nil",
			source:               nil,
			second:               nil,
			wantPanicIntersect:   true,
			equalityComparer:     equalityComparer,
			wantPanicIntersectBy: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+"_Intersect", func(t *testing.T) {
			bSource := backupForAssetUnchanged(tt.source)
			bSecond := backupForAssetUnchanged(tt.second)

			defer deferWantPanicDepends(t, tt.wantPanicIntersect)

			// Intersect
			resultOfIntersect2 := tt.source.Intersect(tt.second)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), resultOfIntersect2.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)
		})
		t.Run(tt.name+"_IntersectBy", func(t *testing.T) {
			bSource := backupForAssetUnchanged(tt.source)
			bSecond := backupForAssetUnchanged(tt.second)

			defer deferWantPanicDepends(t, tt.wantPanicIntersectBy)

			// Intersect
			resultOfIntersect2 := tt.source.IntersectBy(tt.second, tt.equalityComparer)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), resultOfIntersect2.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)
		})
	}

	t.Run("intersect returns distinct", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](5, 2, 2, 6).WithDefaultComparers()
		ieSecond := NewIEnumerable[int](1, 2, 2, 3)
		bSource := backupForAssetUnchanged(ieSrc)
		bSecond := backupForAssetUnchanged(ieSecond)

		ieGot := ieSrc.Intersect(ieSecond)
		assert.Equal(t, 1, ieGot.Count())
		assert.Equal(t, 2, ieGot.ToArray()[0])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)

		ieGot = ieSrc.IntersectBy(ieSecond, nil)
		assert.Equal(t, 1, ieGot.Count())
		assert.Equal(t, 2, ieGot.ToArray()[0])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)

		ieGot = ieSrc.IntersectBy(ieSecond, func(v1, v2 int) bool {
			return v1 == v2
		})
		assert.Equal(t, 1, ieGot.Count())
		assert.Equal(t, 2, ieGot.ToArray()[0])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)
	})
}
