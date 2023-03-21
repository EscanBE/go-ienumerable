package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"testing"
)

func Test_enumerable_Intersect(t *testing.T) {
	fEquals := func(v1, v2 int) bool {
		return comparers.NumericComparer.CompareAny(v1, v2) == 0
	}

	tests := []struct {
		name    string
		source  IEnumerable[int]
		second  IEnumerable[int]
		want    IEnumerable[int]
		fEquals OptionalEqualsFunc[int]
		panic   bool
	}{
		{
			name:    "intersect not any",
			source:  NewIEnumerable[int](1, 2, 3),
			second:  NewIEnumerable[int](4, 5, 6, 7),
			want:    NewIEnumerable[int](),
			fEquals: fEquals,
		},
		{
			name:    "no comparer",
			source:  NewIEnumerable[int](1, 2, 3),
			second:  NewIEnumerable[int](3, 4, 5, 6, 7),
			want:    NewIEnumerable[int](3),
			fEquals: nil,
		},
		{
			name:    "intersect one",
			source:  NewIEnumerable[int](1, 2, 3, 4),
			second:  NewIEnumerable[int](4, 5, 6, 7),
			want:    NewIEnumerable[int](4),
			fEquals: fEquals,
		},
		{
			name:    "intersect some",
			source:  NewIEnumerable[int](1, 2, 3, 5, 6),
			second:  NewIEnumerable[int](4, 5, 6, 7),
			want:    NewIEnumerable[int](5, 6),
			fEquals: fEquals,
		},
		{
			name:    "intersect all",
			source:  NewIEnumerable[int](1, 2, 3, 5, 6, 6),
			second:  NewIEnumerable[int](1, 1, 1, 3, 3, 3, 6, 6, 6, 5, 5, 5, 2, 2, 2),
			want:    NewIEnumerable[int](1, 2, 3, 5, 6),
			fEquals: fEquals,
		},
		{
			name:    "intersect when source empty",
			source:  NewIEnumerable[int](),
			second:  NewIEnumerable[int](4, 5, 6, 7),
			want:    NewIEnumerable[int](),
			fEquals: fEquals,
		},
		{
			name:    "intersect when second empty",
			source:  NewIEnumerable[int](1, 2, 3),
			second:  NewIEnumerable[int](),
			want:    NewIEnumerable[int](),
			fEquals: fEquals,
		},
		{
			name:    "intersect when both empty",
			source:  NewIEnumerable[int](),
			second:  NewIEnumerable[int](),
			want:    NewIEnumerable[int](),
			fEquals: fEquals,
		},
		{
			name:    "panic with nil src",
			source:  nil,
			second:  NewIEnumerable[int](4, 5, 6, 7),
			want:    NewIEnumerable[int](),
			fEquals: fEquals,
			panic:   true,
		},
		{
			name:    "panic with nil second",
			source:  NewIEnumerable[int](1, 2, 3),
			second:  nil,
			want:    NewIEnumerable[int](),
			fEquals: fEquals,
			panic:   true,
		},
		{
			name:    "panic with both nil",
			source:  nil,
			second:  nil,
			want:    NewIEnumerable[int](),
			fEquals: fEquals,
			panic:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+"_Intersect", func(t *testing.T) {
			bSource := backupForAssetUnchanged(tt.source)
			bSecond := backupForAssetUnchanged(tt.second)

			if tt.panic && tt.source == nil {
				return
			}

			defer deferWantPanicDepends(t, tt.panic)

			// nil
			result := tt.source.Intersect(tt.second, nil)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), result.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)

			// EqualsFunc
			result = tt.source.Intersect(tt.second, tt.fEquals)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), result.ToArray()))

			// nil
			result = tt.source.IntersectBy(tt.second, SelfSelector[int](), nil)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), result.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)

			// EqualsFunc
			var isEquals OptionalEqualsFunc[any]
			if tt.fEquals != nil {
				isEquals = func(v1, v2 any) bool {
					return tt.fEquals(v1.(int), v2.(int))
				}
			}
			result = tt.source.IntersectBy(tt.second, SelfSelector[int](), isEquals)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), result.ToArray()))

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

		// Intersect
		ieGot := ieSrc.Intersect(ieSecond, nil)
		assert.Equal(t, 1, ieGot.Count(nil))
		assert.Equal(t, 2, ieGot.ToArray()[0])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)
	})

	t.Run("intercept of nil", func(t *testing.T) {
		var nilStr *string
		ieSrc := NewIEnumerable[*string](nilStr, nilStr, nilStr)
		ieSecond := NewIEnumerable[*string](nilStr, nilStr, nilStr)

		ieSrc.WithDefaultComparer(nil)
		ieSecond.WithDefaultComparer(nil)

		bSource := backupForAssetUnchanged(ieSrc)
		bSecond := backupForAssetUnchanged(ieSecond)

		// Intersect
		ieGot := ieSrc.Intersect(ieSecond, nil)
		assert.Equal(t, 1, ieGot.Count(nil))
		assert.Equal(t, nilStr, ieGot.ToArray()[0])

		// IntersectBy
		ieGot = ieSrc.IntersectBy(ieSecond, SelfSelector[*string](), nil)
		assert.Equal(t, 1, ieGot.Count(nil))
		assert.Equal(t, nilStr, ieGot.ToArray()[0])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)
	})

	t.Run("intersect returns distinct", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](5, 2, 2, 6)
		ieSecond := NewIEnumerable[int](1, 2, 2, 3)
		bSource := backupForAssetUnchanged(ieSrc)
		bSecond := backupForAssetUnchanged(ieSecond)

		// Intersect
		ieGot := ieSrc.Intersect(ieSecond, nil)
		assert.Equal(t, 1, ieGot.Count(nil))
		assert.Equal(t, 2, ieGot.ToArray()[0])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)
	})

	t.Run("found comparer from second sequence", func(t *testing.T) {
		type MyInt64 struct {
			Value int
			Nil   bool
		}
		rad := rand.Intn(1_000_000_000 + 99)
		ieSrc := NewIEnumerable[MyInt64](
			MyInt64{
				Value: rad,
				Nil:   true,
			},
			MyInt64{
				Value: 2,
				Nil:   true,
			},
		)
		ieSecond := NewIEnumerable[MyInt64](
			MyInt64{
				Value: 1,
				Nil:   true,
			},
			MyInt64{
				Value: 2,
				Nil:   false,
			},
		)
		bSource := backupForAssetUnchanged(ieSrc)
		bSecond := backupForAssetUnchanged(ieSecond)

		// Intersect
		ieGot := ieSrc.IntersectBy(ieSecond, func(mi MyInt64) any {
			if mi.Nil {
				return nil
			} else {
				return mi.Value
			}
		}, nil)
		assert.Equal(t, 1, ieGot.Count(nil))
		assert.Equal(t, rad, ieGot.ToArray()[0].Value)

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)
	})

	t.Run("panic if no default resolver (Intersect)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.Intersect(ieSrc, nil)
	})

	t.Run("panic if no default resolver (IntersectBy)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64](MyInt64{})

		defer deferExpectPanicContains(t, getErrorFailedCompare2ElementsInArray().Error(), true)

		ieSrc.IntersectBy(ieSrc, SelfSelector[MyInt64](), nil)
	})

	t.Run("panic if no key selector (IntersectBy)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, getErrorKeySelectorNotNil().Error(), true)

		ieSrc.IntersectBy(ieSrc, nil, nil)
	})
}
