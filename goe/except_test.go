package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_Except_ExceptBy(t *testing.T) {
	fEquals := func(i1, i2 int) bool {
		return comparers.NumericComparer.CompareTyped(i1, i2) == 0
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
			name:    "except not any",
			source:  NewIEnumerable[int](1, 2, 3),
			second:  NewIEnumerable[int](4, 5, 6, 7),
			want:    NewIEnumerable[int](1, 2, 3),
			fEquals: fEquals,
		},
		{
			name:    "returns distinct",
			source:  NewIEnumerable[int](1, 2, 2, 2, 3),
			second:  NewIEnumerable[int](1, 3),
			want:    NewIEnumerable[int](2),
			fEquals: fEquals,
		},
		{
			name:    "returns distinct",
			source:  NewIEnumerable[int](1, 2, 2, 2, 3),
			second:  NewIEnumerable[int](),
			want:    NewIEnumerable[int](1, 2, 3),
			fEquals: fEquals,
		},
		{
			name:    "except all",
			source:  NewIEnumerable[int](1, 2, 3, 1, 2, 3),
			second:  NewIEnumerable[int](3, 1, 2),
			want:    NewIEnumerable[int](),
			fEquals: nil,
		},
		{
			name:    "auto-resolve comparer",
			source:  NewIEnumerable[int](1, 2, 3, 4),
			second:  NewIEnumerable[int](4, 5, 6, 7),
			want:    NewIEnumerable[int](1, 2, 3),
			fEquals: nil,
		},
		{
			name:    "except one",
			source:  NewIEnumerable[int](1, 2, 3, 4),
			second:  NewIEnumerable[int](4, 5, 6, 7),
			want:    NewIEnumerable[int](1, 2, 3),
			fEquals: fEquals,
		},
		{
			name:    "except some",
			source:  NewIEnumerable[int](1, 2, 3, 5, 6),
			second:  NewIEnumerable[int](4, 5, 6, 7),
			want:    NewIEnumerable[int](1, 2, 3),
			fEquals: fEquals,
		},
		{
			name:    "except when source empty",
			source:  NewIEnumerable[int](),
			second:  NewIEnumerable[int](4, 5, 6, 7),
			want:    NewIEnumerable[int](),
			fEquals: fEquals,
		},
		{
			name:    "except when second empty",
			source:  NewIEnumerable[int](1, 2, 3),
			second:  NewIEnumerable[int](),
			want:    NewIEnumerable[int](1, 2, 3),
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

			// EqualsFunc
			result = tt.source.Except(tt.second, tt.fEquals)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), result.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)
		})

		t.Run(tt.name+"_ExceptBy", func(t *testing.T) {
			anySecond := asIEnumerableAny(tt.second)
			bSource := backupForAssetUnchanged(tt.source)
			bSecond := backupForAssetUnchanged(tt.second)

			if tt.panic && tt.source == nil {
				return
			}
			defer deferWantPanicDepends(t, tt.panic)

			// nil
			result := tt.source.ExceptBy(anySecond, SelfSelector[int](), nil)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), result.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)

			// EqualsFunc
			var optionalEqualsFunc OptionalEqualsFunc[any]
			if tt.fEquals != nil {
				optionalEqualsFunc = func(v1, v2 any) bool {
					return tt.fEquals(v1.(int), v2.(int))
				}
			} else {
				optionalEqualsFunc = nil
			}
			result = tt.source.ExceptBy(anySecond, SelfSelector[int](), optionalEqualsFunc)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), result.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)
		})
	}

	t.Run("auto-resolve comparer if default comparer not set", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](5, 2, 2, 6).
			WithDefaultComparer(nil)
		ieSecond := NewIEnumerable[int](5, 6, 7, 8)
		anySecond := asIEnumerableAny(ieSecond)
		ieWant := NewIEnumerable[int](2)

		bSrc := backupForAssetUnchanged(ieSrc)

		got := ieSrc.Except(ieSecond, nil)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		got = ieSrc.ExceptBy(anySecond, SelfSelector[int](), nil)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		var cft1 OptionalEqualsFunc[int]
		got = ieSrc.Except(ieSecond, cft1)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		var cft2 OptionalEqualsFunc[any]
		got = ieSrc.ExceptBy(anySecond, SelfSelector[int](), cft2)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		bSrc.assertUnchanged(t, ieSrc)

		assert.Nil(t, e[int](ieSrc).defaultComparer)
	})

	t.Run("returns distinct with key selector", func(t *testing.T) {
		type MyInt64 struct {
			Value int
		}
		ieSrc := NewIEnumerable[MyInt64](MyInt64{Value: 1}, MyInt64{Value: 2}, MyInt64{Value: 2}, MyInt64{Value: 3}, MyInt64{Value: 4})
		ieSecond := asIEnumerableAny(NewIEnumerable[int](1, 3))
		got := ieSrc.ExceptBy(ieSecond, func(v MyInt64) any {
			return v.Value
		}, nil).ToArray()
		assert.Equal(t, 2, len(got))
		assert.Equal(t, 2, got[0].Value)
		assert.Equal(t, 4, got[1].Value)
	})

	t.Run("panic if no default resolver (Except)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.Except(ieSrc, nil)
	})

	t.Run("panic if no default resolver (ExceptBy)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, getErrorFailedCompare2ElementsInArray().Error(), true)

		ieSrc.ExceptBy(asIEnumerableAny(ieSrc), SelfSelector[MyInt64](), nil)
	})
}
