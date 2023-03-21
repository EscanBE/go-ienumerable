package goe

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_Distinct_DistinctBy(t *testing.T) {
	fEquals := func(v1, v2 int) bool {
		return comparers.NumericComparer.CompareAny(v1, v2) == 0
	}
	var tests = []struct {
		name    string
		source  IEnumerable[int]
		fEquals OptionalEqualsFunc[int]
		want    IEnumerable[int]
	}{
		{
			name:    "empty source",
			source:  createEmptyIntEnumerable(),
			fEquals: fEquals,
			want:    createEmptyIntEnumerable(),
		},
		{
			name:    "distinct",
			source:  NewIEnumerable[int](2),
			fEquals: fEquals,
			want:    NewIEnumerable[int](2),
		},
		{
			name:    "distinct",
			source:  NewIEnumerable[int](2, 2),
			fEquals: fEquals,
			want:    NewIEnumerable[int](2),
		},
		{
			name:    "no equality comparer still ok since int has default comparer",
			source:  NewIEnumerable[int](2),
			fEquals: nil,
			want:    NewIEnumerable[int](2),
		},
		{
			name:    "keep the same order",
			source:  NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			fEquals: fEquals,
			want:    NewIEnumerable[int](1, 2, 3, 6, 5, 4),
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Distinct-%s", tt.name), func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			got := tt.source.Distinct(tt.fEquals)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))

			bSrc.assertUnchanged(t, tt.source)
		})
		t.Run(fmt.Sprintf("DistinctBy-%s", tt.name), func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			var equalsFuncAny OptionalEqualsFunc[any]
			if tt.fEquals != nil {
				equalsFuncAny = func(v1, v2 any) bool {
					return tt.fEquals(v1.(int), v2.(int))
				}
			}

			got := tt.source.DistinctBy(test_getSelfSelector[int](), equalsFuncAny)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))

			bSrc.assertUnchanged(t, tt.source)
		})
	}

	t.Run("auto-resolve comparer if default comparer not set", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4).
			WithDefaultComparer(nil)
		ieWant := NewIEnumerable[int](1, 2, 3, 6, 5, 4)

		bSrc := backupForAssetUnchanged(ieSrc)

		got := ieSrc.Distinct(nil)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		var cff func(v1, v2 int) bool
		got = ieSrc.Distinct(OptionalEqualsFunc[int](cff))
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))
		var cft OptionalEqualsFunc[int]
		got = ieSrc.Distinct(cft)
		assert.True(t, reflect.DeepEqual(ieWant.ToArray(), got.ToArray()))

		assert.Nil(t, e[int](ieSrc).defaultComparer)

		bSrc.assertUnchanged(t, ieSrc)
	})

	t.Run("panic if no default resolver (Distinct)", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.Distinct(nil)
	})
}

func Test_enumerable_DistinctBy(t *testing.T) {
	t.Run("optional equality comparer", func(t *testing.T) {
		eSrc := NewIEnumerable[int](2, 2)
		bSrc := backupForAssetUnchanged(eSrc)

		got := eSrc.DistinctBy(test_getSelfSelector[int](), nil).ToArray()
		assert.Len(t, got, 1)

		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("optional equality comparer", func(t *testing.T) {
		type MyType struct {
			Value int
		}
		eSrc := NewIEnumerable[MyType](MyType{Value: 1}, MyType{Value: 1})
		bSrc := backupForAssetUnchanged(eSrc)
		got := eSrc.DistinctBy(func(v MyType) any {
			return v.Value
		}, nil).ToArray()
		assert.Len(t, got, 1)
		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("custom equality comparer", func(t *testing.T) {
		type MyType struct {
			Value int
		}
		eSrc := NewIEnumerable[MyType](MyType{Value: 1}, MyType{Value: 1})
		bSrc := backupForAssetUnchanged(eSrc)
		got := eSrc.DistinctBy(test_getSelfSelector[MyType](), func(v1, v2 any) bool {
			return v1.(MyType).Value == v2.(MyType).Value
		}).ToArray()
		assert.Len(t, got, 1)
		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("panic if not able to resolve comparer when absent", func(t *testing.T) {
		type MyType struct{}
		eSrc := NewIEnumerable[MyType]()
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		defer deferExpectPanicContains(t, getErrorFailedCompare2ElementsInArray().Error(), true)
		eSrc.DistinctBy(test_getSelfSelector[MyType](), nil)
	})

	t.Run("panic if missing key selector", func(t *testing.T) {
		eSrc := NewIEnumerable[int](2, 2)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()
		defer deferExpectPanicContains(t, getErrorKeySelectorNotNil().Error(), true)
		_ = eSrc.DistinctBy(nil, nil)
	})

	feq := OptionalEqualsFunc[any](func(v1, v2 any) bool {
		return v1 == v2
	})

	t.Run("distinct empty", func(t *testing.T) {
		r := NewIEnumerable[int]().DistinctBy(test_getSelfSelector[int](), feq).ToArray()
		assert.Len(t, r, 0)
	})

	t.Run("distinct one", func(t *testing.T) {
		r := NewIEnumerable[int](9).DistinctBy(test_getSelfSelector[int](), feq).ToArray()
		assert.Len(t, r, 1)
	})

	t.Run("distinct two", func(t *testing.T) {
		r := NewIEnumerable[int](99, 99).DistinctBy(test_getSelfSelector[int](), feq).ToArray()
		assert.Len(t, r, 1)
		assert.Equal(t, 99, r[0])
	})

	t.Run("distinct many", func(t *testing.T) {
		r := NewIEnumerable[int](99, 99, 66, 66).DistinctBy(test_getSelfSelector[int](), feq).ToArray()
		assert.Len(t, r, 2)
		assert.Equal(t, 99, r[0])
		assert.Equal(t, 66, r[1])
	})
}

func Test_distinct(t *testing.T) {
	t.Run("optional equality comparer", func(t *testing.T) {
		got := distinct[int]([]int{2, 2}, nil)
		assert.Len(t, got, 1)
	})

	t.Run("panic if not able to resolve comparer when absent", func(t *testing.T) {
		type MyType struct{}
		defer deferExpectPanicContains(t, getErrorFailedCompare2ElementsInArray().Error(), true)
		distinct[MyType]([]MyType{}, nil)
	})

	feq := OptionalEqualsFunc[int](func(v1, v2 int) bool {
		return v1 == v2
	})

	t.Run("distinct empty", func(t *testing.T) {
		r := distinct[int]([]int{}, feq)
		assert.Len(t, r, 0)
	})

	t.Run("distinct one", func(t *testing.T) {
		r := distinct[int]([]int{9}, feq)
		assert.Len(t, r, 1)
	})

	t.Run("distinct two", func(t *testing.T) {
		r := distinct[int]([]int{99, 99}, feq)
		assert.Len(t, r, 1)
		assert.Equal(t, 99, r[0])
	})

	t.Run("distinct many", func(t *testing.T) {
		r := distinct[int]([]int{99, 99, 66, 66}, feq)
		assert.Len(t, r, 2)
		assert.Equal(t, 99, r[0])
		assert.Equal(t, 66, r[1])
	})
}

func Test_distinctByKeySelector(t *testing.T) {
	t.Run("optional equality comparer", func(t *testing.T) {
		got := distinctByKeySelector[int]([]int{2, 2}, test_getSelfSelector[int](), nil)
		assert.Len(t, got, 1)
	})

	t.Run("optional equality comparer", func(t *testing.T) {
		type MyType struct {
			Value int
		}
		got := distinctByKeySelector[MyType]([]MyType{
			{Value: 1}, {Value: 1},
		}, func(v MyType) any {
			return v.Value
		}, nil)
		assert.Len(t, got, 1)
	})

	t.Run("custom equality comparer", func(t *testing.T) {
		type MyType struct {
			Value int
		}
		got := distinctByKeySelector[MyType]([]MyType{
			{Value: 1}, {Value: 1},
		}, test_getSelfSelector[MyType](), func(v1, v2 any) bool {
			return v1.(MyType).Value == v2.(MyType).Value
		})
		assert.Len(t, got, 1)
	})

	t.Run("panic if not able to resolve comparer when absent", func(t *testing.T) {
		type MyType struct{}
		defer deferExpectPanicContains(t, getErrorFailedCompare2ElementsInArray().Error(), true)
		distinctByKeySelector[MyType]([]MyType{}, test_getSelfSelector[MyType](), nil)
	})

	t.Run("panic if missing key selector", func(t *testing.T) {
		defer deferExpectPanicContains(t, getErrorKeySelectorNotNil().Error(), true)
		_ = distinctByKeySelector[int]([]int{2, 2}, nil, nil)
	})

	feq := OptionalEqualsFunc[any](func(v1, v2 any) bool {
		return v1 == v2
	})

	t.Run("distinct empty", func(t *testing.T) {
		r := distinctByKeySelector[int]([]int{}, test_getSelfSelector[int](), feq)
		assert.Len(t, r, 0)
	})

	t.Run("distinct one", func(t *testing.T) {
		r := distinctByKeySelector[int]([]int{9}, test_getSelfSelector[int](), feq)
		assert.Len(t, r, 1)
	})

	t.Run("distinct two", func(t *testing.T) {
		r := distinctByKeySelector[int]([]int{99, 99}, test_getSelfSelector[int](), feq)
		assert.Len(t, r, 1)
		assert.Equal(t, 99, r[0])
	})

	t.Run("distinct many", func(t *testing.T) {
		r := distinctByKeySelector[int]([]int{66, 66, 99, 99}, test_getSelfSelector[int](), feq)
		assert.Len(t, r, 2)
		assert.Equal(t, 66, r[0])
		assert.Equal(t, 99, r[1])
	})
}
