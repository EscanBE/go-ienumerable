package goe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_Union_UnionBy(t *testing.T) {
	fEquals := func(i1, i2 int) bool {
		return i1 == i2
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
			name:    "union not any duplicated",
			source:  NewIEnumerable[int](1, 2, 3),
			second:  NewIEnumerable[int](4, 5, 6, 7),
			want:    NewIEnumerable[int](1, 2, 3, 4, 5, 6, 7),
			fEquals: fEquals,
		},
		{
			name:    "union returns distinct",
			source:  NewIEnumerable[int](5, 2, 2, 6),
			second:  NewIEnumerable[int](1, 2, 2, 3),
			want:    NewIEnumerable[int](5, 2, 6, 1, 3),
			fEquals: fEquals,
		},
		{
			name:    "union with duplicated",
			source:  NewIEnumerable[int](1, 2, 2, 3),
			second:  NewIEnumerable[int](1, 3),
			want:    NewIEnumerable[int](1, 2, 3),
			fEquals: fEquals,
		},
		{
			name:    "no comparer",
			source:  NewIEnumerable[int](1, 2, 3),
			second:  NewIEnumerable[int](4, 5, 6, 7),
			want:    NewIEnumerable[int](1, 2, 3, 4, 5, 6, 7),
			fEquals: nil,
		},
		{
			name:    "union one",
			source:  NewIEnumerable[int](1, 1, 1, 1, 1),
			second:  NewIEnumerable[int](1, 1, 1, 1),
			want:    NewIEnumerable[int](1),
			fEquals: fEquals,
		},
		{
			name:    "union some",
			source:  NewIEnumerable[int](2, 2, 2, 2, 3, 3, 3, 3),
			second:  NewIEnumerable[int](1, 1, 1, 1, 1, 2, 3),
			want:    NewIEnumerable[int](2, 3, 1),
			fEquals: fEquals,
		},
		{
			name:    "union all",
			source:  NewIEnumerable[int](1, 2, 3, 5, 6, 6),
			second:  NewIEnumerable[int](1, 1, 1, 3, 3, 3, 6, 6, 6, 5, 5, 5, 2, 2, 2),
			want:    NewIEnumerable[int](1, 2, 3, 5, 6),
			fEquals: fEquals,
		},
		{
			name:    "union when source empty",
			source:  NewIEnumerable[int](),
			second:  NewIEnumerable[int](4, 5, 6, 7),
			want:    NewIEnumerable[int](4, 5, 6, 7),
			fEquals: fEquals,
		},
		{
			name:    "union when second empty",
			source:  NewIEnumerable[int](1, 2, 3),
			second:  NewIEnumerable[int](),
			want:    NewIEnumerable[int](1, 2, 3),
			fEquals: fEquals,
		},
		{
			name:    "union when both empty",
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
		t.Run(tt.name+"_Union", func(t *testing.T) {
			bSource := backupForAssetUnchanged(tt.source)
			bSecond := backupForAssetUnchanged(tt.second)

			if tt.panic && tt.source == nil {
				return
			}
			defer deferWantPanicDepends(t, tt.panic)

			// Union
			got := tt.source.Union(tt.second, tt.fEquals)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)
		})

		// TODO
		//t.Run(tt.name+"_UnionBy", func(t *testing.T) {
		//	bSource := backupForAssetUnchanged(tt.source)
		//	bSecond := backupForAssetUnchanged(tt.second)
		//
		//	if tt.panic && tt.source == nil {
		//		return
		//	}
		//	defer deferWantPanicDepends(t, tt.panic)
		//
		//	// EqualsFunc
		//	got := tt.source.UnionBy(tt.second, tt.fEquals)
		//
		//	assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))
		//
		//	bSource.assertUnchanged(t, tt.source)
		//	bSecond.assertUnchanged(t, tt.second)
		//
		//	got = tt.source.UnionBy(tt.second, EqualsFunc[int](tt.fEquals))
		//
		//	assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))
		//
		//	bSource.assertUnchanged(t, tt.source)
		//	bSecond.assertUnchanged(t, tt.second)
		//
		//	// CompareFunc
		//	got = tt.source.UnionBy(tt.second, tt.fCompare)
		//
		//	assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))
		//
		//	bSource.assertUnchanged(t, tt.source)
		//	bSecond.assertUnchanged(t, tt.second)
		//
		//	got = tt.source.UnionBy(tt.second, CompareFunc[int](tt.fCompare))
		//
		//	assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))
		//
		//	bSource.assertUnchanged(t, tt.source)
		//	bSecond.assertUnchanged(t, tt.second)
		//
		//	// IComparer
		//	got = tt.source.UnionBy(tt.second, tt.comparer)
		//
		//	assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))
		//
		//	bSource.assertUnchanged(t, tt.source)
		//	bSecond.assertUnchanged(t, tt.second)
		//})
	}

	t.Run("auto-resolve comparer if default comparer is nil", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](5, 2, 2, 6).WithDefaultComparer(nil)
		ieSecond := NewIEnumerable[int](1, 2, 2, 3).WithDefaultComparer(nil)

		bSource := backupForAssetUnchanged(ieSrc)
		bSecond := backupForAssetUnchanged(ieSecond)

		// Union
		ieGot := ieSrc.Union(ieSecond, nil)
		assert.Equal(t, 5, ieGot.Count(nil))
		assert.Equal(t, 5, ieGot.ToArray()[0])
		assert.Equal(t, 2, ieGot.ToArray()[1])
		assert.Equal(t, 6, ieGot.ToArray()[2])
		assert.Equal(t, 1, ieGot.ToArray()[3])
		assert.Equal(t, 3, ieGot.ToArray()[4])

		bSource.assertUnchanged(t, ieSrc)
		bSecond.assertUnchanged(t, ieSecond)

		// TODO
		// UnionBy
		//ieGot = ieSrc.UnionBy(ieSecond, nil)
		//assert.Equal(t, 5, ieGot.Count(nil))
		//assert.Equal(t, 5, ieGot.ToArray()[0])
		//assert.Equal(t, 2, ieGot.ToArray()[1])
		//assert.Equal(t, 6, ieGot.ToArray()[2])
		//assert.Equal(t, 1, ieGot.ToArray()[3])
		//assert.Equal(t, 3, ieGot.ToArray()[4])
		//
		//bSource.assertUnchanged(t, ieSrc)
		//bSecond.assertUnchanged(t, ieSecond)
	})

	t.Run("panic if no default resolver (Union)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.Union(ieSrc, nil)
	})

	// TODO
	//t.Run("panic if no default resolver (UnionBy)", func(t *testing.T) {
	//	type MyInt64 struct{}
	//	ieSrc := NewIEnumerable[MyInt64]()
	//
	//	defer deferExpectPanicContains(t, "no default comparer registered", true)
	//
	//	ieSrc.UnionBy(ieSrc, nil)
	//})

	// TODO
	//t.Run("panic if not supported comparer", func(t *testing.T) {
	//	ieSrc := NewIEnumerable[int]()
	//
	//	defer deferExpectPanicContains(t, "comparer must be", true)
	//
	//	var badFunc func(v int) bool
	//	ieSrc.UnionBy(ieSrc, badFunc)
	//})

	// TODO
	//t.Run("panic if not supported comparer", func(t *testing.T) {
	//	ieSrc := NewIEnumerable[int](1)
	//
	//	defer deferExpectPanicContains(t, "comparer must be", true)
	//
	//	var badFunc LessFunc[int]
	//	ieSrc.UnionBy(ieSrc, badFunc)
	//})

	// TODO
	//t.Run("panic if not supported comparer", func(t *testing.T) {
	//	ieSrc := NewIEnumerable[int](1)
	//
	//	defer deferExpectPanicContains(t, "comparer must be", true)
	//
	//	var badFunc GreaterFunc[int]
	//	ieSrc.UnionBy(ieSrc, badFunc)
	//})
}
