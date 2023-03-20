package goe

import (
	"fmt"
	comparers "github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_Order_OrderBy_OrderByComparer(t *testing.T) {
	fCompareFunc := func(t1, t2 any) int {
		return comparers.NumericComparer.CompareTyped(t1, t2)
	}
	var tests = []struct {
		name     string
		source   IEnumerable[int]
		fCompare func(t1, t2 any) int
		want     IEnumerable[int]
	}{
		{
			name:     "empty source returns empty",
			source:   createEmptyIntEnumerable(),
			fCompare: fCompareFunc,
			want:     createEmptyIntEnumerable(),
		},
		{
			name:     "not distinct (1)",
			source:   NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			fCompare: fCompareFunc,
			want:     NewIEnumerable[int](1, 2, 2, 3, 3, 4, 4, 5, 6, 6, 6),
		},
		{
			name:     "not distinct (2)",
			source:   NewIEnumerable[int](2, 2),
			fCompare: fCompareFunc,
			want:     NewIEnumerable[int](2, 2),
		},
		{
			name:     "automatically resolve comparer",
			source:   NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			fCompare: nil,
			want:     NewIEnumerable[int](1, 2, 2, 3, 3, 4, 4, 5, 6, 6, 6),
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Order_%s", tt.name), func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			got := tt.source.Order().GetOrderedEnumerable()

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: %v\nGot: %v\n", tt.want.ToArray(), got.ToArray())
			}

			bSrc.assertUnchanged(t, tt.source)

			// auto resolve when comparer not set
			cpSrc := tt.source.CastInt().
				WithDefaultComparer(nil)

			bSrc = backupForAssetUnchanged(cpSrc)

			got = cpSrc.Order().GetOrderedEnumerable()

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: %v\nGot: %v\n", tt.want.ToArray(), got.ToArray())
			}

			bSrc.assertUnchanged(t, cpSrc)

			assert.Nil(t, e[int](cpSrc).defaultComparer)
		})
		t.Run(fmt.Sprintf("OrderBy_%s", tt.name), func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			got := tt.source.OrderBy(test_getSelfSelector[int](), tt.fCompare).GetOrderedEnumerable()

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: %v\nGot: %v\n", tt.want.ToArray(), got.ToArray())
			}

			bSrc.assertUnchanged(t, tt.source)

			// auto resolve when comparer not set
			cpSrc := tt.source.CastInt()

			bSrc = backupForAssetUnchanged(cpSrc)

			got = cpSrc.OrderBy(test_getSelfSelector[int](), nil).GetOrderedEnumerable()

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: %v\nGot: %v\n", tt.want.ToArray(), got.ToArray())
			}

			bSrc.assertUnchanged(t, cpSrc)

			assert.Nil(t, e[int](cpSrc).defaultComparer)
		})
	}

	t.Run("default comparer resolved based on key selector", func(t *testing.T) {
		type MyInt64 struct {
			Value int
		}
		ieSrc := NewIEnumerable[MyInt64](MyInt64{
			Value: 2,
		}, MyInt64{
			Value: 1,
		}, MyInt64{
			Value: 4,
		}, MyInt64{
			Value: 3,
		})

		got := ieSrc.OrderBy(func(my MyInt64) any {
			return my.Value
		}, nil).GetOrderedEnumerable().ToArray()

		assert.Equal(t, 4, len(got))
		assert.Equal(t, 1, got[0].Value)
		assert.Equal(t, 2, got[1].Value)
		assert.Equal(t, 3, got[2].Value)
		assert.Equal(t, 4, got[3].Value)
	})

	t.Run("use custom comparer for type without default comparer", func(t *testing.T) {
		type MyInt64 struct {
			Value int
		}
		ieSrc := NewIEnumerable[MyInt64](MyInt64{
			Value: 2,
		}, MyInt64{
			Value: 1,
		}, MyInt64{
			Value: 4,
		}, MyInt64{
			Value: 3,
		})

		got := ieSrc.OrderBy(func(me MyInt64) any {
			return me
		}, func(x, y any) int {
			return comparers.NumericComparer.CompareTyped(x.(MyInt64).Value, y.(MyInt64).Value)
		}).GetOrderedEnumerable().ToArray()

		assert.Equal(t, 4, len(got))
		assert.Equal(t, 1, got[0].Value)
		assert.Equal(t, 2, got[1].Value)
		assert.Equal(t, 3, got[2].Value)
		assert.Equal(t, 4, got[3].Value)
	})

	t.Run("panic if no default comparer", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.Order()
	})

	t.Run("panic if no default comparer", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.OrderBy(test_getSelfSelector[MyInt64](), nil)
	})
}

func Test_enumerable_OrderByDescending_OrderByDescendingBy_OrderByDescendingByComparer(t *testing.T) {
	fGreater := func(t1, t2 int) bool {
		return t1 > t2
	}
	var tests = []struct {
		name     string
		source   IEnumerable[int]
		fGreater func(t1, t2 int) bool
		comparer comparers.IComparer[any]
		want     IEnumerable[int]
	}{
		{
			name:     "empty source returns empty",
			source:   createEmptyIntEnumerable(),
			fGreater: fGreater,
			comparer: comparers.NumericComparer,
			want:     createEmptyIntEnumerable(),
		},
		{
			name:     "not distinct (1)",
			source:   NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			fGreater: fGreater,
			comparer: comparers.NumericComparer,
			want:     NewIEnumerable[int](6, 6, 6, 5, 4, 4, 3, 3, 2, 2, 1),
		},
		{
			name:     "not distinct (2)",
			source:   NewIEnumerable[int](2, 2),
			fGreater: fGreater,
			comparer: comparers.NumericComparer,
			want:     NewIEnumerable[int](2, 2),
		},
		{
			name:     "no equality comparer still ok since int has default comparer",
			source:   NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			fGreater: nil,
			comparer: nil,
			want:     NewIEnumerable[int](6, 6, 6, 5, 4, 4, 3, 3, 2, 2, 1),
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("OrderByDescending_%s", tt.name), func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			got := tt.source.OrderByDescending().GetOrderedEnumerable()

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: %v\nGot: %v\n", tt.want.ToArray(), got.ToArray())
			}

			bSrc.assertUnchanged(t, tt.source)

			// auto resolve when comparer not set
			cpSrc := tt.source.CastInt().
				WithDefaultComparer(nil)

			bSrc = backupForAssetUnchanged(cpSrc)

			got = cpSrc.OrderByDescending().GetOrderedEnumerable()

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: %v\nGot: %v\n", tt.want.ToArray(), got.ToArray())
			}

			bSrc.assertUnchanged(t, cpSrc)

			assert.Nil(t, e[int](cpSrc).defaultComparer)
		})
		t.Run(fmt.Sprintf("OrderByDescendingBy_%s", tt.name), func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			got := tt.source.OrderByDescendingBy(test_getSelfSelector[int](), nil).GetOrderedEnumerable()

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: %v\nGot: %v\n", tt.want.ToArray(), got.ToArray())
			}

			bSrc.assertUnchanged(t, tt.source)
		})
	}

	t.Run("panic if no default comparer", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.OrderByDescending().GetOrderedEnumerable()
	})

	t.Run("panic if no default comparer", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.OrderByDescendingBy(test_getSelfSelector[MyInt64](), nil)
	})

	t.Run("default comparer resolved based on key selector", func(t *testing.T) {
		type MyInt64 struct {
			Value int
		}
		ieSrc := NewIEnumerable[MyInt64](MyInt64{
			Value: 2,
		}, MyInt64{
			Value: 1,
		}, MyInt64{
			Value: 4,
		}, MyInt64{
			Value: 3,
		})

		got := ieSrc.OrderByDescendingBy(func(my MyInt64) any {
			return my.Value
		}, nil).GetOrderedEnumerable().ToArray()

		assert.Equal(t, 4, len(got))
		assert.Equal(t, 4, got[0].Value)
		assert.Equal(t, 3, got[1].Value)
		assert.Equal(t, 2, got[2].Value)
		assert.Equal(t, 1, got[3].Value)
	})

	t.Run("use custom comparer for type without default comparer", func(t *testing.T) {
		type MyInt64 struct {
			Value int
		}
		ieSrc := NewIEnumerable[MyInt64](MyInt64{
			Value: 2,
		}, MyInt64{
			Value: 1,
		}, MyInt64{
			Value: 4,
		}, MyInt64{
			Value: 3,
		})

		got := ieSrc.OrderByDescendingBy(func(me MyInt64) any {
			return me
		}, func(x, y any) int {
			return comparers.NumericComparer.CompareTyped(x.(MyInt64).Value, y.(MyInt64).Value)
		}).GetOrderedEnumerable().ToArray()

		assert.Equal(t, 4, len(got))
		assert.Equal(t, 4, got[0].Value)
		assert.Equal(t, 3, got[1].Value)
		assert.Equal(t, 2, got[2].Value)
		assert.Equal(t, 1, got[3].Value)
	})
}
