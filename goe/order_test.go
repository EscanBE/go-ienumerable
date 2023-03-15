package goe

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_Order_OrderBy_OrderByComparer(t *testing.T) {
	fLess := func(t1, t2 int) bool {
		return t1 < t2
	}
	var tests = []struct {
		name     string
		source   IEnumerable[int]
		fLess    func(t1, t2 int) bool
		comparer comparers.IComparer[int]
		want     IEnumerable[int]
	}{
		{
			name:     "empty source returns empty",
			source:   createEmptyIntEnumerable(),
			fLess:    fLess,
			comparer: comparers.IntComparer,
			want:     createEmptyIntEnumerable(),
		},
		{
			name:     "not distinct (1)",
			source:   NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			fLess:    fLess,
			comparer: comparers.IntComparer,
			want:     NewIEnumerable[int](1, 2, 2, 3, 3, 4, 4, 5, 6, 6, 6),
		},
		{
			name:     "not distinct (2)",
			source:   NewIEnumerable[int](2, 2),
			fLess:    fLess,
			comparer: comparers.IntComparer,
			want:     NewIEnumerable[int](2, 2),
		},
		{
			name:     "no equality comparer still ok since int has default comparer",
			source:   NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			fLess:    nil,
			comparer: nil,
			want:     NewIEnumerable[int](1, 2, 2, 3, 3, 4, 4, 5, 6, 6, 6),
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Order_%s", tt.name), func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			got := tt.source.Order()

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: %v\nGot: %v\n", tt.want.ToArray(), got.ToArray())
			}

			bSrc.assertUnchanged(t, tt.source)

			// auto resolve when comparer not set
			cpSrc := tt.source.CastInt()
			e[int](cpSrc).defaultComparer = nil

			bSrc = backupForAssetUnchanged(cpSrc)

			got = cpSrc.Order()

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: %v\nGot: %v\n", tt.want.ToArray(), got.ToArray())
			}

			bSrc.assertUnchanged(t, cpSrc)
		})
		t.Run(fmt.Sprintf("OrderBy_%s", tt.name), func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			got := tt.source.OrderBy(tt.fLess)

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: %v\nGot: %v\n", tt.want.ToArray(), got.ToArray())
			}

			bSrc.assertUnchanged(t, tt.source)

			// auto resolve when comparer not set
			cpSrc := tt.source.CastInt()
			e[int](cpSrc).defaultComparer = nil

			bSrc = backupForAssetUnchanged(cpSrc)

			got = cpSrc.OrderBy(nil)

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: %v\nGot: %v\n", tt.want.ToArray(), got.ToArray())
			}

			bSrc.assertUnchanged(t, cpSrc)
		})
		t.Run(fmt.Sprintf("OrderByComparer_%s", tt.name), func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			got := tt.source.OrderByComparer(tt.comparer)

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: %v\nGot: %v\n", tt.want.ToArray(), got.ToArray())
			}

			bSrc.assertUnchanged(t, tt.source)

			// auto resolve when comparer not set
			cpSrc := tt.source.CastInt()
			e[int](cpSrc).defaultComparer = nil

			bSrc = backupForAssetUnchanged(cpSrc)

			got = cpSrc.OrderByComparer(nil)

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: %v\nGot: %v\n", tt.want.ToArray(), got.ToArray())
			}

			bSrc.assertUnchanged(t, cpSrc)
		})
	}

	t.Run("panic if no default resolver", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.Order()
	})

	t.Run("panic if no default resolver", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.OrderBy(nil)
	})

	t.Run("panic if no default resolver", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.OrderByComparer(nil)
	})
}
